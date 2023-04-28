package avs

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	searchtree "our_antivirus/av/search-tree"
	"our_antivirus/av/utils"

	"github.com/beevik/prefixtree"
)

var ErrSignatureFoundInFile = errors.New("Signature was found in file")

func FindInFile(fpath string, signTree *prefixtree.Tree) (*searchtree.SignTree, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return FindS(f.Name(), f, signTree)
}

func FindS(filename string, f io.ReaderAt, signTree *prefixtree.Tree) (*searchtree.SignTree, error) {
	// создаем буфер на 2 символа
	byteSlice := make([]byte, 2)
	// начинаем с начала файла
	curOffset := 0
	for {
		// читаем байты в буфер
		_, err := f.ReadAt(byteSlice, int64(curOffset))
		if err != nil {
			if err == utils.ErrInvalidOffset {
				// смещаем указатель на 1 байт
				curOffset++
				// и возвращаемся вверх цикла
				continue
			}

			return nil, err
		}

		// ищем 2 символа из буфера в префиксном дереве
		s, err := signTree.Find(string(byteSlice))
		if err != nil {
			// ошибка конец файла -> выходим из цикла
			if err == io.EOF {
				break
			}
			// ошибка префикс не найден
			if err == prefixtree.ErrPrefixNotFound {
				// смещаем указатель на 1 байт
				curOffset++
				// и возвращаемся вверх цикла
				continue
			}

			if err == prefixtree.ErrPrefixAmbiguous {
				break
			}

			// неизвестная ошибка выходим
			return nil, err
		}

		// чет найдено, пытаемся преобразовать
		if d, ok := s.(*searchtree.SignTree); ok {

			if err := checkFoundedSignatureInFileSlice(d, f, int64(curOffset)); err != nil {
				if err == ErrSignatureFoundInFile {
					return d, ErrSignatureFoundInFile
				}

				log.Println("error", err)

				curOffset++
				continue
			}

			offset, err := d.Offset()

			if err != nil {
				curOffset++
				continue
			}

			// Определяем, что оффсет/тип файла совпадает
			fileExtension := strings.Trim(filepath.Ext(filename), ".")
			if offset != int64(curOffset) || TypeDefinition(fileExtension) != d.DType() {
				// смещаем указатель на 1 байт
				curOffset++
				// и возвращаемся вверх цикла
				continue
			}

			return d, ErrSignatureFoundInFile // yшиб очка
		}

		// не получилось преобразовать
		return nil, errors.New("signature data corrupted")
	}
	return nil, nil
}

func checkFoundedSignatureInFileSlice(d *searchtree.SignTree, f io.ReaderAt, curOffset int64) error {
	offsetStart, err := d.Offset()
	if err != nil {
		return err
	}
	offsetEnd, err := d.OffsetEnd()
	if err != nil {
		return err
	}
	signLength := offsetEnd - offsetStart

	signBytes := make([]byte, signLength)
	var n int
	for {
		n, err = f.ReadAt(signBytes, curOffset)
		if err != nil {
			if err == utils.ErrInvalidOffset {
				// смещаем указатель на 1 байт
				curOffset++
				continue
			}

			return err
		}

		break
	}

	if n < int(signLength) {
		return errors.New("read less bytes than sign")
	}

	if bytes.Equal(d.B, signBytes) {
		return ErrSignatureFoundInFile
	}

	return nil
}

func TypeDefinition(fileExtension string) string {
	types := map[string]string{
		"EXE": "PE",
		"COM": "COM",
	}

	value, _ := types[strings.ToUpper(fileExtension)]

	return value
}

var signSearchStats map[string][]string

func SearchResults() map[string][]string {
	return signSearchStats
}

func generateFilepathWalkFunction(tree *prefixtree.Tree, signSearchStats map[string][]string) func(path string, info fs.FileInfo, err error) error {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		log.Println("scan file: ", path)
		// Определяем расширение файла
		fileExtension := strings.ToLower(strings.Trim(filepath.Ext(path), "."))

		if fileExtension == "zip" {
			if scanResult, err := ScanZipFile(path, tree); err != nil {
				return err
			} else {
				for p, signs := range scanResult {
					signSearchStats[p] = signs
				}
			}

			return nil
		} else if fileExtension == "7z" {
			if scanResult, err := Scan7ZFile(path, tree); err != nil {
				return err
			} else {
				for p, signs := range scanResult {
					signSearchStats[p] = signs
				}
			}

			return nil
		}

		detectedSignature, err := FindInFile(path, tree)
		if handleError(err) != nil {
			return err
		}

		if detectedSignature != nil {
			if _, ok := signSearchStats[path]; !ok {
				signSearchStats[path] = []string{}
			}

			signSearchStats[path] = append(signSearchStats[path], detectedSignature.Name())
		} else {
			// файл чист
			signSearchStats[path] = nil
		}

		return nil
	}
}

func handleError(err error) error {
	if err != nil {
		switch err {
		case ErrSignatureFoundInFile, io.EOF:
			return nil
		}

		return err
	}

	return nil
}

func FindSignatures(searchlocation string, tree *prefixtree.Tree) error {
	signSearchStats = make(map[string][]string)

	err := filepath.Walk(searchlocation, generateFilepathWalkFunction(tree, signSearchStats))

	if err != nil {
		return err
	}

	return nil
}

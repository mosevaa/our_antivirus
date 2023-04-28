package avs

import (
	"archive/zip"
	"log"
	"our_antivirus/av/utils"

	"github.com/beevik/prefixtree"
	"github.com/bodgit/sevenzip"
)

func Scan7ZFile(
	filePath string,
	signTree *prefixtree.Tree,
) (map[string][]string, error) {
	document, err := sevenzip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}
	defer document.Close()

	log.Printf("[%v] Scanning file as %s...\n", filePath, "7z")

	// Создаем слайс найденых сигнатур в файлах
	signSearchStats := map[string][]string{}

	for _, sevenZFile := range document.File {
		if !sevenZFile.FileInfo().IsDir() {
			log.Printf("[%v] Scan file in %s: %v\n", filePath, "7z", sevenZFile.Name)
			fname := filePath + "://" + sevenZFile.Name

			file, err := sevenZFile.Open()
			if err != nil {
				log.Printf("[%v] %v\n", filePath, err.Error())
				continue
			}
			defer file.Close()

			detectedSignature, err := FindS(sevenZFile.FileInfo().Name(), utils.NewUnbufferedReaderAt(file), signTree)
			if handleError(err) != nil {
				log.Printf("[%v] %v\n", filePath, err.Error())
				continue
			}

			if detectedSignature != nil {
				if _, ok := signSearchStats[fname]; !ok {
					signSearchStats[fname] = []string{}
				}

				signSearchStats[fname] = append(signSearchStats[fname], detectedSignature.Name())
			} else {
				// файл чист
				signSearchStats[fname] = nil
			}
		}
	}

	return signSearchStats, nil
}

func ScanZipFile(
	filePath string,
	signTree *prefixtree.Tree,
) (map[string][]string, error) {
	document, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}
	log.Printf("[%v] Scanning file as zip...\n", filePath)
	defer document.Close()

	// Создаем слайс найденых сигнатур в файлах
	signSearchStats := map[string][]string{}

	for _, zipFile := range document.File {
		if !zipFile.FileInfo().IsDir() {
			log.Printf("[%v] Scan file in zip: %v\n", filePath, zipFile.Name)
			fname := filePath + "://" + zipFile.Name

			file, err := zipFile.Open()
			if err != nil {
				log.Printf("[%v] %v\n", filePath, err.Error())
				continue
			}
			defer file.Close()

			detectedSignature, err := FindS(zipFile.FileInfo().Name(), utils.NewUnbufferedReaderAt(file), signTree)
			if handleError(err) != nil {
				log.Printf("[%v] %v\n", filePath, err.Error())
				continue
			}

			if detectedSignature != nil {
				if _, ok := signSearchStats[fname]; !ok {
					signSearchStats[fname] = []string{}
				}

				signSearchStats[fname] = append(signSearchStats[fname], detectedSignature.Name())
			} else {
				// файл чист
				signSearchStats[fname] = nil
			}
		}
	}

	return signSearchStats, nil
}

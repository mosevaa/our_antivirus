package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"our_antivirus/av/avs"
	"our_antivirus/av/database"
	searchtree "our_antivirus/av/search-tree"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// основная часть программы
func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}
	//scan.Scanning()

	//pipe.Pipe()

	// положить файл в карантин
	//scan.Quarantine("./../folder/file.txt", "quarantine")

	// убрать файл из карантина
	//scan.Restore("quarantine/file.txt", "../folder/file.txt")
}

func run() error {

	log.Println("Programm started")
	if err := database.Open(); err != nil {
		return err
	}
	log.Println("Database connection opened")
	defer database.Close()

	if err := searchtree.BuildSearchTree(); err != nil {
		return err
	}

	ptree := searchtree.GetSearchTree()
	log.Println("Prefix tree loaded")

	cmd := flag.String("cmd", "", "")
	flag.Parse()

	if strings.Contains(string(*cmd), "scan_all") {
		fmt.Println("scan_all")

		link := string(*cmd)[9:len(string(*cmd))]
		fmt.Println(link)
		// вот тут указать  файл
		if err := avs.FindSignatures(link, ptree); err != nil {
			return err
		}
		printResults()

	} else if strings.Contains(string(*cmd), "quarantine") {
		fmt.Println("quarantine")

		quarantineDir := "../quarantine"
		whitelistDir := "../quarantine_white_list"

		// Проверяем, существует ли каталог карантины
		if _, err := os.Stat(quarantineDir); os.IsNotExist(err) {
			fmt.Printf("Каталог %s не существует\n", quarantineDir)
			return err
		}

		// Создаем каталог quarantine_white_list, если он еще не существует
		if _, err := os.Stat(whitelistDir); os.IsNotExist(err) {
			os.Mkdir(whitelistDir, 0755)
		}

		// Получаем список файлов в каталоге карантины
		files, err := ioutil.ReadDir(quarantineDir)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Переносим каждый файл в папку quarantine_white_list и устанавливаем разрешение 0644
		for _, file := range files {
			oldPath := quarantineDir + "/" + file.Name()
			newPath := whitelistDir + "/" + file.Name()
			err = os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Printf("Не удалось переместить файл %s: %v\n", file.Name(), err)
			} else {
				err = os.Chmod(newPath, 0644)
				if err != nil {
					fmt.Printf("Не удалось установить разрешение на файл %s: %v\n", file.Name(), err)
				} else {
					fmt.Printf("Файл %s перемещен в %s и установлено разрешение 0644\n", file.Name(), newPath)
				}
			}
		}

	} else {
		fmt.Println("error")
	}

	return nil
}

func printResults() {
	searchResults := avs.SearchResults()

	if b, err := json.MarshalIndent(searchResults, "", "\t"); err == nil {
		log.Printf("Scan verbose results: %+v\n", string(b))
	}

	infectedFilesScan := map[string]struct{}{}
	for file, signStats := range searchResults {
		if len(signStats) > 0 {
			infectedFilesScan[file] = struct{}{}
		}
	}
	infectedFiles := []string{}
	for file, _ := range infectedFilesScan {
		infectedFiles = append(infectedFiles, file)
	}
	valid := len(searchResults) - len(infectedFiles)

	fmt.Printf(
		color.New(color.FgGreen).Add(color.Bold).Sprintf("Scanned files: ") + strconv.Itoa(len(searchResults)) +
			color.New(color.FgGreen).Add(color.Bold).Sprintf("\nValid documents: ") + strconv.Itoa(valid) +
			color.New(color.FgRed).Add(color.Bold).Sprintf("\nInfected files (%v): ", len(infectedFiles)) + strings.Join(infectedFiles, ", ") + "\n",
	)
}

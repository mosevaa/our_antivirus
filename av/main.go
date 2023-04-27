package main

import (
	"flag"
	"fmt"
	"our_antivirus/av/db"
	"our_antivirus/av/scan"
	"strings"
)

// основная часть программы
func main() {

	cmd := flag.String("cmd", "", "")
	flag.Parse()
	// fmt.Printf("my cmd: \"%v\"\n", string(*cmd))

	if strings.Contains(string(*cmd), "scan_all") {
		fmt.Println("scan_all")
		scan.Scanning()
	} else if strings.Contains(string(*cmd), "quarantine") {
		fmt.Println("quarantine")
		scan.Restore("quarantine/file.txt", "../folder/file.txt")
	} else if strings.Contains(string(*cmd), "scan_one_file") {
		fmt.Println("scan_one_file")
		scan.ScanFile("//path_to_file")
	} else if strings.Contains(string(*cmd), "start") {
		fmt.Println("start")
	} else if strings.Contains(string(*cmd), "stop") {
		fmt.Println("stop")
	} else {
		fmt.Println("error")
	}

	db.Connect_db()

	//scan.Scanning()

	//pipe.Pipe()

	// положить файл в карантин
	//scan.Quarantine("./../folder/file.txt", "quarantine")

	// убрать файл из карантина
	//scan.Restore("quarantine/file.txt", "../folder/file.txt")

	//--------------------
	// signatures := []string{
	// 	"Signature 1",
	// 	"Signature 2",
	// 	"Signature 3",
	// }

	// // Открыть файл для записи
	// file, err := os.Create("../folder/signatures.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// // Записать каждую сигнатуру в файл
	// for _, sig := range signatures {
	// 	_, err := file.WriteString(sig + "\n")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// fmt.Println("Сигнатуры записаны в файл signatures.txt")
	//--------------------
}

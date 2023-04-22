package main

import (
	"our_antivirus/av/db"
	"our_antivirus/av/pipe"
	"our_antivirus/av/scan"
)

// основная часть программы
func main() {
	db.Connect_db()

	scan.Scanning()

	pipe.Pipe()

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

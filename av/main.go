package main

import (
	"our_antivirus/av/db"
	"our_antivirus/av/scan"
)

// основная часть программы
func main() {
	db.Connect_db()

	// обход всех файлов и директорий
	scan.Scanning()
}

package scan

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// ScanFile scans th	e specified file and performs actions on it.
func ScanFile(filePath string) {
	fmt.Println("Scanning file:", filePath)

	// проверка если файл содержит сигнатуры
	//Quarantine(filePath, "quarantine")
	// Список сигнатур
	signatures := []string{
		"Signature 1",
		"Signature 2",
		"Signature 3",
	}

	// Проверяем файл на наличие сигнатур
	if checkSignatures(filePath, signatures) {
		fmt.Println("File contains one of the signatures")
		Quarantine(filePath, "quarantine")
	} else {
		fmt.Println("File does not contain any of the signatures")
	}
}

// ScanFolder scans the specified folder recursively and performs actions on files.
func ScanFolder(folderPath string) {
	fmt.Println("Scanning folder:", folderPath)
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}
		if !info.IsDir() {
			// Perform scanning action on the file
			ScanFile(path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

// MonitorDirectory monitors the specified directory for new files or files copied/moved into it.
func MonitorDirectory(dirPath string) {
	fmt.Println("Monitoring directory:", dirPath)
	for {
		time.Sleep(1 * time.Second) // Sleep for 1 second
		files, err := ioutil.ReadDir(dirPath)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, file := range files {
			// Perform scanning action on the new file
			ScanFile(filepath.Join(dirPath, file.Name()))
		}
	}
}

// ScheduleScan scans the specified folder according to a schedule.
func ScheduleScan(folderPath string, schedule time.Duration) {
	fmt.Println("Scheduling scan for folder:", folderPath)
	for {
		time.Sleep(schedule)
		// Perform scanning action on the folder
		ScanFolder(folderPath)
	}
}

// Quarantine moves the specified file to a quarantine directory.
func Quarantine(filePath, quarantinePath string) error {
	// Получаем информацию о файле.
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	// Создаем папку карантина, если она еще не создана.
	if err := os.MkdirAll(quarantinePath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create quarantine directory: %v", err)
	}

	// Получаем путь для перемещения файла в папку карантина.
	newPath := filepath.Join(quarantinePath, fileInfo.Name())

	// Перемещаем файл в папку карантина.
	if err := os.Rename(filePath, newPath); err != nil {
		return fmt.Errorf("failed to move file to quarantine: %v", err)
	}

	// Устанавливаем права доступа только на чтение для файла.
	if err := os.Chmod(newPath, 0444); err != nil {
		return fmt.Errorf("failed to set read-only permission for file: %v", err)
	}

	fmt.Println("File was successfully quarantined:", newPath)
	return nil
}

// Restore restores a file from quarantine to its original location
// and restores its original permissions.
func Restore(quarantinePath, restorePath string) error {
	// Construct the quarantine file path.
	quarantineFilePath := quarantinePath

	// Move the file from quarantine to the restore path.
	if err := os.Rename(quarantineFilePath, restorePath); err != nil {
		return fmt.Errorf("failed to restore file from quarantine: %v", err)
	}

	// Restore the original permissions for the file.
	if err := os.Chmod(restorePath, 0644); err != nil {
		return fmt.Errorf("failed to restore file permissions: %v", err)
	}

	fmt.Println("File was successfully resotored:")
	return nil
}

// Функция для проверки сигнатур в файле
func checkSignatures(filePath string, signatures []string) bool {
	// Открываем файл для чтения
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	// Читаем файл построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Проверяем каждую сигнатуру
		for _, signature := range signatures {
			match, _ := regexp.MatchString(signature, line)
			if match {
				return true
			}
		}
	}

	// Если файл не содержит ни одной из сигнатур, то возвращаем false
	return false
}

func Scanning() {

	//ScanFile("file.txt")

	ScanFolder("../folder")

	//go MonitorDirectory("/Users/andreysergeev/Documents/GitHub/antivirus/our_antivirus/folder")

	//go ScheduleScan("../folder", 5*time.Minute)

	//--------------
	// работает
	//Quarantine("./../folder/file.txt", "quarantine")
	// работает
	//Restore("quarantine/file.txt", "../folder/file.txt")
	//--------------

	//select {}
}

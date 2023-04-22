package pipe

import (
	"fmt"
	"os"
	"syscall"
)

func Pipe() {
	// Создаем именованный канал
	err := syscall.Mkfifo(`\\.\pipe\mychannel1`, 0666)
	if err != nil {
		fmt.Println("Ошибка создания канала:", err)
		return
	} else {
		fmt.Println("Создали канал:", err)
	}

	// Открываем канал на чтение
	file, err := os.Open(`\\.\pipe\mychannel1`)
	if err != nil {
		fmt.Println("Ошибка открытия канала:", err)
		return
	} else {
		fmt.Println("Открыли канал:", err)
	}
	defer file.Close()

	// Читаем сообщения из канала и выводим их на экран
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка чтения из канала:", err)
			return
		}
		fmt.Println("Получено сообщение:", string(buffer[:n]))
	}
}

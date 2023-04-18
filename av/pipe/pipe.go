package pipe

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func main() {
	// Указываем путь и имя именованного канала
	pipePath := `\\.\pipe\antivs`

	// Создаем именованный канал с использованием CreateNamedPipe
	hPipe, err := windows.CreateNamedPipe(
		pipePath,
		windows.PIPE_ACCESS_DUPLEX,
		windows.PIPE_TYPE_BYTE|windows.PIPE_READMODE_BYTE|windows.PIPE_WAIT,
		windows.PIPE_UNLIMITED_INSTANCES,
		512,
		512,
		0,
		nil,
	)
	if err != nil {
		fmt.Println("Failed to create named pipe:", err)
		return
	}
	defer windows.CloseHandle(hPipe)

	// Ожидаем подключения клиента
	fmt.Println("Waiting for client to connect...")
	err = windows.ConnectNamedPipe(hPipe, nil)
	if err != nil {
		fmt.Println("Failed to connect named pipe:", err)
		return
	}

	// Чтение сообщений из именованного канала
	var buffer [512]byte
	for {
		// Чтение данных из именованного канала
		n, err := windows.ReadFile(hPipe, buffer[:])
		if err != nil {
			fmt.Println("Failed to read from named pipe:", err)
			return
		}

		// Выводим прочитанные данные
		fmt.Println("Received message:", string(buffer[:n]))

		// Ожидаем следующее сообщение
	}
}

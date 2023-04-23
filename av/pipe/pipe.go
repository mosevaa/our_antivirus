package pipe

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Pipe() {
	// создаем новый сканер для чтения stdin
	scanner := bufio.NewScanner(os.Stdin)

	// читаем stdin до тех пор, пока не будет получено сообщение
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "stop" {
			break
		}
		fmt.Println("Received:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
		return
	}
}

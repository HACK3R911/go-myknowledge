package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var tasks []string

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите номер команды")

		fmt.Println("Команды: 1. create, 2. view all, 3. update, 4. delete, 0. exit")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "1":
			task, _ := reader.ReadString('\n')

			tasks = append(tasks, strings.TrimSpace(task))
			fmt.Println("Задача добавлена")

		case "2":
			if len(tasks) == 0 {
				fmt.Println("Задач нет")
				continue
			}

			fmt.Println("Список всех задач: ")
			for i, task := range tasks {
				fmt.Printf("%v\t%v\n", i+1, task)
			}

		case "3":

		case "0":
			return

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

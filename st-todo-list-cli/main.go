package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the TO DO List CLI app!")

	var tasks []string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите команду (create, read, update, delete, exit): ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "create":
			fmt.Println("Введите задачу: ")
			task, _ := reader.ReadString('\n')
			tasks = append(tasks, strings.TrimSpace(task))
			fmt.Println("Задача добавлена")

		case "read":
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст")
				continue
			}
			fmt.Println("Список задач: ")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}

		case "update":
			for {
				fmt.Println("Введите задачу для изменения: ")
				oldTask, _ := reader.ReadString('\n')
				oldTask = strings.TrimSpace(oldTask)

				idx := -1
				for i, task := range tasks {
					if task == oldTask {
						idx = i
						break
					}
				}
				if idx == -1 {
					fmt.Println("Задача не найдена! Попробуйте снова.")
					continue
				}
				for {
					fmt.Println("Введите новое название: ")
					newTask, _ := reader.ReadString('\n')
					newTask = strings.TrimSpace(newTask)

					if len(newTask) > 0 {
						tasks[idx] = newTask
						fmt.Println("Задача обновлена!")
						break
					}
					fmt.Println("Названме должно быть не короче 3 символов! ")
				}
				break
			}

		case "delete":
			for {
				fmt.Println("Введите задачу для удаления: ")
				task, _ := reader.ReadString('\n')
				task = strings.TrimSpace(task)

				idx := -1
				for i, t := range tasks {
					if t == task {
						idx = i
						break
					}
				}
				if idx == -1 {
					fmt.Println("Задача не найдена ! Попробуйте снова ")
					continue
				}

				tasks = append(tasks[:idx], tasks[idx+1:]...)
				fmt.Println("Задача обновлена")
				break
			}

		case "exit":
			fmt.Println("Выходим!")
			return

		default:
			fmt.Println("Неизвесная команда")

		}
		fmt.Println()
	}
}

package task_tracker_cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func StartCLI() {
	fmt.Println("Добро пожаловать в трекер задач!")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nВведите команду: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка: ошибка ввода")
			continue
		}

		text := scanner.Text()
		cmds_args := fieldsWithQuotes(text)

		if len(cmds_args) == 0 {
			continue
		}

		cmd := cmds_args[0]

		switch cmd {
		case "exit":
			return
		case "add":
			id, err := addTask(cmds_args[1:])
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			fmt.Printf("Задача создана (ID: %d)\n", id)
		default:
			fmt.Printf("Ошибка: неизвестная команда %v\n", cmd)
			fmt.Println("Для отображения списка команда введите help")
		}
	}
}

func fieldsWithQuotes(str string) []string {
	var (
		fields   []string
		current  strings.Builder
		inQuotes bool
	)

	for _, r := range str {
		switch r {
		case '"':
			inQuotes = !inQuotes
		case ' ', '\t':
			if inQuotes {
				current.WriteRune(r)
			} else if current.Len() > 0 {
				fields = append(fields, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		fields = append(fields, current.String())
	}

	return fields
}

func addTask(args []string) (int, error) {
	if len(args) == 0 || len(args) > 1 {
		return 0, errors.New("для команды add нужен один аргумент, заключенный в \"\"")
	}

	id, err := createTask(args[0])
	return id, err
}

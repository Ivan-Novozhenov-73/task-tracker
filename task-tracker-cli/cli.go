package task_tracker_cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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
		case "list":
			err := listTasks(cmds_args[1:])
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		case "remove":
			err := removeTask(cmds_args[1:])
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			fmt.Println("Задача успешно удалена")
		case "update":
			err := updateTask(cmds_args[1:])
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			fmt.Println("Задача успешно обнавлена")
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
		return 0, errors.New("неправильное количество аргументов для команды add")
	}

	return postTask(args[0])
}

func listTasks(args []string) error {
	if len(args) > 1 {
		return errors.New("слишком большое количество аргументов для команды list")
	}

	var tasks map[int]*Task
	var err error
	if len(args) == 0 {
		tasks, err = getTasks("")
	} else {
		switch args[0] {
		case "todo":
			tasks, err = getTasks(TASK_STATUS_TODO)
		case "in-progress":
			tasks, err = getTasks(TASK_STATUS_IN_PROGRESS)
		case "done":
			tasks, err = getTasks(TASK_STATUS_DONE)
		default:
			return errors.New("неизвестный аргумент для команды list")
		}
	}

	if err != nil {
		return err
	}

	for id, task := range tasks {
		fmt.Println(
			"ID:", id,
			"\nОписание:", task.Description,
			"\nСтатус:", task.statusToString(),
			"\nДата создания:", task.CreatedAt.Format("15:04 02-01-2006"),
			"\nДата обновления:", task.UpdatedAt.Format("15:04 02-01-2006"),
		)
		fmt.Println("--------------------------------")
	}

	return nil
}

func removeTask(args []string) error {
	if len(args) == 0 {
		return errors.New("отсутствует ID")
	} else if len(args) > 1 {
		return errors.New("слишком большое количество аргументов для команды remove")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil || id < 1 {
		return errors.New("ID должен быть целым положительным числом")
	}

	return deleteTask(id)
}

func updateTask(args []string) error {
	if len(args) == 0 {
		return errors.New("отсутствует ID и новое описание для задачи")
	} else if len(args) == 1 {
		return errors.New("отсутствует новое описание для задачи")
	} else if len(args) > 2 {
		return errors.New("слишком большое количество аргументов для команды update")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil || id < 1 {
		return errors.New("ID должен быть целым положительным числом")
	}

	return putTask(id, args[1])
}

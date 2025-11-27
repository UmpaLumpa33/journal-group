package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunMenu(journal *Journal) {
	for {
		fmt.Println("\nЖУРНАЛ ГРУППЫ")
		fmt.Println("1. Добавить студента")
		fmt.Println("2. Показать всех студентов")
		fmt.Println("3. Фильтр по среднему баллу")
		fmt.Println("4. Изменить оценки студента")
		fmt.Println("5. Статистика группы")
		fmt.Println("6. Выход")
		fmt.Print("Выбор: ")

		choice := readInt()

		switch choice {
		case 1:
			journal.AddStudent()
		case 2:
			journal.ShowAllStudents()
		case 3:
			fmt.Print("Минимальный средний балл: ")
			min := float64(readInt())
			fmt.Print("Максимальный средний балл: ")
			max := float64(readInt())
			journal.FilterByAverage(min, max)
		case 4:
			journal.ModifyStudentGrades()
		case 5:
			journal.ShowStatistics()
		case 6:
			fmt.Println("Выход из программы")
			return
		default:
			fmt.Println("Неверный выбор!")
		}
	}
}

func readInt() int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err == nil {
			return num
		}
		fmt.Print("Ошибка! Введите число: ")
	}
}

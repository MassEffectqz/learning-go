package practics

import (
	"fmt"
)

func Cycles() {
	var j int
	for j = 100; j >= 10; j-- {
		fmt.Println("1 -", j)
	}
	sum := 1 
	for ;sum < 10; {
		sum += sum 
		fmt.Println("2 -", sum)
	}
	for sum <= 100 {
		sum += 10
		fmt.Println("3 - ", sum)
	}
	// for { 
	// 	fmt.Print("Бесконечный цикл")
	// }
}
func Task1() {
	var userInput int
	sumPositive := 0
    countNegative := 0
	for {
		fmt.Println("Введите число: ")
		_,err := fmt.Scan(&userInput)
		if err != nil {
			fmt.Println("Ошибка ввода")
			continue
		}
	if userInput < 0 {
		countNegative++ 
	} else if userInput > 0 {
		sumPositive++
	}
	if userInput == 0 {
		fmt.Println("Положительных: ", sumPositive)
		fmt.Println("Отрицательных: ", countNegative)
		break 
	}

	}
}

func Task2() {
    var userBalls int
    var ones, deuces, threes, fours, fives int
    totalCount := 0
    sumBalls := 0
    for {
        fmt.Println("Введите вашу оценку (1-5), для выхода введите -1: ")
        fmt.Scan(&userBalls)
        if userBalls == -1 {
            break
        }
        if userBalls < 1 || userBalls > 5 {
            fmt.Println("Ошибка! Введите число от 1 до 5")
            continue
		}
        switch userBalls {
        case 1:
            ones++
        case 2:
            deuces++
        case 3:
            threes++
        case 4:
            fours++
        case 5:
            fives++
        }   
        totalCount++
        sumBalls += userBalls
    }
    averageBalls := float64(sumBalls) / float64(totalCount)
    successCount := threes + fours + fives
    successPercent := float64(successCount) / float64(totalCount) * 100
    maxCount := ones
    mostFrequent := 1
    
    if deuces > maxCount {
        maxCount = deuces
        mostFrequent = 2
    }
    if threes > maxCount {
        maxCount = threes
        mostFrequent = 3
    }
    if fours > maxCount {
        maxCount = fours
        mostFrequent = 4
    }
    if fives > maxCount {
        maxCount = fives
        mostFrequent = 5
    }
    fmt.Println("\nРезультаты:")
    fmt.Printf("Средний балл: %.2f\n", averageBalls)
    fmt.Println("Пятерок:", fives)
    fmt.Println("Двоек:", deuces)
    fmt.Println("Самая частая оценка:", mostFrequent)
    fmt.Printf("Процент троек и выше: %.1f%%\n", successPercent)
}

func Task3() {  
	var num int 
	fmt.Println("Введите число: ")
	fmt.Scan(&num) 
	for i := 1; i <= 10; i++ {
		result := num * i 
		fmt.Println(num, "*", i, "=", result)
	}
}
func Task4() {
    type Book struct { 
        Title string 
        Author string 
        Year int 
    }
    book1 := Book {
        Title: "1",
        Author: "Gay",
        Year: 2025,
    }
    book2 := Book {
        Title: "2",
        Author: "Gay",
        Year: 2025,
    }
    book3 := Book {
        Title: "3",
        Author: "Gay",
        Year: 2025,
    }
    fmt.Println("Книжки: ")
    fmt.Printf("Название: %s Автор: %s Год: %d \n", book1.Title, book1.Author, book1.Year)
    fmt.Printf("Название: %s Автор: %s Год: %d \n", book2.Title, book2.Author, book2.Year)
    fmt.Printf("Название: %s Автор: %s Год: %d \n", book3.Title, book3.Author, book3.Year)
}
// решение Task4 с поомщью цикла 
// func Task4() {
//     type Book struct {
//         Title  string
//         Author string // Опечатка: Autor -> Author
//         Year   int
//     }
    
//     books := []Book{
//         {Title: "Война и мир", Author: "Лев Толстой", Year: 1867},
//         {Title: "1984", Author: "Джордж Оруэлл", Year: 1949},
//         {Title: "Мастер и Маргарита", Author: "Михаил Булгаков", Year: 1966},
//     }
    
//     fmt.Println("Книги:")
//     for i, book := range books {
//         fmt.Printf("Книга %d: \"%s\", %s, %dг.\n", i+1, book.Title, book.Author, book.Year)
//     }
// }


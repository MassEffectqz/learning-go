package practics

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)
type BankAccount struct {
	Owner         	string
	Balance       	float64
	Password 		string  
	AccountNumber 	string 
}
func MainBank() {	
	display()
}
func createAccount() BankAccount {
	var bankUser BankAccount 
	bankUser.AccountNumber = generateAccountNumber()
	bankUser.Balance = 0.0
	CreateUserName:
		for {
			fmt.Println("Введите имя пользователя")
			fmt.Scan(&bankUser.Owner)
			if len(bankUser.Owner) < 6 && bankUser.Owner == strings.ToUpper(bankUser.Owner) {
				fmt.Println("Имя пользователя должно содержать больше 6 символов и иметь заглавные буквы")
				continue CreateUserName
				} else {
				fmt.Println("Аккаунт создан, имя пользователя ", bankUser.Owner) 
				break CreateUserName
		}
	}
	CreatePassword:
		for { 
			fmt.Println("Введите пароль")
			fmt.Scan(&bankUser.Password) 
			if len(bankUser.Password) < 8 { 
				fmt.Println("Пароль слишком короткий!")
				continue CreatePassword
				} else {
				fmt.Println("Пароль успешно создан! Ваш пароль", bankUser.Password)
				break CreatePassword
			}
		}
	return bankUser		
}

func generateAccountNumber() string{
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%08d", rand.Intn(100000000))
}
func (b *BankAccount) Deposit() {
	var addUserBalance int  
	fmt.Println("Введите сумму пополнения: ")
	fmt.Scan(&addUserBalance)
	if addUserBalance > 0 {
		b.Balance += float64(addUserBalance)
	} else {
		fmt.Println("Сумма должна быть больше 0!")
	}
}
func (b *BankAccount) Withdraw() {
	var withdrawPrint int 
	fmt.Println("Ввведите желаемую сумму снятия: ")
	fmt.Scan(&withdrawPrint)
	if float64(withdrawPrint) <= b.Balance && b.Balance >= 0 {
		b.Balance -= float64(withdrawPrint)
		fmt.Println("Операция прошла успешно, ваш баланс: ", b.Balance)
	} else {
		fmt.Println("Ошибка")
	}
}
func (b *BankAccount) displayBalance() { 
	fmt.Printf("Ваш баланс: %.2f руб. \n", b.Balance)
}
func (b *BankAccount) renameUser() {
	fmt.Println("Введите новое имя: ")
	fmt.Scan(&b.Owner) 
}
func display() {
	bankUser := createAccount()
	var userChoice string 
Display:
	for {
	fmt.Println("========")
	fmt.Println("Банковский счёт - Афера.com")
	fmt.Println("Имя пользователя: ", bankUser.Owner)
	fmt.Println("Пароль пользователя: ", bankUser.Password)
	fmt.Println("ID пользователя: ", bankUser.AccountNumber)
	fmt.Printf("Баланс пользователя: %.2f руб. \n", bankUser.Balance)
	fmt.Println("========")
	fmt.Printf("Выберите действие: \n 1) Пополнить баланс \n 2) Снять деньги \n 3) Отобразить баланс \n 4) Сменить имя пользователя \n 5) Выйти \n")
	fmt.Scan(&userChoice)
	switch userChoice {
	case "1", "Пополнить баланс":
		bankUser.Deposit()
	case "2", "Снять деньги":
		bankUser.Withdraw()
	case "3", "Отобразить баланс":
		bankUser.displayBalance()
	case "4", "Сменить имя пользователя":
		bankUser.renameUser()
	case "5", "Выйти":
		break Display
	default:
		fmt.Println("Неизвестная команда")
	}
}
}

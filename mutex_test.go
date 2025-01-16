package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_Mutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Counter:", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func Test_RWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(10)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance:", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (u *UserBalance) Lock() {
	u.Mutex.Lock()
}

func (u *UserBalance) Unlock() {
	u.Mutex.Unlock()
}

func (u *UserBalance) Change(amount int) {
	u.Balance = u.Balance + amount
}

func Transfer(u1 *UserBalance, u2 *UserBalance, amount int) {
	u1.Lock()
	fmt.Println("Lock user1:", u1.Name)
	u1.Change(-amount)

	time.Sleep(1 * time.Second)

	u2.Lock()
	fmt.Println("Lock user2:", u2.Name)
	u2.Change(amount)

	time.Sleep(1 * time.Second)

	u1.Unlock()
	u2.Unlock()
}

func Test_Deadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Fatih",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Syilfa",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
}

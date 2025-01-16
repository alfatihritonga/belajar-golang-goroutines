package belajargolanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_CreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Fatih"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)

	channel <- "Hello"
}

func Test_ChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)

	channel <- "Fatih"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func Test_InOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func Test_BufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Data Pertama"
		channel <- "Data Kedua"
		channel <- "Data Ketiga"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "Data Keempat"
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}

func Test_RangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- "Data ke-" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai")
}

func Test_SelectChannel(t *testing.T) {
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)

	go func() {
		channel1 <- 10000
	}()

	go func() {
		channel2 <- 1
	}()

	go func() {
		channel3 <- 500
	}()

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2: ", data)
			counter++
		case data := <-channel3:
			fmt.Println("Data dari Channel 3: ", data)
			counter++
		}

		if counter == 3 {
			break
		}
	}
}

func Test_DefaultSelectChannel(t *testing.T) {
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)

	go func() {
		channel1 <- 10000
	}()

	go func() {
		channel2 <- 1
	}()

	go func() {
		channel3 <- 500
	}()

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2: ", data)
			counter++
		case data := <-channel3:
			fmt.Println("Data dari Channel 3: ", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 3 {
			break
		}
	}
}

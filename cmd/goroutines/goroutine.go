package goroutines

import (
	"fmt"
	"time"
)

func TestGoroutine() {
	go count(10)
	time.Sleep(time.Second * 8)
}

func count(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func TestChannel() {
	c := make(chan string)
	go countChannel("test", c)
	for msg := range c {
		fmt.Println(msg)
	}
}

func countChannel(s string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- s
		time.Sleep(time.Second)
	}
	close(c)
}

func TestSelect() {
	c1 := make(chan string)
	c2 := make(chan string)
	go countChannel("test1", c1)
	go countChannel("test2", c2)
	for { // lặp vô hạn
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func TestBufferedChannel() {
	c := make(chan string, 2)
	c <- "test1"
	c <- "test2"
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func TestChannelSynchronization() {
	c := make(chan string)
	go worker(c)
	c <- "test"
	time.Sleep(time.Second * 8)
}

func worker(c chan string) {
	time.Sleep(time.Second * 2)
	msg := <-c
	fmt.Println(msg)
}

func TestChannelDirection() {
	c := make(chan string)
	go ping(c)
	go pong(c)
	time.Sleep(time.Second * 2)
}

func ping(c chan<- string) {
	c <- "ping"
}

func pong(c <-chan string) {
	msg := <-c
	fmt.Println(msg)
}

func TestChannelSelect() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "test1"
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			c2 <- "test2"
			time.Sleep(time.Second * 2)
		}
	}()
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func TestChannelTimeout() {
	c := make(chan string)
	go func() {
		for {
			c <- "test"
			time.Sleep(time.Second * 2)
		}
	}()
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		case <-time.After(time.Second):
			fmt.Println("timeout")
		}
	}
}

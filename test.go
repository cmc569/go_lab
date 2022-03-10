package main

import (
	"fmt"
	"time"
)

type person struct {
	name string
	age  int
}

func main() {
	man := person{name: "Jason", age: 45}
	man.name = "AAA"
	man.age = 12
	fmt.Println(man)
}

func mainA() {
	x := 5
	zero(&x)
	fmt.Println(x)
}

func zero(x *int) {
	*x = 0
}

func main9() {
	defer func() {
		fmt.Println("first")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("end")
	}()

	f1()
}

func f1() {
	fmt.Println("test")
	panic(false)
	fmt.Println("test2")
}

func main8() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func main7() {
	message := make(chan string)
	go func() {
		message <- "ping data"
	}()

	msg := <-message
	fmt.Println(msg)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func f2(n int) {
	arr := map[string]string{
		"head": "A",
		"body": "B",
		"tail": "C",
	}

	for _, v := range arr {
		fmt.Println(n, ":", v)
	}
}

func main6() {
	go f(0)
	go f2(1)
	time.Sleep(time.Second)
}

func main5() {
	fmt.Println(add(1, 7))
}

func add(a, b int) int {
	return a + b
}

func main4() {
	// elements := make(map[string]string)
	elements := map[string]string{
		"data":  "a",
		"title": "b",
	}

	// elements["data"] = "A"
	// elements["title"] = "B"
	// elements["head"] = "C"
	// elements["body"] = "D"

	fmt.Println(elements)

	// name, ok := elements["body1"]
	// fmt.Println(name, ok)

	if _, ok := elements["body"]; ok {
		// fmt.Println(name, ok)
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}

func main3() {
	_map := make(map[string]int)

	_map["data"] = 1
	_map["title"] = 2
	_map["head"] = 3
	fmt.Println(_map)
	fmt.Println(len(_map))

	for k, _ := range _map {
		fmt.Println(k)
	}
	delete(_map, "data")
	fmt.Println(_map)
	fmt.Println(len(_map))
}

func main2() {
	arr := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(arr)
	fmt.Println(len(arr))
	// x := arr[1:4]
	// fmt.Println(x)
	var x []int
	x = arr[1:2]
	fmt.Println(x)

	x = append(x, 6)
	fmt.Println(x)
}

func main1() {
	for i := 0; i < 10; i++ {
		if i > 5 {
			fmt.Println("Hi")
		}
		fmt.Println(i)
	}
}

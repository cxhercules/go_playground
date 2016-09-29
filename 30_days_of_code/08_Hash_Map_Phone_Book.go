package main

import (
	"errors"
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	myMap := make(map[string]string)
	var input, name, number string
	var readSize int
	var err error

	//fmt.Println(myMap)
	for i := 0; i < T; i++ {
		fmt.Scan(&input)
		name = input
		fmt.Scan(&input)
		number = input
		myMap[name] = number
		//fmt.Println(name, number)
	}

	//fmt.Println(myMap)

	for {

		readSize, err = fmt.Scanln(&input)
		if err != nil {
			return
		} else if readSize < 1 {
			err = errors.New("Can't read input")
			return
		}

		mapvalue, exist := myMap[input]

		//fmt.Println("Printing len  ", len(input))
		//fmt.Println("mapvalue:", mapvalue)
		//fmt.Println("exist:", exist)

		if exist == false {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%s=%s\n", input, mapvalue)
		}

	}

}

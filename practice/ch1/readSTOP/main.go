package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var f *os.File
	f=os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	stopVar := "STOP"
	for scanner.Scan() {
		in := scanner.Text()
		if in == stopVar {
			os.Exit(0)
		}
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter number !!")
			//return
		} else {
			fmt.Println("Input number = ",n)
		}
	}
}
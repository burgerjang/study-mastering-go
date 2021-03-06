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

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err == nil {
			sum += n
			fmt.Println("Thu sum is = ",sum)
		}
	}
}
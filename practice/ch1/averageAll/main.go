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

	avg, sum := float64(0), float64(0)
	cnt := float64(1)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.ParseFloat(scanner.Text(),64)
		if err == nil {
			sum += n
			avg = sum/cnt
			fmt.Println("Thu average is = ",avg)
			cnt+=1
		}
	}
}
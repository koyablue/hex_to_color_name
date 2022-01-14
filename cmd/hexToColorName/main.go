package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//TODO: get input value
	//TODO: validation
	//TODO: http GET request to API endpoint with query params (to get color name)
	//TODO: return(print)color name

	fmt.Println("ENTER HEX CODE:")

	//get stdIn value
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	hex := scanner.Text()
	fmt.Println(hex)
}

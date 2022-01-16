package main

import (
	"bufio"
	"fmt"
	"hex_to_color_name/controllers"
	"os"
)

func main() {
	fmt.Println("ENTER HEX CODE:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	hexColorCode := scanner.Text()

	fmt.Println(controllers.GetColorName(hexColorCode))
}

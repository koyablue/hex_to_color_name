package main

import (
	"bufio"
	"fmt"
	"hex_to_color_name/controllers"
	"hex_to_color_name/services"
	"os"
)

func main() {
	fmt.Println(`ENTER HEX CODE WITHOUT "#":`)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	hexColorCode := scanner.Text()

	apiResponse := controllers.GetColorName(hexColorCode)

	fmt.Println(services.GetColorName(&apiResponse))
}

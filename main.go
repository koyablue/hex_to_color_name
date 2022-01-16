package main

import (
	"bufio"
	"fmt"
	"hex_to_color_name/controllers"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	BaseUrl string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		BaseUrl: cfg.Section("api").Key("base_url").String(),
	}
}

func main() {
	fmt.Println(`ENTER HEX CODE WITHOUT "#":`)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	hexColorCode := scanner.Text()

	fmt.Println(controllers.GetColorName(hexColorCode))
}

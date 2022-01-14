package colorName

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Rgb struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type Color struct {
	Name         string  `json:"name"`
	Hex          string  `json:"hex"`
	Rgb          Rgb     `json:"rgb"`
	Distance     float64 `json:"distance"`
	Luminance    float64 `json:"luminance"`
	RequestedHex string  `json:"requestedHex"`
}

// struct for API return value
type Colors struct {
	Colors []Color `json:"colors"`
}

//https://api.color.pizza/v1/?values=

//TODO: move this variable to const file
var API_BASE_URL string = "https://api.color.pizza/v1/"

//TODO: arg colorCode string
func GetColorName() {
	colorCode := "FF5733"
	baseUrl, err := url.Parse(API_BASE_URL)
	if err != nil {
		log.Fatalln(err)
	}

	reference, err := url.Parse("?values=" + colorCode)
	if err != nil {
		log.Fatalln(err)
	}

	endpoint := baseUrl.ResolveReference(reference).String()

	response, err := http.Get(endpoint)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	var colors Colors
	if err := json.Unmarshal(responseBody, &colors); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(colors.Colors[0].Name)

	//TODO: convert return value to struct then return color name(str)

}

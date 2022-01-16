package controllers

import (
	"encoding/json"
	"hex_to_color_name/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var API_BASE_URL string = "https://api.color.pizza/v1/"

func GetColorName(hexColorCode string) models.ColorNameApiResponse {
	baseUrl, err := url.Parse(API_BASE_URL)
	if err != nil {
		log.Fatalln(err)
	}

	reference, err := url.Parse("?values=" + hexColorCode)
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

	var apiResponse models.ColorNameApiResponse
	if err := json.Unmarshal(responseBody, &apiResponse); err != nil {
		log.Fatalln(err)
	}

	return apiResponse
}

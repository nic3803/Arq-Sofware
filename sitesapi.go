package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//type Sites []Sites

type Sites []struct {
	default_currency_id string `json:"default_currency_id"`
	Name                string `json:"name"`
	id                  string `json:"id"`
}

func main() {

	cats, err := GetCategories("MLA")

	if err != nil {

		fmt.Println("Error: ", err.Error())

		return

	}

	fmt.Println("Los paises son: ")

	for i := 0; i < len(cats); i++ {

		fmt.Println(cats[i].Name)
	}

}

func GetCategories(siteID string) (Sites, error) {

	response, err1 := http.Get("https://api.mercadolibre.com/sites#json")

	if err1 != nil {

		fmt.Println("Error", err1.Error())

		return nil, nil

	}

	bytes, err2 := ioutil.ReadAll(response.Body)

	if err2 != nil {

		fmt.Println("Error: ", err2.Error())
		//exit(1)
		return nil, nil

	}

	var cats Sites

	err := json.Unmarshal(bytes, &cats)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		//exit(1)
		return nil, nil

	}

	return cats, nil

}

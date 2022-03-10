package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	//get
	url := "https://accu-dev.azurewebsites.net/web/get.php?param1=AAA&param2=BBB"
	response := get(url)
	fmt.Println(response)

	//post
	data := map[string]string{
		"name":   "Jason",
		"gender": "male",
		"age":    "45",
	}

	str := postDataConvert(data)

	// url = "https://accu-dev.azurewebsites.net/web/post.php"
	// response = post(url, str)
	// fmt.Println(response)

	//request
	url = "https://accu-dev.azurewebsites.net/web/post.php"
	response = request(url, "POST", str)
	fmt.Println(response)
}

func get(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func postDataConvert(data map[string]string) string {
	var str string
	for k, v := range data {
		if len(str) > 0 {
			str = str + "&"
		}

		str = str + k + "=" + v
	}

	return str
}

func post(url string, values string) string {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(values))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func request(url string, method string, values string) string {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(values))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "name=test")
	req.Header.Set("Authorization", "Bearer AAAAABBBBBB")
	req.Header.Set("x-api-key", "AAAAABBBBBB")

	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	
	defer resp.Body.Close()

	return string(body)
}

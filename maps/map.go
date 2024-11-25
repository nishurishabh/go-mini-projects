package main

import "fmt"

func main() {
	websites := map[string]string{
		"Zepto":               "https://zepto.com",
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["Airbnb"] = "https://airbnb.com"
	fmt.Println(websites)

	delete(websites, "Airbnb")
	fmt.Println(websites)
}

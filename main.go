package main

import (
	"fmt"
	"github.com/dominicphillips/amazing"
	"net/url"
)

func main() {

	client, err := amazing.NewAmazing("IN", "", "", "")
	if err != nil {
		fmt.Println(err)
	}

	params := url.Values{
		"IdType":    []string{"ASIN"},
		"ItemId":    []string{"B00T9LN3XC"},
		"Operation": []string{"ItemLookup"},
	}

	result, err := client.ItemLookup(params)

	fmt.Println("Hello world", result.AmazonItems.Items[0], err)
}

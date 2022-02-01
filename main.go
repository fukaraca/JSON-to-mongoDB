package main

import (
	"context"
	"fmt"
	"get-json-to-mongoDB/lib"
)

func main() {

	//lib.CreateDocument()

	count := 0
	heroes := lib.ExcelToSlice()
	for _, hero := range heroes {
		for _, tag := range lib.RetrieveDocument(hero).Tags {
			if tag == "Mage" {
				fmt.Printf("%s is a Mage\n", hero)
				count++
			}
		}

	}
	fmt.Printf("There is total %d number of Mage", count)

	defer func() {
		if err := lib.Client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
}

func init() {
	lib.Connect()
}

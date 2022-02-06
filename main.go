package main

import (
	"context"
	"fmt"
	"get-json-to-mongoDB/lib"
)

func main() {

	heroes := lib.ExcelToSlice()
	for _, hero := range heroes {
		heroObj := lib.GetHeroInfo(hero)
		if lib.CreateDocument(heroObj) {
			fmt.Println("hero informations fetched and inserted succesfully")
		} else {
			break
		}

	}

	defer func() {
		if err := lib.Client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
}

func init() {
	lib.Connect()
}

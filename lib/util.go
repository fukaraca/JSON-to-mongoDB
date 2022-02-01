package lib

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

var lolHeroes []string

//get hero names from the excel file
func ExcelToSlice() []string {
	f, err := excelize.OpenFile("lolheroes.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}

	}()
	n := 158 //total count of lol heroes
	for i := 1; i <= n; i++ {
		heroName, err := f.GetCellValue("Sayfa1", fmt.Sprintf("E%d", i))
		if err != nil {
			fmt.Println("read from xlsx failed:", err)
		}
		heroName = strings.TrimSpace(heroName)
		lolHeroes = append(lolHeroes, heroName)

	}
	return lolHeroes
}

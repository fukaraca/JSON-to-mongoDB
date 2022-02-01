package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetHeroInfo(heroName string) *HeroInfoStruct {

	getUrl := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/12.2.1/data/en_US/champion/%s.json", url.PathEscape(heroName))
	resp, err := http.Get(getUrl)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	read, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	heroGoverted := LolHeroInfo{}
	err = json.Unmarshal(read, &heroGoverted)
	if err != nil {
		fmt.Println("unmarshall failed:", err)
		return nil
	}
	tempHeroInfo := HeroInfoStruct{}
	for _, v := range heroGoverted.Data {
		tempHeroInfo = v
	}
	return &tempHeroInfo

}

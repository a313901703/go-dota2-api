package spider

import (
	"fmt"
	"net/http"
)

const (
	itemURL       = "/fundata-dota2-free/v2/raw/item"
	leadguesURL   = "/fundata-dota2-free/v2/league/list"
	itemsURL      = "http://dotamax.com/item/"
	keyItemsInfo  = "redis_items_list"
	keyHeroesInfo = "redis_heroes_info"

	maxBaseURL   = "http://dotamax.com/"
	leagueImgURL = "http://cdn.dota2.com/apps/dota2/images/leagues/%s/images/image_8.png"
	keyLeagues   = "redis_leagues_list"

	newsURL = "https://www.dota2.com.cn/news/gamenews/index%s.htm"
	keyNews = "redis_news_list"

	abilityURL      = "/constants/abilities"
	abilitiesIdsURL = "/constants/ability_ids"
	keyAbility      = "redis_ability"
)

func init() {
	fmt.Println("spider GetHeroInfo")
	GetHeroInfo()
	fmt.Println("spider GetItems")
	GetItems()
	fmt.Println("spider GetLeagues")
	GetLeagues()
	fmt.Println("spider GetNews")
	GetNews()
	fmt.Println("spider GetAbility")
	GetAbility()
}

func fetchImg(url string) bool {
	// Request the HTML page.
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}

	res, err := client.Do(req)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return false
	}
	return true
}

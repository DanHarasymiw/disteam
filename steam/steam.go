package steam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VanityUrlJson struct {
	Response VanityUrlResponseJson `json:"response"`
}

type VanityUrlResponseJson struct {
	Success int    `json:"success"`
	SteamID string `json:"steamid"`
}

func ResolveVanityUrl(username, key string) (VanityUrlResponseJson, error) {
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=%s&vanityurl=%s",
		key, username)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	body, _ := ioutil.ReadAll(resp.Body)
	vanityUrl := new(VanityUrlJson)
	json.Unmarshal(body, vanityUrl)

	return vanityUrl.Response, nil
}

type GetOwnedGamesJson struct {
	Response GetOwnedGamesResponseJson `json:"response"`
}

type GetOwnedGamesResponseJson struct {
	Count int        `json:"game_count"`
	Games []GameJson `json:"games"`
}

type GameJson struct {
	Name string `json:"name"`
}

func GetOwnedGames(steamId, key string) (GetOwnedGamesResponseJson, error) {
	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&include_appinfo=1", key, steamId)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err2)
	}
	getOwnedGames := new(GetOwnedGamesJson)
	json.Unmarshal(body, getOwnedGames)

	return getOwnedGames.Response, nil
}

package main

import (
	"fmt"
	"net/url"
	"github.com/Philipp15b/go-steamapi"
)



func main() {
	key := "API_KEY_HERE"
	var names = []string{"test"}

	findCommonSteamGames(names, key);
}

func findCommonSteamGames(usernames []string, key string) {
	count := 0
	var games [1][]string;
	for _, username := range usernames {
		steamId, err := usernameToSteamId(username, key)

		if err != nil {
			count = count + 1
			games[count] = getOwnedGames(steamId, key)
			print (games[count])
		}

	}


}

func getOwnedGames(steamId, key string) ([]string, error){
	var getGames = steamapi.NewSteamMethod("IPlayerService", "GetOwnedGames", 1)

	data := url.Values{}
	data.Add("key", key)
	data.Add("steamid", steamId)
	data.Add("include_appInfo", "1")

	var resp GetOwnedGamesResponse
	err := getGames.Request(data, &resp)
	if err != nil {
		return nil, err
	}

	var gameNames [resp.Game_count]string
	for index, name := range resp.Games {
		gameNames[index] = name
	}

	return gameNames, nil
}

func usernameToSteamId(username, key string) (steamapi.ResolveVanityURLResponse, error) {
	response, err := steamapi.ResolveVanityURL(username, key)

	if err != nil || response.Success != 1 {
		return nil, err
	}

	return response.SteamID, nil
}

type GetOwnedGamesResponse struct {
	Game_count int
	Games []Game
}

type Game struct {
	Success int
	AppId int
	Name string
	Playtime_forever int
	Img_icon_url string
	Img_logo_url string
	Has_community_visible_stats bool
}
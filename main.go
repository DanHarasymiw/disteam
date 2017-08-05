package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/danharasymiw/disteam/steam"
)

func main() {
	key := os.Getenv("STEAM_KEY")

	if key == "" {
		fmt.Println("STEAM_KEY ENV VAR NOT SET")
	}
	fmt.Println(getCommonGames(os.Args[1:], key))
}

func getCommonGames(usernames []string, key string) string {

	steamIds := make([]string, 0, len(usernames))
	unfoundUsers := make(map[string]bool)

	// try and find people
	for _, username := range usernames {
		steamId, err := usernameToSteamId(username, key)

		if err == nil && steamId != "" {
			steamIds = append(steamIds, steamId)
		} else {
			unfoundUsers[username] = true
		}

		for unfoundUser, _ := range unfoundUsers {
			for _, steamId := range steamIds {
				foundSteamId := checkFriendsForUser(steamId, unfoundUser)

				if foundSteamId != "" {
					steamIds = append(steamIds, foundSteamId)
					delete(unfoundUsers, unfoundUser)
				}
			}
		}

		if len(unfoundUsers) > 0 {
			output := "Unable to find: "
			for user, _ := range unfoundUsers {
				output += user + ", "
			}
			fmt.Println(output[:len(output)-2])
		}
	}

	allGames := getAllUsersGames(steamIds, key)
	commonGames := findIntersection(allGames)
	gamesString := strings.Join(commonGames, ", ")

	return fmt.Sprintf("There are %d games in common: %s", len(commonGames), gamesString)
}

func checkFriendsForUser(steamId string, unfoundName string) string {
	// get friends
	// check each friends persona name
	// if we find a match return the steam id, else 0
	return ""
}

func findIntersection(allGames [][]string) []string {
	var gamesMap = make(map[string]int)

	for _, games := range allGames {
		for _, game := range games {
			if val, ok := gamesMap[game]; ok {
				gamesMap[game] = val + 1
			} else {
				gamesMap[game] = 1
			}
		}
	}

	count := 0
	var commonGames = make([]string, len(gamesMap))
	for gameName, numOfOwned := range gamesMap {
		if numOfOwned == len(allGames) {
			commonGames[count] = gameName
			count++
		}
	}

	commonGames = commonGames[:count]

	return commonGames
}

func getAllUsersGames(steamIds []string, key string) [][]string {

	count := 0
	allGames := make([][]string, len(steamIds))

	for _, steamId := range steamIds {
		userGames, err := getOwnedGameNames(steamId, key)

		if err != nil {
			fmt.Printf("Couldn't find games for steam id: %s", steamId)
			return nil
		}

		allGames[count] = userGames
		count = count + 1
	}

	return allGames[:count]
}

func getOwnedGameNames(steamId string, key string) ([]string, error) {
	games, _ := steam.GetOwnedGames(steamId, key)

	var names = make([]string, games.Count)
	for index, game := range games.Games {
		names[index] = game.Name
	}

	return names, nil
}

func usernameToSteamId(username, key string) (string, error) {
	response, err := steam.ResolveVanityUrl(username, key)

	if err != nil || response.Success != 1 {
		return "", err
	}

	return response.SteamID, err
}

package groupietrackers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type SpotifyArtistsInformationsInfo struct {
	Error struct {
		Status int
	}
	Artists struct {
		Items []struct {
			Followers struct {
				Total int
			}
			Genres []string
			Id     string
		}
	}
}

var SpotifyInfo SpotifyArtistsInformationsInfo

func Token(SpotifyToken *string) {
	/*
	* This function create a id for
	 */
	clientID := "aa2a6c4a4c2f4b4aa318a8d8c8bed839"
	clientSecret := "d925b1ccbf974c99914751d3b1d15dc1"
	var tokenResponse map[string]interface{}
	authURL := "https://accounts.spotify.com/api/token"
	data := strings.NewReader("grant_type=client_credentials")

	req, err := http.NewRequest("POST", authURL, data)
	if err != nil {
		fmt.Println("Request Creation failed : Spotify.go  | line : 23")
		os.Exit(0)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed : Spotify.go  | line : 32")
		os.Exit(0)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil failed : Spotify.go  | line : 39")
		os.Exit(0)
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		fmt.Println("Unmarshal failed : Spotify.go  | line : 45")
		os.Exit(0)
	}

	*SpotifyToken = tokenResponse["access_token"].(string)
}

func RemoveSpace(entry string) string {
	ByteEntry := []byte(entry)
	for i, value := range ByteEntry {
		if value == 32 {
			ByteEntry[i] = 45
		}
	}
	return string(ByteEntry)
}
func GetArtist(Artists string, SpotifyToken *string) {
	Artists = RemoveSpace(Artists)
	url := "https://api.spotify.com/v1/search?query=artist%3A" + Artists + "&type=artist&locale=fr%2Cfr-FR%3Bq%3D0.8%2Cen-US%3Bq%3D0.5%2Cen%3Bq%3D0.3&offset=0&limit=1&access_token=" + *SpotifyToken
	APICall(url , &SpotifyInfo)
	if SpotifyInfo.Error.Status == 401 {
		SpotifyInfo.Error.Status = 0
		Token(SpotifyToken)
		url = "https://api.spotify.com/v1/search?query=artist%3A" + Artists + "&type=artist&locale=fr%2Cfr-FR%3Bq%3D0.8%2Cen-US%3Bq%3D0.5%2Cen%3Bq%3D0.3&offset=0&limit=1&access_token=" + *SpotifyToken
		GetArtist(Artists, SpotifyToken)
	}
}

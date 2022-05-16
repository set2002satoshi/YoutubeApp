package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Repstypes struct {
	SearchURL string   `json:"kind"`
	Item      []Items  `json:"items"`
	Branding  Channels `json:"brandingSettings"`
}
type Items struct {
	Snippet Channels `json:"snippet"`
}
type Channels struct {
	Id   string   `json:"channelId"`
	Img  []string `json:"default"`
	Name string   `json:"title"`
}
type Branding struct {
	Channels   []string   `json:"channel"`
}


func SearchChannels(Keyword string) (resultChannels interface{}, err error) {
	youtubeKey := YoutubeApiKey()
	fmt.Println(youtubeKey)
	ChannelList, err := GetChannels(youtubeKey, Keyword)
	if err != nil {
		return "", err
	}
	resultChannels = ChannelList
	return resultChannels, nil
}

func YoutubeApiKey() string {
	//.envファイルの読み込み
	err := godotenv.Load("./youtube.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	APIKEY := os.Getenv("YoutubeKey")
	fmt.Println(APIKEY)
	return APIKEY
}

func GetChannels(KEY, Keyword string) (interface{}, error) {
	url := "https://www.googleapis.com/youtube/v3/search"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(request)

	//クエリパラメータ
	params := request.URL.Query()
	params.Add("key", KEY)
	params.Add("part", "snippet")
	params.Add("type", "channel")
	params.Add("q", Keyword)
	// params.Add("q", "ヒカキン")
	params.Add("maxResults", "5")

	request.URL.RawQuery = params.Encode()

	// fmt.Println(request.URL.String()) //https://jsonplaceholder.typicode.com/todos?userId=1

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer response.Body.Close()

	channels, err := ioutil.ReadAll(response.Body)
	// resp := response.Body

	if err != nil {
		log.Fatal(err)
		return "", err
	}
	// fmt.Println(channels.json())

	var te Repstypes

	err = json.Unmarshal(channels, &te)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	fmt.Println(te)

	return te, nil
}

package service   

import (
	"os"
	"fmt"
	"log"
	"time"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/joho/godotenv"
)

type TestID struct {
	ChannelId string `json:"channelId"`
}



func SearchChannels(Keyword string) (resultChannels string, err error) {
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
	if(err != nil){
		log.Fatal("Error loading .env file")
	}
	APIKEY := os.Getenv("YoutubeKey")
	fmt.Println(APIKEY)
	return APIKEY
}

func GetChannels(KEY, Keyword string) (resultChannels string, err error) {
	url := "https://www.googleapis.com/youtube/v3/search"
 
	request, err := http.NewRequest("GET", url, nil)
	if err != nil{
		log.Fatal(err)
		return "", err
	}
	
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
	if err != nil{
		log.Fatal(err)
		return "", err
	}
	
	defer response.Body.Close()
 
	channels, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	fmt.Println(string(channels))

	var te []TestID

	err = json.Unmarshal(channels, &te)

	fmt.Println(te)

 
	return resultChannels, nil
}
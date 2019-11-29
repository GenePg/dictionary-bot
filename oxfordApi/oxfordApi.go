package oxfordApi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var baseUrl = "https://od-api.oxforddictionaries.com/api/v2"

var app_id = os.Getenv("APP_ID")
var app_key = os.Getenv("APP_KEY")
var endpoint = "entries"
var language_code = "en-us"

func init() {
	log.Println("oxfordApi")
}

func fetchDataByUrl(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	log.Println("err1:", err)
	req.Header.Add("app_id", app_id)
	req.Header.Add("app_key", app_key)

	resp, err := client.Do(req)

	log.Println("err2:", err)
	log.Printf("resp: %#v", resp)
	return resp, err
}

func GetDefinition(text string) string {
	word_id := text
	url := fmt.Sprintf("%s/%s/%s/%s", baseUrl, endpoint, language_code, word_id)

	resp, err := fetchDataByUrl(url)

	result := EntriesDef{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result: %+v", result)

	word := result.Word
	log.Printf("word: %v", word)

	definition := result.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0]
	log.Printf("definition: %+v", definition)

	msg := fmt.Sprintf("word: %s\ndefinition: %s", word, definition)

	return msg
}

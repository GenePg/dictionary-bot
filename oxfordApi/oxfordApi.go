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

func GetDefinition(text string) string {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	app_id = os.Getenv("APP_ID")
	app_key = os.Getenv("APP_KEY")

	word_id = text

	url := fmt.Sprintf("%s/%s/%s/%s", baseUrl, endpoint, language_code, word_id)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	log.Println("err1:", err)

	req.Header.Add("app_id", app_id)
	req.Header.Add("app_key", app_key)
	resp, err := client.Do(req)

	log.Println("err2:", err)
	log.Printf("resp: %#v", resp)

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)

	// log.Println("body:", string(body))

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

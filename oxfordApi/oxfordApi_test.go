package oxfordApi

import (
	"net/http"
	"testing"
)

func TestFetchDataByUrl(t *testing.T) {
	url := "https://od-api.oxforddictionaries.com/api/v2/entries/en-us/adorable"
	resp, err := fetchDataByUrl(url)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("resp status is not OK!")
	}
	if err != nil {
		t.Error(nil)
	}
}

package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Phonetic2 struct {
	Text  string `json:"text"`
	Audio string `json:"audio"`
}



type Meaning2 struct{
	Definition string `json:"definition"`
	Example string `json:"example"`
}

type Meaning struct {
	PartofSpeech string `json:"partOfSpeech"`
	Definitions []Meaning2 `json:"definitions"`

}

type Dictionary struct {
	Word      string     `json:"word"`
	Phonetic  string     `json:"phonetic"`
	Phonetics []Phonetic2 `json:"phonetics"`
	Meaning string `json:"meaning,omitempty"`
	Meanings []Meaning `json:"meanings"`
}

func GetWord(lang, word string) []Dictionary {
	// akses url
	resp, err := http.Get(fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/%s/%s", lang, word))
	if err != nil {
		panic(err)
	}

	// ambil data dari response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// masukkan data ke struct
	res := make([]Dictionary, 0)
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	return res
}

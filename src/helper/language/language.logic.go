package main

import (
	"encoding/json"
	"fmt"
	"os"

	LanguageInterface "expense-tracker-api/src/helper/interface/language"
)

func LoadConfiguration (filename string) (LanguageInterface.Word, error) {

	var word LanguageInterface.Word
	readWord, err := os.Open(filename)
	defer readWord.Close()
	if err != nil {
	   return word, err
	}
	jsonParser := json.NewDecoder(readWord)
	err = jsonParser.Decode(&word)
	return word, err
}

func main() {
	// default language code
	languageCode := "en"

	fmt.Println("Starting to read json file...")
	word, _ := LoadConfiguration("dictionary/"+languageCode+".json")
	fmt.Println(word)
}

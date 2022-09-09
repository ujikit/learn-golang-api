package language

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"context"

	// Interface
	LanguageInterface "expense-tracker-api/src/interface/helper/language"
)

func LoadConfiguration (filename string) (LanguageInterface.Word, error) {

	var word LanguageInterface.Word
	readWord, err := os.Open(filename)
	if err != nil {
		return word, err
	}
	defer readWord.Close()
	
	jsonParser := json.NewDecoder(readWord)
	err = jsonParser.Decode(&word)
	if err != nil {
		fmt.Println("error", err)
		return word, err
	}
	return word, err
}

func Translate(ctx context.Context, word string) (string) {
	fmt.Println("Starting to read json file...")

	languageCode := ctx.Value("languageCode").(string)
	var allWords LanguageInterface.Word
	allWords, err := LoadConfiguration("src/helper/language/dictionary/"+languageCode+".json")
	
	if err != nil {
		fmt.Println("error", err)
		return err.Error()
	}

	var reflectValue = reflect.ValueOf(allWords)

    if reflectValue.Kind() == reflect.Ptr {
        reflectValue = reflectValue.Elem()
    }

    for i := 0; i < reflectValue.NumField(); i++ {
		varName := reflectValue.Type().Field(i).Name
		if varName ==  word{
			return reflectValue.Field(i).String()
		}
    }

	return word
}

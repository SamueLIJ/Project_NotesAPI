package main

import (
	"NotesAPI/config"
	"NotesAPI/middleware"
	"NotesAPI/route"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	config.InitLog()
	config.InitMigration()

	// fmt.Println("Calling API...")

	// result := lib.GetWord("en", "hello")

	// fmt.Println(result)
	// fmt.Println(hello.Meanings[1])
	// fmt.Println(hello.Phonetics[1])

	// contoh kode 2
	// client := &http.Client{}
	// req, err := http.NewRequest("GET", "https://api.dictionaryapi.dev/api/v2/entries/en/hello", nil)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }
	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("Content-Type", "application/json")
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }
	// defer resp.Body.Close()
	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	// var responseObject dictionary
	// json.Unmarshal(bodyBytes, &responseObject)
	// fmt.Printf("API Response as struct %+v\n", responseObject)

	e := echo.New()
	e.Use(middleware.Log)
	route.NewUser(e)
	route.NewNotes(e)
	route.NewLabel(e)
	route.NewReminders(e)
	route.NewPictures(e)
	route.Dict(e)
	e.Start(":8080")
}

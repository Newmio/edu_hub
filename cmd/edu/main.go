package main

import (
	"ed"
	"ed/internal/app"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// content, err := ioutil.ReadFile("test.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// file, err := os.Create("t.txt")

	// data := map[string]interface{}{
	// 	"id_account": 2,
	// 	"id_chank":   3,
	// 	"data":       string(content),
	// 	"file_type":  "png",
	// 	"last":       true,
	// }

	// body, err := json.Marshal(data)
	// if err != nil{
	// 	log.Fatal(err)
	// }

	// file.Write(body)

	// fmt.Println("_______________________________________________________________________________")
	// fmt.Println(content)
	// fmt.Println("_______________________________________________________________________________")
	// fmt.Println(string(content))
	// fmt.Println("_______________________________________________________________________________")

	r, err := app.InitEngine()
	if err != nil {
		log.Panic(err)
	}

	srv := new(ed.Server)
	if err := srv.Run("4041", r); err != nil {
		log.Panic(err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	types "github.com/VitorEstevam/igb_editor/types"
)

type TodoPageData struct {
	Participantes []types.Participante
}

func main() {
	jsonFile, err := os.Open("elenco.igelenco")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var elenco types.Elenco

	json.Unmarshal([]byte(byteValue), &elenco)

	var participantes = elenco.Participantes

	jsonFile.Close()

	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			Participantes: participantes,
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}

package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
	"encoding/json"
	"io/ioutil"

	"github.com/brownlow2/gohttp/pkg/character"
)

var html_resp = template.Must(template.New("html_resp").Parse(`
<h1>{{.Name}}</h1>
<table>
  <tr>
    <th>Race</th>
    <th>Class</th>
    <th>Level</th>
    <th>Total HP</th>
    <th>Current HP</th>
    <th>Base AC</th>
    <th>Initiative</th>
    <th>Speed</th>
    <th>Hit Dice</th>
    <th>Inspiration</th>
    <th>Profiency Bonus</th>
  </tr>
  <tr>
    <td>{{.Race}}</td>
    <td>{{.Class}}</td>
    <td>{{.Level}}</td>
    <td>{{.TotalHP}}</td>
    <td>{{.CurrentHP}}</td>
    <td>{{.BaseAC}}</td>
    <td>{{.Initiative}}</td>
    <td>{{.Speed}}</td>
    <td>{{.HitDice}}</td>
    <td>{{.Inspiration}}</td>
    <td>{{.ProficiencyBonus}}</td>
  </tr>
</table>
<h2>SKILLS</h2>
{{range .Skills.Skills}}
<h3>{{.SkillName}} ({{.Value}})</h3>
<table>
  <tr>
    {{range .Skills}}
    <th>{{.}}</th>
    {{end}}
  </tr>
  <tr>
    {{range $key, $value := .SkillValues}}
    <td>{{$value}}</td>
    {{end}}
  </tr>
</table>
{{end}}
`))

func main() {
	http.HandleFunc("/", HelloHandler)

	fmt.Println("Listening on port 4567...")
	err := http.ListenAndServe(":4567", nil)
	if err != nil {
		panic(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//char := Character{"Uldar the Cursed", "Elf", "Blood Hunter", 5, stats.Stats{46, 38, 17, 1, 30, 8, false, 3}}
	char := build_char()
	if err := html_resp.Execute(w, char); err != nil {
		log.Fatal(err)
	}
	//fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func build_char() *character.Character {
	file, _ := ioutil.ReadFile("chars/uldar/character.json")
	char := &character.Character{}
	err := json.Unmarshal([]byte(file), char)
	if err != nil {
		fmt.Println(err)
	}
	return char
}

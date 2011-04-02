package main

import (
	"http"
	"fmt"
	"strings"
	"έντομο"
	"github.com/droundy/gui"
	"github.com/droundy/goopt"
)

var port = goopt.Int([]string{"--port"}, 8080, "port on which to serve")

var bug = έντομο.Type("bug")
var todo = έντομο.Type("todo")

func main() {
	http.HandleFunc("/style.css", styleServer)

	bugs := []gui.Widget{gui.Text("This will be a bug browser, some day!")}
	bl, err := bug.List()
	if err != nil {
		panic("bug.List: " + err.String())
	}
	bugtable := [][]gui.Widget{{
		gui.Text("id"), gui.Text("status"), gui.Text("date"), gui.Text("bug"),
	}}
	for bnum, b := range bl {
		b.Comments() // to get attributes
		bugname := fmt.Sprint(bug, "-", bnum)
		// bugs = append(bugs, gui.Text(""), gui.Text(bugname))
		cs, err := b.Comments()
		if err != nil {
		 	continue
		}
		// for _, c := range cs {
		// 	bugs = append(bugs, gui.Text(c.Author), gui.Text(c.Date), gui.Text(c.Text))
		// }
		lines := strings.Split(cs[0].Text, "\n", 2)
		status,_ := b.Attributes["status"]
		bugtable = append(bugtable, []gui.Widget{
			gui.Button(bugname), gui.Text(status), gui.Text(cs[0].Date), gui.Text(lines[0])})
	}
	bugs = append(bugs, gui.Empty(), gui.Empty(), gui.Table(bugtable...))
	err = gui.Run(*port,
		gui.Column(bugs...))
	if err != nil {
		panic("ListenAndServe: " + err.String())
	}
}

package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/teris-io/cli"
	"net/http"
	"os"
	"strings"
	"time"
)

const access_token string = "YOUR_OP_USER_ACCESS_TOKEN"

type Payload struct {
	Comment Comment `json:"comment"`
	SpentOn string  `json:"spentOn"`
	Hours   string  `json:"hours"`
	Links   Links   `json:"_links"`
}
type Comment struct {
	Format string      `json:"format"`
	Raw    interface{} `json:"raw"`
	HTML   string      `json:"html"`
}
type WorkPackage struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}
type Activity struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}
type Self struct {
	Href interface{} `json:"href"`
}
type Links struct {
	WorkPackage WorkPackage `json:"workPackage"`
	Activity    Activity    `json:"activity"`
	Self        Self        `json:"self"`
}

func main() {
	time := cli.NewCommand("time", "idő rögzítése").
		WithShortcut("t").
		WithArg(cli.NewArg("work_package", "wp id-ja (2408)")).
		WithArg(cli.NewArg("hours", "logolt órák (1.5)")).
		WithOption(cli.NewOption("date", "iso dátum (2020-01-01)").WithChar('d').WithType(cli.TypeString)).
		WithOption(cli.NewOption("ma", "mai napra").WithChar('m').WithType(cli.TypeBool)).
		WithOption(cli.NewOption("tegnap", "tegnapi napra").WithChar('t').WithType(cli.TypeBool)).
		WithAction(func(args []string, options map[string]string) int {
			var wp = args[0]
			var wtime = strings.Split(args[1], ".")
			var date = options["date"]
			var h = wtime[0]
			var time_to_log = "PT" + h + "H"
			if len(wtime) == 2 {
				time_to_log += "30M"
			}
			if len(date) == 0 && options["ma"] == "true" {
				date = time.Now().Format("2006-01-02")
			}
			if len(date) == 0 && options["tegnap"] == "true" {
				date = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
			}
			data := Payload{
				SpentOn: date,
				Hours:   time_to_log,
				Comment: Comment{Format: "plain", HTML: "GoLang"},
				Links: Links{
					WorkPackage: WorkPackage{Href: "/api/v3/work_packages/" + wp},
					Activity:    Activity{Href: "/api/v3/time_entries/activities/7"},
				},
			}
			payloadBytes, err := json.Marshal(data)
			if err != nil {
				fmt.Print(err)
				return -1
			}
			body := bytes.NewReader(payloadBytes)
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: tr}
			req, err := http.NewRequest("POST", "https://openproject.guidance.hu/api/v3/time_entries", body)
			if err != nil {
				fmt.Print(err)
				return -1
			}
			req.SetBasicAuth("apikey", access_token)
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				fmt.Print(err)
				return -1
			}
			defer resp.Body.Close()
			return 0
		})
	ma := cli.NewCommand("ma", "mit logoltam ma?").
		WithAction(func(args []string, options map[string]string) int {
			//todo
			return 0
		})
	_ = ma

	app := cli.New("OpenProject idő rögzítés cli").
		WithCommand(time)

	os.Exit(app.Run(os.Args, os.Stdout))
}

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	traceCmd = &cobra.Command{
		Use:   "trace",
		Short: "IP Tracing",
		Long:  "IP Tracing Long",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				for _, ip := range args {
					showData(ip)
				}
			} else {
				fmt.Println("Please provide some IP")
			}
		},
	}
)

type IP struct {
	IP       string `json::"ip"`
	City     string `json::"city"`
	Region   string `json::"region"`
	Country  string `json::"country"`
	Loc      string `json::"loc"`
	Org      string `json::"org"`
	Postal   string `json::"postal"`
	TimeZone string `json::"timezone"`
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

func showData(ip string) {
	bytes := getData(ip)

	data := IP{}
	err := json.Unmarshal(bytes, &data)

	if err != nil {
		log.Println(err.Error())
	}

	title := `
	IP Tracked!
	`

	d := color.New(color.FgCyan, color.Bold)
	d.Printf("\n\tKarakorum IP Tracker\n")

	d = color.New(color.FgRed, color.Bold)
	d.Printf(title)

	message := `
	IP: ` + data.IP + `
	Company: ` + data.Org + `
	City: ` + data.City + `
	Zipcode: ` + data.Postal + `
	Region: ` + data.Region + `
	Country: ` + data.Country + `
	TimeZone: ` + data.TimeZone + `
	Coordinates: ` + data.Loc + `
	`
	color.Green(message)
}

func getData(ip string) []byte {
	url := "https://ipinfo.io/" + ip + "/geo"
	response, err := http.Get(url)

	if err != nil {
		log.Println(err.Error())
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
	}

	return responseBytes
}

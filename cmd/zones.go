/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thspinto/isecnet-go/pkg/alarm"
)

// zonesCmd represents the zones command
var zonesCmd = &cobra.Command{
	Use:   "zones",
	Short: "Get Zone status",
	Long: `Gets the status for all zones in the alarm central.

	You can configure the zone names in '.isecnet-go'.
	If zone names are set, all unamed zones will be ignored.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		client := alarm.NewClient(viper.GetString("alarm_host"), viper.GetString("alarm_port"), viper.GetString("alarm_password"))
		zones, err := client.GetZones(context.Background())
		if err != nil {
			log.Fatalln("Failed to get zone status: ", err)
		}
		if viper.GetBool("watch") {
			watchStatus(client, zones)
		} else {
			printZones(client, zones)
		}

	},
}

func watchStatus(c alarm.AlarmClient, zones []alarm.ZoneModel) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	updateUI(c, zones)

	tick := time.Tick(2 * time.Second)
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-tick:
			updateUI(c, zones)
		}
	}
}

func updateUI(c alarm.AlarmClient, zones []alarm.ZoneModel) {
	viewCount := 0
	for _, z := range zones {
		if z.Name == "" && !viper.GetBool("all") {
			continue
		}
		p := widgets.NewParagraph()
		p.SetRect(0+(15*int(viewCount/4)), 0+(5*(viewCount%4)), 15+(15*int(viewCount/4)), 5+(5*(viewCount%4)))
		p.BorderStyle.Fg = ui.ColorGreen
		p.TitleStyle.Bg = ui.ColorClear
		p.Text = z.Name
		p.Title = z.Status
		switch z.Status {
		case "Open":
			p.BorderStyle.Fg = ui.ColorCyan
		case "Violated":
			p.BorderStyle.Fg = ui.ColorRed
		case "Anulated":
			p.BorderStyle.Fg = ui.ColorWhite
		}
		viewCount++
		ui.Render(p)
	}
}

func printZones(c alarm.AlarmClient, zones []alarm.ZoneModel) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Zone", "Name", "Status"})

	for _, z := range zones {
		if z.Name == "" && !viper.GetBool("all") {
			continue
		}
		table.Append([]string{
			strconv.Itoa(z.Id),
			z.Name,
			z.Status,
		})
	}
	table.Render()
}

func init() {
	rootCmd.AddCommand(zonesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zonesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zonesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	zonesCmd.Flags().BoolP("watch", "w", false, "Watch Zone Status")
	zonesCmd.Flags().Bool("all", false, "Show all zones")
	if err := viper.BindPFlags(zonesCmd.LocalFlags()); err != nil {
		panic(err)
	}
}

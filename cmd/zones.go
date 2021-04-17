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
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thspinto/isecnet-go/pkg/client"
)

type ZonesDescription struct {
	Id          int    `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

// zonesCmd represents the zones command
var zonesCmd = &cobra.Command{
	Use:   "zones",
	Short: "Get Zone status",
	Long: `Gets the status for all zones in the alarm central.

	You can configure the zone names in '.isecnet-go'.
	If zone names are set, all unamed zones will be ignored.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.NewClient(viper.GetString("host"), viper.GetString("port"), viper.GetString("password"))
		if err != nil {
			log.Fatal(err)
		}
		var zonesDesc []ZonesDescription
		err = viper.UnmarshalKey("zones", &zonesDesc)
		if err != nil {
			log.Fatalf("unable to decode into struct: %v", err)
		}

		if viper.GetBool("watch") {
			watchStatus(client, zonesDesc)
		} else {
			printZones(client, zonesDesc)
		}

	},
}

func watchStatus(c *client.Client, zonesDesc []ZonesDescription) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	updateUI(c, zonesDesc)

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
			updateUI(c, zonesDesc)
		}
	}
}

func updateUI(c *client.Client, zonesDesc []ZonesDescription) {
	status, err := c.GetPartialStatus()
	if err != nil {
		log.Fatal(err)
	}
	zones := status.Zones

	for i, z := range zonesDesc {
		p := widgets.NewParagraph()
		p.SetRect(0+(15*int(i/3)), 0+(5*(i%3)), 15+(15*int(i/3)), 5+(5*(i%3)))
		p.BorderStyle.Fg = ui.ColorGreen
		p.TitleStyle.Bg = ui.ColorClear
		p.Text = z.Name

		switch {
		case zones[z.Id-1].Open:
			p.Title = "Open"
			p.BorderStyle.Fg = ui.ColorCyan
		case zones[z.Id-1].Violated:
			p.Title = "Violated"
			p.BorderStyle.Fg = ui.ColorRed
		case zones[z.Id-1].Anulated:
			p.Title = "Anulated"
			p.BorderStyle.Fg = ui.ColorWhite
		}
		ui.Render(p)
	}
}

func printZones(c *client.Client, zonesDesc []ZonesDescription) {
	status, err := c.GetPartialStatus()
	if err != nil {
		log.Fatal(err)
	}

	if len(zonesDesc) > 0 {
		showConfiguredZones(zonesDesc, status.Zones)
	} else {
		showAllZones(status.Zones)
	}

}

func showConfiguredZones(zoneDesc []ZonesDescription, zones []client.Zone) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Zone", "Anulated", "Open", "Violated", "LowBattery", "Tamper", "Short Circuit"})

	for _, z := range zoneDesc {
		table.Append([]string{
			z.Name,
			strconv.FormatBool(zones[z.Id-1].Anulated),
			strconv.FormatBool(zones[z.Id-1].Open),
			strconv.FormatBool(zones[z.Id-1].Violated),
			strconv.FormatBool(zones[z.Id-1].LowBattery),
			strconv.FormatBool(zones[z.Id-1].Tamper),
			strconv.FormatBool(zones[z.Id-1].ShortCircuit),
		})
	}
	table.Render()
}

func showAllZones(zones []client.Zone) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Zone", "Anulated", "Open", "Violated", "LowBattery", "Tamper", "Short Circuit"})

	for i, z := range zones {
		table.Append([]string{
			"Zone " + fmt.Sprint(i+1),
			strconv.FormatBool(z.Anulated),
			strconv.FormatBool(z.Open),
			strconv.FormatBool(z.Anulated),
			strconv.FormatBool(z.LowBattery),
			strconv.FormatBool(z.Tamper),
			strconv.FormatBool(z.ShortCircuit),
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
	if err := viper.BindPFlags(zonesCmd.LocalFlags()); err != nil {
		panic(err)
	}
}

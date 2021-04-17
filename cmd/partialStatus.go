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

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thspinto/isecnet-go/pkg/client"
)

// partialStatusCmd represents the partialStatus command
var partialStatusCmd = &cobra.Command{
	Use:   "partialStatus",
	Short: "Get partial central status",
	Long:  `Get partial central status. This returns all info about the central and the connected peripherics.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.NewClient(viper.GetString("host"), viper.GetString("port"), viper.GetString("password"))
		if err != nil {
			log.Fatal(err)
		}
		status, err := client.GetPartialStatus()
		if err != nil {
			log.Fatal(err)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Zone", "Anulated", "Open", "Violated", "LowBattery", "Tamper", "Short Circuit"})
		for i, z := range status.Zones {
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

		table = tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Keyboard", "Issue", "Tamper", "Receiver Issue"})
		for i, k := range status.Keyboards {
			table.Append([]string{
				"Keyboard " + fmt.Sprint(i+1),
				strconv.FormatBool(k.Issue),
				strconv.FormatBool(k.Tamper),
				strconv.FormatBool(k.ReceiverIssue),
			})
		}
		table.Render()

		table = tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Partition", "Enabled"})
		for i, p := range status.Partitions {
			table.Append([]string{
				"Partition " + fmt.Sprint(i+1),
				strconv.FormatBool(p.Enabled),
			})
		}
		table.Render()

		fmt.Println("YYYY-MM-DD: ", status.Date)
		fmt.Printf("Model: %s, Firmware: %s\n", status.Central.Model, status.Central.Firmware)
		fmt.Printf("Central \n Activated: %v\n Alerting: %v\n IssueWarning: %v\n",
			status.Central.Activated,
			status.Central.Alerting,
			status.Central.IssueWarn)
		fmt.Printf("Siren\n Enabled: %v\n WireCut: %v\n ShortCircuit %v\n",
			status.Central.Siren.Enabled,
			status.Central.Siren.WireCut,
			status.Central.Battery.ShortCircuit)
		fmt.Printf("External Power Failure: %v\nPhoneLineCut: %v\n",
			status.Central.ExternalPowerFault,
			status.Central.PhoneLineCut)
	},
}

func init() {
	rootCmd.AddCommand(partialStatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// partialStatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// partialStatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

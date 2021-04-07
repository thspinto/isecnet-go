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
		client := client.NewClient(viper.GetString("host"), viper.GetString("port"), viper.GetString("password"))
		status, err := client.GetPartialStatus()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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

		fmt.Printf("%+v\n", status.Central)
		fmt.Println("YYYY-MM-DD: ", status.Date)
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

// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (

	"github.com/spf13/cobra"
//	"golang.org/x/net/context"

	// api "github.com/ploio/ploio/cmd/ploioctl/apiclient"
	// pp "github.com/ploio/ploio/pkg/api/ploioproto"
)

// applicationGetCmd represents the applicationGet command
var applicationGetCmd = &cobra.Command{
	Use:   "application",
	Short: "",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("applicationGet called")

		// ag := &pp.ApplicationGet{}

		// al, err := api.Client.ListApplications(context.Background(), ag)
		// if err != nil {
		// 	fmt.Print(err)
		// }
		// fmt.Println("ID\tName\tOwner\tRepo")
		// for _, a := range al.Applications {

		// 	fmt.Println(strconv.Itoa(int(a.ID)) + "\t" + a.Name + "\t" + a.Owner + "\t" + a.Repo)
		// }
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	getCmd.AddCommand(applicationGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applicationGetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applicationGetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

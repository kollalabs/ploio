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
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/context"

	"github.com/spf13/cobra"

	api "github.com/ploio/ploio/cmd/ploioctl/apiclient"
	pp "github.com/ploio/ploio/pkg/api/ploioproto"
)

// applicationCreateCmd represents the applicationCreate command
var applicationCreateCmd = &cobra.Command{
	Use:   "application",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating application: " + args[0])
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Email address of owner: ")
		owner, _ := reader.ReadString('\n')
		owner = strings.TrimSpace(owner)
		fmt.Print("Container Repo: ")
		repo, _ := reader.ReadString('\n')
		repo = strings.TrimSpace(repo)

		fmt.Println("\nCreating application:")
		fmt.Println("Name: " + args[0])
		fmt.Println("Owner: " + owner)
		fmt.Println("Repo: " + repo)
		ac := &pp.ApplicationCreate{
			Name:  args[0],
			Owner: owner,
			Repo:  repo,
		}
		_, err := api.Client.CreateApplication(context.Background(), ac)
		if err != nil {
			fmt.Print(err)
			return
		}
	},

	Args: cobra.ExactArgs(1),
}

func init() {
	createCmd.AddCommand(applicationCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applicationCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applicationCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

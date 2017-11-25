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
	"log"
	"fmt"

	"github.com/spf13/cobra"
	"io/ioutil"
	"golang.org/x/net/context"

	api "github.com/ploio/ploio/cmd/ploioctl/apiclient"
	//pp "github.com/ploio/ploio/pkg/api/ploioproto"
	"github.com/ploio/ploio/pkg/parser"
)

var yamlFile string

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply ploio yaml config",
	Long: `To manage infrastructure as code, you can apply ploio YAML files to add/update applications
	Example: ploioctl apply -f myapp.yaml
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apply called")
		fmt.Println("Applying config file: " + yamlFile)
		yaml, err := ioutil.ReadFile(yamlFile)
		if err != nil {
			log.Fatalf("yamlFile.Get err   #%v ", err)
		}

		a, err := parser.Marshal(yaml)
		if err != nil {
			log.Fatalf("Couldn't parse YAML: %s", err)
		}
		_, err = api.Client.UpsertApplication(context.Background(), a)
		if err != nil {
			log.Fatalf("Error applying config: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(applyCmd)

	applyCmd.Flags().StringVarP(&yamlFile, "file", "f", "", "ploio application file in YAML format")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

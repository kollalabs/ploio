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
	"context"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"

	api "github.com/ploio/ploio/cmd/ploioctl/apiclient"
	pp "github.com/ploio/ploio/pkg/api/ploioproto"
)

var pipelinefile string

// pipelineCreateCmd represents the pipelineCreate command
var pipelineCreateCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Create a pipeline for an application",
	Long:  `ploioctl create pipeline [application-name] -f mytomlfile.toml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pipelineCreate called")
		var err error
		pc := &pp.PipelineCreate{}
		if pipelinefile != "" {
			err = createPipelineFromFile(pipelinefile, pc)
			if err != nil {
				fmt.Print(err)
				return
			}
			pc.Application = args[0]
			fmt.Printf("%+v", pc)
		} else {
			fmt.Println("Sorry, pipelines have to be created with the -f flag currently")
			return
		}
		_, err = api.Client.CreatePipeline(context.Background(), pc)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
	Args: cobra.ExactArgs(1),
}

func createPipelineFromFile(file string, pc *pp.PipelineCreate) error {
	_, err := toml.DecodeFile(file, pc)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func init() {
	pipelineCreateCmd.Flags().StringVarP(&pipelinefile, "file", "f", "", "Build pipeline from a TOML file")
	createCmd.AddCommand(pipelineCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pipelineCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pipelineCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
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
	"fmt"

	"github.com/spf13/cobra"
)

// applicationDescribeCmd represents the applicationDescribe command
var applicationDescribeCmd = &cobra.Command{
	Use:   "application",
	Short: "describe an application in more detail",
	Long:  `This will list out all the details of your application in more detail`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("applicationDescribe called")

	},
	Args: cobra.ExactArgs(1),
}

func init() {
	describeCmd.AddCommand(applicationDescribeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applicationDescribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applicationDescribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

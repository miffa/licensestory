/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

// genuuidCmd represents the genuuid command
var genuuidCmd = &cobra.Command{
	Use:   "genuuid",
	Short: "生成uuid",
	Long:  `根据 sa token生成uuid`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("genuuid called")
		b := sha256.Sum256([]byte(uuid))

		myuuid := base64.StdEncoding.EncodeToString(b[:])
		fmt.Printf("uuid from token is :%s", myuuid)
	},
}

func init() {
	rootCmd.AddCommand(genuuidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genuuidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genuuidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	genuuidCmd.Flags().StringVarP(&uuid, "token", "t", "", "get  uuid from token")
	genuuidCmd.MarkFlagRequired("token")
}

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
	"fmt"

	"github.com/spf13/cobra"
)

// genkeyCmd represents the genkey command
var (
	genkeyCmd = &cobra.Command{
		Use:   "genkey",
		Short: "生成rsa公钥私钥",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generate rsa key")
			fmt.Printf("private key: %s%s\n", prefix, PrivateKey)
			fmt.Printf("public key: %s%s\n", prefix, PublicKey)
			GenerateKey(prefix)
		},
	}
)

func init() {
	rootCmd.AddCommand(genkeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genkeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genkeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	genkeyCmd.Flags().StringVarP(&prefix, "prefix", "p", "rsa", "key's prefix")
	genkeyCmd.MarkFlagRequired("prefix")
}

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
	"encoding/json"
	"fmt"

	license "code.troila.io/license/licenseplate"
	utils "code.troila.io/license/utils"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "验证license",
	Long:  `./tpaaslicense verify -f testlicense.txt -u 5555666`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("license验证开始")
		ss, err := license.NewLicenseSigner(uuid, prefix+PublicKey, prefix+PrivateKey)
		if err != nil {
			fmt.Printf("签名算法出现问题:%v\n", err)
			return
		}
		licensebyte, err := utils.LoadFile(licensename)
		if err != nil {
			fmt.Printf("读取license文件%s失败:%v", err)
			return
		}
		fmt.Println("license:", string(licensebyte))

		id, err := ss.Verify(string(licensebyte))
		if err != nil {
			fmt.Printf("verify license err:%v \n", err)
			return
		}
		ep, _ := json.Marshal(id)
		fmt.Printf("license校验成功:\n%s\n", string(ep))
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	verifyCmd.Flags().StringVarP(&prefix, "prefix", "p", "rsa", "key's prefix")
	verifyCmd.MarkFlagRequired("prefix")
	verifyCmd.Flags().StringVarP(&licensename, "file", "f", "license.txt", "license file name")
	verifyCmd.MarkFlagRequired("file")

	verifyCmd.Flags().StringVarP(&uuid, "uuid", "u", "top500", "uuid")
	verifyCmd.MarkFlagRequired("uuid")
}

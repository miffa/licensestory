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

	license "code.troila.io/license/licenseplate"

	"github.com/spf13/cobra"
)

// getuuidCmd represents the getuuid command
var getuuidCmd = &cobra.Command{
	Use:   "getuuid",
	Short: "解密uuid",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("解密uuid开始")

		ss, err := license.NewLicenseDecryptor(prefix + "_private.pem")
		if err != nil {
			fmt.Printf("RSA解密算法出现问题:%v\n", err)
			return
		}

		youruuid, err := ss.Decrypt(uuid)
		if err != nil {
			fmt.Printf("解密失败:%v\n", err)
			return
		}
		//var d1 = []byte(LicenseText)
		//err = ioutil.WriteFile(licensename, d1, 0666)
		//if err != nil {
		//	fmt.Printf("license文件生成失败:%v\n", err)
		//	return
		//}
		fmt.Printf("uuid解密成功：%s\n", youruuid)
	},
}

func init() {
	rootCmd.AddCommand(getuuidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getuuidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getuuidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getuuidCmd.Flags().StringVarP(&uuid, "uuid", "u", "", "uuid encrypt")
	getuuidCmd.MarkFlagRequired("uuid")
	getuuidCmd.Flags().StringVarP(&prefix, "prefix", "p", "rsa", "key's prefix")
	getuuidCmd.MarkFlagRequired("prefix")
}

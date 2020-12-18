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
	"io/ioutil"

	"github.com/spf13/cobra"

	license "code.miffa.io/license/licenseplate"
)

// signCmd represents the sign command
var (
	signCmd = &cobra.Command{
		Use:   "sign",
		Short: "生成企业license",
		Long:  `./tpaaslicense sign  -f testlicense.txt -u 5555666 -c '''{"corporation":"badluckin.com.ltd","quota":40,"expired_time":"2022-02-03 23:59:59","extension":"thi is a test","version":"v2.34"}'''`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("签名服务开始运行")

			ss, err := license.NewLicenseSigner(uuid, prefix+"_public.pem", prefix+"_private.pem")
			if err != nil {
				fmt.Printf("签名算法出现问题:%v\n", err)
				return
			}

			var mayday license.TpaasLicenseMeta
			err = json.Unmarshal([]byte(corporation), &mayday)
			if err != nil {
				fmt.Printf("企业信息格式非法，不是json\n")
				return
			}

			//mayday := license.TpaasLicenseMeta{
			//	Corporation: "badluckin.com.ltd",
			//	Quota:       40,
			//	ExpiredTime: "2022-02-03 23:59:59",
			//	Extension:   "thi is a test",
			//	Version:     "v2.34",
			//}
			LicenseText, err := ss.Sign(mayday)
			if err != nil {
				fmt.Printf("签名失败:%v\n", err)
				return
			}
			var d1 = []byte(LicenseText)
			err = ioutil.WriteFile(licensename, d1, 0666)
			if err != nil {
				fmt.Printf("license文件生成失败:%v\n", err)
				return
			}
			fmt.Printf("license文件生成成功：%s\n", licensename)
		},
	}
)

func init() {
	rootCmd.AddCommand(signCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	signCmd.Flags().StringVarP(&uuid, "uuid", "u", "top500", "uuid")
	signCmd.MarkFlagRequired("uuid")
	signCmd.Flags().StringVarP(&licensename, "file", "f", "license.txt", "license file name")
	signCmd.MarkFlagRequired("file")

	signCmd.Flags().StringVarP(&corporation, "corporation", "c", "", "corporation information")
	signCmd.MarkFlagRequired("corporation")

	signCmd.Flags().StringVarP(&prefix, "prefix", "p", "rsa", "key's prefix")
	signCmd.MarkFlagRequired("prefix")
}

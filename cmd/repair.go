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
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"code.troila.io/license/crypt"
	utils "code.troila.io/license/utils"

	"github.com/spf13/cobra"
)

// repairCmd represents the repair command
var repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "修复license时间戳数据",
	Run: func(cmd *cobra.Command, args []string) {

		pubkey, err := utils.LoadFile(prefix + "_public.pem")
		if err != nil {
			fmt.Printf("read public key err:%v", err)
			return
		}

		my := strconv.FormatInt(time.Now().Unix(), 10)
		tpmt, _ := crypt.AesEncrypt([]byte(my), pubkey)
		mydata := base64.StdEncoding.EncodeToString(tpmt)
		fmt.Println("repair called")
		fmt.Printf("now sec:%s, mydata:%s\n", my, mydata)
		fmt.Printf(`insert into t_tpaas_ntp values(0, '%s', %s)\n`, mydata, my)

		vtpmt, err := base64.StdEncoding.DecodeString(mydata)
		if err != nil {
			fmt.Printf("repair (DecodeString) err:%v", err)
			return
		}
		myss, err := crypt.AesDecrypt(vtpmt, pubkey)
		if err != nil {
			fmt.Printf("repair (RsaDecrypt) err:%v", err)
			return
		}
		fmt.Printf("now sec:%s, de mydata:%s\n", my, string(myss))
	},
}

func init() {
	rootCmd.AddCommand(repairCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repairCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repairCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	repairCmd.Flags().StringVarP(&prefix, "prefix", "p", "rsa", "key's prefix")

}

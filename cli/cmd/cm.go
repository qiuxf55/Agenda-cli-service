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
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/spf13/cobra"
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "createmeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetString("participators")
		startDate, _ := cmd.Flags().GetString("startDate")
		endDate, _ := cmd.Flags().GetString("endDate")

		data := make(url.Values)
		data["title"] = []string{title}
		data["participators"] = []string{participators}
		data["start"] = []string{startDate}
		data["end"] = []string{endDate}
	//把post表单发送给目标服务器

		res, err := http.PostForm("http://localhost:8080/v1/meetings", data)
		//设置http中header参数，可以再此添加cookie等值
		//res.Header.Add("User-Agent", "***")
		//res.Header.Add("http.socket.timeou", 5000)

		if err != nil {
			fmt.Println("create meeting failed")
		    fmt.Println(err.Error())
		    return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		fmt.Println("create meeting success")
		fmt.Println(string(body))

	},
}

func init() {
	RootCmd.AddCommand(cmCmd)
	cmCmd.Flags().StringP("title", "t", "Anonymous", "提供会议主题名字")
	cmCmd.Flags().StringP("participators", "p","Anonymous", "提供一个会议的参与者")
	cmCmd.Flags().StringP("startDate", "s", "Anonymous", "会议开始时间，格式为YYYY-MM-DD/HH:MM,")
	cmCmd.Flags().StringP("endDate", "e", "Anonymous", "会议结束时间，格式为YYYY-MM-DD/HH:MM,")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

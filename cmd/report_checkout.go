package cmd

// Copyright © 2019 oleg2807@gmail.com
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

import (
	"fmt"
	"os"
	"time"

	"github.com/olegsu/timewatch/pkg/timewatch"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var checkoutCmdOptions struct {
	time string
}

// checkoutCmd represents the status command
var checkoutCmd = &cobra.Command{
	Use: "checkout",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		if reportCmdOptions.date == "" {
			t := now.Format("2006-02-01")
			reportCmdOptions.date = t
		}
		if checkinCmdOptions.time == "" {
			t := now.Format("15:04")
			checkinCmdOptions.time = t
		}
		tw, err := timewatch.RestoreLogin()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = tw.Login()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = tw.Report(timewatch.ActionCheckOUT, &timewatch.ReportOptions{
			Report: timewatch.Report{
				Checkout: checkoutCmdOptions.time,
			},
			Date: reportCmdOptions.date,
		})
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	reportCmd.AddCommand(checkoutCmd)
	checkoutCmd.Flags().StringVar(&checkoutCmdOptions.time, "time", viper.GetString("time"), "Set the date in format HH-MM [$TW_TIME]")
}

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
	"strings"
	"time"

	"github.com/olegsu/timewatch/pkg/timewatch"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var checkinCmdOptions struct {
	time string
}

// checkinCmd represents the status command
var checkinCmd = &cobra.Command{
	Use: "checkin",
	Run: func(cmd *cobra.Command, args []string) {
		log := buildLogger("report checkin")
		now := time.Now()
		if reportCmdOptions.date == "" {
			t := now.Format("2006-01-02")
			log.Debug("Date is not set, setting default", "date", t)
			reportCmdOptions.date = t
		}
		if checkinCmdOptions.time == "" {
			t := now.Format("15:04")
			t = strings.Replace(t, ":", "-", 1)
			log.Debug("Time is not set, setting default", "time", t)
			checkinCmdOptions.time = t
		}
		tw, err := timewatch.RestoreLogin(log)
		dieOnError(err, log)
		err = tw.Login()
		dieOnError(err, log)
		err = tw.Report(timewatch.ActionCheckIN, &timewatch.ReportOptions{
			Report: timewatch.Report{
				Checkin: checkinCmdOptions.time,
			},
			Date: reportCmdOptions.date,
		})
		dieOnError(err, log)
	},
}

func init() {
	reportCmd.AddCommand(checkinCmd)
	viper.BindEnv("time", "TW_TIME")
	checkinCmd.Flags().StringVar(&checkinCmdOptions.time, "time", viper.GetString("time"), "Set the date in format HH-MM [$TW_TIME]")
}

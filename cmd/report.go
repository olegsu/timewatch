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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var reportCmdOptions struct {
	date string
}

// reportCmd represents the status command
var reportCmd = &cobra.Command{
	Use: "report",
}

func init() {
	rootCmd.AddCommand(reportCmd)

	viper.BindEnv("date", "TW_DATE")

	reportCmd.PersistentFlags().StringVar(&reportCmdOptions.date, "date", viper.GetString("date"), "Set the date in format YYYY-MM-DD [$TW_DATE]")
}

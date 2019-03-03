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

	"github.com/olegsu/timewatch/pkg/timewatch"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmdOptions struct {
	compony  string
	user     string
	password string
}

// initCmd represents the status command
var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(initCmdOptions)
		tw := timewatch.New(&timewatch.TimewatchOptions{
			User:     initCmdOptions.user,
			Compony:  initCmdOptions.compony,
			Password: initCmdOptions.password,
		})
		err := tw.Login()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	viper.BindEnv("compony", "TW_COMPONY_ID")
	viper.BindEnv("user", "TW_USER_ID")
	viper.BindEnv("password", "TW_PASSWORD")

	initCmd.Flags().StringVar(&initCmdOptions.compony, "compony-id", viper.GetString("compony"), "Set the id of the compony [$TW_COMPONY_ID]")
	initCmd.Flags().StringVar(&initCmdOptions.user, "user-id", viper.GetString("user"), "Set the id of the user [$TW_USER_ID]")
	initCmd.Flags().StringVar(&initCmdOptions.password, "password", viper.GetString("password"), "Set the password [$TW_PASSWORD]")
}

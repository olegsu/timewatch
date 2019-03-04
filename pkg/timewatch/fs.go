package timewatch

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
	"io/ioutil"
	"os"

	"github.com/olegsu/timewatch/pkg/logger"
	"gopkg.in/yaml.v2"
)

const path = ".timewatch.yaml"

func persistLogin(tw *tw) error {
	res, err := yaml.Marshal(tw)
	if err != nil {
		return err
	}
	p := fmt.Sprintf("%s/%s", os.Getenv("HOME"), path)
	err = ioutil.WriteFile(p, res, 0644)
	if err != nil {
		return err
	}
	return nil
}

func RestoreLogin(log logger.Logger) (Timewatch, error) {
	p := fmt.Sprintf("%s/%s", os.Getenv("HOME"), path)
	log.Debug("Reading config file from", "path", p)
	f, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	tw := &tw{}
	tw.log = log
	err = yaml.Unmarshal(f, tw)
	if err != nil {
		return nil, err
	}
	return tw, nil
}

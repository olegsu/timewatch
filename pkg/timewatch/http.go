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
	"net/http"
	"strings"

	"github.com/olegsu/timewatch/pkg/logger"
)

const (
	baseURL = "https://checkin.timewatch.co.il/punch"
	login   = "punch2.php"
	edit    = "editwh3.php"
	daily   = "editwh2.php"
)

type (
	requestOptions struct {
		data    map[string]string
		qs      map[string]string
		cookies []*http.Cookie
		url     string
		method  string
		log     logger.Logger
	}
)

func doAPICall(opt *requestOptions) (*http.Response, error) {
	client := &http.Client{}
	requestURL := fmt.Sprintf("%s/%s", baseURL, opt.url)
	if opt.qs != nil {
		requestURL += "?"
	}
	for k, v := range opt.qs {
		requestURL += fmt.Sprintf("%s=%s&", k, v)
	}
	opt.log.Debug("Final request URL", "url", requestURL)
	finalPayloadString := ""
	for k, v := range opt.data {
		d := fmt.Sprintf("%s=%s", k, v)
		opt.log.Debug("Adding data to payload", "data", d)
		finalPayloadString += d + "&"
	}
	payload := strings.NewReader(finalPayloadString)

	request, err := http.NewRequest(opt.method, requestURL, payload)
	if err != nil {
		return nil, err
	}
	for _, c := range opt.cookies {
		request.AddCookie(c)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", "http://checkin.timewatch.co.il/punch/editwh2.php")
	return client.Do(request)
}

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
	"bytes"
	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func findEmployeeIDFromHTML(htm []byte) (string, error) {
	r := bytes.NewReader(htm)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}
	selection := doc.Find("input[id='ixemplee']")
	res, found := selection.Attr("value")
	if !found {
		return "", errors.New("Not foun value of input id=ixemplee")
	}
	return res, nil
}

func getInputValue(htm []byte, input string) (string, error) {
	r := bytes.NewReader(htm)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}
	selection := doc.Find(fmt.Sprintf("input[name='%s']", input))
	res, found := selection.Attr("value")
	if !found {
		return "", errors.New(fmt.Sprintf("Not found value of input id=%s", input))
	}
	return res, nil
}

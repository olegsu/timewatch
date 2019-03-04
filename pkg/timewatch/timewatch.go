package timewatch

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/olegsu/timewatch/pkg/logger"
)

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

const (
	ActionCheckIN  = "checkin"
	ActionCheckOUT = "checkout"
)

type (
	Timewatch interface {
		Login() error
		Report(action string, opt *ReportOptions) error
	}

	TimewatchOptions struct {
		User     string
		Compony  string
		Password string
		Log      logger.Logger
	}

	ReportOptions struct {
		Date   string
		Report Report
	}

	tw struct {
		cookies    []*http.Cookie
		User       string `yaml:"user"`
		Compony    string `yaml:"compony"`
		Password   string `yaml:"password"`
		EmployeeID string `yaml:"employee_id"`
		log        logger.Logger
	}

	dailyReportOptions struct {
		date string
	}

	dailyReport struct {
		report Report
	}

	Report struct {
		Checkin  string
		Checkout string
	}
)

func New(opt *TimewatchOptions) Timewatch {
	return &tw{
		User:     opt.User,
		Compony:  opt.Compony,
		Password: opt.Password,
		log:      opt.Log,
	}
}

func (t *tw) Login() error {
	resp, err := doAPICall(&requestOptions{
		url:    login,
		method: "POST",
		data: map[string]string{
			"comp": t.Compony,
			"name": t.User,
			"pw":   t.Password,
		},
		log: t.log,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	t.cookies = resp.Cookies()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if strings.Contains(string(body), "The login details you entered are incorrect") {
		return errors.New("Failed to login")
	}
	t.log.Debug("Logged in")
	id, err := findEmployeeIDFromHTML(body)
	if err != nil {
		return err
	}
	t.EmployeeID = id
	return persistLogin(t)
}

func (t *tw) Report(action string, opt *ReportOptions) error {
	err := t.buildDailyReport(opt)
	if err != nil {
		return err
	}
	if action == ActionCheckIN {
		return t.doCheck(opt)
	}
	if action == ActionCheckOUT {
		return t.doCheck(opt)
	}
	return errors.New("Action not found")
}
func (t *tw) doCheck(opt *ReportOptions) error {
	data := map[string]string{
		"e":  t.EmployeeID,
		"tl": t.EmployeeID,
		"c":  t.Compony, // compony
		"d":  opt.Date,  // date to change
	}
	if opt.Report.Checkout != "" {
		t := strings.Split(opt.Report.Checkout, "-")
		data["xhh0"] = t[0]
		data["xmm0"] = t[1]
	}
	if opt.Report.Checkin != "" {
		t := strings.Split(opt.Report.Checkin, "-")
		data["ehh0"] = t[0]
		data["emm0"] = t[1]
	}
	reqOpt := &requestOptions{
		url:     edit,
		method:  "POST",
		data:    data,
		cookies: t.cookies,
		log:     t.log,
	}
	resp, err := doAPICall(reqOpt)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyStr := string(body)
	if strings.Contains(bodyStr, "TimeWatch - Reject") {
		return errors.New("TimeWatch - Reject")
	}
	if bodyStr == "error" {
		return errors.New("error")
	}
	return nil
}
func (t *tw) buildDailyReport(opt *ReportOptions) error {
	resp, err := doAPICall(&requestOptions{
		url:    daily,
		method: "GET",
		qs: map[string]string{
			"ie": t.Compony,
			"e":  t.EmployeeID,
			"tl": t.EmployeeID,
			"d":  opt.Date,
		},
		cookies: t.cookies,
		log:     t.log,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if opt.Report.Checkin == "" {
		ehh0, err := getInputValue(body, "ehh0")
		emm0, err := getInputValue(body, "emm0")
		if err != nil {
			return err
		}
		if ehh0 != "" && emm0 != "" {
			opt.Report.Checkin = fmt.Sprintf("%s-%s", ehh0, emm0)
		}
	}
	if opt.Report.Checkout == "" {
		xhh0, err := getInputValue(body, "xhh0")
		xmm0, err := getInputValue(body, "xmm0")
		if err != nil {
			return err
		}
		if xhh0 != "" && xmm0 != "" {
			opt.Report.Checkout = fmt.Sprintf("%s-%s", xhh0, xmm0)
		}
	}
	return nil
}

/*
Copyright © 2024 Taisuke Miyazaki <imishinist@gmail.com>

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
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/imishinist/tmt/internal"
)

func parseDate(date string) (time.Time, error) {
	ret, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}
	return ret, nil
}

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test recurrence rule",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		recurrence := args[0]
		dateStr := args[1]

		schedule, err := internal.ParseRecurrence(recurrence)
		if err != nil {
			return err
		}
		date, err := parseDate(dateStr)
		if err != nil {
			return err
		}

		if internal.Match(schedule, date) {
			fmt.Printf("rule: %q\n\t%q => matched\n", recurrence, dateStr)
		} else {
			fmt.Printf("rule: %q\n\t%q => not matched\n", recurrence, dateStr)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.SetUsageTemplate(`Usage:
  tmt test <recurrence> <date> [flags]

Flags:
  -h, --help   help for test

Example:
  tmt test "* * 1-5" 2024-01-01
`)
}

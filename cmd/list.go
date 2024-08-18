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
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/imishinist/tmt/internal"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := internal.LoadTasks(taskFile)
		if err != nil {
			if !errors.Is(err, internal.ErrFileNotFound) {
				return err
			}
			tasks = []internal.Task{}
		}

		header := []string{"ID", "Title", "Recurrence", "Description", "next"}
		taskData := make([][]string, 0, len(tasks))
		for i, task := range tasks {
			next, err := task.Next(time.Now())
			if err != nil {
				return err
			}
			taskData = append(taskData, []string{fmt.Sprintf("#%d", i+1), task.Title, task.Recurrence, task.Description, next.Format("2006-01-02")})
		}
		PrintAsTable(header, taskData)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

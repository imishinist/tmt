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

	"github.com/spf13/cobra"

	"github.com/imishinist/tmt/internal"
)

var (
	addingTask = internal.Task{
		Title:       "",
		Recurrence:  "* * 1-5", // @daily
		Description: "",
	}

	// addCmd represents the add command
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add the task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := internal.InitTaskFile(taskFile); err != nil {
				return err
			}
			if err := addingTask.Verify(); err != nil {
				return err
			}

			tasks, err := internal.LoadTasks(taskFile)
			if err != nil {
				if !errors.Is(err, internal.ErrFileNotFound) {
					return err
				}
				tasks = []internal.Task{}
			}
			tasks = append(tasks, addingTask)
			if err := internal.SaveTasks(taskFile, tasks); err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addingTask.Title, "title", "t", addingTask.Title, "task title")
	addCmd.Flags().StringVarP(&addingTask.Recurrence, "recurrence", "r", addingTask.Recurrence, "task recurrence")
	addCmd.Flags().StringVarP(&addingTask.Description, "description", "d", addingTask.Description, "task description")
}

/*

@Time : 2019/5/28
@Author : Jiangs

*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Services",
	Run: func(cmd *cobra.Command, args []string) {
		d, err := ioutil.ReadFile("1.lock")
		if err != nil {
			fmt.Println(err)
		} else {
			var id = string(d)
			//command:=exec.Command("tskill",id)  //windows
			command := exec.Command("kill", "-9", id) //linux
			err = command.Start()
			if err != nil {
				fmt.Println("Unable to kill the process")
			} else {
				del := os.Remove("1.lock")
				if del != nil {
					fmt.Println("Unable to delete .lock file")
				}

			}

		}
		fmt.Println("stop called")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

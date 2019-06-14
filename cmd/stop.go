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
		d1, err := ioutil.ReadFile("1.lock")
		if err != nil {
			fmt.Println(err)
		} else {
			var id = string(d1)
			command1:=exec.Command("tskill",id)  //windows
			//command1 := exec.Command("kill", "-9", id) //linux
			err = command1.Start()
			if err != nil {
				fmt.Println("Unable to kill the font process")
			} else {
				del := os.Remove("1.lock")
				if del != nil {
					fmt.Println("Unable to delete .lock file")
				}

			}

		}
		d2, err := ioutil.ReadFile("grafana-server.pid")
		if err != nil {
			fmt.Println(err)
		} else {
			var id = string(d2)
			command2:=exec.Command("tskill",id)  //windows
			//command2 := exec.Command("kill", "-9", id) //linux
			err = command2.Start()
			if err != nil {
				fmt.Println("Unable to kill the server process")
			} else {
				del := os.Remove("grafana-server.pid")
				if del != nil {
					fmt.Println("Unable to delete .pid file")
				}

			}

		}
		fmt.Println("stop called")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

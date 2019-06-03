/*

@Time : 2019/5/25
@Author : Jiangs

*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os/exec"
	"strconv"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Services",
	Run: func(cmd *cobra.Command, args []string) {
		//command:=exec.Command("./grafana")
		command := exec.Command("monitor")

		err := command.Start()
		if err != nil {
			fmt.Println(err)
		}
		var id = command.Process.Pid
		data := []byte(strconv.Itoa(id))
		err = ioutil.WriteFile("1.lock", data, 0777)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

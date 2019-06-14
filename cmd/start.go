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
		command1:=exec.Command("grafana-server.exe","--pidfile=./grafana-server.pid")
		command2:= exec.Command("statichtml.exe")
		err := command1.Start()
		if err != nil {
			fmt.Println(err)
		}
		err = command2.Start()
		if err != nil {
			fmt.Println(err)
		}
		var id = command2.Process.Pid
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

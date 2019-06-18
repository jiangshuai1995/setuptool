/*

@Time : 2019/5/25
@Author : Jiangs

*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Services",
	Run: func(cmd *cobra.Command, args []string) {
		command1 := exec.Command("./grafana-server", "--pidfile=./grafana-server.pid")
		command2 := exec.Command("./statichtml")
		if Exists("./grafana-server.pid") || Exists("./1.lock") {
			fmt.Println("请检查服务是否已经启动")
			return
		}
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
		fmt.Println("start called")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func Exists(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

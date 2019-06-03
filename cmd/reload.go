/*

@Time : 2019/6/3
@Author : Jiangs

*/
package cmd

import (
	"demo/model"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// reloadCmd represents the reload command
var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "reload config( take effect when start up )",
	Run: func(cmd *cobra.Command, args []string) {
		viper1 := viper.New() //读取用户修改的文件
		viper1.AddConfigPath("../conf")
		viper1.SetConfigType("yaml")
		viper1.SetConfigName("config")
		err := viper1.ReadInConfig()
		if err != nil {
			fmt.Println(err)
		}
		//var DB_CONF model.Db_conf
		host := viper1.GetString("database_host")
		port := viper1.GetString("database_port")
		user := viper1.GetString("database_user")
		pass := viper1.GetString("database_pass")

		d1, err := ioutil.ReadFile("../conf/provisioning/datasources/datasource.yaml") //此处输入相对路径，datsource.yaml
		if err != nil {
			fmt.Println(err)
		}
		var DS model.PreDatasource
		err = yaml.Unmarshal(d1, &DS)
		if err != nil {
			fmt.Println(err)
		}
		if DS.Datasources != nil {

			if host == "" || port == "" {
				fmt.Println("数据库地址配置出错")
			} else {
				DS.Datasources[0].Url = host + ":" + port
			}
			if user == "" {
				fmt.Println("缺少用户名信息")
			} else {
				DS.Datasources[0].User = user
			}
			if pass == "" {
				fmt.Println("密码信息为空")
			} else {
				DS.Datasources[0].SecureJsonData.Password = pass
			}
			data1, err := yaml.Marshal(&DS)
			if err != nil {
				fmt.Println(err)
			}
			err = ioutil.WriteFile("../conf/provisioning/datasources/datasource.yaml", data1, 0777) //此处填yaml相对路径及文件名
		} else {
			fmt.Println("数据源配置出错")
		}

		d2, err := ioutil.ReadFile("../conf/provisioning/notifiers/notifier.yaml") //此处输入相对路径 notifier.yaml
		if err != nil {
			fmt.Println(err)
		}
		var N model.PreNotifier
		err = yaml.Unmarshal(d2, &N)
		if err != nil {
			fmt.Println(err)
		}
		if N.Notifiers != nil {

			email := viper1.Get("email").(string)
			if email == "" {
				fmt.Println("请检查告警邮箱设置")
			} else {
				N.Notifiers[0].Settings.Addresses = email
			}
			data2, err := yaml.Marshal(&N)
			if err != nil {
				fmt.Println(err)
			}
			err = ioutil.WriteFile("../conf/provisioning/notifiers/notifier.yaml", data2, 0777) //此处填yaml相对路径及文件名
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("警报通道配置出错")
		}

		var cpu_dashboard model.Dashboard
		d3, err := ioutil.ReadFile("") //此处填包含CPU告警信息的dashboard.json
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(d3, &cpu_dashboard)
		if err != nil {
			fmt.Println(err)
		}
		cpu_threshold := viper1.Get("cpu_alert").(int)
		if cpu_threshold == 0 {
			fmt.Println("CPU报警阈值未被正确设置")
		} else {
			for i := 0; i < len(cpu_dashboard.Panels); i++ {
				if cpu_dashboard.Panels[i].Title == "CPU使用率" {
					cpu_dashboard.Panels[i].Alert.Conditions[0].Evaluator.Params = []int{cpu_threshold}
				}
			}
		}
		data3, err := json.Marshal(&cpu_dashboard)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile("", data3, 0777) //此处填json相对路径及文件名
		if err != nil {
			fmt.Println(err)
		}

		var mem_dashboard model.Dashboard
		d4, err := ioutil.ReadFile("") //此处填包含内存告警信息的dashboard.json
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(d4, &mem_dashboard)
		if err != nil {
			fmt.Println(err)
		}
		mem_threshold := viper1.Get("mem_alert").(int)
		if mem_threshold == 0 {
			fmt.Println("内存报警阈值未被正确设置")
		} else {
			for i := 0; i < len(cpu_dashboard.Panels); i++ {
				if mem_dashboard.Panels[i].Title == "CPU使用率" {
					mem_dashboard.Panels[i].Alert.Conditions[0].Evaluator.Params = []int{mem_threshold}
				}
			}
		}
		data4, err := json.Marshal(&mem_dashboard)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile("", data4, 0777) //此处填json相对路径及文件名
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("reload called")
	},
}

func init() {
	rootCmd.AddCommand(reloadCmd)
}

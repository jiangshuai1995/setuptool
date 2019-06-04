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

		var alert_dashboard model.Dashboard
		d3, err := ioutil.ReadFile("../conf/predashboards/general_view.json") //此处填包含CPU告警信息的dashboard.json
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(d3, &alert_dashboard)
		if err != nil {
			fmt.Println(err)
		}
		cpu_threshold := viper1.Get("cpu_alert").(int)
		men_threshold := viper1.Get("men_alert").(int)
		disk_threshold := viper1.Get("disk_alert").(int)
		if cpu_threshold == 0 {
			fmt.Println("CPU报警阈值未被正确设置")
		} else {
			for i := 0; i < len(alert_dashboard.Panels); i++ {
				if alert_dashboard.Panels[i].Title == "CPU使用率" {
					alert_dashboard.Panels[i].Alert.Conditions[0].Evaluator.Params = []int{cpu_threshold}
				}
			}
		}
		if men_threshold == 0 {
			fmt.Println("内存报警阈值未被正确设置")
		} else {
			for i := 0; i < len(alert_dashboard.Panels); i++ {
				if alert_dashboard.Panels[i].Title == "内存使用率" {
					alert_dashboard.Panels[i].Alert.Conditions[0].Evaluator.Params = []int{men_threshold}
				}
			}
		}

		if disk_threshold == 0 {
			fmt.Println("磁盘报警阈值未被正确设置")
		} else {
			for i := 0; i < len(alert_dashboard.Panels); i++ {
				if alert_dashboard.Panels[i].Title == "磁盘使用率" {
					alert_dashboard.Panels[i].Alert.Conditions[0].Evaluator.Params = []int{disk_threshold}
				}
			}
		}

		data3, err := json.Marshal(&alert_dashboard)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile("../conf/predashboards/general_view.json", data3, 0777) //此处填json相对路径及文件名
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("reload called")
	},
}

func init() {
	rootCmd.AddCommand(reloadCmd)
}

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

		cpu_alert:=viper1.GetInt("cpu_alert")
		mem_alert:=viper1.GetInt("mem_alert")
		disk_alert:=viper1.GetInt("disk_alert")
		bytevalue,err:=ioutil.ReadFile("../conf/predashboards/general_view.json")
		var result map[string]interface{}
		err =json.Unmarshal(bytevalue,&result)
		panels:=result["panels"].([]interface{})
		for _ ,panel:=range panels{
			m:=panel.(map[string]interface{})
			if alert,exists:=m["alert"];exists{
				if m["title"]=="CPU使用率"{
					a:=alert.(map[string]interface{})
					conds:=a["conditions"].([]interface{})
					c:=conds[0].(map[string]interface{})
					e:=c["evaluator"].(map[string](interface{}))
					e["params"]=[]int{cpu_alert} //修改params
					ts:=m["thresholds"].([]interface{})
					t:=ts[0].(map[string]interface{})
					t["value"]=cpu_alert //修改thresholds
					fmt.Printf("cpu告警阈值已修改为 %d",cpu_alert)
				}else if m["title"]=="内存使用率"{
					a:=alert.(map[string]interface{})
					conds:=a["conditions"].([]interface{})
					c:=conds[0].(map[string]interface{})
					e:=c["evaluator"].(map[string](interface{}))
					e["params"]=[]int{mem_alert}
					ts:=m["thresholds"].([]interface{})
					t:=ts[0].(map[string]interface{})
					t["value"]=mem_alert
					fmt.Printf("内存告警阈值已修改为 %d",mem_alert)
				}else if m["title"]=="磁盘使用率"{
					a:=alert.(map[string]interface{})
					conds:=a["conditions"].([]interface{})
					c:=conds[0].(map[string]interface{})
					e:=c["evaluator"].(map[string](interface{}))
					e["params"]=[]int{disk_alert}
					ts:=m["thresholds"].([]interface{})
					t:=ts[0].(map[string]interface{})
					t["value"]=disk_alert
					fmt.Printf("磁盘告警阈值已修改为 %d",disk_alert)
				}
			}

		}
		bytevalue2,err:=json.MarshalIndent(result," "," ") //注意修改encode.go文件 escapeHTML: false 不然会出现特殊字符的转义问题
		err=ioutil.WriteFile("../conf/predashboards/general_view.json",bytevalue2,0777)


		fmt.Println("reload called")
	},
}

func init() {
	rootCmd.AddCommand(reloadCmd)
}


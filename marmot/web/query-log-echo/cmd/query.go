package cmd

import (
	"log"
	"strings"
	"github.com/spf13/cobra"
	"toy/marmot/web/query-log-echo/launch"
)

var (
	confRoot string
	env     string
)

var desc = strings.Join([]string{
	"confRoot 表示配置的路基",
	"env区分配置的环境，默认是dev环境	",
}, "\n")

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "web日志查询系统",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {

		if confRoot == ""{

			log.Fatalf("暂不支持该转换模式，请执行 help query 查看帮助文档")
			return
		}

		launch.InitLaunchWithParam(confRoot, env)
	},
}

func init()  {
	queryCmd.Flags().StringVarP(&confRoot, "confRoot", "c", "", "root path of configuration")
	queryCmd.Flags().StringVarP(&env, "env", "e", "dev", "prod or dev or test")
}
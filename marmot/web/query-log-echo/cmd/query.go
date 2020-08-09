package cmd

import (
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
		launch.InitLaunchWithParam(confRoot, env)
	},
}

func init()  {
	queryCmd.Flags().StringVarP(&confRoot, "confRoot", "c", "./conf", "root path of configuration")
	queryCmd.Flags().StringVarP(&env, "env", "e", "dev", "prod or dev or test")
}
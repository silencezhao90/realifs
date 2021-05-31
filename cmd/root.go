package cmd

import (
	"filesystem/config"
	"filesystem/router"
	"filesystem/storage"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var cfgFile string
var serverPort int

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "realifs server",
	Long:  `realifs server use help get more info`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("启动参数: ", args)
		runServer()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $CURRENT_DIR/config/config.yaml)")
	rootCmd.Flags().IntVarP(&serverPort, "port", "p", 8080, "port on which the server will listen")
	rootCmd.AddCommand(versionCmd)
}

// Execute rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 初始化配置
func initConfig() {
	config.Init(cfgFile)
}

func runServer() {
	storage.Load()
	g := gin.Default()
	router.Load(g)
	g.Run(fmt.Sprintf(":%d", serverPort))
}

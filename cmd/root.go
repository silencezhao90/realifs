package cmd

import (
	"filesystem/router"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $CURRENT_DIR/config.yaml)")
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
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		path, _ := os.Getwd()
		configFile := path + "/config.yaml"
		fmt.Println(configFile)
		viper.SetConfigFile(configFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
}

func runServer() {
	g := gin.Default()
	router.Load(g)
	g.Run(fmt.Sprintf(":%d", serverPort))
}

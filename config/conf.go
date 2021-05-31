/*
 * @Author: zhaobo
 * @Date: 2020-07-24 14:33:37
 * @Last Modified by: zhaobo
 * @Last Modified time: 2020-07-27 17:19:12
 */

package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Init(filePath string) {
	if filePath == "" {
		path, _ := os.Getwd()
		filePath = path + "/config/config.yaml"
	}
	fmt.Println(filePath)

	viper.SetConfigFile(filePath)
	viper.SetEnvPrefix("ENV")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
}

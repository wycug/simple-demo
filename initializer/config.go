/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package initializer

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("viper readinconfig error: %s\n", err))
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("viper unmarshal err: %s\n", err))
	}
}

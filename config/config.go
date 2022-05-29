/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package config

// Config 组合全部配置模型
type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
}

// Server 服务启动端口号配置
type Server struct {
	Protocol string `mapstructure:"protocol"`
	Port     string `mapstructure:"port"`
}

// Mysql MySQL数据源配置
type Mysql struct {
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Url             string `mapstructure:"url"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

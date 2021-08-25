package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"time"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, nil
}

type ServerSettings struct {
	RunMode  string
	HttpPort string
}

var config *viper.Viper

func Test() {
	config = WatchEtcd()
	if config == nil {
		return
	}

	r := gin.Default()
	r.GET("/getConfig", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"config": config.AllSettings(),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func WatchEtcd() *viper.Viper {
	// 或者你可以创建一个新的viper实例
	var v = viper.New()

	err := v.AddRemoteProvider("etcd", "http://127.0.0.1:2379", "/config/lua.json")
	if err != nil {
		fmt.Println("链接etcd失败：", err)
		return nil
	}
	v.SetConfigType("yaml")

	// 第一次从远程读取配置
	err = v.ReadRemoteConfig()
	if err != nil {
		fmt.Println("第一次从远程读取配置失败", err)
		return nil
	}

	// 反序列化
	//runtime_viper.Unmarshal(&runtime_conf)

	// 开启一个单独的goroutine一直监控远端的变更
	go func() {
		for {
			time.Sleep(time.Second * 5) // 每次请求后延迟一下

			// 目前只测试了etcd支持
			err := v.WatchRemoteConfig()
			if err != nil {

				continue
			}
			v.OnConfigChange(func(e fsnotify.Event) {
				fmt.Println("文件背该懂了")
			})

			//// 将新配置反序列化到我们运行时的配置结构体中。你还可以借助channel实现一个通知系统更改的信号
			//runtime_viper.Unmarshal(&runtime_conf)
		}
	}()
	return v
}

func initConfigure() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")   // 设置文件名称（无后缀）
	v.SetConfigType("yaml")     // 设置后缀名 {"1.6以后的版本可以不设置该后缀"}
	v.AddConfigPath("version/") // 设置文件所在路径
	v.Set("verbose", true)      // 设置默认参数

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(" Config file not found; ignore error if desired")
		} else {
			panic("Config file was found but another error was produced")
		}
	}
	// 监控配置和重新获取配置
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("文件背该懂了")
	})
	return v
}

func main() {

}

//3,15,10 首位3表示每月，1-28表示1-28号（注意有些月份没有29，30，31号所以不能配置），10表示10点

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var jsonData map[string]interface{}

//获取初始化的JSON文件内容
func initJSON() {

	bytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("ReadFile:", err.Error())
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("Invalid config: ", err.Error())
		os.Exit(-1)
	}

}

//数据库配置结构体
type DBConfig struct {
	Dialect            string
	Database           string
	User               string
	Password           string
	Host               string
	Port               string
	Charset            string
	URL                string
	MaxIdleConnections int
	MaxOpenConnections int
}

//dBconfig 数据库配置对象
var dBconfig DBConfig

//数据库初始化
func initDB() {

}

//Redis配置结构体
type RedisConfig struct {
	Host      string
	Port      string
	Password  string
	URL       string
	MaxIdle   int
	MaxActive int
}

var redisConfig RedisConfig

//redis初始化
func initRedis() {

}

//Mongo配置结构体
type MongoConfig struct {
	URL      string
	Database string
}

var mongoConfig MongoConfig

//mongo初始化
func initMongo() {

}

//服务配置结构体
type ServerConfig struct {
	APIPoweredBy       string
	SiteName           string
	Host               string
	ImgHost            string
	Env                string
	LogDir             string
	LogFile            string
	APIPrefix          string
	UploadImgDir       string
	ImgPath            string
	MaxMultipartMemory int
	Port               int
	StatsEnabled       bool
	TokenSecret        string
	TokenMaxAge        int
	PassSalt           string
	LuosimaoVerifyURL  string
	LuosimaoAPIKey     string
	CrawlerName        string
	MailUser           string //域名邮箱账号
	MailPass           string //域名邮箱密码
	MailHost           string //smtp邮箱域名
	MailPort           int    //smtp邮箱端口
	MailFrom           string //邮件来源
	Github             string
	BaiduPushLink      string
}

var serverConfig ServerConfig

func initServer() {

}

type StatsConfig struct {
	URL    string
	Prefix string
}

var statsConfig StatsConfig

func initStats() {

}

func init() {

}

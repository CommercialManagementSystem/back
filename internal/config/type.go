package config

import (
	"fmt"
)

// HTTPType http 配置信息
type HTTPType struct {
	Port             string `yaml:"Port"`
	MaxContentLength int    `yaml:"MaxContentLength"`
	ShutdownTimeout  int    `yaml:"ShutdownTimeout"`
	MaxLoggerLength  int    `yaml:"MaxLoggerLength"`
}

// CORSType 跨域设置
type CORSType struct {
	Enable           bool     `yaml:"Enable"`
	AllowOrigins     []string `yaml:"AllowOrigins"`
	AllowMethods     []string `yaml:"AllowMethods"`
	AllowHeaders     []string `yaml:"AllowHeaders"`
	AllowCredentials bool     `yaml:"AllowCredentials"`
	MaxAge           int      `yaml:"MaxAge"`
}

// GORMType gorm 配置信息
type GORMType struct {
	Debug             bool `yaml:"Debug"`
	MaxLifetime       int  `yaml:"MaxLifetime"`
	MaxOpenConns      int  `yaml:"MaxOpenConns"`
	MaxIdleConns      int  `yaml:"MaxIdleConns"`
	EnableAutoMigrate bool `yaml:"EnableAutoMigrate"`
}

// DBType 数据库配置定义
type DBType struct {
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	User       string `yaml:"User"`
	Password   string `yaml:"Password"`
	DBName     string `yaml:"DBName"`
	Parameters string `yaml:"Parameters"`
}

// DSN 得到数据库连接
func (d *DBType) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DBName,
		d.Parameters,
	)
}

func (d *DBType) PgDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		d.Host, d.User, d.Password, d.DBName, d.Port)
}

// LogType 日志配置类型定义
type LogType struct {
	Level  int8   `yaml:"Level"`
	Output string `yaml:"Output"`
}

// LogFileHookType 文件归档钩子配置
type LogFileHookType struct {
	Filename   string `yaml:"Filename"`
	MaxSize    int    `yaml:"Maxsize"`
	MaxBackups int    `yaml:"MaxBackups"`
	MaxAge     int    `yaml:"Maxage"`
	Compress   bool   `yaml:"Compress"`
}

// JWTType 文件归档钩子配置
type JWTType struct {
	Enable  bool   `yaml:"Enable"`
	Secret  string `yaml:"Secret"`
	Expires int    `yaml:"Expires"`
	Issuer  string `yaml:"Issuer"`
}

type OSSType struct {
	CRC             bool   `yaml:"CRC"`
	Endpoint        string `yaml:"Endpoint"`
	AccessKeyId     string `yaml:"AccessKeyId"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
}

type OCRType struct {
	GrantType    string `yaml:"GrantType"`
	ClientID     string `yaml:"ClientID"`
	ClientSecret string `yaml:"ClientSecret"`
}

func (o *OCRType) TokenUrl() string {
	return fmt.Sprintf(
		"https://aip.baidubce.com/oauth/2.0/token?grant_type=%s&client_id=%s&client_secret=%s",
		o.GrantType,
		o.ClientID,
		o.ClientSecret,
	)
}

// CType 配置文件类型定义
type CType struct {
	Mode        string          `yaml:"Mode"`
	HTTP        HTTPType        `yaml:"HTTP"`
	CORS        CORSType        `yaml:"CORS"`
	GORM        GORMType        `yaml:"GORM"`
	DB          DBType          `yaml:"DB"`
	Log         LogType         `yaml:"Log"`
	LogFileHook LogFileHookType `yaml:"LogFileHook"`
	JWT         JWTType         `yaml:"JWT"`
	OSS         OSSType         `yaml:"OSS"`
	OCR         OCRType         `yaml:"OCR"`
}

// Package config  uses Viper to load all configuration from the yml config file.
package config

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
)

// Configuration represents the composition of yml settings.
// Viper has a built-in mapstructure, and the yml file uses "-" to distinguish words.
type Configuration struct {
	System   SystemConfiguration   `mapstructure:"system" json:"system"`
	Logs     LogsConfiguration     `mapstructure:"logs" json:"logs"`
	Database DatabaseConfiguration `mapstructure:"database" json:"database"`
	Auth     AuthConfiguration     `mapstructure:"auth" json:"auth"`
	Consul   ConsulConfiguration   `mapstructure:"consul" json:"consul"`
}

type SystemConfiguration struct {
	ApiVersion  string `mapstructure:"api-version" json:"apiVersion"`
	WaitTimeout int    `mapstructure:"wait-timeout" json:"waitTimeout"`
}

type LogsConfiguration struct {
	Encoding   string        `mapstructure:"encoding" json:"encoding"`
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}

type DatabaseConfiguration struct {
	Dialect     string `mapstructure:"dialect" json:"dialect"`
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	Dbname      string `mapstructure:"dbname" json:"dbname"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Query       string `mapstructure:"query" json:"query"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode"`
	TablePrefix string `mapstructure:"table-prefix" json:"tablePrefix"`
	Charset     string `mapstructure:"charset" json:"charset"`
	Collation   string `mapstructure:"collation" json:"collation"`
	Migration   bool   `mapstructure:"migration" json:"migration"`
	InitData    bool   `mapstructure:"init-data" json:"initData"`
}

type AuthConfiguration struct {
	JwtKey    string `mapstructure:"jwt-key" json:"jwtKey"`
	Expires   uint   `mapstructure:"expires" json:"expires"`
	RuleTable string `mapstructure:"rule-table" json:"ruleTable"`
}

type ConsulConfiguration struct {
	Address string `mapstructure:"address" json:"address"`
	Port    int    `mapstructure:"port" json:"port"`
}

//go:embed config.dev.yml
var defaultConfigFile embed.FS

// LoadAppConfig reads the settings written to the yml file
func LoadAppConfig(filename string) (*Configuration, error) {
	v := viper.New()
	// load default config file
	file, _ := defaultConfigFile.ReadFile(developmentConfig)
	if err := readConfig(v, file); err != nil {
		fmt.Printf("Failed to read %s: %v", filename, err)
		return nil, err
	}
	settings := v.AllSettings()
	for index, s := range settings {
		v.SetDefault(index, s)
	}

	// Read file
	if filename == "" {
		fmt.Println("No configuration file provided, default dev config will load.")
	} else {
		if strings.HasPrefix(filename, "/") {
			// Get readFile absolute path
			pwd, _ := os.Getwd()
			filename = filepath.Join(pwd, filename)
		}
		readFile, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("%s path not exists: %v", filename, err)
			return nil, err
		}
		err = readConfig(v, readFile)
		if err != nil {
			fmt.Printf("Failed to read %s: %v", filename, err)
			return nil, err
		}
	}

	// Convert config struct
	config := &Configuration{}
	if err := v.Unmarshal(config); err != nil {
		fmt.Printf("Failed to load config: %v", err)
		return nil, err
	}

	return config, nil
}

// readConfig reads settings by Viper
func readConfig(v *viper.Viper, b []byte) error {
	v.SetConfigType(configType)
	err := v.ReadConfig(bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("load config file failed: %v", err)
	}
	return nil
}

package configs

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

// region structure
// App config struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Logger   Logger
	Email    Email
}

// Server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	PprofPort         string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
	Retries           int
	RetryTime         time.Duration
	Debug             bool
	IsIIS             bool
	SSL               bool
}

// Postgresql config
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Schema   string
	SSLMode  string
	PgDriver string
}

// Redis config
type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultdb string
	MinIdleConns   int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
	CacheTTL       time.Duration
}

// Logger config
type Logger struct {
	Development   bool
	DisableCaller bool
	Encoding      string
	Level         string
	EnableFile    bool
	InfoFileName  string
	MaxBackups    int
	MaxAge        int
	MaxSize       int
}

// Email config
type Email struct {
	Account  string
	Password string
}

// endregion

// region public methods
// Load config file from given path
// @param file name
// @result viper and error
func LoadConfig(filename string) (*viper.Viper, error) {
	instance := viper.New()
	instance.SetConfigName(filename)
	instance.AddConfigPath(".") // search setting in the working directory
	instance.AutomaticEnv()     // read in environment variables that match
	if err := instance.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return instance, nil
}

// Parse config file
// @param viper
func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg) // 反序列化
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}

// endregion

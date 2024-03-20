package config

import (
	"flag"
	"os"
)

type Config struct {
	ServerAddress  string
	DatabaseDSN    string
	AccrualSysAddr string
	LogLevel       string
	ServiceName    string
}

const (
	defaultServiceName = "gophermart"
	defaultLogLevel    = "info"
)

func Parse(c *Config) {
	serverAddress := flag.String("a", "", "адрес и порт запуска сервиса")
	dsn := flag.String("d", "", "адрес подключения к базе данных")
	accrualSysAddr := flag.String("r", "", "адрес системы расчёта начислений")

	flag.Parse()

	c.ServerAddress = getEnvString("RUN_ADDRESS", serverAddress)
	c.DatabaseDSN = getEnvString("DATABASE_URI", dsn)
	c.AccrualSysAddr = getEnvString("ACCRUAL_SYSTEM_ADDRESS", accrualSysAddr)

	c.LogLevel = defaultLogLevel
	c.ServiceName = defaultServiceName
}

func getEnvString(key string, argumentValue *string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return *argumentValue
}

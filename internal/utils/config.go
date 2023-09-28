package utils

import (
	"os"
)

// type Conf models.DBConfig
var DefaultDBFile string

func GetConfig() string {
	sqlFile, exists := os.LookupEnv("DB_FILE")
	if !exists {
		return DefaultDBFile
	}
	if DBFileExists(sqlFile) {
		return sqlFile
	}
	return DefaultDBFile
}

func DBFileExists(f string) bool {
	if _, err := os.Stat(f); err == nil {
		return true
	}
	return false
}

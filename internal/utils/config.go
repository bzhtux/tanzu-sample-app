package utils

import (
	"fmt"
	"os"
)

// type Conf models.DBConfig
var DefaultDBFile string

func GetConfig() string {
	fmt.Printf("DefaultDBFile: %s\n", DefaultDBFile)
	sqlFile, exists := os.LookupEnv("DB_FILE")
	if !exists {
		if !DBFileExists(DefaultDBFile) {
			DBFileCreate(DefaultDBFile)
		}
		return DefaultDBFile
	}
	if DBFileExists(sqlFile) {
		return sqlFile
	} else {
		DBFileCreate(sqlFile)
	}
	if !DBFileExists(DefaultDBFile) {
		DBFileCreate(DefaultDBFile)
	}
	return DefaultDBFile
}

func DBFileExists(f string) bool {
	if _, err := os.Stat(f); err == nil {
		return true
	}
	return false
}

func DBFileCreate(f string) bool {
	if _, err := os.Create(f); err == nil {
		return false
	}
	return true
}

func GetProjectDir() string {
	path, err := os.Getwd()
	if err != nil {
		return "nil"
	}
	return path
}

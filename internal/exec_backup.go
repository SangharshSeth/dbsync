package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ExecuteBackupCommand(config BackupConfig) {
	var bin string
	var hostName string
	var file *os.File

	hostName = strings.Split(config.Address, ":")[0]

	if hostName == "localhost" {
		hostName = "127.0.0.1"
	}
	//Make a backup directory if not exists
	err := os.MkdirAll("backups", os.ModePerm)
	if err != nil {
		panic(err)
	}

	if config.Output == "" {
		backupFileName := fmt.Sprintf("%s-%s-%s.sql", config.Type, time.Now().Format("20060102-150405"), config.Database)
		file, err = os.Create("backups/" + backupFileName)
		if err != nil {
			panic(err)
		}
	} else {
		file, err = os.Create("backups/" + fmt.Sprintf("%s.sql", config.Output))
		if err != nil {
			panic(err)
		}
	}

	switch config.Type {
	case "mysql":
		bin = "mysqldump"
		cmd := exec.Command(bin, "-u", config.Username, "-p"+config.Password, "-h", hostName, config.Database)
		cmd.Stdout = file
		err := cmd.Run()
		if err != nil {
			panic(err)
		}

	case "postgres":
		bin = "pg_dump"

	case "mongodb":
		bin = "mongodump"
	}
}

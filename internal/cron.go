package internal

import "github.com/robfig/cron/v3"

func ScheduleBackup(schedule string, config BackupConfig) {
	c := cron.New()

	c.AddFunc(schedule, func() {
		ExecuteBackupCommand(config)
	})
	c.Run()

}

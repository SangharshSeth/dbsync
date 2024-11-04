package cmd

import (
	"dbsync/internal"
	"fmt"
	"github.com/spf13/cobra"
	"log/slog"
	"time"
)

type Address struct {
	Country string `json:"country"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

var logger *slog.Logger
var backupConfig internal.BackupConfig

func init() {
	backupConfig = internal.BackupConfig{}
	backupCommand.Flags().StringVarP(
		(*string)(&backupConfig.Type),
		"type",
		"t",
		"postgres",
		"Database dialect(mysql, postgres, mongodb)")
	backupCommand.Flags().StringVarP(
		&backupConfig.Address,
		"address",
		"a",
		"hostname:port",
		"Database address with hostname and port")
	backupCommand.Flags().StringVarP(
		&backupConfig.Username,
		"username",
		"u",
		"",
		"Database username")
	backupCommand.Flags().StringVarP(
		&backupConfig.Password,
		"password",
		"p",
		"",
		"Database password",
	)
	backupCommand.Flags().StringVarP(
		&backupConfig.Database,
		"database",
		"d",
		"",
		"Database name")
	backupCommand.Flags().StringVarP(
		&backupConfig.Output,
		"output",
		"o",
		"",
		"Output file name",
	)

	requiredFlags := []string{"type", "address", "username", "password", "database"}
	for _, flag := range requiredFlags {
		err := backupCommand.MarkFlagRequired(flag)
		if err != nil {
			logger.Error("Flag Required", err, err.Error())
		}
	}
}

var backupCommand = &cobra.Command{
	Use:   "backup",
	Short: "Backup a database",
	Long:  "Backup a database using a specific backup strategy",
	RunE:  BackupCommandRunner,
}

func BackupCommandRunner(cmd *cobra.Command, args []string) error {
	logger := internal.InitLogger()
	// Build connection string
	dsn := internal.FormatDSN(backupConfig)

	//Connect to database
	databaseInstance, err := internal.NewDatabaseConnection(dsn, internal.Dialect(backupConfig.Type))
	if err != nil {
		logger.Error("Database Initialization ", "error", err)
		return fmt.Errorf("database Initialization Failed %s", err)
	}

	connected, err := databaseInstance.Connect()
	if err != nil || !connected {
		logger.Error("Database Connection Failed ", err, err)
		return fmt.Errorf("database Connection Failed %s", err)

	}
	startTime := time.Now()
	internal.ExecuteBackupCommand(backupConfig)
	backupTime := time.Since(startTime)
	logger.Info("Backup Completed in", "time", backupTime)
	return nil
}

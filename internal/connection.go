package internal

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DBConnection interface {
	Connect() (bool, error)
	ConnectNoSQL() (bool, error)
}

type DatabaseConnection struct {
	DSN     string
	Dialect Dialect
}

func (conn *DatabaseConnection) Connect() (bool, error) {

	logger := InitLogger()
	// Validate dialect
	if conn.Dialect != MySQL && conn.Dialect != Postgres {
		return false, fmt.Errorf("unsupported dialect: %s", conn.Dialect)
	}
	db, err := sql.Open(string(conn.Dialect), conn.DSN)
	if err != nil {
		logger.Error("Error connecting to database:", "error", err.Error(), "dsn", conn.DSN, "dialect", conn.Dialect)
		return false, fmt.Errorf("error opening database connection: %w", err)
	}
	defer db.Close() // Close connection at the end of function execution

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	err = db.Ping()
	if err != nil {
		return false, fmt.Errorf("error pinging database: %w", err)
	}

	return true, nil
}
func (conn *DatabaseConnection) ConnectNoSQL() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(conn.DSN)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return false, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return false, fmt.Errorf("error pinging MongoDB: %w", err)
	}

	return true, nil
}

func NewDatabaseConnection(dsn string, dialect Dialect) (*DatabaseConnection, error) {
	return &DatabaseConnection{DSN: dsn, Dialect: dialect}, nil
}

func FormatDSN(config BackupConfig) string {
	logger := InitLogger()
	fmt.Println(config.Type)
	switch config.Type {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Username,
			config.Password,
			config.Address,
			config.Database)

	case "postgres":
		return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
			config.Username,
			config.Password,
			config.Address,
			config.Database)

	case "mongodb":
		return fmt.Sprintf("mongodb://%s:%s@%s/%s",
			config.Username,
			config.Password,
			config.Address,
			config.Database)

	default:
		// Handle unsupported dialects
		logger.Error("unsupported database dialect", "dialect", config.Type)
		return ""
	}
}

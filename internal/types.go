package internal

type Dialect string

const (
	MySQL    Dialect = "mysql"
	Postgres Dialect = "postgres"
	MongoDB  Dialect = "mongodb"
)

type BackupConfig struct {
	Type     Dialect
	Address  string
	Username string
	Password string
	Database string
	Output   string
}

package mysqlhelper

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// MySQLConfig holds the configuration for MySQL connection
type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	// Connection pool settings
	MaxOpenConns    int // Max open connections
	MaxIdleConns    int // Max idle connections
	ConnMaxLifetime int // Connection max lifetime in seconds
}

// GetEnvWithDefault reads an environment variable and returns its value.
// If the variable is empty, it returns the default value.
func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// ConnectToMySQL establishes a connection to MySQL with connection pooling.
func GetDB() *sql.DB {
	// Read connection details from environment variables or use default values
	username := GetEnvWithDefault("MYSQL_USER", "root")
	password := GetEnvWithDefault("MYSQL_PASSWORD", "password")
	host := GetEnvWithDefault("MYSQL_HOST", "localhost")
	port := GetEnvWithDefault("MYSQL_PORT", "3306")
	database := GetEnvWithDefault("MYSQL_DATABASE", "mydb")
	maxOpenConns, _ := strconv.Atoi(GetEnvWithDefault("MYSQL_MAX_OPEN_CONNS", "25"))
	maxIdleConns, _ := strconv.Atoi(GetEnvWithDefault("MYSQL_MAX_IDLE_CONNS", "25"))
	connMaxLifetime, _ := strconv.Atoi(GetEnvWithDefault("MYSQL_CONN_MAX_LIFETIME", "300")) // In seconds

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}

	// Set database connection pool parameters
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	// Verify the connection is successful by pinging the database
	if err := db.Ping(); err != nil {
		log.Fatalf("Could not ping database: %v", err)
	}

	fmt.Println("Successfully connected to MySQL")
	return db
}

// ConnectToMySQL establishes a connection to MySQL database using given configuration
func GetDBWithConfig(config MySQLConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.User, config.Password, config.Host, config.Port, config.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)

	// Try to establish a connection to the database
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MySQL")
	return db, nil
}

package postgreshelper

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgresConfig holds the configuration for the PostgreSQL connection
type PostgresConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	MaxConnLifetime time.Duration // Maximum connection lifetime
	MaxConns        int32         // Maximum number of connections
	MinConns        int32         // Minimum number of connections
}

// NewPostgresConfig creates and returns a new PostgresConfig with either environment variables or default values
func NewPostgresConfig() PostgresConfig {
	port, _ := strconv.Atoi(getEnvOrDefault("PG_PORT", "5432"))
	maxConns, _ := strconv.Atoi(getEnvOrDefault("PG_MAX_CONNS", "25"))
	minConns, _ := strconv.Atoi(getEnvOrDefault("PG_MIN_CONNS", "5"))
	maxConnLifetime, _ := strconv.Atoi(getEnvOrDefault("PG_MAX_CONN_LIFETIME", "300"))

	return PostgresConfig{
		Host:            getEnvOrDefault("PG_HOST", "localhost"),
		Port:            port,
		User:            getEnvOrDefault("PG_USER", "postgres"),
		Password:        getEnvOrDefault("PG_PASSWORD", ""),
		Database:        getEnvOrDefault("PG_DATABASE", "postgres"),
		MaxConnLifetime: time.Duration(maxConnLifetime) * time.Second,
		MaxConns:        int32(maxConns),
		MinConns:        int32(minConns),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetDB establishes a connection to PostgreSQL using environment variables or default configuration
func GetDB() (*pgxpool.Pool, error) {
	config := NewPostgresConfig()
	return GetDBWithConfig(config)
}

// GetDBWithConfig establishes a connection to PostgreSQL using the provided configuration
func GetDBWithConfig(config PostgresConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?pool_max_conns=%d&pool_min_conns=%d&pool_max_conn_lifetime=%s",
		config.User, config.Password, config.Host, config.Port, config.Database,
		config.MaxConns, config.MinConns, config.MaxConnLifetime)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pool config: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL")
	return pool, nil
}

package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	"github.com/cardenasrjl/emem/pkg/config"

	// mysql driver

	_ "github.com/go-sql-driver/mysql"

	"github.com/cardenasrjl/emem/pkg/logger"
	"github.com/cardenasrjl/emem/pkg/protocol/grpc"
	"github.com/cardenasrjl/emem/pkg/protocol/rest"
	v1 "github.com/cardenasrjl/emem/pkg/service/v1"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string

	//redis
	RedisPort string
	RedisHost string

	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	envConfig, err := config.New()
	if err != nil {
		return err
	}

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", envConfig.GrpcPort, "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", envConfig.HttpPort, "HTTP port to bind")

	flag.StringVar(&cfg.DatastoreDBHost, "db-host", envConfig.DatabaseHost, "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", envConfig.DatabaseUser, "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", envConfig.DatabasePass, "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", envConfig.DatabaseSchema, "Database schema")

	flag.IntVar(&cfg.LogLevel, "log-level", envConfig.LogLevel, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "2006-01-02T15:04:05Z07:00",
		"Print time format for logger e.g. 2006-01-02T15:04:05Z07:00")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	//mems service
	v1API := v1.NewMemService(db)
	//volume service
	v1VAPI := v1.NeVolumeService(db)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, v1VAPI, cfg.GRPCPort)

}

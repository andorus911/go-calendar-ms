package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/andorus911/go-calendar-ms/api-ms/api"
	"github.com/andorus911/go-calendar-ms/api-ms/logger"
	"github.com/andorus911/go-calendar-ms/api-ms/tools/domain/services"
	"github.com/andorus911/go-calendar-ms/api-ms/tools/postgres"
	"github.com/spf13/viper"
	"log"
)

type config struct {
	HttpListen string `mapstructure:"http_listen"`
	LogFile    string `mapstructure:"log_file"`
	LogLevel   string `mapstructure:"log_level"`

	SqlUser     string `mapstructure:"sql_user"`
	SqlPassword string `mapstructure:"sql_pass"`
	SqlHost     string `mapstructure:"sql_host"`
	SqlPort     string `mapstructure:"sql_port"`
	DbName      string `mapstructure:"db_name"`
}

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "", "path to the config")
	flag.Parse()

	if configFilePath == "" {
		log.Fatal("Config is not presented")
	}

	viper.SetConfigFile(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	cfg := &config{}

	err = viper.Unmarshal(cfg)
	if err != nil {
		fmt.Printf("Unable to decode into config struct, %v", err)
	}

	// logger init
	lg := logger.NewLogger(cfg.LogFile, cfg.LogLevel)
	defer lg.Sync()

	// db init
	db, err := postgres.InitDB(context.Background(), lg, cfg.SqlUser, cfg.SqlPassword, cfg.SqlHost, cfg.SqlPort, cfg.DbName)
	if err != nil {
		lg.Fatal(err.Error())
	}
	defer func() {
		if err := postgres.CloseDBCxn(); err != nil {
			lg.Error(err.Error())
		}
	}()

	eventService := services.EventService{EventStorage: db}

	api.StartServer(cfg.HttpListen, *lg, &eventService)
}

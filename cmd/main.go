package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	sneakerq "github.com/ukurysheva/sneaker-q"
	"github.com/ukurysheva/sneaker-q/internal/cron"
	"github.com/ukurysheva/sneaker-q/internal/parser"
	"github.com/ukurysheva/sneaker-q/pkg/handler"
	"github.com/ukurysheva/sneaker-q/pkg/repository"
	"github.com/ukurysheva/sneaker-q/pkg/services"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializaing configs: %s", err.Error())
		fmt.Println(err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env cars: %s", err.Error())
		fmt.Println(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error initializaing db: %s", err.Error())
		fmt.Println(err)
	}
	repo := repository.NewRepository(db)

	parserTask := parser.NewParserTask(repo)
	// implimenting cronjob
	_, err = cron.RunCron(parserTask.ParseTask)
	if err != nil {
		return
	}

	services := services.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(sneakerq.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	return
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

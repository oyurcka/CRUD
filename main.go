package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	_ "github.com/oyurcka/CRUD/model/app"
	_personLogic "github.com/oyurcka/CRUD/person/logic"
	_personRepository "github.com/oyurcka/CRUD/person/repository"
	"github.com/sirupsen/logrus"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "admin"
	dbname   = "person"
	timeout  = 20
	adress   = "localhost:9080"
)

func main() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname)
	conn, err := dbr.Open("postgres", dsn, nil)
	if err != nil {
		logrus.Error(err)
	}
	err = conn.Ping()
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()

	e := echo.New()
	perRepo := _personRepository.NewPostgresqlPersonRepository(conn.DB)

	timeoutCtx := time.Duration(timeout) * time.Second
	_personLogic.NewPersonLogic(perRepo, timeoutCtx)

	logrus.Fatal(e.Start(adress))

}

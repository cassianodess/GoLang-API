package pkg

import (
	"database/sql"
	conf "main/config"

	"fmt"

	_ "github.com/lib/pq"
)

var psqlConnection string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.LoadDotEnv("DB_HOST"), conf.LoadDotEnv("DB_PORT"), conf.LoadDotEnv("DB_USER"), conf.LoadDotEnv("DB_PASSWORD"), conf.LoadDotEnv("DB_NAME"))

var Db, Err = sql.Open("postgres", psqlConnection)

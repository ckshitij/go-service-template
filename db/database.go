package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.io/ckshitij/go-service-template/config"
)

func NewPostgresDB(conf *config.Config) (*sql.DB, error) {
	sqlConf := mysql.Config{
		User:   conf.Databases["users"].User,
		Passwd: conf.Databases["users"].Password,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%d", conf.Databases["users"].Host, conf.Databases["users"].Port),
		DBName: conf.Databases["users"].Database,
	}

	db, err := sql.Open("postgres", sqlConf.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}

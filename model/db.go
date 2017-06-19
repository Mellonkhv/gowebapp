package model

import (
	"database/sql"

	"github.com/mellonkhv/gowebapp/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	ConnectString string
}

func InitDb(cfg Config) (*pgDb, error)  {
	if dbConn, err := sqlx.Connect("postgres", cfg.ConnectString); err != nil{
		return nil, err
	} else {
		p := &pgDb{dbConn: dbConn}
		if err := p.dbConn.Ping(); err != nill{
			return nil, err
		}
	}
}

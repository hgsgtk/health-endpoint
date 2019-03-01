package api

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func NewMySQL(conf DBConfig) (*sql.DB, error) {
	ds := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name,
	)
	con, err := sql.Open("mysql", ds)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get connection to database")
	}
	if err := con.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to ping database")
	}

	return con, nil
}

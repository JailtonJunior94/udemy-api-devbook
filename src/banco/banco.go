package banco

import (
	"database/sql"

	"github.com/JailtonJunior94/udemy-api-devbook/src/config"
	_ "github.com/denisenkom/go-mssqldb"
)

func Conectar() (*sql.DB, error) {
	db, err := sql.Open("sqlserver", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

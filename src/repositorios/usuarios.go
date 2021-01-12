package repositorios

import (
	"database/sql"
	"time"

	"github.com/JailtonJunior94/udemy-api-devbook/src/modelos"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	criadoEm := time.Now()
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO Usuarios ([Name], [Nick], [Email], [Senha], [CriadoEm]) VALUES (@nome, @nick, @email, @senha, @criadoEm); SELECT SCOPE_IDENTITY()")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	lastInsertId := 0
	erro = statement.QueryRow(sql.Named("nome", usuario.Nome),
		sql.Named("nick", usuario.Nick),
		sql.Named("email", usuario.Email),
		sql.Named("senha", usuario.Senha),
		sql.Named("criadoEm", criadoEm)).Scan(&lastInsertId)

	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertId), nil
}

package database

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

type sqliteRepo struct {
	*Queries
}

func NewSqliteRepo(dbPath string) (Repository, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("PRAGMA journal_mode = WAL;")
	if err != nil {
		return nil, err
	}

	return &sqliteRepo{
		Queries: New(db),
	}, nil
}

func (r *sqliteRepo) InsertUser(ctx context.Context, user User) error {
	return r.insertUser(ctx, insertUserParams(user))
}

func (r *sqliteRepo) UpdateUser(ctx context.Context, username string, user User) error {
	params := updateUserParams{
		Username:   user.Username,
		Password:   user.Password,
		Salt:       user.Salt,
		Username_2: username,
	}
	return r.updateUser(ctx, params)
}

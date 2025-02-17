package data

import (
    "errors"
    "database/sql"
)

var (
    ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
    Movies MovieModel
}

func NewModels(db *sql.DB) Models {
    return Models {
        Movies: MovieModel{DB: db},
    }
}

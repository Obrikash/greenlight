package main

import (
    "errors"
    "net/http"
    "strconv"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
    param := r.PathValue("id")

    id, err := strconv.ParseInt(param, 10, 64)
    if err != nil || id < 1{
        return 0, errors.New("invalid id parameter")
    }

    return id, nil
}

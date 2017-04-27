// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

// Crypt calls the stored procedure 'public.crypt(text, text) text' on db.
func Crypt(db XODB, v0 string, v1 string) (string, error) {
	var err error

	// sql query
	const sqlstr = `SELECT public.crypt($1, $2)`

	// run query
	var ret string
	XOLog(sqlstr, v0, v1)
	err = db.QueryRow(sqlstr, v0, v1).Scan(&ret)
	if err != nil {
		return "", err
	}

	return ret, nil
}

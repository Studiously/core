// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

// Armor calls the stored procedure 'public.armor(bytea, bytea, text[], text[]) text' on db.
func Armor(db XODB, v0 []byte, v1 []byte, v2 StringSlice, v3 StringSlice) (string, error) {
	var err error

	// sql query
	const sqlstr = `SELECT public.armor($1, $2, $3, $4)`

	// run query
	var ret string
	XOLog(sqlstr, v0, v1, v2, v3)
	err = db.QueryRow(sqlstr, v0, v1, v2, v3).Scan(&ret)
	if err != nil {
		return "", err
	}

	return ret, nil
}

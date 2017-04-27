// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

// PgpSymDecrypt calls the stored procedure 'public.pgp_sym_decrypt(bytea, text, bytea, text, text) text' on db.
func PgpSymDecrypt(db XODB, v0 []byte, v1 string, v2 []byte, v3 string, v4 string) (string, error) {
	var err error

	// sql query
	const sqlstr = `SELECT public.pgp_sym_decrypt($1, $2, $3, $4, $5)`

	// run query
	var ret string
	XOLog(sqlstr, v0, v1, v2, v3, v4)
	err = db.QueryRow(sqlstr, v0, v1, v2, v3, v4).Scan(&ret)
	if err != nil {
		return "", err
	}

	return ret, nil
}

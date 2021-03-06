// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

// GenerateID calls the stored procedure 'public.generate_id(integer) bigint' on db.
func GenerateID(db XODB, v0 int) (int64, error) {
	var err error

	// sql query
	const sqlstr = `SELECT public.generate_id($1)`

	// run query
	var ret int64
	XOLog(sqlstr, v0)
	err = db.QueryRow(sqlstr, v0).Scan(&ret)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

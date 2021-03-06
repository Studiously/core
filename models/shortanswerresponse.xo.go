// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// ShortAnswerResponse represents a row from 'public.short_answer_responses'.
type ShortAnswerResponse struct {
	ID         int64          `json:"id"`          // id
	QuestionID sql.NullInt64  `json:"question_id"` // question_id
	AuthorID   int64          `json:"author_id"`   // author_id
	Response   sql.NullString `json:"response"`    // response

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ShortAnswerResponse exists in the database.
func (sar *ShortAnswerResponse) Exists() bool {
	return sar._exists
}

// Deleted provides information if the ShortAnswerResponse has been deleted from the database.
func (sar *ShortAnswerResponse) Deleted() bool {
	return sar._deleted
}

// Insert inserts the ShortAnswerResponse to the database.
func (sar *ShortAnswerResponse) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if sar._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.short_answer_responses (` +
		`id, question_id, author_id, response` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`

	// run query
	XOLog(sqlstr, sar.ID, sar.QuestionID, sar.AuthorID, sar.Response)
	err = db.QueryRow(sqlstr, sar.ID, sar.QuestionID, sar.AuthorID, sar.Response).Scan(&sar.ID)
	if err != nil {
		return err
	}

	// set existence
	sar._exists = true

	return nil
}

// Update updates the ShortAnswerResponse in the database.
func (sar *ShortAnswerResponse) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !sar._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if sar._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.short_answer_responses SET (` +
		`question_id, author_id, response` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, sar.QuestionID, sar.AuthorID, sar.Response, sar.ID)
	_, err = db.Exec(sqlstr, sar.QuestionID, sar.AuthorID, sar.Response, sar.ID)
	return err
}

// Save saves the ShortAnswerResponse to the database.
func (sar *ShortAnswerResponse) Save(db XODB) error {
	if sar.Exists() {
		return sar.Update(db)
	}

	return sar.Insert(db)
}

// Upsert performs an upsert for ShortAnswerResponse.
//
// NOTE: PostgreSQL 9.5+ only
func (sar *ShortAnswerResponse) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if sar._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.short_answer_responses (` +
		`id, question_id, author_id, response` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, question_id, author_id, response` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.question_id, EXCLUDED.author_id, EXCLUDED.response` +
		`)`

	// run query
	XOLog(sqlstr, sar.ID, sar.QuestionID, sar.AuthorID, sar.Response)
	_, err = db.Exec(sqlstr, sar.ID, sar.QuestionID, sar.AuthorID, sar.Response)
	if err != nil {
		return err
	}

	// set existence
	sar._exists = true

	return nil
}

// Delete deletes the ShortAnswerResponse from the database.
func (sar *ShortAnswerResponse) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !sar._exists {
		return nil
	}

	// if deleted, bail
	if sar._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.short_answer_responses WHERE id = $1`

	// run query
	XOLog(sqlstr, sar.ID)
	_, err = db.Exec(sqlstr, sar.ID)
	if err != nil {
		return err
	}

	// set deleted
	sar._deleted = true

	return nil
}

// Profile returns the Profile associated with the ShortAnswerResponse's AuthorID (author_id).
//
// Generated from foreign key 'short_answer_author_fkey'.
func (sar *ShortAnswerResponse) Profile(db XODB) (*Profile, error) {
	return ProfileByID(db, sar.AuthorID)
}

// ShortAnswerQuestion returns the ShortAnswerQuestion associated with the ShortAnswerResponse's QuestionID (question_id).
//
// Generated from foreign key 'short_answer_question_fkey'.
func (sar *ShortAnswerResponse) ShortAnswerQuestion(db XODB) (*ShortAnswerQuestion, error) {
	return ShortAnswerQuestionByID(db, sar.QuestionID.Int64)
}

// ShortAnswerResponseByID retrieves a row from 'public.short_answer_responses' as a ShortAnswerResponse.
//
// Generated from index 'short_answer_responses_pkey'.
func ShortAnswerResponseByID(db XODB, id int64) (*ShortAnswerResponse, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, question_id, author_id, response ` +
		`FROM public.short_answer_responses ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	sar := ShortAnswerResponse{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&sar.ID, &sar.QuestionID, &sar.AuthorID, &sar.Response)
	if err != nil {
		return nil, err
	}

	return &sar, nil
}

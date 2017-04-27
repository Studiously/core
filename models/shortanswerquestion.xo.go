// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

// ShortAnswerQuestion represents a row from 'public.short_answer_questions'.
type ShortAnswerQuestion struct {
	ID            uuid.UUID     `json:"id"`             // id
	ClassID       uuid.UUID     `json:"class_id"`       // class_id
	AuthorID      uuid.UUID     `json:"author_id"`      // author_id
	Prompt        string        `json:"prompt"`         // prompt
	MinimumLength sql.NullInt64 `json:"minimum_length"` // minimum_length
	MaximumLength sql.NullInt64 `json:"maximum_length"` // maximum_length
	UnitID        uuid.UUID     `json:"unit_id"`        // unit_id
	Timestamp     time.Time     `json:"timestamp"`      // timestamp

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ShortAnswerQuestion exists in the database.
func (saq *ShortAnswerQuestion) Exists() bool {
	return saq._exists
}

// Deleted provides information if the ShortAnswerQuestion has been deleted from the database.
func (saq *ShortAnswerQuestion) Deleted() bool {
	return saq._deleted
}

// Insert inserts the ShortAnswerQuestion to the database.
func (saq *ShortAnswerQuestion) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if saq._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.short_answer_questions (` +
		`id, class_id, author_id, prompt, minimum_length, maximum_length, unit_id, timestamp` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`)`

	// run query
	XOLog(sqlstr, saq.ID, saq.ClassID, saq.AuthorID, saq.Prompt, saq.MinimumLength, saq.MaximumLength, saq.UnitID, saq.Timestamp)
	err = db.QueryRow(sqlstr, saq.ID, saq.ClassID, saq.AuthorID, saq.Prompt, saq.MinimumLength, saq.MaximumLength, saq.UnitID, saq.Timestamp).Scan(&saq.ID)
	if err != nil {
		return err
	}

	// set existence
	saq._exists = true

	return nil
}

// Update updates the ShortAnswerQuestion in the database.
func (saq *ShortAnswerQuestion) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !saq._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if saq._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.short_answer_questions SET (` +
		`class_id, author_id, prompt, minimum_length, maximum_length, unit_id, timestamp` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) WHERE id = $8`

	// run query
	XOLog(sqlstr, saq.ClassID, saq.AuthorID, saq.Prompt, saq.MinimumLength, saq.MaximumLength, saq.UnitID, saq.Timestamp, saq.ID)
	_, err = db.Exec(sqlstr, saq.ClassID, saq.AuthorID, saq.Prompt, saq.MinimumLength, saq.MaximumLength, saq.UnitID, saq.Timestamp, saq.ID)
	return err
}

// Save saves the ShortAnswerQuestion to the database.
func (saq *ShortAnswerQuestion) Save(db XODB) error {
	if saq.Exists() {
		return saq.Update(db)
	}

	return saq.Insert(db)
}

// Upsert performs an upsert for ShortAnswerQuestion.
//
// NOTE: PostgreSQL 9.5+ only
func (saq *ShortAnswerQuestion) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if saq._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.short_answer_questions (` +
		`id, class_id, author_id, prompt, minimum_length, maximum_length, unit_id, timestamp` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, class_id, author_id, prompt, minimum_length, maximum_length, unit_id, timestamp` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.class_id, EXCLUDED.author_id, EXCLUDED.prompt, EXCLUDED.minimum_length, EXCLUDED.maximum_length, EXCLUDED.unit_id, EXCLUDED.timestamp` +
		`)`

	// run query
	XOLog(sqlstr, saq.ID, saq.ClassID, saq.AuthorID, saq.Prompt, saq.MinimumLength, saq.MaximumLength, saq.UnitID, saq.Timestamp)
	_, err = db.Exec(sqlstr, saq.ID, saq.ClassID, saq.AuthorID, saq.Prompt, saq.MinimumLength, saq.MaximumLength, saq.UnitID, saq.Timestamp)
	if err != nil {
		return err
	}

	// set existence
	saq._exists = true

	return nil
}

// Delete deletes the ShortAnswerQuestion from the database.
func (saq *ShortAnswerQuestion) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !saq._exists {
		return nil
	}

	// if deleted, bail
	if saq._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.short_answer_questions WHERE id = $1`

	// run query
	XOLog(sqlstr, saq.ID)
	_, err = db.Exec(sqlstr, saq.ID)
	if err != nil {
		return err
	}

	// set deleted
	saq._deleted = true

	return nil
}

// Profile returns the Profile associated with the ShortAnswerQuestion's AuthorID (author_id).
//
// Generated from foreign key 'short_answer_author_fkey'.
func (saq *ShortAnswerQuestion) Profile(db XODB) (*Profile, error) {
	return ProfileByID(db, saq.AuthorID)
}

// Class returns the Class associated with the ShortAnswerQuestion's ClassID (class_id).
//
// Generated from foreign key 'short_answer_class_fkey'.
func (saq *ShortAnswerQuestion) Class(db XODB) (*Class, error) {
	return ClassByID(db, saq.ClassID)
}

// Unit returns the Unit associated with the ShortAnswerQuestion's UnitID (unit_id).
//
// Generated from foreign key 'short_answer_unit_id_fkey'.
func (saq *ShortAnswerQuestion) Unit(db XODB) (*Unit, error) {
	return UnitByID(db, saq.UnitID)
}

// ShortAnswerQuestionByID retrieves a row from 'public.short_answer_questions' as a ShortAnswerQuestion.
//
// Generated from index 'short_answer_questions_pkey'.
func ShortAnswerQuestionByID(db XODB, id uuid.UUID) (*ShortAnswerQuestion, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, class_id, author_id, prompt, minimum_length, maximum_length, unit_id, timestamp ` +
		`FROM public.short_answer_questions ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	saq := ShortAnswerQuestion{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&saq.ID, &saq.ClassID, &saq.AuthorID, &saq.Prompt, &saq.MinimumLength, &saq.MaximumLength, &saq.UnitID, &saq.Timestamp)
	if err != nil {
		return nil, err
	}

	return &saq, nil
}

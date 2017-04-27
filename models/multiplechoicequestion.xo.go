// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

// MultipleChoiceQuestion represents a row from 'public.multiple_choice_questions'.
type MultipleChoiceQuestion struct {
	ID            uuid.UUID `json:"id"`             // id
	ClassID       uuid.UUID `json:"class_id"`       // class_id
	AuthorID      uuid.UUID `json:"author_id"`      // author_id
	Question      string    `json:"question"`       // question
	PartialCredit bool      `json:"partial_credit"` // partial_credit
	Timestamp     time.Time `json:"timestamp"`      // timestamp
	UnitID        uuid.UUID `json:"unit_id"`        // unit_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the MultipleChoiceQuestion exists in the database.
func (mcq *MultipleChoiceQuestion) Exists() bool {
	return mcq._exists
}

// Deleted provides information if the MultipleChoiceQuestion has been deleted from the database.
func (mcq *MultipleChoiceQuestion) Deleted() bool {
	return mcq._deleted
}

// Insert inserts the MultipleChoiceQuestion to the database.
func (mcq *MultipleChoiceQuestion) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if mcq._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.multiple_choice_questions (` +
		`id, class_id, author_id, question, partial_credit, timestamp, unit_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`)`

	// run query
	XOLog(sqlstr, mcq.ID, mcq.ClassID, mcq.AuthorID, mcq.Question, mcq.PartialCredit, mcq.Timestamp, mcq.UnitID)
	err = db.QueryRow(sqlstr, mcq.ID, mcq.ClassID, mcq.AuthorID, mcq.Question, mcq.PartialCredit, mcq.Timestamp, mcq.UnitID).Scan(&mcq.ID)
	if err != nil {
		return err
	}

	// set existence
	mcq._exists = true

	return nil
}

// Update updates the MultipleChoiceQuestion in the database.
func (mcq *MultipleChoiceQuestion) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !mcq._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if mcq._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.multiple_choice_questions SET (` +
		`class_id, author_id, question, partial_credit, timestamp, unit_id` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6` +
		`) WHERE id = $7`

	// run query
	XOLog(sqlstr, mcq.ClassID, mcq.AuthorID, mcq.Question, mcq.PartialCredit, mcq.Timestamp, mcq.UnitID, mcq.ID)
	_, err = db.Exec(sqlstr, mcq.ClassID, mcq.AuthorID, mcq.Question, mcq.PartialCredit, mcq.Timestamp, mcq.UnitID, mcq.ID)
	return err
}

// Save saves the MultipleChoiceQuestion to the database.
func (mcq *MultipleChoiceQuestion) Save(db XODB) error {
	if mcq.Exists() {
		return mcq.Update(db)
	}

	return mcq.Insert(db)
}

// Upsert performs an upsert for MultipleChoiceQuestion.
//
// NOTE: PostgreSQL 9.5+ only
func (mcq *MultipleChoiceQuestion) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if mcq._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.multiple_choice_questions (` +
		`id, class_id, author_id, question, partial_credit, timestamp, unit_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, class_id, author_id, question, partial_credit, timestamp, unit_id` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.class_id, EXCLUDED.author_id, EXCLUDED.question, EXCLUDED.partial_credit, EXCLUDED.timestamp, EXCLUDED.unit_id` +
		`)`

	// run query
	XOLog(sqlstr, mcq.ID, mcq.ClassID, mcq.AuthorID, mcq.Question, mcq.PartialCredit, mcq.Timestamp, mcq.UnitID)
	_, err = db.Exec(sqlstr, mcq.ID, mcq.ClassID, mcq.AuthorID, mcq.Question, mcq.PartialCredit, mcq.Timestamp, mcq.UnitID)
	if err != nil {
		return err
	}

	// set existence
	mcq._exists = true

	return nil
}

// Delete deletes the MultipleChoiceQuestion from the database.
func (mcq *MultipleChoiceQuestion) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !mcq._exists {
		return nil
	}

	// if deleted, bail
	if mcq._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.multiple_choice_questions WHERE id = $1`

	// run query
	XOLog(sqlstr, mcq.ID)
	_, err = db.Exec(sqlstr, mcq.ID)
	if err != nil {
		return err
	}

	// set deleted
	mcq._deleted = true

	return nil
}

// Profile returns the Profile associated with the MultipleChoiceQuestion's AuthorID (author_id).
//
// Generated from foreign key 'multiple_choice_author_fkey'.
func (mcq *MultipleChoiceQuestion) Profile(db XODB) (*Profile, error) {
	return ProfileByID(db, mcq.AuthorID)
}

// Class returns the Class associated with the MultipleChoiceQuestion's ClassID (class_id).
//
// Generated from foreign key 'multiple_choice_class_fkey'.
func (mcq *MultipleChoiceQuestion) Class(db XODB) (*Class, error) {
	return ClassByID(db, mcq.ClassID)
}

// Unit returns the Unit associated with the MultipleChoiceQuestion's UnitID (unit_id).
//
// Generated from foreign key 'multiple_choice_unit_id_fkey'.
func (mcq *MultipleChoiceQuestion) Unit(db XODB) (*Unit, error) {
	return UnitByID(db, mcq.UnitID)
}

// MultipleChoiceQuestionByID retrieves a row from 'public.multiple_choice_questions' as a MultipleChoiceQuestion.
//
// Generated from index 'multiple_choice_questions_pkey'.
func MultipleChoiceQuestionByID(db XODB, id uuid.UUID) (*MultipleChoiceQuestion, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, class_id, author_id, question, partial_credit, timestamp, unit_id ` +
		`FROM public.multiple_choice_questions ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	mcq := MultipleChoiceQuestion{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&mcq.ID, &mcq.ClassID, &mcq.AuthorID, &mcq.Question, &mcq.PartialCredit, &mcq.Timestamp, &mcq.UnitID)
	if err != nil {
		return nil, err
	}

	return &mcq, nil
}

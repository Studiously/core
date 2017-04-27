// Package models contains the types for schema 'public'.
package models

import uuid "github.com/satori/go.uuid"

// GENERATED BY XO. DO NOT EDIT.

// QuestionRating represents a row from 'public.question_ratings'.
type QuestionRating struct {
	AuthorID   uuid.UUID `json:"author_id"`   // author_id
	QuestionID uuid.UUID `json:"question_id"` // question_id
	Vote       bool      `json:"vote"`        // vote
}

// Profile returns the Profile associated with the QuestionRating's AuthorID (author_id).
//
// Generated from foreign key 'ratings_rater_fkey'.
func (qr *QuestionRating) Profile(db XODB) (*Profile, error) {
	return ProfileByID(db, qr.AuthorID)
}

// QuestionRatingByQuestionIDAuthorID retrieves a row from 'public.question_ratings' as a QuestionRating.
//
// Generated from index 'ratings_question_id_author_id_key'.
func QuestionRatingByQuestionIDAuthorID(db XODB, questionID uuid.UUID, authorID uuid.UUID) (*QuestionRating, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`author_id, question_id, vote ` +
		`FROM public.question_ratings ` +
		`WHERE question_id = $1 AND author_id = $2`

	// run query
	XOLog(sqlstr, questionID, authorID)
	qr := QuestionRating{}

	err = db.QueryRow(sqlstr, questionID, authorID).Scan(&qr.AuthorID, &qr.QuestionID, &qr.Vote)
	if err != nil {
		return nil, err
	}

	return &qr, nil
}
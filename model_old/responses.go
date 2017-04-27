package model_old

import "github.com/satori/go.uuid"

type MultipleChoiceResponse struct {
	ID uuid.UUID `db:"id" json:"id"`
	QuestionID uuid.UUID `db:"question_id" json:"question_id"`
	AuthorID uuid.UUID `db:"author_id" json:"author_id"`
	ResponseID uuid.UUID `db:"response_id" json:"response_id"`
}

type ShortAnswerResponse struct {
	ID uuid.UUID `db:"id" json:"id"`
	QuestionID uuid.UUID `db:"question_id" json:"question_id"`
	AuthorID uuid.UUID `db:"author_id" json:"author_id"`
	Response string `db:"response" json:"response"`
}

type TrueFalseResponse struct {
	ID uuid.UUID `db:"id" json:"id"`
	QuestionID uuid.UUID `db:"question_id" json:"question_id"`
	AuthorID uuid.UUID `db:"author_id" json:"author_id"`
	Response bool `db:"response" json:"response"`
}
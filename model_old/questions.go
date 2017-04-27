package model_old

import (
	"time"

	"github.com/satori/go.uuid"
)

type Question struct {
	ID           uuid.UUID `db:"id" json:"id"`
	AuthorID     uuid.UUID `db:"author_id" json:"author_id"`
	QuestionType string `db:"question_type" json:"question_type"`
	UnitID       uuid.UUID `db:"unit_id" json:"unit_id"`
	VoteCount    int `db:"vote_count" json:"vote_count"`
}

type MultipleChoiceQuestion struct {
	ID            uuid.UUID `db:"id" json:"id"`
	ClassID       uuid.UUID `db:"class_id" json:"class_id"`
	AuthorID      uuid.UUID `db:"author_id" json:"author_id"`
	Question      string `db:"question" json:"question"`
	PartialCredit bool `db:"partial_credit" json:"partial_credit"`
	Timestamp     time.Time `db:"timestamp" json:"timestamp"`
	UnitID        uuid.UUID `db:"unit_id" json:"unit_id"`
}

type MultipleChoiceAnswer struct {
	ID         uuid.UUID `db:"id" json:"id"`
	QuestionID uuid.UUID `db:"question_id" json:"question_id"`
	Answer     string `db:"answer" json:"answer"`
	Correct    bool `db:"correct" json:"correct"`
}

type ShortAnswerQuestion struct {
	ID            uuid.UUID `db:"id" json:"id"`
	ClassID       uuid.UUID `db:"class_id" json:"class_id"`
	AuthorID      uuid.UUID `db:"author_id" json:"author_id"`
	Prompt        string `db:"prompt" json:"prompt"`
	MinimumLength int `db:"minimum_length" json:"minimum_length"`
	MaximumLength int `db:"maximum_length" json:"maximum_length"`
	UnitID        uuid.UUID `db:"unit_id" json:"unit_id"`
}

type TrueFalseQuestion struct {
	ID            uuid.UUID `db:"id" json:"id"`
	ClassID       uuid.UUID `db:"class_id" json:"class_id"`
	AuthorID      uuid.UUID `db:"author_id" json:"author_id"`
	Question      string `db:"question" json:"question"`
	CorrectAnswer bool `db:"correct_answer" json:"correct_answer"`
	UnitID        uuid.UUID `db:"unit_id" json:"unit_id"`
}

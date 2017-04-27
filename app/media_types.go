// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "studiously": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/studiously/core/design
// --out=$(GOPATH)/src/github.com/studiously/core
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
	uuid "github.com/satori/go.uuid"
)

// StudiouslyClass media type (default view)
//
// Identifier: application/studiously.class+json; view=default
type StudiouslyClass struct {
	// Current unit of study
	CurrentUnit *uuid.UUID `form:"current_unit,omitempty" json:"current_unit,omitempty" xml:"current_unit,omitempty"`
	ID          uuid.UUID  `form:"id" json:"id" xml:"id"`
	Name        string     `form:"name" json:"name" xml:"name"`
}

// Validate validates the StudiouslyClass media type instance.
func (mt *StudiouslyClass) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// StudiouslyClassCollection is the media type for an array of StudiouslyClass (default view)
//
// Identifier: application/studiously.class+json; type=collection; view=default
type StudiouslyClassCollection []*StudiouslyClass

// Validate validates the StudiouslyClassCollection media type instance.
func (mt StudiouslyClassCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// StudiouslyMember media type (default view)
//
// Identifier: application/studiously.member+json; view=default
type StudiouslyMember struct {
	ID   uuid.UUID `form:"id" json:"id" xml:"id"`
	Name string    `form:"name" json:"name" xml:"name"`
	Role string    `form:"role" json:"role" xml:"role"`
}

// Validate validates the StudiouslyMember media type instance.
func (mt *StudiouslyMember) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Role == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "role"))
	}
	if !(mt.Role == "student" || mt.Role == "moderator" || mt.Role == "teacher" || mt.Role == "administrator") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.role`, mt.Role, []interface{}{"student", "moderator", "teacher", "administrator"}))
	}
	return
}

// StudiouslyMemberCollection is the media type for an array of StudiouslyMember (default view)
//
// Identifier: application/studiously.member+json; type=collection; view=default
type StudiouslyMemberCollection []*StudiouslyMember

// Validate validates the StudiouslyMemberCollection media type instance.
func (mt StudiouslyMemberCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// StudiouslyQuestion media type (by_author view)
//
// Identifier: application/studiously.question+json; view=by_author
type StudiouslyQuestionByAuthor struct {
	Answered     *bool        `form:"answered,omitempty" json:"answered,omitempty" xml:"answered,omitempty"`
	Data         *interface{} `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	ID           uuid.UUID    `form:"id" json:"id" xml:"id"`
	QuestionType string       `form:"question_type" json:"question_type" xml:"question_type"`
	UnitID       uuid.UUID    `form:"unit_id" json:"unit_id" xml:"unit_id"`
	Votes        *int         `form:"votes,omitempty" json:"votes,omitempty" xml:"votes,omitempty"`
}

// Validate validates the StudiouslyQuestionByAuthor media type instance.
func (mt *StudiouslyQuestionByAuthor) Validate() (err error) {

	if mt.QuestionType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "question_type"))
	}

	return
}

// StudiouslyQuestion media type (by_type view)
//
// Identifier: application/studiously.question+json; view=by_type
type StudiouslyQuestionByType struct {
	Answered *bool        `form:"answered,omitempty" json:"answered,omitempty" xml:"answered,omitempty"`
	AuthorID *uuid.UUID   `form:"author_id,omitempty" json:"author_id,omitempty" xml:"author_id,omitempty"`
	Data     *interface{} `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	ID       uuid.UUID    `form:"id" json:"id" xml:"id"`
	UnitID   uuid.UUID    `form:"unit_id" json:"unit_id" xml:"unit_id"`
	Votes    *int         `form:"votes,omitempty" json:"votes,omitempty" xml:"votes,omitempty"`
}

// Validate validates the StudiouslyQuestionByType media type instance.
func (mt *StudiouslyQuestionByType) Validate() (err error) {

	return
}

// StudiouslyQuestion media type (by_unit view)
//
// Identifier: application/studiously.question+json; view=by_unit
type StudiouslyQuestionByUnit struct {
	Answered     *bool        `form:"answered,omitempty" json:"answered,omitempty" xml:"answered,omitempty"`
	AuthorID     *uuid.UUID   `form:"author_id,omitempty" json:"author_id,omitempty" xml:"author_id,omitempty"`
	Data         *interface{} `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	ID           uuid.UUID    `form:"id" json:"id" xml:"id"`
	QuestionType string       `form:"question_type" json:"question_type" xml:"question_type"`
	Votes        *int         `form:"votes,omitempty" json:"votes,omitempty" xml:"votes,omitempty"`
}

// Validate validates the StudiouslyQuestionByUnit media type instance.
func (mt *StudiouslyQuestionByUnit) Validate() (err error) {

	if mt.QuestionType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "question_type"))
	}
	return
}

// StudiouslyQuestion media type (default view)
//
// Identifier: application/studiously.question+json; view=default
type StudiouslyQuestion struct {
	Answered     *bool        `form:"answered,omitempty" json:"answered,omitempty" xml:"answered,omitempty"`
	AuthorID     *uuid.UUID   `form:"author_id,omitempty" json:"author_id,omitempty" xml:"author_id,omitempty"`
	Data         *interface{} `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	ID           uuid.UUID    `form:"id" json:"id" xml:"id"`
	QuestionType string       `form:"question_type" json:"question_type" xml:"question_type"`
	UnitID       uuid.UUID    `form:"unit_id" json:"unit_id" xml:"unit_id"`
	Votes        *int         `form:"votes,omitempty" json:"votes,omitempty" xml:"votes,omitempty"`
}

// Validate validates the StudiouslyQuestion media type instance.
func (mt *StudiouslyQuestion) Validate() (err error) {

	if mt.QuestionType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "question_type"))
	}

	return
}

// StudiouslyQuestion media type (feed view)
//
// Identifier: application/studiously.question+json; view=feed
type StudiouslyQuestionFeed struct {
	Answered *bool      `form:"answered,omitempty" json:"answered,omitempty" xml:"answered,omitempty"`
	AuthorID *uuid.UUID `form:"author_id,omitempty" json:"author_id,omitempty" xml:"author_id,omitempty"`
	ID       uuid.UUID  `form:"id" json:"id" xml:"id"`
	UnitID   uuid.UUID  `form:"unit_id" json:"unit_id" xml:"unit_id"`
	Votes    *int       `form:"votes,omitempty" json:"votes,omitempty" xml:"votes,omitempty"`
}

// Validate validates the StudiouslyQuestionFeed media type instance.
func (mt *StudiouslyQuestionFeed) Validate() (err error) {

	return
}

// StudiouslyQuestionCollection is the media type for an array of StudiouslyQuestion (by_author view)
//
// Identifier: application/studiously.question+json; type=collection; view=by_author
type StudiouslyQuestionByAuthorCollection []*StudiouslyQuestionByAuthor

// Validate validates the StudiouslyQuestionByAuthorCollection media type instance.
func (mt StudiouslyQuestionByAuthorCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// StudiouslyQuestionCollection is the media type for an array of StudiouslyQuestion (by_type view)
//
// Identifier: application/studiously.question+json; type=collection; view=by_type
type StudiouslyQuestionByTypeCollection []*StudiouslyQuestionByType

// Validate validates the StudiouslyQuestionByTypeCollection media type instance.
func (mt StudiouslyQuestionByTypeCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// StudiouslyQuestionCollection is the media type for an array of StudiouslyQuestion (by_unit view)
//
// Identifier: application/studiously.question+json; type=collection; view=by_unit
type StudiouslyQuestionByUnitCollection []*StudiouslyQuestionByUnit

// Validate validates the StudiouslyQuestionByUnitCollection media type instance.
func (mt StudiouslyQuestionByUnitCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// StudiouslyQuestionCollection is the media type for an array of StudiouslyQuestion (default view)
//
// Identifier: application/studiously.question+json; type=collection; view=default
type StudiouslyQuestionCollection []*StudiouslyQuestion

// Validate validates the StudiouslyQuestionCollection media type instance.
func (mt StudiouslyQuestionCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// StudiouslyQuestionCollection is the media type for an array of StudiouslyQuestion (feed view)
//
// Identifier: application/studiously.question+json; type=collection; view=feed
type StudiouslyQuestionFeedCollection []*StudiouslyQuestionFeed

// Validate validates the StudiouslyQuestionFeedCollection media type instance.
func (mt StudiouslyQuestionFeedCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

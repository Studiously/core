// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "studiously": CLI Commands
//
// Command:
// $ goagen
// --design=github.com/studiously/core/design
// --out=$(GOPATH)/src/github.com/studiously/core
// --version=v1.2.0-dirty

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"github.com/studiously/core/client"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// ListClassCommand is the command line data structure for the list action of class
	ListClassCommand struct {
		PrettyPrint bool
	}

	// ShowClassCommand is the command line data structure for the show action of class
	ShowClassCommand struct {
		// Class ID
		ClassID     int
		PrettyPrint bool
	}

	// ShowMembersClassCommand is the command line data structure for the show_members action of class
	ShowMembersClassCommand struct {
		// Class ID
		ClassID     int
		PrettyPrint bool
	}

	// ShowQuestionsClassCommand is the command line data structure for the show_questions action of class
	ShowQuestionsClassCommand struct {
		// Class ID
		ClassID int
		// Filter by whether the question has been answered by the member
		Answered string
		// Filter by author
		AuthorID int
		// Filter by question type
		QuestionType string
		// Filter by unit
		UnitID      int
		PrettyPrint bool
	}

	// ShowQuestionCommand is the command line data structure for the show action of question
	ShowQuestionCommand struct {
		// Question ID
		QuestionID  int
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "list",
		Short: `Get all classes a user is in`,
	}
	tmp1 := new(ListClassCommand)
	sub = &cobra.Command{
		Use:   `class ["/classes"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp2 := new(ShowClassCommand)
	sub = &cobra.Command{
		Use:   `class ["/classes/CLASS_ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp3 := new(ShowQuestionCommand)
	sub = &cobra.Command{
		Use:   `question ["/questions/QUESTIONID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show-members",
		Short: `Get members of a class`,
	}
	tmp4 := new(ShowMembersClassCommand)
	sub = &cobra.Command{
		Use:   `class ["/classes/CLASS_ID/members"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show-questions",
		Short: `Get questions for a class`,
	}
	tmp5 := new(ShowQuestionsClassCommand)
	sub = &cobra.Command{
		Use:   `class ["/classes/CLASS_ID/questions"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the ListClassCommand command.
func (cmd *ListClassCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/classes"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListClass(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListClassCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowClassCommand command.
func (cmd *ShowClassCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/classes/%v", cmd.ClassID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowClass(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowClassCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var classID int
	cc.Flags().IntVar(&cmd.ClassID, "class_id", classID, `Class ID`)
}

// Run makes the HTTP request corresponding to the ShowMembersClassCommand command.
func (cmd *ShowMembersClassCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/classes/%v/members", cmd.ClassID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowMembersClass(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowMembersClassCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var classID int
	cc.Flags().IntVar(&cmd.ClassID, "class_id", classID, `Class ID`)
}

// Run makes the HTTP request corresponding to the ShowQuestionsClassCommand command.
func (cmd *ShowQuestionsClassCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/classes/%v/questions", cmd.ClassID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	var tmp6 *bool
	if cmd.Answered != "" {
		var err error
		tmp6, err = boolVal(cmd.Answered)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *bool value", "flag", "--answered", "err", err)
			return err
		}
	}
	resp, err := c.ShowQuestionsClass(ctx, path, tmp6, intFlagVal("author_id", cmd.AuthorID), stringFlagVal("question_type", cmd.QuestionType), intFlagVal("unit_id", cmd.UnitID))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowQuestionsClassCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var classID int
	cc.Flags().IntVar(&cmd.ClassID, "class_id", classID, `Class ID`)
	var answered string
	cc.Flags().StringVar(&cmd.Answered, "answered", answered, `Filter by whether the question has been answered by the member`)
	var authorID int
	cc.Flags().IntVar(&cmd.AuthorID, "author_id", authorID, `Filter by author`)
	var questionType string
	cc.Flags().StringVar(&cmd.QuestionType, "question_type", questionType, `Filter by question type`)
	var unitID int
	cc.Flags().IntVar(&cmd.UnitID, "unit_id", unitID, `Filter by unit`)
}

// Run makes the HTTP request corresponding to the ShowQuestionCommand command.
func (cmd *ShowQuestionCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/questions/%v", cmd.QuestionID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowQuestion(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowQuestionCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var questionID int
	cc.Flags().IntVar(&cmd.QuestionID, "questionID", questionID, `Question ID`)
}

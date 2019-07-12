package cmd

import "github.com/pkg/errors"

const (
	Version = "1.0.0"
	ProjectConfigFile = ".taskmaster.json"

	AppDataKey = "appdata"
	DefaultData = `{
	"mit": {
		"tasks_type": "MIT",
		"tasks": []
	},
	"tasks": {
		"tasks_type": "TODO",
		"tasks": []
	}
}`

	MITLimitKey = "config.mit_limit"
	TodoLimitKey = "config.todo_limit"
)

var (
	ErrOnlyOneRm = errors.New("only one task can be removed from list at a time")
	ErrOnlyOneMv = errors.New("only one task can be moved from list at a time")
	ErrInvalidMv = errors.New("cannot move that task")
	ErrNAN = errors.New("not a number")
)

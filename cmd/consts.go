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
)

var (
	ErrOnlyOneRm = errors.New("only one task can be removed from list at a time")
	ErrNAN = errors.New("not a number")
)

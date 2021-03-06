package cli

import (
	"io"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	catCommand     = app.Command("cat", "Displays contents of a repository object.")
	catCommandPath = catCommand.Arg("path", "Path").Required().String()
)

func runCatCommand(context *kingpin.ParseContext) error {
	rep := mustOpenRepository(nil)
	defer rep.Close()

	oid, err := parseObjectID(*catCommandPath, rep)
	if err != nil {
		return err
	}
	r, err := rep.Open(oid)
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, r)
	return nil
}

func init() {
	catCommand.Action(runCatCommand)
}

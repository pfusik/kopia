package cli

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/kopia/kopia/repo"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	showCommand = objectCommands.Command("show", "Show contents of a repository object.")

	showObjectIDs = showCommand.Arg("id", "IDs of objects to show").Required().Strings()
	showJSON      = showCommand.Flag("json", "Pretty-print JSON content").Short('j').Bool()
	showRaw       = showCommand.Flag("raw", "Show raw content (disables format auto-detection)").Short('r').Bool()
)

func runShowCommand(context *kingpin.ParseContext) error {
	rep := mustOpenRepository(nil)
	defer rep.Close()

	for _, oidString := range *showObjectIDs {
		oid, err := parseObjectID(oidString, rep)
		if err != nil {
			return err
		}

		if err := showObject(rep, oid); err != nil {
			return err
		}
	}

	return nil
}

func showObject(r *repo.Repository, oid repo.ObjectID) error {
	rd, err := r.Open(oid)
	if err != nil {
		return err
	}
	defer rd.Close()

	rawdata, err := ioutil.ReadAll(rd)
	if err != nil {
		return err
	}

	format := "raw"

	if len(rawdata) >= 2 && rawdata[0] == '{' && rawdata[len(rawdata)-1] == '}' {
		format = "json"
	}

	if *showJSON {
		format = "json"
	}

	if *showRaw {
		format = "raw"
	}

	switch format {
	case "json":
		var buf bytes.Buffer

		json.Indent(&buf, rawdata, "", "  ")
		os.Stdout.Write(buf.Bytes())

	default:
		os.Stdout.Write(rawdata)
	}
	return nil
}

func init() {
	showCommand.Action(runShowCommand)
}

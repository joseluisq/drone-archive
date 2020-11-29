package main

import (
	"log"
	"os"

	cli "github.com/joseluisq/cline"
	archive "github.com/joseluisq/drone-archive"
)

var (
	versionNumber string = "devel"
	buildTime     string
)

func main() {
	app := cli.New()
	app.Name = "archive plugin"
	app.Summary = "Archive a file or directory using Tar/GZ or Zip with checksum computation."
	app.Version = versionNumber
	app.BuildTime = buildTime
	app.Flags = []cli.Flag{
		cli.FlagString{
			Name:    "source",
			Summary: "file or directory to archive and compress",
			EnvVar:  "PLUGIN_SOURCE",
		},
		cli.FlagString{
			Name:    "destination",
			Summary: "path of archived/compressed output file",
			EnvVar:  "PLUGIN_DESTINATION",
		},
	}
	app.Handler = appHandler
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func appHandler(ctx *cli.AppContext) error {
	plugin := archive.Plugin{
		Src: ctx.Flags.String("source"),
		Dst: ctx.Flags.String("destination"),
	}
	return plugin.Exec()
}

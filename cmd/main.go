package main

import (
	"fmt"
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
	app.Summary = "Archive a file or directory using Tar/GZ or Zip with optional checksum computation."
	app.Version = versionNumber
	app.BuildTime = buildTime
	app.Flags = []cli.Flag{
		cli.FlagString{
			Name:    "source",
			Summary: "File or directory to archive and compress.",
			Aliases: []string{"s"},
			EnvVar:  "PLUGIN_SOURCE",
		},
		cli.FlagString{
			Name:    "destination",
			Summary: "File path to save the archived and compressed file.",
			Aliases: []string{"d"},
			EnvVar:  "PLUGIN_DESTINATION",
		},
		cli.FlagString{
			Name:    "format",
			Summary: "Define a `tar` and `zip` archiving format with compression. tar format uses Gzip compression.",
			Value:   "tar",
			Aliases: []string{"f"},
			EnvVar:  "PLUGIN_FORMAT",
		},
		cli.FlagBool{
			Name:    "checksum",
			Summary: "Enable checksum file computation. File'll be saved on same base path like destination.",
			Value:   false,
			Aliases: []string{"c"},
			EnvVar:  "PLUGIN_CHECKSUM",
		},
		cli.FlagString{
			Name:    "checksum-algo",
			Summary: "Define the checksum `md5`, `sha1`, `sha256` or `sha512` algorithm.",
			Value:   "sha256sum",
			Aliases: []string{"a"},
			EnvVar:  "PLUGIN_CHECKSUM_ALGO",
		},
		cli.FlagString{
			Name:    "checksum-destination",
			Summary: "Define the checksum file destination path.",
			Aliases: []string{"e"},
			EnvVar:  "PLUGIN_CHECKSUM_DESTINATION",
		},
	}
	app.Handler = appHandler
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func appHandler(ctx *cli.AppContext) error {
	source := ctx.Flags.String("source")
	if source == "" {
		return fmt.Errorf("source has an empty value")
	}
	destination := ctx.Flags.String("destination")
	if destination == "" {
		return fmt.Errorf("destination has an empty value")
	}
	format := ctx.Flags.String("format")
	if format == "" {
		return fmt.Errorf("format has an empty value")
	}
	checksum, err := ctx.Flags.Bool("checksum")
	if err != nil {
		return err
	}
	checksumAlgo := ctx.Flags.String("checksum-algo")
	if checksumAlgo == "" {
		return fmt.Errorf("checksum-algo has an empty value")
	}
	plugin := archive.Plugin{
		Source:       source,
		Destination:  destination,
		Format:       format,
		Checksum:     checksum,
		ChecksumAlgo: checksumAlgo,
	}
	return plugin.Exec()
}

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
			Name:    "src",
			Summary: "File or directory to archive and compress.",
			Aliases: []string{"s"},
			EnvVar:  "PLUGIN_SOURCE",
		},
		cli.FlagString{
			Name:    "dest",
			Summary: "File destination path to save the archived/compressed file.",
			Aliases: []string{"d"},
			EnvVar:  "PLUGIN_DESTINATION",
		},
		cli.FlagString{
			Name:    "format",
			Summary: "Define a `tar` and `zip` archiving format with compression. Tar format uses Gzip compression.",
			Value:   "tar",
			Aliases: []string{"f"},
			EnvVar:  "PLUGIN_FORMAT",
		},
		cli.FlagBool{
			Name:    "checksum",
			Summary: "Enable checksum file computation.",
			Value:   false,
			Aliases: []string{"c"},
			EnvVar:  "PLUGIN_CHECKSUM",
		},
		cli.FlagString{
			Name:    "checksum-algo",
			Summary: "Define the checksum `md5`, `sha1`, `sha256` or `sha512` algorithm.",
			Value:   "sha256",
			Aliases: []string{"a"},
			EnvVar:  "PLUGIN_CHECKSUM_ALGO",
		},
		cli.FlagString{
			Name:    "checksum-dest",
			Summary: "File destination path of the checksum.",
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
	source := ctx.Flags.String("src").Value()
	if source == "" {
		return fmt.Errorf("source file or directory path was not provided")
	}
	destination := ctx.Flags.String("dest").Value()
	if destination == "" {
		return fmt.Errorf("archive file destination path was not provided")
	}
	format := ctx.Flags.String("format").Value()
	if format == "" {
		return fmt.Errorf("archive format was not provided or unsupported")
	}
	checksum, err := ctx.Flags.Bool("checksum").Value()
	if err != nil {
		return err
	}
	checksumAlgo := ctx.Flags.String("checksum-algo").Value()
	if checksum && checksumAlgo == "" {
		return fmt.Errorf("checksum algorithm was not provided")
	}
	checksumDest := ctx.Flags.String("checksum-dest").Value()
	if checksum && checksumDest == "" {
		return fmt.Errorf("checksum file destination path was not provided")
	}
	plugin := archive.Plugin{
		Source:       source,
		Destination:  destination,
		Format:       format,
		Checksum:     checksum,
		ChecksumAlgo: checksumAlgo,
		ChecksumDest: checksumDest,
	}
	return plugin.Exec()
}

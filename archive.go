package archive

import (
	"fmt"
	"log"

	"github.com/joseluisq/compactor"
)

// Plugin defines the Archive plugin parameters.
type Plugin struct {
	BasePath     string
	Source       string
	Destination  string
	Format       string
	Checksum     bool
	ChecksumAlgo string
	ChecksumDest string
}

// Exec executes the plugin step
func (p Plugin) Exec() (err error) {
	var checksumFile string
	switch p.Format {
	case "tar":
		if p.Checksum {
			checksumFile, err = compactor.CreateTarballWithChecksum(
				p.BasePath,
				p.Source,
				p.Destination,
				p.ChecksumAlgo,
				p.ChecksumDest,
			)
		} else {
			err = compactor.CreateTarball(p.BasePath, p.Source, p.Destination)
		}
		if err != nil {
			return err
		}
		log.Printf("tar/gzip file saved on %s\n", p.Destination)
		if p.Checksum {
			log.Printf("checksum file saved on %s\n", checksumFile)
		}
	case "zip":
		if p.Checksum {
			checksumFile, err = compactor.CreateZipballWithChecksum(
				p.BasePath,
				p.Source,
				p.Destination,
				p.ChecksumAlgo,
				p.ChecksumDest,
			)
		} else {
			err = compactor.CreateZipball(p.BasePath, p.Source, p.Destination)
		}
		if err != nil {
			return err
		}
		log.Printf("zip file saved on %s\n", p.Destination)
		if p.Checksum {
			log.Printf("checksum file saved on %s\n", checksumFile)
		}
	default:
		return fmt.Errorf("archive format `%s` is not supported", p.Format)
	}
	return err
}

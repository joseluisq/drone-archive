package archive

import (
	"fmt"
	"log"

	"github.com/joseluisq/compactor"
)

// Plugin defines the Archive plugin parameters.
type Plugin struct {
	Source       string
	Destination  string
	Format       string
	Checksum     bool
	ChecksumAlgo string
	ChecksumDest string
}

// Exec executes the plugin step
func (p Plugin) Exec() error {
	switch p.Format {
	case "tar":
		if p.Checksum {
			checksum, err := compactor.CreateTarballWithChecksum(
				p.Source,
				p.Destination,
				p.ChecksumAlgo,
				p.ChecksumDest,
			)
			if err != nil {
				return err
			}
			log.Printf("checksum file saved at %s.\n", checksum)
		} else {
			return compactor.CreateTarball(p.Source, p.Destination)
		}
	case "zip":
		if p.Checksum {
			checksum, err := compactor.CreateZipballWithChecksum(
				p.Source,
				p.Destination,
				p.ChecksumAlgo,
				p.ChecksumDest,
			)
			if err != nil {
				return err
			}
			log.Printf("checksum file saved at %s.\n", checksum)
		} else {
			return compactor.CreateZipball(p.Source, p.Destination)
		}
	default:
		return fmt.Errorf("format `%s` is not supported", p.Format)
	}
	return nil
}

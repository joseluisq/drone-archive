package archive

// Plugin defines the Archive plugin parameters.
type Plugin struct {
	Src string
	Dst string
}

// Exec executes the plugin step
func (p Plugin) Exec() error {
	return nil
}

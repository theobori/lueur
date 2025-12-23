package gophermap

import "fmt"

type FileFormat int

const (
	FileFormatGPH FileFormat = iota
	FileFormatGophermap
)

func NewFileFormatFromString(s string) (FileFormat, error) {
	switch s {
	case "gph":
		return FileFormatGPH, nil
	case "gophermap":
		return FileFormatGophermap, nil
	default:
		return FileFormatGophermap, fmt.Errorf(
			"'%s' is not available, it must be '%s' or '%s'",
			s,
			"gph",
			"gophermap",
		)
	}
}

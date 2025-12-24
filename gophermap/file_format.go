package gophermap

import "fmt"

type FileFormat int

const (
	FileFormatGPH FileFormat = iota
	FileFormatGophermap
	FileFormatTxt
)

func NewFileFormatFromString(s string) (FileFormat, error) {
	switch s {
	case "gph":
		return FileFormatGPH, nil
	case "gophermap":
		return FileFormatGophermap, nil
	case "txt":
		return FileFormatTxt, nil
	default:
		return FileFormatGophermap, fmt.Errorf(
			"'%s' is not available, it must be '%s' or '%s'",
			s,
			"gph",
			"gophermap",
		)
	}
}

func (f *FileFormat) String() string {
	switch *f {
	case FileFormatGPH:
		return "gph"
	case FileFormatTxt:
		return "txt"
	case FileFormatGophermap:
		return "gophermap"
	// Cannot reach this block
	default:
		return "unknown"
	}
}

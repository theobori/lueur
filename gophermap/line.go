package gophermap

import (
	"fmt"
)

type Line struct {
	ItemType    ItemType
	Description string
	Path        string
	Domain      string
	Port        int
}

func NewLine(
	itemType ItemType,
	description string,
	path string,
	domain string,
	port int,
) (*Line, error) {
	if port < 0 {
		return nil, fmt.Errorf("%d is negative, the port must be positive", port)
	}

	return &Line{itemType, description, path, domain, port}, nil
}

func (l *Line) StringGPHFormat() string {
	return fmt.Sprintf(
		"[%s|%s|%s|%s|%d]",
		l.ItemType.String(), l.Description, l.Path, l.Domain, l.Port,
	)
}

func (l *Line) StringGophermapFormat() string {
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s%d",
		l.ItemType.String(), l.Description, DefaultSeparator, l.Path,
		DefaultSeparator, l.Domain, DefaultSeparator, l.Port,
	)
}

func (l *Line) StringTextFormat() string {
	// The description should contains everything
	return l.Description
}

func (l *Line) String() string {
	return l.StringGophermapFormat()
}

func (l *Line) StringFromFileFormat(fileFormat FileFormat) string {
	switch fileFormat {
	case FileFormatGophermap:
		return l.StringGophermapFormat()
	case FileFormatGPH:
		return l.StringGPHFormat()
	case FileFormatTxt:
		return l.StringTextFormat()
	default:
		return l.String()
	}
}

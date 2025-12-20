package gophermap

import "fmt"

type Line struct {
	itemType    ItemType
	description string
	path        string
	domain      string
	port        int
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

	return &Line{
		itemType,
		description,
		path,
		domain,
		port,
	}, nil
}

func (l *Line) StringWithSep(sep string) string {
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s%d",
		l.itemType.String(),
		l.description,
		sep,
		l.path,
		sep,
		l.domain,
		sep,
		l.port,
	)
}

func (l *Line) String() string {
	return l.StringWithSep(DefaultSeparator)
}

func (l *Line) StringGPHFormat() string {
	return fmt.Sprintf(
		"[%s|%s|%s|%s|%d]",
		l.itemType.String(),
		l.description,
		l.path,
		l.domain,
		l.port,
	)
}

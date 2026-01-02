package walker

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/internal/common"
)

// Dirty shit, don't try to understand this function
//
// It's just handling the different edge cases with
// specific walker reference position
func (w *Walker) processReferenceLineEdgeCases(line *gophermap.Line, destination string) string {
	var inlineAnswer string

	switch w.options.ReferencePosition() {
	case AfterTraverse:
		number := len(w.ctx.ReferencesQueue) + 1
		inlineAnswer = fmt.Sprintf("(%s)[%d]", line.Description, number)

		if w.options.FileFormat() == gophermap.FileFormatTxt {
			line.Description = destination
		}

		line.Description = fmt.Sprintf("[%d] %s", number, line.Description)
	case AfterBlocks:
		inlineAnswer = line.Description
	}

	return inlineAnswer
}

func (w *Walker) referenceLine(description string, destination string) (*gophermap.Line, error) {
	line := gophermap.Line{
		Description: description,
	}

	if common.IsURL(destination) {
		u, err := url.Parse(destination)
		if err != nil {
			return nil, err
		}

		line.Port, err = gophermap.PortFromURL(u)
		if err != nil {
			return nil, err
		}

		line.ItemType = gophermap.NewItemTypeFromURL(u)
		line.Domain = u.Host
		line.Path = gophermap.PathFromURL(u)
	} else {
		line.Port = w.options.Port()
		line.ItemType = gophermap.NewItemTypeFromPath(destination)
		line.Domain = w.options.Domain()
		line.Path = filepath.Join("/", w.options.PathPrefix, "/", destination)
	}

	return &line, nil
}

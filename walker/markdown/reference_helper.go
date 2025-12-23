package markdown

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/internal/common"
	"github.com/yuin/goldmark/ast"
)

func (w *Walker) ContainsOnlyRefs(node ast.Node) (bool, error) {
	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		switch c.(type) {
		case *ast.AutoLink, *ast.Link, *ast.Image:
			continue
		case *ast.Text, *ast.TextBlock:
			s, err := w.Walk(c)
			if err != nil {
				return false, err
			}

			s = strings.TrimSpace(s)
			if len(s) > 0 {
				return false, nil
			}
		default:
			return false, nil
		}
	}

	return true, nil
}

func pathFromURL(u *url.URL) string {
	var path string

	switch u.Scheme {
	case "https", "http":
		path = "URL:" + u.String()
	case "telnet", "tn3270":
		path = u.User.Username()
		if path == "" {
			path = "user"
		}
	// TODO: research on the others ?
	default:
		if u.Path == "" {
			path = "/"
		} else {
			path = u.Path
		}
	}

	return path
}

func portFromURL(u *url.URL) (int, error) {
	portString := u.Port()

	var (
		port int
		err  error
	)

	if portString == "" {
		port = gophermap.PortFromProtocol(u.Scheme)
	} else {
		port, err = strconv.Atoi(portString)
		if err != nil {
			return 0, err
		}
	}

	return port, nil
}

func (w *Walker) walkReferenceHelper(
	node ast.Node, title string, destination string,
) (string, error) {
	var (
		description string
		err         error
	)

	if title != "" {
		description = title
	} else {
		description, err = w.walkIteratorHelper(node)
		if err != nil {
			return "", err
		}
	}

	var (
		itemType gophermap.ItemType
		domain   string
		path     string
		port     int
	)

	if common.IsURL(destination) {
		u, err := url.Parse(destination)
		if err != nil {
			return "", err
		}

		itemType = gophermap.NewItemTypeFromURL(u)
		domain = u.Host

		port, err = portFromURL(u)
		if err != nil {
			return "", err
		}

		path = pathFromURL(u)
		// TODO: handle gopher://host and modify path if needed
	} else {
		itemType = gophermap.NewItemTypeFromPath(destination)
		domain = w.options.Domain
		port = w.options.Port
		path = "/" + strings.TrimLeft(destination, "/")
	}

	line, err := w.createLineString(
		itemType,
		description,
		path,
		domain,
		port,
	)
	if err != nil {
		return "", err
	}

	w.referencesQueue = append(w.referencesQueue, line)

	containsOnlyRefs, err := w.ContainsOnlyRefs(node.Parent())
	if err != nil {
		return "", err
	}

	if containsOnlyRefs {
		return "", nil
	}

	return description, nil
}

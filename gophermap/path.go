package gophermap

import "net/url"

func SafePath(path string) string {
	n := len(path)
	if n < 2 {
		return path
	}

	// /item_type
	if (n == 2 && path[0] == '/' && IsByteItemType(path[1])) ||
		// item_type/path_rest
		(IsByteItemType(path[0]) && path[1] == '/') {
		return "/" + path[2:]
	}

	if n < 3 {
		return path
	}

	// /item_type/path_rest
	if path[0] == '/' && IsByteItemType(path[1]) && path[2] == '/' {
		return path[2:]
	}

	return path
}

func PathFromURL(u *url.URL) string {
	var path string

	switch u.Scheme {
	case "https", "http":
		path = "URL:" + u.String()
	case "telnet", "tn3270":
		path = u.User.Username()
		if path == "" {
			path = "user"
		}
	case "gopher":
		return SafePath(u.Path)
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

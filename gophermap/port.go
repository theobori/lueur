package gophermap

import (
	"net/url"
	"strconv"

	"github.com/theobori/lueur/internal/common"
)

var PortFromProtocol = common.CreatePortFromProtocolFunc(DefaultGopherPort)

func PortFromURL(u *url.URL) (int, error) {
	portString := u.Port()

	if portString == "" {
		return PortFromProtocol(u.Scheme), nil
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return 0, err
	}

	return port, nil
}

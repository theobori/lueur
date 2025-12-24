package gophermap

import (
	"net/url"
	"strings"
)

type ItemType int

// See gopher://baud.baby/0/phlog/fs20181102.txt
const (
	ItemTypeTextFile   ItemType = iota
	ItemTypeGopherMenu          // Submenu or directory
	ItemTypeCCSONameserver
	ItemTypeErrorCode
	ItemTypeBinHexEncodedFile
	ItemTypeDOSBinary
	ItemTypeUNIXUuencodedFile
	ItemTypeGopherFullTextSearch
	ItemTypeTelnet
	ItemTypeBinaryFile
	ItemTypeMirrorOrAlternateServer
	ItemTypeGIFFile
	ItemTypeOtherImageFile
	ItemTypeTelnet3270
	ItemTypeHTML
	ItemTypeInlineText
	ItemTypeSoundFile
)

func (i *ItemType) String() string {
	switch *i {
	case ItemTypeTextFile:
		return "0"
	case ItemTypeGopherMenu:
		return "1"
	case ItemTypeCCSONameserver:
		return "2"
	case ItemTypeErrorCode:
		return "3"
	case ItemTypeBinHexEncodedFile:
		return "4"
	case ItemTypeDOSBinary:
		return "5"
	case ItemTypeUNIXUuencodedFile:
		return "6"
	case ItemTypeGopherFullTextSearch:
		return "7"
	case ItemTypeTelnet:
		return "8"
	case ItemTypeBinaryFile:
		return "9"
	case ItemTypeMirrorOrAlternateServer:
		return "+"
	case ItemTypeGIFFile:
		return "g"
	case ItemTypeOtherImageFile:
		return "I"
	case ItemTypeTelnet3270:
		return "T"
	case ItemTypeHTML:
		return "h"
	case ItemTypeInlineText:
		return "i"
	case ItemTypeSoundFile:
		return "s"
	default:
		return "1"
	}
}

func NewItemTypeFromURL(u *url.URL) ItemType {
	switch u.Scheme {
	case "http", "https":
		return ItemTypeHTML
	case "telnet":
		return ItemTypeTelnet
	case "tn3270": // See http://pages.upf.pf/Patrick.Capolsini/cours_internet/URLs.htm
		return ItemTypeTelnet3270
	case "file", "gopher":
		return NewItemTypeFromPath(u.Path)
	default: // Return HTML item type by default
		return ItemTypeHTML
	}
}

func NewItemTypeFromExtension(extension string) ItemType {
	switch extension {
	case "txt":
		return ItemTypeTextFile
	case "gif":
		return ItemTypeGIFFile
	case "html", "css", "js":
		return ItemTypeHTML
	case "bin", "img", "iso":
		return ItemTypeBinaryFile
	case "hqx":
		return ItemTypeBinHexEncodedFile
	case "ph":
		return ItemTypeCCSONameserver
	case "exe", "com", "dll", "sys":
		return ItemTypeDOSBinary
	case "uu", "uue":
		return ItemTypeUNIXUuencodedFile
	// See https://en.wikipedia.org/wiki/Audio_file_format
	case "3gp", "aa", "aac", "aax", "act", "aiff", "alac",
		"amr", "ape", "au", "awb", "dss", "dvf", "flac",
		"gsm", "iklax", "ivs", "m4a", "m4b", "m4p", "mmf",
		"movpkg", "mp1", "mp2", "mp3", "mpc", "msv",
		"nmf", "ogg", "opus", "ra", "raw", "rf64", "sln",
		"tta", "voc", "vox", "wav", "wma", "wv", "webm",
		"8svx", "cda":
		return ItemTypeSoundFile
	// See https://developer.mozilla.org/fr/docs/Web/Media/Guides/Formats/Image_types
	case "apng", "avif", "jpg", "jpeg", "jfif", "pjpeg",
		"pjp", "png", "svg", "webp", "bmp",
		"ico", "cur", "tif", "tiff":
		return ItemTypeOtherImageFile
	default:
		// Similar to https://github.com/khoulihan/gopher-render/blob/master/gopher_render/_parser.py#L134
		return ItemTypeGopherMenu
	}
}

func NewItemTypeFromPath(path string) ItemType {
	if strings.HasSuffix(path, "/") {
		return ItemTypeGopherMenu
	}

	pathParts := strings.Split(path, ".")
	extension := pathParts[len(pathParts)-1]

	return NewItemTypeFromExtension(extension)
}

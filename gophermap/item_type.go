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

var itemtypeToTag = map[ItemType]byte{
	ItemTypeTextFile:                '0',
	ItemTypeGopherMenu:              '1',
	ItemTypeCCSONameserver:          '2',
	ItemTypeErrorCode:               '3',
	ItemTypeBinHexEncodedFile:       '4',
	ItemTypeDOSBinary:               '5',
	ItemTypeUNIXUuencodedFile:       '6',
	ItemTypeGopherFullTextSearch:    '7',
	ItemTypeTelnet:                  '8',
	ItemTypeBinaryFile:              '9',
	ItemTypeMirrorOrAlternateServer: '+',
	ItemTypeGIFFile:                 'g',
	ItemTypeOtherImageFile:          'I',
	ItemTypeTelnet3270:              'T',
	ItemTypeHTML:                    'h',
	ItemTypeInlineText:              'i',
	ItemTypeSoundFile:               's',
}

func NewItemTypeFromURL(u *url.URL) ItemType {
	switch u.Scheme {
	case "http", "https":
		return ItemTypeHTML
	case "telnet":
		return ItemTypeTelnet
	case "tn3270": // See http://pages.upf.pf/Patrick.Capolsini/cours_internet/URLs.htm
		return ItemTypeTelnet3270
	case "file":
		return NewItemTypeFromPath(u.Path)
	default: // Return HTML item type by default
		return ItemTypeHTML
	}
}

func NewItemTypeFromPath(path string) ItemType {
	if strings.HasSuffix(path, "/") {
		return ItemTypeGopherMenu
	}

	pathParts := strings.Split(path, ".")
	extension := pathParts[len(pathParts)-1]

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

func (i *ItemType) String() string {
	return string(itemtypeToTag[*i])
}

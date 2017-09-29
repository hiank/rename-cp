package rc

import (
	"os"
	"strings"
)

// NameFilter is the Filter for code file
func NameFilter(s *string, b *bool) {

	name := (*s)[strings.LastIndexByte(*s, os.PathSeparator)+1:]
	// f := config.GetConfig().F
	
	// *b = false
	// for _, v := range f.Ignore {

	// 	if v == name {
	// 		return
	// 	}
	// }

	// for _, v := range f.IgnorePre {

	// 	if strings.HasPrefix(name, v) {
	// 		return
	// 	}
	// }

	// for _, v := range f.MatchSuf {

	// 	if strings.HasSuffix(name, v) {

	// 		*b = true
	// 		return
	// 	}
	// }
	

	switch name {
		case "iTunesArtwork":
			*b = false
		default:
			*b = true
	}

	// switch {
	// case strings.HasPrefix(name, "CC"): *b = false
	// // case strings.HasSuffix(*s, ".h"): fallthrough
	// case strings.HasSuffix(*s, ".m"): fallthrough
	// case strings.HasSuffix(*s, ".mm"): fallthrough
	// case strings.HasSuffix(*s, ".cpp"): *b = true
	// default: *b = false
	// }
}

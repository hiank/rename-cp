package rc

import (
	"github.com/hiank/core"
	"math/rand"
	"fmt"
	"strings"
	"os"
)



// RenameFile used to rename the file to rand name, return new name
func RenameFile(name string) string {
	
	idx := strings.LastIndexByte(name, os.PathSeparator) + 1
	onlyName := name[idx:]

	rn := name[:idx] + RandName(onlyName)
	err := os.Rename(name, rn)
	if err != nil {

		fmt.Println("rename file " + name + " error :" + err.Error())
		return name
	}
	return rn
}

// RandName used to rand new name
func RandName(name string) string {

	switch {
	case strings.HasPrefix(name, "."):
		return name
	}

	preArr := []string{"xgcode", "sufixa", "wyzh", "llwy"}

	var end string
	idx := strings.LastIndexByte(name, '.')
	if idx > 0 {
		end = name[idx:]
		name = name[:idx]
	}

	pre := preArr[rand.Intn(4)]
	suf := core.RandBytes(rand.Intn(5) + 2)

	return pre + "_" + name + "_" + string(suf) + end
}
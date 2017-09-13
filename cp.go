package rc

import (
	"time"
	"math/rand"
	"os"
	"fmt"
	"github.com/hiank/core"
	"strings"
)

// DuplicateDir used to copy file to dst site
func DuplicateDir(lDir string, rDir string) {

	lDir, rDir = core.AddApart(lDir), core.AddApart(rDir)
	core.DuplicateDir(lDir, rDir)

	d, err := core.NewDirInfo(rDir, nil)
	if err != nil {

		fmt.Println("read dir error :" + err.Error())
		return
	}

	rand.Seed(time.Now().UnixNano())
	for {

		name := d.NextFile()
		if name == nil {
			break
		}

		renameFile(*name)
	}
}


func renameFile(name string) {

	idx := strings.LastIndexByte(name, os.PathSeparator) + 1
	onlyName := name[idx:]

	err := os.Rename(name, name[:idx] + randName(onlyName))
	if err != nil {

		fmt.Println("rename file " + name + " error :" + err.Error())
	}
}


func randName(name string) string {

	switch {
	case strings.HasPrefix(name, "."):
		return name
	}

	preArr := []string{"xgcode", "sufixa", "wyzh", "llwy"}

	var end string
	idx := strings.LastIndexByte(name, '.')
	if idx > 1 {
		end = name[idx:]
		name = name[:idx]
	}

	pre := preArr[rand.Intn(4)]
	suf := core.RandBytes(rand.Intn(5) + 2)

	return pre + "_" + name + "_" + string(suf) + end
}

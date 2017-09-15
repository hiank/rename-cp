package rc

import (
	"time"
	"math/rand"
	"fmt"
	"github.com/hiank/core"
)

// DuplicateDirRenameFile used to copy file to dst site, and rename file, and return filename map
func DuplicateDirRenameFile(lDir string, rDir string) (fileNameMap map[string]string) {

	lDir, rDir = core.AddApart(lDir), core.AddApart(rDir)
	core.DuplicateDir(lDir, rDir)

	d, err := core.NewDirInfo(rDir, nil)
	if err != nil {

		fmt.Println("read dir error :" + err.Error())
		return
	}

	rand.Seed(time.Now().UnixNano())
	dirLen := len(d.RootPath())
	fileNameMap = make(map[string]string)
	for {

		name := d.NextFile()
		if name == "" {
			break
		}

		newName := RenameFile(name)
		fileNameMap[name[dirLen:]] = newName[dirLen:]
	}
	return
}

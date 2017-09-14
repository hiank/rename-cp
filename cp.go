package rc

import (
	"time"
	"math/rand"
	"fmt"
	"github.com/hiank/core"
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
		if name == "" {
			break
		}

		RenameFile(name)
	}
}


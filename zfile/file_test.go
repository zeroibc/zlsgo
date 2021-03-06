package zfile

import (
	"os"
	"testing"
	"time"

	. "github.com/sohaha/zlsgo"
)

func TestFile(T *testing.T) {
	t := NewTest(T)

	filePath := "../doc.go"
	tIsFile := FileExist(filePath)
	t.Equal(true, tIsFile)

	notPath := "zlsgo.php"
	status, _ := PathExist(notPath)
	t.Equal(0, status)

	size := FileSize("../doc.go")
	t.Equal("0 B" != size, true)

	size = FileSize("../_doc.go")
	t.Equal("0 B" == size, true)

	dirPath := RealPathMkdir("../zfile")
	tIsDir := DirExist(dirPath)
	t.Equal(true, tIsDir)

	dirPath = SafePath("../zfile/ok")
	t.Equal("ok", dirPath)

	path := RealPathMkdir("../tmp")
	RealPathMkdir(path + "/ooo")
	t.Log(path)
	t.Equal(true, Rmdir(path, true))
	t.Equal(true, Rmdir(path))
	_ = ProgramPath(true)
}

func TestCopy(tt *testing.T) {
	t := NewTest(tt)
	dest := RealPathMkdir("../tmp", true)
	defer Rmdir(dest)
	err := CopyFile("../doc.go", dest+"tmp.tmp")
	t.Equal(nil, err)
	err = CopyDir("../znet", dest, func(srcFilePath, destFilePath string) bool {
		return srcFilePath == "../znet/timeout/timeout.go"
	})
	t.Equal(nil, err)
}

func TestPut(t *testing.T) {
	var err error
	tt := NewTest(t)
	defer os.Remove("./text.txt")
	err = PutOffset("./text.txt", []byte(time.Now().String()+"\n"), 0)
	tt.EqualNil(err)
	err = PutAppend("./text.txt", []byte(time.Now().String()+"\n"))
	tt.EqualNil(err)
	os.Remove("./text.txt")
	err = PutAppend("./text.txt", []byte(time.Now().String()+"\n"))
	tt.EqualNil(err)
	err = PutOffset("./text.txt", []byte("\n(ok)\n"), 5)
	tt.EqualNil(err)
}

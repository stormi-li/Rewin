package rewin

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Install() {
	var rootpath string
	_, rootpath, _, _ = runtime.Caller(0)       //获取本代码文件地址
	binpath := filepath.Dir(rootpath) + "\\bin" //获取redis所在文件夹
	gopath := os.Getenv("GOPATH")               //获取gopath地址
	copyAllFiles(binpath, gopath)               //将redis复制到gopath目录下

}

func copyAllFiles(srcDir, dstDir string) error {
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path != srcDir {
			dstPath := filepath.Join(dstDir, info.Name())
			if info.Mode().IsRegular() {
				err := copyFile(path, dstPath)
				if err != nil {
					return err
				}
				fmt.Print("Install successfully -- path:" + dstPath + " -- ")
				output, _ := exec.Command("cmd", "/C", "redis-server --version").CombinedOutput()
				fmt.Print(string(output))
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	err = dstFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

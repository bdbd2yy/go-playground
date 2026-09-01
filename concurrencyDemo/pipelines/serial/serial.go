package serial

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func MD5ALL(root string) (map[string][md5.Size]byte, error) {
    m := make(map[string][md5.Size]byte)
    err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            fmt.Println(err)
            return err
        }
        if !info.Mode().IsRegular() {
            return nil
        }
        data, err := os.ReadFile(path)
        if err != nil {
            return nil
        }
        m[path] = md5.Sum(data)
        return nil
    })
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    return m, nil
}

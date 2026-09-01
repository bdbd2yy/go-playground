package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
    // exit by closing paths
    for path := range paths {
        data, err := os.ReadFile(path)
        select {
        case c <- result{path, md5.Sum(data), err}:
        // NOTE: the done is set to prevent goroutine leaks
        case <-done:
            return
        }
    }
}

func MD5ALL(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

    paths, errc := walkFiles(done, root)

    c := make(chan result)
    var wg sync.WaitGroup
    const numDigesters = 20
    wg.Add(numDigesters)
    for i := 0; i < numDigesters; i++ {
        go func() {
            digester(done, paths, c)
            wg.Done()
        }()
    }
    go func() {
        // NOTE: the wg and c escaped memory
        wg.Wait()
        close(c)
    }()

    m := make(map[string][md5.Size]byte)
    // if you don't close c, the for loop of range channel here will be blocked
    for r := range c {
        // NOTE: if MD5ALL receive the err and exit early, it will cause a goroutine leak without the done in digester.
        if r.err != nil {
            return nil, r.err
        }
        m[r.path] = r.sum
    }
    if err := <- errc; err != nil {
        return nil, err
    }
    return m, nil
}

//                     ┌── digester 1 ──┐
//                     │                │
//                     ├── digester 2 ──┤
// walkFiles → paths ──┼── digester 3 ──┼──→ c → MD5ALL → map
//                     │                │
//                     ├── ...          ┤
//                     │                │
//                     └── digester 20 ─┘

func main() {
	m, err := MD5ALL(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	// traverse the map keys
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}

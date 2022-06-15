package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)
func explorer(p string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return err
	}
	if !info.IsDir() {
		if path.Ext(p) == *ext {
			f, err := os.Open(p)
			if err != nil {
				log.Println(err)
				return nil
			}
			buf := bufio.NewReader(f)
			for {
				line, _, err := buf.ReadLine()
				if err != nil {
					break
				}
				_, err = wf.Write(line)
				if err != nil {
					break
				}
				_, _ = wf.Write([]byte("\n"))
				if count >= 3000 {
					return errors.New("finish")
				}
				count++
			}
			_ = f.Close()
		}
	}
	return nil
}

var (
	ext = flag.String("e", ".vue", "要取代码的扩展名")
)

var count = 0
var wf *bufio.Writer

func main() {
	flag.Parse()
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	f, err := os.Create(dir+"/out.txt")
	if err != nil {
		panic(err)
	}
	wf = bufio.NewWriter(f)
	log.Println("work started...")
	_ = filepath.Walk(dir, explorer)
	log.Println("work finish...")
	_ = wf.Flush()
	_ = f.Close()
}
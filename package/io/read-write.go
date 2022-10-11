package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	logpkg()
}

// 读文件
func readFile() {
	if fin, err := os.Open("go.mod"); err != nil {
		fmt.Println(err)
		return
	} else {
		defer fin.Close()
		reader := bufio.NewReader(fin)
		for {
			if line, err := reader.ReadString('\n'); err == nil {
				line = strings.TrimRight(line, "\n")
				fmt.Println(line)
			} else {
				//读到末尾
				if err == io.EOF {
					if len(line) > 0 {
						fmt.Println(line)
					}
					break
				} else {
					fmt.Println(err)
					break
				}
			}
		}
	}
}

// 写文件
func writeFile() {
	if fo, err := os.OpenFile("a.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
		return
	} else {
		defer fo.Close()
		writer := bufio.NewWriter(fo)
		writer.WriteString("我是写文件")
		//强制清空缓存，把缓存写入磁盘
		writer.Flush()
	}
}

// 创建文件或者目录
// 目录的遍历
func rangeDir(path string) error {
	if subFiles, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, file := range subFiles {
			fmt.Println(file.Name())
			if file.IsDir() {
				if err := rangeDir(filepath.Join(path, file.Name())); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// 日志
func logpkg() {
	log.Printf("%d\n", 5)
	fo, err := os.OpenFile("me.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer fo.Close()
	//日期
	logWriter := log.New(fo, "china", log.Ldate|log.Lmicroseconds)
	logWriter.Printf("%s\n", "abc")
}

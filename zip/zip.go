package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	// "strconv"
)

func main() {
	flag.Parse()

	filename := flag.Arg(0)

	if filename == "" {
		fmt.Println("Usage go-zip sourcefile")
		os.Exit(1)
	}

	fmt.Printf("Zipping file. \n")

	fmt.Println("args len", len(flag.Args()))
	filelist := make([]string, 0)
	rlist := make([]string, 0)
	for i := 0; i < len(flag.Args()); i++ {
		fmt.Println("t", i)
		filelist = append(filelist, flag.Arg(i))
		file := flag.Arg(i)

		fmt.Println("arg", file)
		filepath := strings.Replace(file, "\\", "/", -1)
		releate := path.Dir(filepath)
		fmt.Println("releate ", releate)
		rlist = append(rlist, releate)
	}
	fmt.Println("list", filelist)
	fmt.Println("releative", rlist)

	err := zipfile(filename, filelist, rlist)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Zipped to ", filename+".zip")
}

func zipfile(filename string, list []string, rlist []string) error {

	newfile, err := os.Create(filename + ".zip")
	if err != nil {
		return err
	}
	defer newfile.Close()
	zipit := zip.NewWriter(newfile)
	defer zipit.Close()

	for i, j := range list {
		fmt.Println("index ", i)
		zipfile, err := os.Open(j)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		// get the file information
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		if info.IsDir() {
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}
			header.Name = j + "/"
			fmt.Println("header folder", header.Name)

			if _, err := zipit.CreateHeader(header); err != nil {
				return err
			}
		} else {

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}
			fmt.Println("before header name", header.Name)
			if rlist[i] != "." {
				header.Name = path.Join(rlist[i], header.Name)
			}

			fmt.Println("header name", header.Name)
			writer, err := zipit.CreateHeader(header)
			if err != nil {
				return err
			}
			_, err = io.Copy(writer, zipfile)
		}

	}
	return err

}

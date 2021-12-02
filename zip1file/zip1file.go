package main

import (
	"archive/zip"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	sourceAbs, err := filepath.Abs(source)
	if info.IsDir() {
		//baseDir = filepath.Base(source)
	}

	logger.Printf("achive here : %s\n", source)
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		curAbs, err := filepath.Abs(path)
		if sourceAbs == curAbs && info.IsDir() {
			if err != nil {
				return err
			}
			logger.Printf("curAbs : [%s] rootAbs : [%s]\n", curAbs, sourceAbs)
			if curAbs == sourceAbs {
				logger.Println("root dir is skipped\n")
				return nil
			}
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		if baseDir != "" {
			//header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

var logger *log.Logger

// exam> .\zipzen.exe -path-source "e:\github\go_exam\zip1file\2021-08-12\helloworld.go" -path-dest "e:\temp"
func main() {
	//initlogger
	binName := "zip1file"
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fpLog, err := os.OpenFile(exPath+"/"+binName+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	logger = log.New(multiWriter, binName+" : ", log.LstdFlags)

	path_source := flag.String("path-source", "./", "source path")
	path_dest := flag.String("path-dest", "./", "dest path")
	//suffix := flag.String("suffix", "", "YYYY-MM-DD*suffix*.zip type for suffix (not have def separator(ex, - or _)")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	from := *path_source
	filename_ext := filepath.Base(from)
	ext := filepath.Ext(filename_ext)
	filename := strings.TrimSuffix(filename_ext, ext)
	to := *path_dest + "/" + filename + ".zip"
	logger.Printf("source path : %s\n", from)
	logger.Printf("output file : %s\n", to)

	//check it
	if _, err := os.Stat(from); os.IsNotExist(err) {
		logger.Printf("source path : %s is not exist.\n", from)
		logger.Println("exit program!")
		return
	}
	if _, err := os.Stat(*path_dest); os.IsNotExist(err) {
		logger.Printf("dest path : %s is not exist.\n", *path_dest)
		logger.Println("exit program!")
		return
	}

	err = zipit(from, to)
	if err == nil {
		logger.Printf("file achive successful dest[%s] source[%s]\n", to, from)
	} else {
		panic(err)
	}
}

package main

import (
	"archive/zip"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
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
		baseDir = filepath.Base(source)
	}

	logger.Printf("achive here : %s\n", source)
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if baseDir == info.Name() && info.IsDir() {
			curAbs, err := filepath.Abs(path)
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

func main() {
	//initlogger
	logger = log.New(os.Stdout, "fileAppender,: ", log.LstdFlags)

	path_source := flag.String("path-source", "./", "source path")
	path_dest := flag.String("path-dest", "./", "dest path")
	suffix := flag.String("suffix", "", "YYYY-MM-DD-*suffix*.zip type for suffix")

	offsetDay := flag.Int("offsetDay", -1, "day offset : int")
	offsetMon := flag.Int("offsetMon", 0, "mon offset : int (default 0)")
	offsetYear := flag.Int("offsetYear", 0, "year offset : int (default 0)")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	now := time.Now().AddDate(*offsetYear, *offsetMon, *offsetDay).Format("2006-01-02")
	logger.Println(now)
	logger.Printf("source path : %s\n", *path_source)
	logger.Printf("output file : %s\n", now+*suffix+".zip")
	err := zipit(*path_source+now+"/", *path_dest+now+*suffix+".zip")
	if err == nil {
		logger.Printf("file achive successful dest[%s] source[%s]\n", *path_dest+now+*suffix+".zip", *path_source+now)
	} else {
		panic(err)
	}
}

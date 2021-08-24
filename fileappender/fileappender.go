package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func fileAppender(source, target string) error {
	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir, err = filepath.Abs(source)
		if err != nil {
			panic(err)
			return err
		}
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		currentname := filepath.Join(baseDir, info.Name())

		if info.IsDir() == false && strings.HasSuffix(currentname, target) == false {
			log.Printf("baseDir: %s \n", baseDir)
			log.Printf("path: %s \n", path)
			log.Printf("filename: %s : size: %d\n", currentname, info.Size())

			//file, err := os.OpenFile(currentname, os.O_RDONLY|os.O_RDWR|os.O_TRUNC, 0644)
			file, err := os.Open(currentname)

			if err != nil {
				panic(err)
			}
			defer file.Close()

			buf := make([]byte, info.Size())
			n := 0
			n, err = file.Read(buf)
			if err != nil {
				panic(err)
			}

			logger.Printf("read file length : %d \n", n)
			//logger.Printf("%s", buf)

			removeindex := strings.LastIndex(currentname, target) + 4
			logger.Printf("find %s from %s - removefrom :%d\n", target, currentname, removeindex)
			pairFileName := currentname[:removeindex]

			if _, err := os.Stat(pairFileName); os.IsNotExist(err) {
				logger.Printf("pairFile path : %s is not exist.\n", pairFileName)
				logger.Println("exit program!")
				return
			}

			logger.Printf("pairFile: %s \n", pairFileName)
			pairFile, err := os.OpenFile(pairFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

			if err != nil {
				panic(err)
			}
			defer pairFile.Close()

			bufWriter := bufio.NewWriter(pairFile)

			currentPos, err := pairFile.Seek(0, 2)

			if err != nil {
				logger.Panicln("unable to seek to the end of")
				os.Exit(3)
			}
			logger.Printf("currnetseekpos: %d \n", currentPos)

			//cnt, err := pairFile.Write(buf)
			cnt, err := bufWriter.Write(buf)

			logger.Printf("write size: %d \n", cnt)

			if err != nil {
				panic(err)
			}
			bufWriter.Flush()

			if err != nil {
				panic(err)
			}

			file.Close()

			err = os.Remove(currentname)

			if err != nil {
				panic(err)
			}
		}

		return err
	})

	return err
}

var logger *log.Logger

func main() {
	//initlogger
	logger = log.New(os.Stdout, "fileAppender,: ", log.LstdFlags)

	path_source := flag.String("path-source", "./", "source path")
	suffix := flag.String("suffix", "", ".log")

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

	err := fileAppender(*path_source+"/"+now, *suffix)
	if err == nil {
	} else {
		panic(err)
	}
}

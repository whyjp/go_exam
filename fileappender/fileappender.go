package main

import (
	"bufio"
	"flag"
	"io"
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
		}
	}
	logger.Printf("baseDir: %s \n", baseDir)

	bFind := false

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		currentname := filepath.Join(baseDir, info.Name())

		if info.IsDir() || strings.HasSuffix(currentname, target) {
			return nil
		}
		bFind = true
		logger.Printf("path: %s \n", path)
		logger.Printf("rolledfile name is : %s : size : %d\n", currentname, info.Size())

		rolledFileName := currentname
		rolledFP, err := os.OpenFile(rolledFileName, os.O_APPEND|os.O_WRONLY, 0600)

		if err != nil {
			panic(err)
		}
		defer rolledFP.Close()

		currentPos, err := rolledFP.Seek(0, 2)
		if err != nil {
			logger.Panicln("unable to seek to the end of")
			os.Exit(3)
		}
		logger.Printf("currnetseekpos: %d \n", currentPos)
		bufWriter := bufio.NewWriter(rolledFP)

		removeindex := strings.LastIndex(rolledFileName, target) + len(target)
		logFileName := rolledFileName[:removeindex]

		if _, err := os.Stat(logFileName); os.IsNotExist(err) {
			logger.Printf("log file path : %s is not exist.\n", logFileName)
			logger.Println("exit program!")
			return nil
		}

		logger.Printf("log fileName is : %s \n", logFileName)

		logFP, err := os.OpenFile(logFileName, os.O_RDONLY|os.O_RDWR, 0644)

		if err != nil {
			panic(err)
		}
		defer logFP.Close()

		bufReader := bufio.NewReader(logFP)

		logger.Printf("log file [%s] append to rolled file [%s]'s end \n", logFileName, rolledFileName)
		written, err := io.Copy(bufWriter, bufReader)

		logger.Printf("write size: %d \n", written)
		if err != nil {
			panic(err)
		}

		logFP.Close()
		logger.Printf("remove log file [%s]\n", logFileName)
		err = os.Remove(logFileName)

		if err != nil {
			panic(err)
		}
		rolledFP.Close()
		logger.Printf("rolled filename [%s] is renamed log filename [%s] \n", rolledFileName, logFileName)
		err = os.Rename(rolledFileName, logFileName)

		if err != nil {
			panic(err)
		}

		return err
	})
	if !bFind {
		logger.Printf("can not found log rolled file \n")
	}
	return err
}

var logger *log.Logger

func main() {
	//initlogger
	binName := "fileAppender"
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

	from := *path_source + "/" + now + "/"
	logger.Printf("source path : %s\n", from)

	if _, err := os.Stat(from); os.IsNotExist(err) {
		logger.Printf("source path : %s is not exist.\n", from)
		logger.Println("exit program!")
		return
	}

	err = fileAppender(from, *suffix)
	if err != nil {
		panic(err)
	}
}

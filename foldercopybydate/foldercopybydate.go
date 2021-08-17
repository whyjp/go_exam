package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

var logger *log.Logger

func main() {
	//initlogger
	logger = log.New(os.Stdout, "foldercopybydate,: ", log.LstdFlags)

	path_source := flag.String("path-source", "./", "source path")
	path_dest := flag.String("path-dest", "./", "dest path")

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
	to := *path_dest + "/" + now + "/"
	logger.Printf("copy source folder [%s] \n", from)
	logger.Printf("to [%s] \n", to)

	//check it
	if _, err := os.Stat(from); os.IsNotExist(err) {
		logger.Printf("source path : %s is not exist.\n", from)
		logger.Println("exit program!")
		return
	}
	if _, err := os.Stat(to); os.IsNotExist(err) {
		logger.Printf("dest path : %s is not exist.\n", from)
		logger.Println("exit program!")
		return
	}

	err := CopyDir(from, to)

	if err == nil {
	} else {
		panic(err)
	}
}

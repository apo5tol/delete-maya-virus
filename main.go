package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/delete-maya-virus/utils"
)

func pressAnyKey() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	recursive := flag.Bool("r", false, "recursive path traversal")
	createBackup := flag.Bool("b", false, "make a backup copy of processed files")

	flag.Parse()
	argPath := flag.Arg(0)

	defer pressAnyKey()

	if argPath == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err.Error())
		}
		argPath = currentDir
	}

	fi, err := os.Stat(argPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	var scannedMayaFiles []string

	switch mode := fi.Mode(); {
	case mode.IsDir():
		if *recursive {
			files, err := utils.ReturnMayaFilesFromDirRecursively(argPath)
			if err != nil {
				log.Fatal(err.Error())
			}
			scannedMayaFiles = files
		} else {
			files, err := utils.ReturnMayaFilesFromDir(argPath)
			if err != nil {
				log.Fatal(err.Error())
			}
			scannedMayaFiles = files
		}
	case mode.IsRegular():
		fileExtension := filepath.Ext(argPath)
		if fileExtension != utils.MaFileExt {
			log.Println("Only .ma files can be processed!")
			return
		}
		scannedMayaFiles = append(scannedMayaFiles, argPath)
	}

	for _, pathToMayaScannedFile := range scannedMayaFiles {
		mayaFileName := filepath.Base(pathToMayaScannedFile)
		backupPath := pathToMayaScannedFile + ".backup"
		mayaFileData, err := ioutil.ReadFile(pathToMayaScannedFile)
		if err != nil {
			log.Fatal("Can`t open maya file! -", mayaFileName, err.Error())
		}
		log.Println("Checking the file -", mayaFileName)

		isVaccineVirusInFile := utils.CheckVaccineVirus(mayaFileData)
		isMayaMelVirusInFile := utils.CheckMayaMelVirus(mayaFileData)

		if isVaccineVirusInFile {
			log.Println("Vaccine virus found! Processing file... ", mayaFileName)
			if *createBackup {
				err := utils.CreateBuckup(backupPath, mayaFileData)
				if err != nil {
					log.Fatal(err.Error())
				}
			}
			mayaFileData = utils.DeleteVaccineVirus(mayaFileData)
		}

		if isMayaMelVirusInFile {
			log.Println("MayaMelUIConfigurationFile virus found! Processing file... ", mayaFileName)
			if *createBackup {
				err := utils.CreateBuckup(backupPath, mayaFileData)
				if err != nil {
					log.Fatal(err.Error())
				}
			}
			mayaFileData = utils.DeleteMayaMelVirus(mayaFileData)
		}

		writeErr := ioutil.WriteFile(pathToMayaScannedFile, mayaFileData, utils.FilePermission)
		if writeErr != nil {
			log.Fatal(writeErr.Error())
		}
	}
	log.Println("All files are checked.")
}

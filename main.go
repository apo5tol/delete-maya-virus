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
		files, err := returnMayaFilesFromDir(argPath)
		if err != nil {
			log.Fatal(err)
		}
		scannedMayaFiles = files
	case mode.IsRegular():
		fileExtension := filepath.Ext(argPath)
		if fileExtension != maFileExt {
			log.Println("Only .ma files can be processed!")
			return
		}
		scannedMayaFiles = append(scannedMayaFiles, argPath)
	}

	for _, pathToMayaScannedFile := range scannedMayaFiles {
		mayaFileName := filepath.Base(pathToMayaScannedFile)
		mayaFileData, err := ioutil.ReadFile(pathToMayaScannedFile)
		if err != nil {
			log.Fatal("Can`t open maya file! -", mayaFileName, err.Error())
		}
		log.Println("Checking the file -", mayaFileName)

		isVaccineVirusInFile := utils.CheckVaccineVirus(mayaFileData)
		isMayaMelVirusInFile := utils.CheckMayaMelVirus(mayaFileData)

		if isVaccineVirusInFile {
			log.Println("Vaccine virus found! Processing file... ", mayaFileName)
			mayaFileData = utils.DeleteVaccineVirus(mayaFileData)
		}

		if isMayaMelVirusInFile {
			log.Println("MayaMelUIConfigurationFile virus found! Processing file... ", mayaFileName)
			mayaFileData = utils.DeleteMayaMelVirus(mayaFileData)
		}

		err = ioutil.WriteFile(pathToMayaScannedFile, mayaFileData, filePermission)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	log.Println("All files are checked.")
}

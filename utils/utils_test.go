package utils

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

const pathToDataVaccineVirus = "../testData/dataVaccineVirus.ma"
const pathToDataWithoutVaccineVirus = "../testData/dataWithouVaccineVirus.ma"
const pathToMayaFiles = "../testData/"
const pathToMayaFilesRec = "../"

func TestCheckVaccineVirusTrue(t *testing.T) {
	mayaFile, err := ioutil.ReadFile(pathToDataVaccineVirus)
	if err != nil {
		t.Error(err.Error())
	}
	isViirusInFile := CheckVaccineVirus(mayaFile)
	if !isViirusInFile {
		t.Error("Expected true, got", isViirusInFile)
	}
}

func TestCheckVaccineVirusFalse(t *testing.T) {
	mayaFile, err := ioutil.ReadFile(pathToDataWithoutVaccineVirus)
	if err != nil {
		t.Error(err.Error())
	}
	isViirusInFile := CheckVaccineVirus(mayaFile)
	if isViirusInFile {
		t.Error("Expected false, got", isViirusInFile)
	}
}

func TestDeleteVaccineVirus(t *testing.T) {
	mayaFileData, err := ioutil.ReadFile(pathToDataVaccineVirus)
	if err != nil {
		t.Error(err.Error())
	}
	clearData := DeleteVaccineVirus(mayaFileData)
	isViirusInFile := CheckVaccineVirus(clearData)
	if isViirusInFile {
		t.Error("Expected true, got", isViirusInFile)
	}
}

func TestReturnMayaFilesFromDirCount(t *testing.T) {
	mayaFiles, err := ReturnMayaFilesFromDir(pathToMayaFiles)
	if err != nil {
		t.Error(err)
	}
	lenMayaFiles := len(mayaFiles)
	if lenMayaFiles != 2 {
		t.Error("Expected 2, got", lenMayaFiles)
	}
}

func TestReturnMayaFilesFromDirOnlyMa(t *testing.T) {
	mayaFiles, err := ReturnMayaFilesFromDir(pathToMayaFiles)
	if err != nil {
		t.Error(err)
	}
	for _, pathToFile := range mayaFiles {
		fileExtension := filepath.Ext(pathToFile)
		if fileExtension != ".ma" {
			t.Error("Expected .ma file, got", pathToFile, fileExtension)
		}
	}
}

func TestReturnMayaFilesFromDirRecursively(t *testing.T) {
	mayaFiles, err := ReturnMayaFilesFromDirRecursively(pathToMayaFilesRec)
	if err != nil {
		t.Error(err)
	}
	lenMayaFiles := len(mayaFiles)
	if lenMayaFiles != 2 {
		t.Error("Expected 2, got", lenMayaFiles)
	}
}

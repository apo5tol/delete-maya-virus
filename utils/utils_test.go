package utils

import (
	"io/ioutil"
	"testing"
)

const pathToDataVaccineVirus = "../testData/dataVaccineVirus.ma"
const pathToDataWithoutVaccineVirus = "../testData/dataWithouVaccineVirus.ma"

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

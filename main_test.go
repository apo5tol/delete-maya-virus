package main

import (
	"path/filepath"
	"testing"
)

const pathToMayaFiles = "testData/"

func TestReturnMayaFilesFromDirCount(t *testing.T) {
	mayaFiles, err := returnMayaFilesFromDir(pathToMayaFiles)
	if err != nil {
		t.Error(err)
	}
	lenMayaFiles := len(mayaFiles)
	if lenMayaFiles != 2 {
		t.Error("Expected 2, got", lenMayaFiles)
	}
}

func TestReturnMayaFilesFromDirOnlyMa(t *testing.T) {
	mayaFiles, err := returnMayaFilesFromDir(pathToMayaFiles)
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

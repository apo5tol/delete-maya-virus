package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

var emptySymbol = []byte("")

var vaccineVirusTemplate = []byte("createNode script -n \"vaccine_gene\";")
var vaccineGeneNodeRegexp = regexp.MustCompile("createNode script -n \"vaccine_gene\";[\\s\\S]+leukocyte\\.antivirus\\(\\)(.+)?\\s")
var breedGeneNodeRegexp = regexp.MustCompile("createNode script -n \"breed_gene\";[\\s\\S]+setAttr \".stp\" 1;(.+)?\\s")

var mayaMelVirusTemplate = []byte("MayaMelUIConfigurationFile")
var mayaMelVirusNodeRegexp = regexp.MustCompile("createNode script -n \"MayaMelUIConfigurationFile\"[\\s\\S]+setAttr \".st\" 1;(.+)?\\s")

// MaFileExt is extension maya file
const MaFileExt = ".ma"

// FilePermission is permissions for file
const FilePermission = 0644

// DeleteVaccineVirus removes the Vaccine node from the input data.
func DeleteVaccineVirus(mayaFile []byte) []byte {
	out := vaccineGeneNodeRegexp.ReplaceAll(mayaFile, emptySymbol)
	out = breedGeneNodeRegexp.ReplaceAll(out, emptySymbol)
	return out
}

// DeleteMayaMelVirus removes the MayaMelUIConfigurationFile node from the input data.
func DeleteMayaMelVirus(mayaFile []byte) []byte {
	out := mayaMelVirusNodeRegexp.ReplaceAll(mayaFile, emptySymbol)
	return out
}

// CheckVaccineVirus check if the input data contains Vaccine node.
func CheckVaccineVirus(mayaFile []byte) bool {
	return bytes.Contains(mayaFile, vaccineVirusTemplate)
}

// CheckMayaMelVirus check if the input data contains MayaMelUIConfigurationFile node.
func CheckMayaMelVirus(mayaFile []byte) bool {
	return bytes.Contains(mayaFile, mayaMelVirusTemplate)
}

// ReturnMayaFilesFromDirRecursively traverse the directory recursively and only return .ma files
func ReturnMayaFilesFromDirRecursively(rootPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileExtension := filepath.Ext(path)
			if fileExtension == MaFileExt {
				files = append(files, path)
			}
		}
		return nil
	})
	return files, err
}

// ReturnMayaFilesFromDir returns .ma files from directory
func ReturnMayaFilesFromDir(rootPath string) ([]string, error) {
	var files []string
	dirFiles, err := ioutil.ReadDir(rootPath)
	if err != nil {
		return files, err
	}
	for _, dirFile := range dirFiles {
		if !dirFile.IsDir() {
			filePath := filepath.Join(rootPath, dirFile.Name())
			fileExtension := filepath.Ext(filePath)
			if fileExtension == ".ma" {
				files = append(files, filePath)
			}
		}
	}
	return files, nil
}

// CreateBuckup create backup of processed the maya file
func CreateBuckup(backupPath string, backupData []byte) error {
	err := ioutil.WriteFile(backupPath, backupData, FilePermission)
	return err
}

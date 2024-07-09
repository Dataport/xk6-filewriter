package fileWriter

import (
	"errors"
	"fmt"
	"go.k6.io/k6/js/modules"
	"os"
)

type FileWriter struct{}

func init() {
	modules.Register("k6/x/filewriter", new(FileWriter))
}

func (d *FileWriter) WriteString(path string, filename string, s string) error {
	//check if path exists
	_, err := os.Stat(path)
	pathString := fmt.Sprintf("%s%s", path, filename)
	if err != nil {
		os.MkdirAll(path, 0750)
		file, pErr := os.Create(pathString)
		if pErr != nil {
			var errorsString string = fmt.Sprintf("Path with file: %s can not be created successfully", pathString)
			return errors.New(errorsString)
		}
		file.WriteString(s)
		return nil
	} else {
		err := d.AppendString(path, filename, s)
		return err
	}
}

func (d *FileWriter) AppendString(path string, filename string, s string) error {
	pathString := fmt.Sprintf("%s%s", path, filename)
	f, err := os.OpenFile(pathString,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0750)
	if err != nil {
		return err
	}
	defer f.Close()
	newLineString := fmt.Sprintf("\n%s", s)
	_, writeErr := f.WriteString(newLineString)
	if writeErr != nil {
		return writeErr
	}
	return nil
}


// note: no check if other fileWriter threads are running
// note: on Windows write protected files can be overwritten when not protected against deletion
func (d *FileWriter) CreateFile(path string, filename string) error {
	
	// check and create path
	_, err := os.Stat(path)	
	if err != nil {
		os.MkdirAll(path, 0750)
	}
	
	filePath := fmt.Sprintf("%s%s", path, filename)
	
	// 	delete existing file	
	_, err = os.Stat(filePath);
	if !errors.Is(err, os.ErrNotExist) {
		err := os.Remove(filePath)
		if(err != nil){
			var errorsString string = fmt.Sprintf("File %s cannot be removed", filePath)
			return errors.New(errorsString)
		}
	}
	
	// create new file
	_, err = os.Create(filePath)
	if err != nil {
		var errorsString string = fmt.Sprintf("File %s cannot be created successfully", filePath)
		return errors.New(errorsString)
	}	
		
	return nil	
}
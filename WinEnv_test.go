package winenv

import (
	"errors"
	"os"
	"testing"
)

func TestShouldReadContentsOfGivenFile(t *testing.T) {
	vars := map[string]string{
		"FOO":  "BAR",
		"TEST": "CASE",
		"INT":  "1",
	}
	fileName := "./.sample_env"

	// Create the ENV File(s)
	err := helperSetupFile(fileName, vars)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(fileName)

	paths := []string{fileName}
	// Parse the file
	if err := Parse(paths...); err != nil {
		t.Error(err)
	}

	// make sure that all of the variables are in the os.ENV scope
	for key, val := range vars {
		if envVal := os.Getenv(key); val != envVal {
			t.Errorf("Invalid ENV VAR; got %s expected %s", envVal, val)
		}
	}
}

func TestShouldReadContentsOfMultipleFiles(t *testing.T) {
	var1 := map[string]string{
		"FOO":  "BAR",
		"TEST": "CASE",
		"INT":  "1",
	}
	var1Name := "./.sample1_env"

	var2 := map[string]string{
		"FIRST_NAME": "Tony",
		"LAST_NAME":  "Danza",
		"AGE":        "19",
	}
	var2Name := "./.sample2_env"

	// Create the ENV File(s)
	err := helperSetupFile(var1Name, var1)
	if err != nil {
		t.Error(err)
	}
	err = helperSetupFile(var2Name, var2)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(var1Name)
	defer os.Remove(var2Name)

	paths := []string{var1Name, var2Name}
	// Parse the file
	if err := Parse(paths...); err != nil {
		t.Error(err)
	}

	// make sure that all of the variables are in the os.ENV scope
	for key, val := range var1 {
		if envVal := os.Getenv(key); val != envVal {
			t.Errorf("Invalid ENV VAR; got %s expected %s", envVal, val)
		}
	}

	for key, val := range var2 {
		if envVal := os.Getenv(key); val != envVal {
			t.Errorf("Invalid ENV VAR; got %s expected %s", envVal, val)
		}
	}
}

func TestShouldGetContentsOfDefaultFile(t *testing.T) {
	expected := map[string]string{
		"FIRST_NAME": "John",
		"LAST_NAME":  "Doe",
		"AGE":        "18",
	}

	paths := []string{}
	if err := Parse(paths...); err != nil {
		t.Error(err)
	}

	// make sure that all of the variables are in the os.ENV scope
	for key, val := range expected {
		if envVal := os.Getenv(key); val != envVal {
			t.Errorf("Invalid ENV VAR; got %s expected %s", envVal, val)
		}
	}
}

func TestShouldFailIfFileDoesNotExist(t *testing.T) {
	paths := []string{"./.not_existant_file"}
	if err := Parse(paths...); err == nil {
		t.Error("Error should not be nil if given file doesn't exists")
	}
}

// given a filename and map of key/val pairs, create a file for use with test functions
func helperSetupFile(filePath string, vars map[string]string) error {
	// make sure the file doesn't already exist
	if _, err := os.Stat(filePath); os.IsExist(err) {
		return errors.New("File Already exists, Please remove it to properly run test")
	}

	// if it doesn't, open a writable stream
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var loc int64
	// place each key/val into the file
	for key, val := range vars {
		// Convert key/val into []byte
		bs := []byte(key + "=" + val + "\n")
		bytes, err := file.WriteAt(bs, loc)
		if err != nil {
			return err
		}
		loc += int64(bytes)
	}

	return nil
}

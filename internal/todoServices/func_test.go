package todoServices

import (
	"os"
	fileStruct "projMod/internal/models"
	"reflect"
	"testing"
)

const fileName = "FileForTests" // not correct, no sufx ".json"

// simple test for all functions (tst = testing, test, or another with "test")
func TestFunctionality(t *testing.T) {
	var err error

	// Clean before testing
	err = os.Remove("FileForTests.json")
	if err != nil {
		t.Fatalf("CleaningSys dead: %v", err)
	}

	tstDataForFile := []fileStruct.FileDataStruct{
		{
			TaskNum: 1, ShortName: "Shopping", Comment: "Oh my dog!",
		},
		{
			TaskNum: 2, ShortName: "Lopping", Comment: "Oh my Lop!",
		},
		{
			TaskNum: 3, ShortName: "Joggin", Comment: "Oh my jog!",
		},
	}

	//tst for todoAdd
	for _, val := range tstDataForFile {
		err = AddTask(val.TaskNum, val.ShortName, val.Comment, fileName)
		if err != nil {
			t.Fatalf("AddTasktst GG: %v", err)
		}
	}

	// tst for readFile
	dataFromFile, err := ReadFromFile(fileName)
	if err != nil {
		t.Fatalf("ReadFiletst 1 GG: %v ", err)
	}
	if !reflect.DeepEqual(dataFromFile, tstDataForFile) {
		t.Fatalf("ReadFiletst 2 GG: In createdfile: %v , In codeCurrData: %v ", dataFromFile, tstDataForFile)
	}

	t.Logf("ReadFiletst all Ok, data from file: %v", dataFromFile)

	// tst for todoDelete
	DeleteTask(1, fileName)
	expectedData, err := ReadFromFile(fileName)
	if err != nil {
		t.Fatalf("ReadFiletst 1 GG: %v ", err)
	}

	if !reflect.DeepEqual(tstDataForFile[1:], expectedData) {
		t.Fatalf("DeleteTasktst GG: In file after delete: %v , In codeCurrData: %v ", expectedData, tstDataForFile[1:])
	}
	t.Logf("DeleteTasktst all ok, data from file: %v", expectedData)

	// Close file

	// Clean after testing
	err = os.Remove("FileForTests.json")
	if err != nil {
		t.Fatalf("CleaningSys dead: %v", err)
	}
}

// Ill try to create tests for all functions in rhis package
// but can i do it in one file, or better to do it in different
// And whats better to do a lot of tests for different cases, or
// do one - two and stop hmmmmm

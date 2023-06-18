package shredder

import (
	"io/ioutil"
	"os"
	"testing"
)

/*
Not 100% certain which is more correct either to write test cases as individual functions
or as 1 function, due to size of test cases choose to write as "1" function but ill add that
should these test cases become more complex or larger or perhaps even more tests be added, would 
probably separate out into multiple functions for better formating
*/
func TestShred(t *testing.T) {

	// small file test
	t.Run("Small File Test", func(t *testing.T) {
		path := "test.txt"
		err := ioutil.WriteFile(path, []byte("Random Text"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		err = Shred(path)
		if err != nil {
			t.Fatal(err)
		}

		_, err = os.Stat(path)
		if !os.IsNotExist(err) {
			t.Fatalf("expected error")
		}
	})

	// large file test
	t.Run("Large File Test", func(t *testing.T) {
		path := "test.bin"
		// 1MB is not actually that much data but is relatively large 
		data := make([]byte, 1024*1024)
		err := ioutil.WriteFile(path, data, 0644)
		if err != nil {
			t.Fatal(err)
		}

		err = Shred(path)
		if err != nil {
			t.Fatal(err)
		}

		_, err = os.Stat(path)
		if !os.IsNotExist(err) {
			t.Fatalf("expected error")
		}
	})

	// file path doesnt exist test
	t.Run("Path Doesnt Exist Test", func(t *testing.T) {
		path := "test.txt"
		err := Shred(path)
		if err == nil {
			t.Fatalf("expected error")
		}
	})

	// directory instead of file test
	t.Run("Directory instead of file Test", func(t *testing.T) {
		path := "testdir"
		err := os.Mkdir(path, 0755)
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(path)
		err = Shred(path)
		if err == nil {
			t.Fatalf("expected error")
		}
	})

	// file does not have write permissions test
	t.Run("Non Writtable File Test", func(t *testing.T) {
		path := "test.txt"
		err := ioutil.WriteFile(path, []byte("I Am Write Protected :) Fight Me"), 0400)
		if err != nil {
			t.Fatal(err)
		}
		defer os.Chmod(path, 0644)
		defer os.Remove(path)
		err = Shred(path)
		if err == nil {
			t.Fatalf("expected error")
		}
	})
}

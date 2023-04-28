package fsx

import "testing"

func Test_Zip(t *testing.T) {
	fs, err := New("../../db")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("PATH: ", fs.Path())

	err = fs.Zip("../../tmp/os.zip")
	if err != nil {
		t.Fatal(err)
	}

	fs, err = New("../../tmp/os.zip")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("PATH: ", fs.Path())
	fs.Unzip("../../tmp")
}

func Test_Base(t *testing.T) {
	fs, err := New("../../")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("PATH: ", fs.Base())
}

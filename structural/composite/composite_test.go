package composite

import "testing"

func TestComposite(t *testing.T) {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}
	folder1 := &folder{
		name: "Folder1",
	}
	folder1.add(file1)

	folder2 := &folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")

	/*
		Searching recursively for keyword rose in folder Folder2
		Searching for keyword rose in file File2
		Searching for keyword rose in file File3
		Searching recursively for keyword rose in folder Folder1
		Searching for keyword rose in file File1
	*/
}

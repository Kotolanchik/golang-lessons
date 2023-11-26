package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func renameFile(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Ошибка при чтении", err)
	}

	re := regexp.MustCompile(`^task-(\d+)$`)

	for _, file := range files {
		if !file.IsDir() {
			matches := re.FindStringSubmatch(file.Name())
			if len(matches) == 2 {
				num, err := strconv.Atoi(matches[1])
				if err == nil {
					newName := fmt.Sprintf("task-%02d", num)
					err = os.Rename(filepath.Join(path, file.Name()), filepath.Join(path, newName)+".go")
					if err != nil {
						fmt.Printf("Ошибка при переименовании файла %s: %v\n", file.Name(), err)
					} else {
						fmt.Printf("Файл %s успешно переименован в %s\n", file.Name(), newName)
					}
				}
			}
		}
	}
}

func createFile(path string, start, end int) {
	for i := start; i < end; i++ {
		filename := fmt.Sprintf("task-%d.go", i)

		file, err := os.Create(path + filename)
		if err != nil {
			fmt.Printf("Ошибка при создании файла %s: %v\n", filename, err)
			return
		}

		defer file.Close()

		content := []byte(fmt.Sprintf("package main\n\nimport \"fmt\"\n\nfunc main() {\n}\n"))
		err = ioutil.WriteFile(filename, content, 0644)
		if err != nil {
			fmt.Printf("Ошибка при записи в файл %s: %v\n", filename, err)
			return
		}
		fmt.Printf("Файл %s успешно создан.\n", filename)
	}
}

func main() {
	path := "C:\\Users\\Konstantin\\Desktop\\e-commerce\\golang-lessons\\module-2\\"
	//renameFile(path)
	createFile(path, 12, 18)
}

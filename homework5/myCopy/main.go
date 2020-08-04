package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func copyPartsFile(s, d string) error {
	file, err := os.Open(s)

	if err != nil {
		return err
	}
	defer file.Close()

	file1, err := os.Create(d)
	defer file1.Close()
	if err != nil {
		return err
	}
	data := make([]byte, 1024)

	for {
		n, err := file.Read(data)

		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}

		_, err = file1.Write(data[:n])
		if err != nil {
			return err
		}

	}
	return nil
}

func copyFullFile(s, d string) error {
	str, err := ioutil.ReadFile(s)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(d, str, 0777)
	if err != nil {
		return err
	}
	return nil
}
func copyFile(s, d string) error {
	file, _ := os.Open(s)
	stat, err := file.Stat()
	if stat != nil {
		return err
	}
	defer file.Close()
	if stat.Size() > 1024 {
		err = copyPartsFile(s, d)
		if err != nil {
			return err
		}
	} else {
		err = copyFullFile(s, d)
		if err != nil {
			return err
		}
	}

	return nil
}

func copyDir(s, d string) error {

	sInfo, err := os.Stat(s)
	if err != nil {
		return err
	}
	err = os.MkdirAll(d, sInfo.Mode())
	if err != nil {
		return err
	}

	dir, _ := os.Open(s)

	objects, err := dir.Readdir(-1)

	for _, obj := range objects {

		sName := s + "/" + obj.Name()

		dName := d + "/" + obj.Name()

		if obj.IsDir() {

			err = copyDir(sName, dName)
			if err != nil {
				return err
			}
		} else {

			err = copyFile(sName, dName)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func main() {
	flag.Parse()
	source := flag.Arg(0)
	dest := flag.Arg(1)
	src, err := os.Stat(source)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Open(dest)
	if !os.IsNotExist(err) {
		fmt.Println("Обьект копирования существует. Переписать? Y - да, N - нет (Y)")
		var n string
		fmt.Scanln(&n)

		if strings.ToLower(n) == "Y" || strings.ToLower(n) == "" {
			if !src.IsDir() {
				err = copyFile(source, dest)
				if err != nil {
					log.Fatal(err)
				}

			} else {
				err = copyDir(source, dest)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}

	fmt.Println("Копирование завершено")

}

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

type repository struct{}

func (r repository) getAlbums() ([]album, error) {
	filename := "albums.json"
	err := checkFile(filename)
	if err != nil {
		logrus.Error(err)
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	var albums []album

	json.Unmarshal(file, &albums)

	return albums, nil
}

func (r repository) addAlbum(element album) error {
	filename := "albums.json"

	oldAlbums, err := r.getAlbums()
	newAlbums := append(oldAlbums, element)

	bytes, err := json.Marshal(newAlbums)

	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		logrus.Error(err)
	}

	return err
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		return err
	}
	return nil
}

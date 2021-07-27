package infra

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"

	"study-webapi/domain"
)

type Repository struct{}

func (r Repository) GetAlbums() ([]domain.Album, error) {
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

	var albums []domain.Album

	json.Unmarshal(file, &albums)

	return albums, nil
}

func (r Repository) AddAlbum(element domain.Album) error {
	filename := "albums.json"

	oldAlbums, err := r.GetAlbums()
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

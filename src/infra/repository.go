package infra

import (
	"encoding/json"
	"io/ioutil"
	"os"

	uuid "github.com/satori/go.uuid"
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

	uniqueId := generateUniqueId(oldAlbums)
	element.Id = uniqueId
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

func generateUniqueId(albums []domain.Album) string {
	var uniqueId string
	for uniqueId == "" {
		testId := uuid.NewV4().String()
		containsId := containsId(albums, testId)

		if !containsId {
			uniqueId = testId
		}
	}
	return uniqueId
}

func containsId(albums []domain.Album, id string) bool {
	for _, album := range albums {
		if album.Id == id {
			return true
		}
	}
	return false
}

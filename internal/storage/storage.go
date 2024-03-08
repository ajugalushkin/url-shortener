package storage

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

const file string = "internal/db/store.db"

type URLData struct {
	Key string
	Url string
}

func Create(url URLData) error {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return err
	}
	defer db.Close()

	_, errCreate := db.Exec("insert into urls (key,url) values ($1,$2)", url.Key, url.Url)
	if errCreate != nil {
		return errCreate
	}

	return nil
}

func GetByUrl(url string) (URLData, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return URLData{}, err
	}
	defer db.Close()

	rows, err := db.Query("select * from urls where url = $1", url)
	if err != nil {
		return URLData{}, err
	}
	defer rows.Close()

	urlData := URLData{}
	for rows.Next() {
		err := rows.Scan(&urlData.Key, &urlData.Url)
		if err != nil {
			return URLData{}, err
		}
		break
	}

	if urlData == (URLData{}) {
		return URLData{}, errors.New("Key not found!")
	}

	return urlData, nil
}

func GetByKey(key string) (string, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return "", err
	}
	defer db.Close()

	rows, err := db.Query("select * from urls where key = $1", key)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	urlData := URLData{}
	for rows.Next() {
		err := rows.Scan(&urlData.Key, &urlData.Url)
		if err != nil {
			return "", err
		}
		break
	}

	if urlData == (URLData{}) {
		return "", errors.New("Key not found!")
	}

	return urlData.Url, nil
}

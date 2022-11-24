package db

import "sync"

type Email struct {
	Name  string
	Email string
	Phone string
	File  *File
}

type File struct {
	FileName string
	File     []byte
}

type DB struct {
	emails []*Email
}

var instance *DB
var once sync.Once

func Db() *DB {
	once.Do(func() {
		instance = &DB{}
	})

	return instance
}

func DBGetEmails() []*Email {
	db := Db()
	return db.emails
}

func DBGetFiles() []*File {
	db := Db()
	files := make([]*File, 0)
	for _, e := range db.emails {
		files = append(files, e.File)
	}
	return files
}

func DBPut(email *Email) {
	db := Db()
	db.emails = append(db.emails, email)
}

func NewEntry(name, email, phone, fname string, file []byte) *Email {
	f := &File{
		FileName: fname,
		File:     file,
	}
	return &Email{
		Name:  name,
		Email: email,
		Phone: phone,
		File:  f,
	}
}

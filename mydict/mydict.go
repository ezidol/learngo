package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not found")
var errAlreadyExists = errors.New("Already exists")

// Search word on Dictionary
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	} else {
		return "", errNotFound
	}
}

// Add word, def on Dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
		return nil
	case nil:
		return errAlreadyExists
	}
	return nil
}
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return err
	case nil:
		d[word] = def
		return nil
	}
	return nil
}
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return err
	case nil:
		delete(d, word)
		return nil
	}
	return nil
}

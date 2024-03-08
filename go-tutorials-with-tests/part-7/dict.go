package main

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err == nil {
		return ErrWordExists
	} else {
		d[word] = definition
		return nil
	}
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err == nil {
		d[word] = definition
		return nil
	} else {
		return ErrWordDoesNotExist
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

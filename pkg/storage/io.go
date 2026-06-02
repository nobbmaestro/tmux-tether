package storage

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func (s *SessionStorage) Read() error {
	f, err := os.Open(s.path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}

	defer func() {
		if cerr := f.Close(); err == nil {
			err = cerr
		}
	}()

	return yaml.NewDecoder(f).Decode(s)
}

func (s *SessionStorage) Write() error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0755); err != nil {
		return err
	}

	f, err := os.Create(s.path)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := f.Close(); err == nil {
			err = cerr
		}
	}()

	return yaml.NewEncoder(f).Encode(s)
}

package persistence

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rogelioConsejo/golibs/file"
	"io"
	"obra-blanca/login/user"
	"os"
	"strings"
)

func New(filename file.Name) (persistence, error) {
	if strings.TrimSpace(string(filename)) == "" {
		filename = defaultFilename
	}
	f, err := os.OpenFile(string(filename), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return persistence{}, fmt.Errorf("could not open file: %w", err)
	}
	return persistence{
		file: f,
	}, nil
}

type persistence struct {
	file *os.File
}

func (p persistence) CheckUser(usr user.Username, pass user.Password) (bool, error) {
	if strings.TrimSpace(string(usr)) == "" {
		return false, errors.New(emptyUserErrorMessage)
	}
	if strings.TrimSpace(string(pass)) == "" {
		return false, errors.New(emptyPasswordErrorMessage)
	}

	savedUsers := make(users)
	err := json.NewDecoder(p.file).Decode(&savedUsers)
	if err != nil && err != io.EOF {
		return false, fmt.Errorf("could not decode file: %w", err)
	}

	var userHash = usernameHash(getHash(string(usr)))
	var passHash = passwordHash(getHash(string(pass)))
	if _, exists := savedUsers[userHash]; !exists || passHash != savedUsers[userHash] {
		return false, nil
	}

	return true, nil
}

func (p persistence) AddUser(usr user.Username, pass user.Password) error {
	savedUsers := make(users)
	err := json.NewDecoder(p.file).Decode(&savedUsers)
	if err != nil && err != io.EOF {
		return fmt.Errorf("could not decode file: %w", err)
	}
	var fileName = p.file.Name()
	p.file.Close()
	err = os.Truncate(fileName, 0)
	if err != nil && err != io.EOF {
		return fmt.Errorf("could not truncate file: %w", err)
	}

	var userHash = usernameHash(getHash(string(usr)))
	var passHash = passwordHash(getHash(string(pass)))

	savedUsers[userHash] = passHash
	p.file, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	err = json.NewEncoder(p.file).Encode(savedUsers)
	if err != nil {
		return fmt.Errorf("could not encode file: %w", err)
	}

	return nil
}

func (p persistence) Close() {
	_ = p.file.Close()
}

func (p persistence) deleteUser(userName user.Username) error {
	savedUsers := make(users)
	err := json.NewDecoder(p.file).Decode(&savedUsers)
	if err != nil && err != io.EOF {
		return fmt.Errorf("could not decode file: %w", err)
	}
	var fileName = p.file.Name()
	p.file.Close()
	err = os.Truncate(fileName, 0)
	if err != nil && err != io.EOF {
		return fmt.Errorf("could not truncate file: %w", err)
	}

	var userHash = usernameHash(getHash(string(userName)))
	delete(savedUsers, userHash)

	p.file, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	err = json.NewEncoder(p.file).Encode(savedUsers)
	if err != nil {
		return fmt.Errorf("could not encode file: %w", err)
	}

	return nil
}

func getHash(str string) string {
	hashed := sha512.New()
	hashed.Write([]byte(str))
	cleanHash := hashed.Sum(nil)
	decoded := hex.EncodeToString(cleanHash)

	return decoded
}

type users map[usernameHash]passwordHash

type usernameHash string
type passwordHash string

const emptyUserErrorMessage = "cannot check empty user"
const emptyPasswordErrorMessage = "cannot check empty password"
const defaultFilename = "credentials.ob"

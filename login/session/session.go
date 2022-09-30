package session

import (
	"errors"
	"fmt"
	"github.com/rogelioConsejo/golibs/file"
	"io"
	"obra-blanca/login/token"
	"obra-blanca/login/user"
	"strings"
	"time"
)

func NewPersistence(fileName file.Name) (Persistence, error) {
	if strings.TrimSpace(string(fileName)) == "" {
		fileName = defaultFileName
	}
	fs := file.New(fileName)

	return &persistence{fileSaver: fs}, nil
}

type Persistence interface {
	New(user.User, time.Time) (DTO, error)
	GetSession(token.Token) (DTO, error)
}

type persistence struct {
	fileSaver file.Persistence
}

func (p *persistence) New(u user.User, expiration time.Time) (DTO, error) {
	sessions, err := p.readSessions()
	if err != nil {
		return DTO{}, err
	}

	dto := DTO{
		Token:      token.MakeToken(token.Size),
		User:       u,
		Expiration: expiration,
	}
	sessions[dto.Token.String()] = dto
	err = p.fileSaver.Save(sessions)
	if err != nil {
		return DTO{}, err
	}

	return dto, nil
}

func (p *persistence) GetSession(t token.Token) (DTO, error) {
	var s, err = p.readSessions()
	if err != nil {
		return DTO{}, fmt.Errorf("could not read sessions: %w", err)
	}

	dto, exists := s[t.String()]
	if !exists {
		return DTO{}, errors.New("session does not exist")
	}
	return dto, nil
}

type Session interface {
	GetToken() token.Token
}

type DTO struct {
	Token      token.Token `json:"token"`
	User       user.User   `json:"user"`
	Expiration time.Time   `json:"expiration"`
}

func (d DTO) GetToken() token.Token {
	return d.Token
}

type Sessions map[string]DTO

func (p *persistence) readSessions() (Sessions, error) {
	var sessions Sessions = make(Sessions)
	err := p.fileSaver.Get(&sessions)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return sessions, nil
}

const defaultFileName file.Name = "session.ob"

package notifer

import (
	"encoding/json"
	"os"
)

func init() {
	Register("default", NewStdNotifer())
}

type DefaultEncoder struct {
}

func (d *DefaultEncoder) Encode(v interface{}) ([]byte, error) {
	bin, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return []byte(string(bin) + "\n"), nil
}

type StdNotifer struct {
	e Encoder
}

func NewStdNotifer() *StdNotifer {
	return &StdNotifer{
		e: &DefaultEncoder{},
	}
}

func (s *StdNotifer) SetEncoder(e Encoder) {
	s.e = e
}

func (s *StdNotifer) Notice(v interface{}) error {
	bin, err := s.e.Encode(v)
	if err != nil {
		return err
	}
	if _, err := os.Stdout.Write(bin); err != nil {
		return err
	}
	return nil
}

type Notifer interface {
	SetEncoder(e Encoder)
	Notice(v interface{}) error
}

type Encoder interface {
	Encode(v interface{}) ([]byte, error)
}

var notifers = map[string]Notifer{}

func Register(key string, n Notifer) {
	notifers[key] = n
}

func GetNotifer(key string) Notifer {
	if key == "" {
		return notifers["default"]
	}
	return notifers[key]
}

package shared

import (
	"encoding/json"
)

func Stob(s *Session) ([]byte, error) {
	return json.Marshal(s)
}

func Btos(b []byte) (*Session, error) {
	s := &Session{}
	err := json.Unmarshal(b, s)
	return s, err
}

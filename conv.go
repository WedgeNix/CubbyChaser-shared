package shared

import (
	"encoding/json"
)

func Stob(s Session) ([]byte, error) {
	return json.Marshal(s)
}

func Btos(b []byte) (Session, error) {
	s := Session{}
	err := json.Unmarshal(b, &s)
	return s, err
}

func Ctob(c Cubby) ([]byte, error) {
	return json.Marshal(c)
}

func Btoc(b []byte) (Cubby, error) {
	c := Cubby{}
	err := json.Unmarshal(b, &c)
	return c, err
}

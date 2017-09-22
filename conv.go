package shared

import (
	"encoding/json"
)

func Stob(s Session) []byte {
	b, _ := json.Marshal(s)
	return b
}

func Btos(b []byte) Session {
	var s Session
	json.Unmarshal(b, &s)
	return s
}

func Otob(o Order) []byte {
	b, _ := json.Marshal(o)
	return b
}

func Btoo(b []byte) Order {
	var o Order
	json.Unmarshal(b, &o)
	return o
}

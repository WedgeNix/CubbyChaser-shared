package shared

func Must(e error) {
	if e != nil {
		panic(e)
	}
}
package panic

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Panic(e interface{}) {
	panic(e)
}

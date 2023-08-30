package print

func CheckString(arg interface{}) bool {
	_, ok := arg.(string)
	return ok
}

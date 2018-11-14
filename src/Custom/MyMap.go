package Custom

var mapA map[string]string

func init() {
	mapA = make(map[string]string)
	mapA["ktds"] = "ktds inc"
	mapA["kt"] = "kt inc"
}

func getValueByKey(key string) string {
	return mapA[key]
}

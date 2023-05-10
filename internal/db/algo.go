package db

var ref = "0289PYLQGRJCUV"

func getTagFromID(id int) string {
	var index int
	var tag string
	size := len(ref)
	for id > 0 {
		index = id % size
		tag = string(ref[index]) + tag
		id -= index
		id /= size
	}
	return tag
}

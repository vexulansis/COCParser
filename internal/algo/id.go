package algo

var Ref = "0289PYLQGRJCUV"

func GetIDFromLH(low int, high int) int {
	return (low << 8) + high
}
func GetTagFromID(id int) string {
	var index int
	var tag string
	size := len(Ref)
	for id > 0 {
		index = id % size
		tag = string(Ref[index]) + tag
		id -= index
		id /= size
	}
	return tag
}

package algo

var ref = "0289CGJLPQRUVY"
var segments = []int{
	0, 14, 210, 2954, 41370, 579194, 8108730, 113522234, 1589311290, 22250358074,
}

func IdFromTag(tag string) int {
	var id int

	return id
}
func TagFromId(id int) string {
	var tag string
	var tagsize int
	size := len(ref)
	for i := 0; id >= segments[i]; i++ {
		tagsize = i + 1
	}
	pos := id - segments[tagsize-1]
	for i := tagsize; i > 0; i-- {
		tag = string(ref[pos%size]) + tag
		pos /= size
	}
	return tag
}

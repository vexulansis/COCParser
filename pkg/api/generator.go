package api

var Ref = "0289PYLQGRJCUV"
var LowMax = 40
var HighMax = 256

func GenerateTags(AC *APIClient) error {
	for high := 0; high < HighMax; high++ {
		go GenerateTagChunk(high, AC)
	}
	return nil
}
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
func GenerateTagChunk(high int, AC *APIClient) {
	for low := 1; low < LowMax; low++ {
		id := GetIDFromLH(low, high)
		tag := GetTagFromID(id)
		task := &Task{
			ID:   id,
			Data: tag,
		}
		AC.TaskChan <- task
	}
}

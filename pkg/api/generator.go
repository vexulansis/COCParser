package api

var Ref = "0289PYLQGRJCUV"
var LowMax = 4000
var HighMax = 256

func GenerateTags(AC *APIClient) error {
	for low := 1; low <= LowMax; low++ {
		for high := 0; high < HighMax; high++ {
			id := GetIDFromLH(low, high)
			tag := GetTagFromID(id)
			task := &Task{
				ID:   id,
				Data: tag,
			}
			AC.TaskChan <- task
			// f := APILoggerFields{
			// 	Source:      "GENERATOR",
			// 	Method:      "SEND",
			// 	Subject:     "#" + tag,
			// 	Destination: "TASKCHANNEL",
			// }
			// AC.Logger.Print(f, 0)
		}
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

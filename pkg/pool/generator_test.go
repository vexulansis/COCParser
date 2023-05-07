package pool

// import (
// 	"testing"
// )

// func TestGenerator(t *testing.T) {
// 	generator := NewGenerator("TEST", 10000)
// 	generator.Prepare()
// 	go generator.Start(0, 10000000)
// 	for {
// 		for _, o := range generator.Outputs {
// 			select {
// 			case <-o:
// 			case <-generator.Quit:
// 				return
// 			}
// 		}
// 	}
// }

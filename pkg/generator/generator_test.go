package generator

import (
	"context"
	"testing"
)

const testDelay = 1

const testSize = 10000
const testAmount = 100000000

var testFunc = GenerateTags

func TestGenerator(t *testing.T) {
	testConfig := &GeneratorConfig{
		Size:     testSize,
		Generate: testFunc,
		Delay:    testDelay,
	}
	testGenerator := NewGenerator("TEST", testConfig)
	testContext := context.Background()
	testGenerator.Start(testContext, testAmount)
}

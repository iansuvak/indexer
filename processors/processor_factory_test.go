package processors

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

var logger *logrus.Logger

func init() {
	logger, _ = test.NewNullLogger()
}

type mockProcessor struct {
	Processor
}

func (m *mockProcessor) Metadata() ProcessorMetadata {
	return MakeProcessorMetadata("foobar", "", false)
}

type mockProcessorConstructor struct {
	me *mockProcessor
}

func (c *mockProcessorConstructor) New() Processor {
	return c.me
}

func TestProcessorBuilderByNameSuccess(t *testing.T) {
	me := mockProcessor{}
	RegisterProcessor("foobar", &mockProcessorConstructor{&me})

	expBuilder, err := ProcessorBuilderByName("foobar")
	assert.NoError(t, err)
	exp := expBuilder.New()
	assert.Implements(t, (*Processor)(nil), exp)
}

func TestProcessorBuilderByNameNotFound(t *testing.T) {
	_, err := ProcessorBuilderByName("barfoo")
	expectedErr := "no Processor Constructor for barfoo"
	assert.EqualError(t, err, expectedErr)
}

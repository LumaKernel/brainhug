package brainhug_test

import (
	"os"
	"testing"

	"github.com/LumaKernel/brainhug"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
    // Before all
	// brainhug.Logging = true

    code := m.Run()

	// After all

    os.Exit(code)
}

func TestBrainhugSimple(t *testing.T) {
	stdout, err := brainhug.Proceed(
		`>+++[<+++>-]<.`,
		[]byte{},
	)
	assert.Nil(t, err)
	assert.Equal(t, stdout, []byte{9}, "")

	stdout, err = brainhug.Proceed(
		`> + + (^ o ^) + [ あ <  +  ++お >-] <.`,
		[]byte{},
	)
	assert.Nil(t, err)
	assert.Equal(t, stdout, []byte{9}, "")
}

func TestBrainhugNegative(t *testing.T) {
	stdout, err := brainhug.Proceed(
		`-.<.+.`,
		[]byte{0, 0, 0},
	)
	assert.Nil(t, err)
	assert.Equal(t, stdout, []byte{255, 0, 1}, "")
}

func TestBrainhugInput(t *testing.T) {
	stdout, err := brainhug.Proceed(
		`,+.<,.>,-.`,
		[]byte{10, 20, 30},
	)
	assert.Nil(t, err)
	assert.Equal(t, stdout, []byte{11, 20, 29}, "")
}

func TestBrainhugNest(t *testing.T) {
	stdout, err := brainhug.Proceed(
		`++++[>++++[>++++<-]<-]>>.`,
		[]byte{},
	)
	assert.Nil(t, err)
	assert.Equal(t, stdout, []byte{4 * 4 * 4}, "")
}

func TestBrainhugHelloWolrd(t *testing.T) {
	const program = `
		>++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<+
		+.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-
		]<+.
	`
	stdout, err := brainhug.Proceed(
		program,
		[]byte{},
	)
	assert.Nil(t, err)
	assert.Equal(t, string(stdout), "Hello, World!", "")
}

package podcast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type enclosureTest struct {
	t        EnclosureType
	expected string
}

var enclosureTests = []enclosureTest{
	enclosureTest{M4A, "audio/x-m4a"},
	enclosureTest{M4V, "video/x-m4v"},
	enclosureTest{MP4, "video/mp4"},
	enclosureTest{MP3, "audio/mpeg"},
	enclosureTest{MOV, "video/quicktime"},
	enclosureTest{PDF, "application/pdf"},
	enclosureTest{EPUB, "document/x-epub"},
	enclosureTest{M4A, "audio/x-m4a"},
	enclosureTest{99, "application/octet-stream"},
}

func TestEnclosureTypes(t *testing.T) {
	t.Parallel()
	for _, et := range enclosureTests {
		et := et
		t.Run(et.t.String(), func(t *testing.T) {
			t.Parallel()

			assert.EqualValues(t, et.expected, et.t.String())
		})
	}
}

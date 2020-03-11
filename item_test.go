package podcast_test

import (
	"fmt"
	"testing"

	"github.com/jaderebrasil/podcast"
	"github.com/stretchr/testify/assert"
)

func TestItemAddDescriptionTooLong(t *testing.T) {
	t.Parallel()

	// arrange
	i := podcast.Item{
		Title:       "item.title",
		Description: &podcast.Description{Text: "item.desc"},
		Link:        "http://example.com/article.html",
	}
	desc := ""
	for {
		if len(desc) >= 4051 {
			break
		}
		desc += "abc ss 5 "
	}

	// act
	i.AddDescription(desc)
	fmt.Print(i.Description.Text)
	// assert
	assert.Len(t, i.Description.Text, 4000)
}

func TestItemAddImageEmptyUrl(t *testing.T) {
	t.Parallel()

	// arrange
	i := podcast.Item{
		Title:       "item.title",
		Description: &podcast.Description{Text: "item.desc"},
		Link:        "http://example.com/article.html",
	}

	// act
	i.AddImage("")

	// assert
	assert.Nil(t, i.IImage)
}

func TestItemAddDurationZero(t *testing.T) {
	t.Parallel()

	// arrange
	i := podcast.Item{
		Title:       "item.title",
		Description: &podcast.Description{Text: "item.desc"},
		Link:        "http://example.com/article.html",
	}
	d := int64(0)

	// act
	i.AddDuration(d)

	// assert
	assert.EqualValues(t, "", i.IDuration)
}

func TestItemAddDurationLessThanZero(t *testing.T) {
	t.Parallel()

	// arrange
	i := podcast.Item{
		Title:       "item.title",
		Description: &podcast.Description{Text: "item.desc"},
		Link:        "http://example.com/article.html",
	}
	d := int64(-13)

	// act
	i.AddDuration(d)

	// assert
	assert.EqualValues(t, "", i.IDuration)
}

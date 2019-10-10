package podcast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemAddSummaryTooLong(t *testing.T) {
	t.Parallel()

	// arrange
	i := Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}
	summary := ""
	for {
		if len(summary) >= 4051 {
			break
		}
		summary = summary + "abc ss 5 "
	}

	// act
	i.AddSummary(summary)

	// assert
	assert.Len(t, i.ISummary.Text, 4000)
}

func TestAddImage(t *testing.T) {
	t.Parallel()

	// arrange
	i := Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}
	const img_url = "https://golang.org/doc/gopher/doc.png"

	// act
	i.AddImage(img_url)

	// assert
	assert.EqualValues(t, img_url, i.IImage.HREF)
}

func TestItemAddImageEmptyUrl(t *testing.T) {
	t.Parallel()

	// arrange
	i := Item{
		Title:       "item.title",
		Description: "item.desc",
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
	i := Item{
		Title:       "item.title",
		Description: "item.desc",
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
	i := Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}
	d := int64(-13)

	// act
	i.AddDuration(d)

	// assert
	assert.EqualValues(t, "", i.IDuration)
}

func TestAddDuration(t *testing.T) {
	t.Parallel()

	// arrange
	i := Item{
		Title:       "item title",
		Description: "item desc",
		Link:        "item link",
	}
	d := int64(533)

	// act
	i.AddDuration(d)

	// assert
	assert.EqualValues(t, "8:53", i.IDuration)
}

func TestAddPubDate(t *testing.T) {
	t.Parallel()

	// arrange
	i := Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}

	d := pubDate.AddDate(0, 0, -11)

	// act
	i.AddPubDate(&d)

	// assert
	assert.EqualValues(t, "2017-01-24 08:21:52 +0000 UTC", i.PubDate.String())
	assert.EqualValues(t, "Tue, 24 Jan 2017 08:21:52 +0000", i.PubDateFormatted)
}

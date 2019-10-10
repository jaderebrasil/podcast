package podcast

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	createdDate = time.Date(2017, time.February, 1, 8, 21, 52, 0, time.UTC)
	updatedDate = createdDate.AddDate(0, 0, 5)
	pubDate     = createdDate.AddDate(0, 0, 3)
)

func TestNewNonNils(t *testing.T) {
	t.Parallel()

	// arrange
	ti, l, d := "title", "link", "description"

	// act
	p := New(ti, l, d, &createdDate, &updatedDate)

	// assert
	assert.EqualValues(t, ti, p.Title)
	assert.EqualValues(t, l, p.Link)
	assert.EqualValues(t, d, p.Description)
	assert.True(t, createdDate.Format(time.RFC1123Z) >= p.PubDate)
	assert.True(t, updatedDate.Format(time.RFC1123Z) >= p.LastBuildDate)
}

func TestNewNils(t *testing.T) {
	t.Parallel()

	// arrange
	ti, l, d := "title", "link", "description"

	// act
	p := New(ti, l, d, nil, nil)

	// assert
	now := time.Now().UTC().Format(time.RFC1123Z)
	assert.EqualValues(t, ti, p.Title)
	assert.EqualValues(t, l, p.Link)
	assert.EqualValues(t, d, p.Description)
	// ensure time.Now().UTC() is set, or close to it
	assert.True(t, now >= p.PubDate)
	assert.True(t, now >= p.LastBuildDate)
}

func TestAddAuthorEmailEmpty(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)

	// act
	p.AddAuthor("", "")

	// assert
	assert.Len(t, p.ManagingEditor, 0)
	assert.Len(t, p.IAuthor, 0)
}

func TestAddAuthorManagingEditor(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)

	// act
	p.AddAuthor("the name", "me@test.com")

	// assert
	assert.EqualValues(t, "me@test.com (the name)", p.ManagingEditor)
	assert.EqualValues(t, "me@test.com (the name)", p.IAuthor)
}

func TestAddAtomLinkHrefEmpty(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)

	// act
	p.AddAtomLink("")

	// assert
	assert.Nil(t, p.AtomLink)
}

func TestAddAtomLink(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)

	// act
	p.AddAtomLink("atom.link")

	// assert
	assert.EqualValues(t, "atom.link", p.AtomLink.HREF)
	assert.EqualValues(t, "self", p.AtomLink.Rel)
	assert.EqualValues(t, "application/rss+xml", p.AtomLink.Type)
}

func TestAddCategoryEmpty(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)

	// act
	p.AddCategory("", nil)

	// assert
	assert.Len(t, p.ICategories, 0)
	assert.Len(t, p.Category, 0)
}

func TestAddCategory2(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)

	// act
	p.AddCategory("cat1", []string{"sub1"})
	p.AddCategory("cat2", nil)

	// assert
	assert.Len(t, p.ICategories, 2)
	assert.Len(t, p.Category, len("cat1,cat2"))
	assert.EqualValues(t, "cat1,cat2", p.Category)
	assert.EqualValues(t, "cat1", p.ICategories[0].Text)
	assert.EqualValues(t, "sub1", p.ICategories[0].ICategories[0].Text)
}

func TestAddCategorySubCatEmpty1(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)

	// act
	p.AddCategory("mycat", []string{""})

	// assert
	assert.Len(t, p.ICategories, 1)
	assert.EqualValues(t, p.Category, "mycat")
	assert.Len(t, p.ICategories[0].ICategories, 0)
}

func TestAddCategorySubCatEmpty2(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)

	// act
	p.AddCategory("mycat", []string{"xyz", "", "abc"})

	// assert
	assert.Len(t, p.ICategories, 1)
	assert.EqualValues(t, p.Category, "mycat")
	assert.Len(t, p.ICategories[0].ICategories, 2)
}

func TestAddImageEmpty(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)

	// act
	p.AddImage("")

	// assert
	assert.Nil(t, p.Image)
	assert.Nil(t, p.IImage)
}

func TestPodcastAddImage(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)

	// act
	p.AddImage("image.link")

	// assert
	assert.EqualValues(t, "image.link", p.Image.URL)
	assert.EqualValues(t, "image.link", p.IImage.HREF)
}

func TestAddItemEmptyTitleDescription(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)
	i := Item{}

	// act
	added, err := p.AddItem(i)

	// assert
	assert.EqualValues(t, 0, added)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Title")
	assert.Contains(t, err.Error(), "Description")
	assert.Contains(t, err.Error(), "required")
}

func TestAddItemEmptyEnclosureURL(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)
	i := Item{Title: "title", Description: "desc"}
	i.AddEnclosure("", MP3, 1)

	// act
	added, err := p.AddItem(i)

	// assert
	assert.EqualValues(t, 0, added)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Enclosure.URL is required")
}

func TestAddItemEmptyEnclosureType(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)
	i := Item{Title: "title", Description: "desc"}
	i.AddEnclosure("http://example.com/1.mp3", 99, 1)

	// act
	added, err := p.AddItem(i)

	// assert
	assert.EqualValues(t, 0, added)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Enclosure.Type is required")
}

func TestAddItemEmptyLink(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)
	i := Item{Title: "title", Description: "desc"}

	// act
	added, err := p.AddItem(i)

	// assert
	assert.EqualValues(t, 0, added)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Link is required")
}

func TestAddItemEnclosureLengthMin(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)
	i := Item{Title: "title", Description: "desc"}
	i.AddEnclosure("http://example.com/1.mp3", MP3, -1)

	// act
	added, err := p.AddItem(i)

	// assert
	assert.EqualValues(t, 1, added)
	assert.NoError(t, err)
	assert.Len(t, p.Items, 1)
	assert.EqualValues(t, 0, p.Items[0].Enclosure.Length)
}

func TestAddItemEnclosureNoLinkOverride(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)
	i := Item{Title: "title", Description: "desc"}
	i.AddEnclosure("http://example.com/1.mp3", MP3, -1)

	// act
	added, err := p.AddItem(i)

	// assert
	assert.EqualValues(t, 1, added)
	assert.NoError(t, err)
	assert.Len(t, p.Items, 1)
	assert.EqualValues(t, i.Enclosure.URL, p.Items[0].Link)
}

func TestAddItemEnclosureLinkPresentNoOverride(t *testing.T) {
	t.Parallel()

	// arrange
	theLink := "http://someotherurl.com/story.html"
	p := New("title", "link", "description", nil, nil)
	i := Item{Title: "title", Description: "desc"}
	i.Link = theLink
	i.AddEnclosure("http://example.com/1.mp3", MP3, -1)

	// act
	added, err := p.AddItem(i)

	// assert
	assert.EqualValues(t, 1, added)
	assert.NoError(t, err)
	assert.Len(t, p.Items, 1)
	assert.EqualValues(t, theLink, p.Items[0].Link)
}

func TestAddItemNoEnclosureGUIDValid(t *testing.T) {
	t.Parallel()

	// arrange
	theLink := "http://someotherurl.com/story.html"
	p := New("title", "link", "description", nil, nil)
	i := Item{Title: "title", Description: "desc"}
	i.Link = theLink

	// act
	added, err := p.AddItem(i)

	// assert
	assert.EqualValues(t, 1, added)
	assert.NoError(t, err)
	assert.Len(t, p.Items, 1)
	assert.EqualValues(t, theLink, p.Items[0].GUID)
}

func TestAddItemAuthor(t *testing.T) {
	t.Parallel()

	// arrange
	theAuthor := Author{Name: "Jane Doe", Email: "me@janedoe.com"}
	p := New("title", "link", "description", nil, nil)
	i := Item{Title: "title", Description: "desc", Link: "http://a.co/"}
	i.Author = &theAuthor

	// act
	added, err := p.AddItem(i)

	// assert
	assert.EqualValues(t, 1, added)
	assert.NoError(t, err)
	assert.Len(t, p.Items, 1)
	assert.EqualValues(t, &theAuthor, p.Items[0].Author)
	assert.EqualValues(t, theAuthor.Email, p.Items[0].IAuthor)
}

func TestAddItemRootManagingEditorSetsAuthorIAuthor(t *testing.T) {
	t.Parallel()

	// arrange
	theAuthor := "me@janedoe.com"
	p := New("title", "link", "description", nil, nil)
	p.ManagingEditor = theAuthor
	i := Item{Title: "title", Description: "desc", Link: "http://a.co/"}

	// act
	added, err := p.AddItem(i)

	// assert
	assert.EqualValues(t, 1, added)
	assert.NoError(t, err)
	assert.Len(t, p.Items, 1)
	assert.EqualValues(t, theAuthor, p.Items[0].Author.Email)
	assert.EqualValues(t, theAuthor, p.Items[0].IAuthor)
}

func TestAddItemRootIAuthorSetsAuthorIAuthor(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "link", "description", nil, nil)
	p.IAuthor = "me@janedoe.com"
	i := Item{Title: "title", Description: "desc", Link: "http://a.co/"}

	// act
	added, err := p.AddItem(i)

	// assert
	assert.EqualValues(t, 1, added)
	assert.NoError(t, err)
	assert.Len(t, p.Items, 1)
	assert.EqualValues(t, "me@janedoe.com", p.Items[0].Author.Email)
	assert.EqualValues(t, "me@janedoe.com", p.Items[0].IAuthor)
}

func TestAddSubTitleEmpty(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "desc", "Link", nil, nil)

	// act
	p.AddSubTitle("")

	// assert
	assert.Len(t, p.ISubtitle, 0)
}

func TestAddSubTitleTooLong(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "desc", "Link", nil, nil)
	subTitle := ""
	for {
		if len(subTitle) >= 80 {
			break
		}
		subTitle = subTitle + "ajd 2 "
	}

	// act
	p.AddSubTitle(subTitle)

	// assert
	assert.Len(t, p.ISubtitle, 64)
}

func TestAddSummaryTooLong(t *testing.T) {
	t.Parallel()

	// arrange
	p := New(
		"title",
		"desc",
		"Link",
		nil, nil)
	summary := ""
	for {
		if len(summary) >= 4051 {
			break
		}
		summary = summary + "jax ss 7 "
	}

	// act
	p.AddSummary(summary)

	// assert
	assert.Len(t, p.ISummary.Text, 4000)
}

func TestAddSummaryEmpty(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("title", "desc", "Link", nil, nil)

	// act
	p.AddSummary("")

	// assert
	assert.Nil(t, p.ISummary)
}

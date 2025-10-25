package mylibrary

import(
	"time"
)

type Book struct {
	Title    string
	Author  string
	Pages   int
	Text    string

	Readed  bool
	Addtime time.Time
	ReadedAt *time.Time
}

func AddBook(title string, author string, pages int, text string) Book{
	return Book{
		Title: title,
		Author: author,
		Pages: pages,
		Text: text,

		Readed: false,
		Addtime: time.Now(),
		ReadedAt: nil,
	}
}

func (b *Book)Read() {
	readedAt := time.Now()

	b.Readed = true
	b.ReadedAt = &readedAt
}
package mylibrary

type List struct{
	Books map[string]Book
}

func NewBook() *List{
	return &List{
		Books: make(map[string]Book),
	}
}

func (l *List)AddBook(book Book) error{
	for k, _ := range l.Books {
		if k == book.Title{
			return ErrBookAlreadyInLibrary
		}
	}
	l.Books[book.Title] = book

	return nil
}

func (l *List)GetBook(title string) (Book, error){
	book, err := l.Books[title]
	if !err{
		return Book{}, ErrBookNotFound
	}

	return book, nil
}

func (l *List) ListBooks() map[string]Book{
	return l.Books
}

func (l *List) ListUnreadedBooks() map[string]Book {
	unreadBooks := make(map[string]Book)
	for title, book := range l.Books{
		if !book.Readed {
			unreadBooks[title] = book
		}
	}

	return unreadBooks
}

func (l *List) ListReadedBooks() map[string]Book {
	readBooks := make(map[string]Book)
	for title, book := range l.Books{
		if book.Readed {
			readBooks[title] = book
		}
	}

	return readBooks
}

func (l *List) ReadBook(title string) (Book, error){
	book, ok := l.Books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}
	book.Read()
	return book, nil
}

func (l *List) DeleteBook(title string) error{
	_, ok := l.Books[title]
	if !ok {
		return ErrBookNotFound
	}

	delete(l.Books, title)

	return nil
}
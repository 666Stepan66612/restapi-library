package mylibrary

import "sync"

type List struct {
	books map[string]Book
	mtx   sync.RWMutex
}

func NewBook() *List {
	return &List{
		books: make(map[string]Book),
	}
}

func (l *List) AddBook(book Book) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.books[book.Title]; ok{
		return ErrBookAlreadyInLibrary
	}
	l.books[book.Title] = book

	return nil
}

func (l *List) GetBook(title string) (Book, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	
	book, err := l.books[title]
	if !err {
		return Book{}, ErrBookNotFound
	}

	return book, nil
}

func (l *List) ListBooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	return l.books
}

func (l *List) ListUnreadedBooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	unreadBooks := make(map[string]Book)
	for title, book := range l.books {
		if !book.Readed {
			unreadBooks[title] = book
		}
	}

	return unreadBooks
}

func (l *List) ListReadedBooks() map[string]Book {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	readBooks := make(map[string]Book)
	for title, book := range l.books {
		if book.Readed {
			readBooks[title] = book
		}
	}

	return readBooks
}

func (l *List) ReadBook(title string) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	book, ok := l.books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	book.Read()

	return book, nil
}

func (l *List) DeleteBook(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	
	_, ok := l.books[title]
	if !ok {
		return ErrBookNotFound
	}

	delete(l.books, title)

	return nil
}
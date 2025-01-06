from models import Book

book_list = []

def add_book(title, author, price, quantity):
    book_list.append(Book(title, author, price, quantity))
    return "Book added successfully!"

def view_books():
    if not book_list:
        return ["No books available."]
    return [book.display_details() for book in book_list]

def search_book(query):
    results = [
        book.display_details()
        for book in book_list
        if query.lower() in book.title.lower() or query.lower() in book.author.lower()
    ]
    return results if results else ["No books found."]

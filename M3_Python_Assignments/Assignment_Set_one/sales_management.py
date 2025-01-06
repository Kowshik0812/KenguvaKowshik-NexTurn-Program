from models import Transaction
from book_management import book_list

sales_records = []

def sell_book(customer_name, customer_email, customer_phone, book_title, quantity):
    for book in book_list:
        if book.title.lower() == book_title.lower():
            if book.quantity >= quantity:
                book.quantity -= quantity
                transaction = Transaction(customer_name, customer_email, customer_phone, book_title, quantity)
                sales_records.append(transaction)
                return f"Sale successful! Remaining quantity: {book.quantity}"
            else:
                return f"Error: Only {book.quantity} copies available."
    return "Error: Book not found."

def view_sales():
    if not sales_records:
        return ["No sales records available."]
    return [sale.display_transaction() for sale in sales_records]

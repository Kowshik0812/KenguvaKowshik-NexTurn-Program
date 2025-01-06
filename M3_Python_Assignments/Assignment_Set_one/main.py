from book_management import add_book, view_books, search_book
from customer_management import add_customer, view_customers
from sales_management import sell_book, view_sales

def main():
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == "1":
            print("\n1. Add Book\n2. View Books\n3. Search Book")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                title = input("Title: ")
                author = input("Author: ")
                try:
                    price = float(input("Price: "))
                    quantity = int(input("Quantity: "))
                    if price > 0 and quantity > 0:
                        print(add_book(title, author, price, quantity))
                    else:
                        print("Price and Quantity must be positive.")
                except ValueError:
                    print("Invalid input. Please enter valid numbers.")
            elif sub_choice == "2":
                print("\nAvailable Books:")
                for book in view_books():
                    print(book)
            elif sub_choice == "3":
                query = input("Enter title or author to search: ")
                print("\nSearch Results:")
                for book in search_book(query):
                    print(book)

        elif choice == "2":
            print("\n1. Add Customer\n2. View Customers")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                name = input("Name: ")
                email = input("Email: ")
                phone = input("Phone: ")
                print(add_customer(name, email, phone))
            elif sub_choice == "2":
                print("\nCustomers:")
                for customer in view_customers():
                    print(customer)

        elif choice == "3":
            print("\n1. Sell Book\n2. View Sales")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                customer_name = input("Customer Name: ")
                customer_email = input("Customer Email: ")
                customer_phone = input("Customer Phone: ")
                book_title = input("Book Title: ")
                try:
                    quantity = int(input("Quantity: "))
                    if quantity > 0:
                        print(sell_book(customer_name, customer_email, customer_phone, book_title, quantity))
                    else:
                        print("Quantity must be positive.")
                except ValueError:
                    print("Invalid input. Please enter a valid quantity.")
            elif sub_choice == "2":
                print("\nSales Records:")
                for sale in view_sales():
                    print(sale)

        elif choice == "4":
            print("Exiting BookMart. Goodbye!")
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main()

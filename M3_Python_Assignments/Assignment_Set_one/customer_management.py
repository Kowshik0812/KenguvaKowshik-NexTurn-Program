from models import Customer

customer_list = []

def add_customer(name, email, phone):
    customer_list.append(Customer(name, email, phone))
    return "Customer added successfully!"

def view_customers():
    if not customer_list:
        return ["No customers available."]
    return [customer.display_details() for customer in customer_list]

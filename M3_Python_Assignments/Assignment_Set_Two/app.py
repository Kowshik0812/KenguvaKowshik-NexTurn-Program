from flask import Flask, request, jsonify
import sqlite3

app = Flask(__name__)

DATABASE = 'books.db'

def query_db(query, args=(), one=False):
    connection = sqlite3.connect(DATABASE)
    cursor = connection.cursor()
    cursor.execute(query, args)
    rows = cursor.fetchall()
    connection.commit()
    connection.close()
    return (rows[0] if rows else None) if one else rows

@app.route('/books', methods=['POST'])
def add_book():
    data = request.json
    if not all(key in data for key in ("title", "author", "published_year", "genre")):
        return jsonify({"error": "Invalid data", "message": "Missing required fields"}), 400

    try:
        query_db(
            'INSERT INTO books (title, author, published_year, genre) VALUES (?, ?, ?, ?)',
            (data['title'], data['author'], data['published_year'], data['genre'])
        )
        return jsonify({"message": "Book added successfully"}), 201
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route('/books', methods=['GET'])
def get_books():
    books = query_db('SELECT * FROM books')
    return jsonify([{
        "id": row[0],
        "title": row[1],
        "author": row[2],
        "published_year": row[3],
        "genre": row[4]
    } for row in books])

@app.route('/books/<int:id>', methods=['GET'])
def get_book_by_id(id):
    book = query_db('SELECT * FROM books WHERE id = ?', (id,), one=True)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404
    return jsonify({
        "id": book[0],
        "title": book[1],
        "author": book[2],
        "published_year": book[3],
        "genre": book[4]
    })

@app.route('/books/<int:id>', methods=['PUT'])
def update_book(id):
    data = request.json
    book = query_db('SELECT * FROM books WHERE id = ?', (id,), one=True)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    try:
        query_db(
            'UPDATE books SET title = ?, author = ?, published_year = ?, genre = ? WHERE id = ?',
            (data.get('title', book[1]), data.get('author', book[2]), data.get('published_year', book[3]),
             data.get('genre', book[4]), id)
        )
        return jsonify({"message": "Book updated successfully"})
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route('/books/<int:id>', methods=['DELETE'])
def delete_book(id):
    book = query_db('SELECT * FROM books WHERE id = ?', (id,), one=True)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    try:
        query_db('DELETE FROM books WHERE id = ?', (id,))
        return jsonify({"message": "Book deleted successfully"})
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True)

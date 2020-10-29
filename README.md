# makemake

A simple, RESTful, Go CRUD application. An attempt was made to use the 'Clean Architecture' pattern. 

## API

### `/books`

Get all books:

```
curl -XGET \
    http://127.0.0.1:8080/books
```

Get book by ID:

```
curl -XGET \
    http://127.0.0.1:8080/books/103
```

Create a book:

```
curl -XPOST \
    -d '{"title":"The Stand","author":"King, Stephen","isbn":"978-0-385-12168-2"}' \
    http://127.0.0.1:8080/books
```

Update a book:

```
curl -XPUT \
    -d '{"title":"The Stand","author":"King, Stephen","isbn":"978-0307743688"}' \
    http://127.0.0.1:8080/books/103
```

Delete a book:

```
curl -XDELETE \
    http://127.0.0.1:8080/books/103
```
# Snippetbox

Snippetbox is a web application developed in Go that allows users to create, view, and share text snippets. This project is built following the *Let's Go* book by Alex Edwards, which provides a comprehensive guide to building web applications with Go.

## Features

- **Create Snippets**: Users can submit snippets of text through a form, which are then stored in the application.
- **View Snippets**: Each snippet has a unique URL, allowing users to view individual snippets.
- **List Recent Snippets**: The homepage displays a list of the most recent snippets.

## Project Structure

The project follows a modular structure as recommended in the *Let's Go* book:

```
.
├── README.md
├── cmd
│   └── web
│       ├── context.go
│       ├── handlers.go
│       ├── handlers_test.go
│       ├── helpers.go
│       ├── main.go
│       ├── middleware.go
│       ├── middleware_test.go
│       ├── routes.go
│       ├── templates.go
│       ├── templates_test.go
│       └── testutils_test.go
├── db
│   └── README.md
├── go.mod
├── go.sum
├── internal
│   ├── assert
│   │   └── assert.go
│   ├── models
│   │   ├── errors.go
│   │   ├── snippets.go
│   │   └── users.go
│   └── validator
│       └── validator.go
├── tls
│   ├── cert.pem
│   └── key.pem
├── tmp
│   ├── error.log
│   └── info.log
└── ui
    ├── efs.go
    ├── html
    │   ├── base.tmpl
    │   ├── pages
    │   │   ├── create.tmpl
    │   │   ├── home.tmpl
    │   │   ├── login.tmpl
    │   │   ├── signup.tmpl
    │   │   └── view.tmpl
    │   └── partials
    │       └── nav.tmpl
    └── static
        ├── css
        │   ├── index.html
        │   └── main.css
        ├── img
        │   ├── favicon.ico
        │   ├── index.html
        │   └── logo.png
        ├── index.html
        └── js
            ├── index.html
            └── main.js
```

## Prerequisites

- **Go**: Ensure you have Go installed. You can download it from the [official website](https://golang.org/dl/).
- **MySQL**: The application uses MySQL as its database. Install MySQL and ensure it's running.

## Setup Instructions

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/jyotirmoydotdev/snippetbox.git
   cd snippetbox
   ```

2. **Set Up the Database**:
   
   - Log in to MySQL:
     ```bash
     mysql -u root -p
     ```
     
   - Create a new database and user:
     ```sql
     CREATE DATABASE snippetbox;
     CREATE USER 'web'@'localhost' IDENTIFIED BY 'pass';
     GRANT ALL PRIVILEGES ON snippetbox.* TO 'web'@'localhost';
     ```
     
   - Create snippets table
     ```sql
     -- Create a `snippets` table.
      CREATE TABLE snippets (
          id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
          title VARCHAR(100) NOT NULL,
          content TEXT NOT NULL,
          created DATETIME NOT NULL,
          expires DATETIME NOT NULL
      );
      
      -- Add an index on the created column.
      CREATE INDEX idx_snippets_created ON snippets(created);
     ```
     
   - Create sessions table
     ```sql
      USE snippetbox;
     
      CREATE TABLE sessions (
          token CHAR(43) PRIMARY KEY,
          data BLOB NOT NULL,
          expiry TIMESTAMP(6) NOT NULL
      );
      
      CREATE INDEX sessions_expiry_idx ON sessions (expiry);
     ```
     
   - Create users table
     ```sql
     USE snippetbox;

     CREATE TABLE users (
        id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        hashed_password CHAR(60) NOT NULL,
        created DATETIME NOT NULL
     );
     ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
     ```
     
   - Exit MySQL:
     ```sql
     EXIT;
     ```

5. **Run the Application**:
   ```bash
   go run ./cmd/web
   ```

6. **Access the Application**:
   Open your browser and navigate to `https://localhost:4000`.

## Configuration

Configuration settings, such as the server port and database credentials, can be adjusted in the `main.go` file or through environment variables as described in the *Let's Go* book.

## Acknowledgements

This project is based on the *Let's Go* book by Alex Edwards. For more information and additional resources, visit the [official website](https://lets-go.alexedwards.net/).

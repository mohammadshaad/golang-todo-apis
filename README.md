# Go Todo App

This is a simple web-based todo application built with Go (Golang) and Fiber. It allows you to create, update, and delete todo items.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Contributing](#contributing)

## Features

- Create new todo items.
- Update existing todo items.
- Delete todo items.
- Simple and clean web interface.
- Uses a PostgreSQL database to store todo items.

## Requirements

Before you begin, ensure you have met the following requirements:

- Go (Golang) installed on your local machine.
- PostgreSQL database server installed and running.
- Dependencies managed using Go Modules.

## Getting Started

1. Clone this repository:

```bash
git clone https://github.com/mohammadshaad golang-todo-app.git
```

Change to the project directory:

```bash
cd golang-todo-app
```

Install project dependencies:

```bash
go mod tidy
```

Set up the PostgreSQL database. Create a database named todos and configure the database connection in the server.go file.

Run the application:

```bash
go run server.go
```

Access the application in your web browser at http://localhost:8080 (or the specified port).

## Usage

Visit the application in your web browser.
Add new todo items using the input field and "Add" button.
Edit existing todo items by clicking the edit button.
Delete todo items by clicking the delete button.
Enjoy organizing your todos!

## Contributing
Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

Fork the project.

Create your feature branch: 

```bash
git checkout -b feature/your-feature-name
```

Commit your changes:

```bash
git commit -m 'Add some feature'
```

Push to the branch: 

```bash
git push origin feature/your-feature-name
```

Submit a pull request.

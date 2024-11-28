# Go Time API

This project provides a simple REST API to get the current time in Toronto and log the time in a MySQL database. The API has two endpoints:

- `/current-time`: Returns the current time in Toronto and logs it to the database.
- `/logged-times`: Retrieves and returns a list of all logged times in the database.

## Requirements

- Go 1.20 or later
- MySQL database
- Docker (optional for running MySQL in a container)
- Environment variables in a `.env` file

# Go Time API

This project provides a simple REST API to get the current time in Toronto and log the time in a MySQL database. The API has two endpoints:

- `/current-time`: Returns the current time in Toronto and logs it to the database.
- `/logged-times`: Retrieves and returns a list of all logged times in the database.

## Requirements

- Go 1.20 or later
- MySQL database
- Docker (optional for running MySQL in a container)
- Environment variables in a `.env` file

##	Set Up MySQL Database
-	Install MySQL and create a new database.
-	Create a table named time_log with at least two fields: id (primary key) and timestamp.
  ![image](https://github.com/user-attachments/assets/1a7b33f2-9074-48bb-929a-e58844aba87a)

- Created Database and Table time_log
  ![image](https://github.com/user-attachments/assets/389e7fb5-7f68-4001-9952-5cde4b76b598)

## API Development
-	Write a Go application with a web server.
-	Create an API endpoint /current-time that returns the current time in Toronto.
  <img width="629" alt="image" src="https://github.com/user-attachments/assets/f0a30d2f-0867-4585-823a-07510b477116">

## Time Zone Conversion:
- Use Go's time package to handle the time zone conversion to Toronto's local time.
  ![image](https://github.com/user-attachments/assets/8d8a215e-2945-465d-8db5-b89076cb7d8b)

## Database Connection:
- Connect to your MySQL database from your Go application.
  ![image](https://github.com/user-attachments/assets/9c8fe7a3-88ef-4a64-a3a9-fdba3e15178c)

-	On each API call, insert the current time into the time_log table.
  ![image](https://github.com/user-attachments/assets/7dedccb1-e848-4d2b-9530-f965768807ea)

  ![image](https://github.com/user-attachments/assets/5f69945f-2577-4456-ad8e-9487e1fc87fb)

  ![image](https://github.com/user-attachments/assets/75057fb5-16d9-4870-bbe5-dadebaee7279)

- Record inserted to table each time API is called.
  ![image](https://github.com/user-attachments/assets/97157d02-c0b6-415d-9d8f-aa840f3ff8a6)

## Return Time in JSON:
![image](https://github.com/user-attachments/assets/7cf08e80-d96a-4ded-9919-100505d92453)

## TESTING
![image](https://github.com/user-attachments/assets/13cd005c-9e3d-4a78-b9d0-967393fb4617)

![image](https://github.com/user-attachments/assets/e6f703af-9e1b-4175-963b-3760311c3f0b)

## Create an additional endpoint to retrieve all logged times from the database.
![image](https://github.com/user-attachments/assets/80c78850-0474-4a84-9be3-e7c5aff67200)

![image](https://github.com/user-attachments/assets/24c28f2c-0168-4418-8c1b-c451374c3794)
















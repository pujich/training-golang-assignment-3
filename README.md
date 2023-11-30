# training-golang-assignment-3

Project Structure

|-- config/
|   |-- config.go
|-- model/
|   |-- installment.go
|   |-- loan.go
|-- services/
|   |-- installment_service.go
|   |-- loan_service.go
|-- .gitignore
|-- go.mod
|-- go.sum
|-- main.go
|-- README.md
config: Configuration files and setup for the database.
model: Database models (Installment and Loan) and functions for interacting with the database.
services: Implementation of services for handling HTTP requests and business logic.
Installation
Make sure you have Golang installed on your machine.

Clone the repository:

bash

git clone https://github.com/pujich/training-golang-assignment-3.git

Change into the project directory:

cd training-golang-assignment-3

Install dependencies:

go mod tidy

Run the application:

go run main.go

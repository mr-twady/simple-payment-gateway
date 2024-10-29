# Payment Gateway Microservice

## About
My solution to the Exinity Senior Backend Assessment.

Design Document and architecture diagram here - https://drive.google.com/file/d/1nTShZ-u7fzEhxksK9p2Ttzse5KWoIfBU/view?usp=sharing

## Overview
The Payment Service is a microservice designed to handle payment transactions through multiple payment gateways. It provides endpoints for deposit and withdrawal operations, along with a callback mechanism to handle updates from payment gateways. Additionally, I implemented a simple health check endpoint and an endpoint to retrieve the updated user balance upon a successful deposit and withdrawal operation.
The Payment Gateway Microservice is a robust, extensible service designed to handle financial transactions such as deposits (cash-in) and withdrawals (cash-out) through various payment gateways. The service is built in Go and utilizes a flexible architecture to support multiple protocols and data formats, including JSON over HTTP and ISO8583 over TCP. 

### Features
- InitiateDeposit and verify deposits and withdrawals.
- Handle asynchronous callbacks for transaction updates from payment gateways.
- Support for multiple payment gateways with minimal changes to the codebase.
- Resilience strategies including circuit breakers, retries, and timeouts for gateway failures.
- Comprehensive error handling and meaningful HTTP status codes.

## Table of Contents
- [Getting Started](#getting-started)
- [Requirements](#requirements)
- [Setup Instructions](#setup-instructions)
- [Directory Structure](#directory-structure)
- [Usage](#usage)
- [Testing](#testing)

## Getting Started
These instructions will help you set up the project locally for development and testing.

### Requirements
- Go 1.20 or higher
- PostgreSQL (or any supported database)
- Docker (optional but recommended, for containerization) and quick setup
- Docker Compose (optional, for local development)


## Setup Instructions

1. **Clone the repository and change directory to the cloned project**
   ```
   git clone https://github.com/mr-twady/simple-payment-gateway.git
   cd simple-payment-gateway


2. **Create a .env file from .env.example file in the root directory to store your environment variables. Example:**
    ```
    DB_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
    HTTP_PORT=8080
    GATEWAY_A_URL=http://payment.gateway-a.com
    GATEWAY_B_URL=http://payment.gateway-b.com
    TIMEOUT=

2. **Initialize Go Modules If you haven't already:**
    ```
    go mod init dunsin-olubobokun/simple-payment-gateway #from your root directory

3. **Once that is done, if you intend to use Docker to run this application, just run the following commands.**
    ```
    docker-compose build --no-cache
    docker-compose up 
    ```
    - **If all was well setup, you should be able to hit http://localhost:8080/health**

4. **If you want to run this application manually i.e without using Docker, update line 16 in internal/config/config.go as follows**
    ```internal/config/config.go
    viper.SetConfigFile(".env") // update from
    viper.SetConfigFile("../.env") // update to

5. **Install Dependencies Use the following command to install the necessary dependencies:**
    ```
    go get # cd into main if doesnt work from root dir

7. **On application start up, I added a simple migration to set up the Database and run any necessary migrations. Ensure your database connection details in .env are correct.**

8. **Run the Application Start the server using:**
    ```
    go run main/main.go # if you are in root dir and run
    go run main.go # if you are in main dir 

9. **For API reference, kindly refer to your /docs of your base server URL. A list of API Endpoints and details are explained there**
    ```
    http://localhost:8080/docs

# Directory Structure
High level overview of project structure

dunsin-olubobokun/simple-payment-gateway/
├── main/                    # Main entry point for the application
├── internal/                # Internal application logic
│   ├── api/                 # API handlers
│   ├── config/              # Configuration management
│   ├── gateways/            # Mock Payment gateway integrations
│   ├── models/              # Data models
│   ├── migrations/          # Database migrations 
│   ├── repository/          # Database interactions
│   ├── service/             # Business logic
│   ├── middleware/          # Middleware functions
│   └── utils/               # Utility functions such as
└── tests/                   # Unit tests


## Usage 
    - Health check 
    - GET /health
        - Response body
        ```
           {
                "message": "Payment service is up and running!"
            }
        ```


    - Deposit Funds (Initiate)
    - POST /api/deposit
        - Request body
        ```
            {
                "type": "deposit",
                "amount": 10,
                "currency": "USD",
                "customerReference": "abd1qas0dd1",
                "email": "test@test.com"
            }
        ```
        - Response body
        ```
            {
                "type": "deposit",
                "amount": 10,
                "currency": "USD",
                "status": "processing",
                "customer_reference": "abd1qas0dd1",
                "email": "test@test.com"
            }
        ```

    - Verify Deposit
    - POST api/deposit/verify
        - Request body
        ```
            {
                "customerReference": "abd1qas0dd1"
            }
        ```
        - Response body
        ```
            {
                "type": "deposit",
                "amount": 10,
                "currency": "USD",
                "status": "processing",
                "customer_reference": "abd1qas0dd1",
                "email": "test@test.com"
            }
        ```

    - Callback (Confirm a deposit transaction)
    - POST /api/callback
        - Request body
        ```
            {
                "customerReference": "abc455456",
                "status": "completed"
            }
        ```
        - Response body
        ```
            {
                "status":"completed"
            }
        ```

    - Withdraw Funds 
    - POST /api/withdrawal
        - Request body
        ```
            {
                "type": "withdrawal",
                "amount": 50.0,
                "currency": "USD",
                "customerReference": "dsdsds999zsds",
                "email": "test@test.com"
            }
        ```
        - Response body
        ```
            {
                "type": "withdrawal",
                "amount": 50,
                "currency": "USD",
                "status": "completed",
                "customer_reference": "dsdsds999zsds",
                "email": "test@test.com"
            }
        ```

    - Get user balance  
    - GET /api/user/balance?email=
        - Request param
        ```
            {
                "email": "test@test.com"
            }
        ```
        - Response body
        ```
            {
                "balance":500
            }
        ```
   
# Testing
    ```
    go test ./tests # from root dir



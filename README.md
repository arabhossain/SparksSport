# SparksSport

## Overview
SparksSport is a Golang-based application that provides a robust and modular foundation for building scalable applications, with a focus on efficient data handling and extensibility.
## Requirements
- Go 1.23
- MySQL Database

## Installation
### 1. Clone the Repository
```bash
git clone https://github.com/arabhossain/SparksSport.git
cd SparksSport
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Set Up Environment Variables
Create a `.env` file in the root directory:
```env
SERVER_PORT=8080
DATABASE_URL=root:password@tcp(localhost:3306)/sparks_sport?charset=utf8mb4&parseTime=True&loc=Local
```

### 4. Run Migrations
When the application starts, it will automatically run the migrations first, ensuring the database is up to date before launching the HTTP server.
### 5. Start the Application
```bash
go run cmd/app/main.go

```

### 6. Generating a New Module

#### 1. Navigate to the `cmd/modular` Directory
Open your terminal and move to the `cmd/modular` directory where the module generation script is located:

```bash
cd cmd/modular
```

#### 2. Run the Module Generation Command
Execute the following command to create a new module, replacing `<module_name>` with your desired module name:

```bash
go run main.go make:module <module_name>
```

## 7. Dependencies
The project relies on the following Go modules:

```go
require (
    github.com/google/uuid v1.6.0
    github.com/gorilla/mux v1.8.1
    github.com/joho/godotenv v1.5.1
    gorm.io/driver/mysql v1.5.7
    gorm.io/gorm v1.25.12

    // Indirect dependencies
    filippo.io/edwards25519 v1.1.0
    github.com/go-sql-driver/mysql v1.8.1
    github.com/jinzhu/inflection v1.0.0
    github.com/jinzhu/now v1.1.5
    golang.org/x/text v0.22.0
)
```  

## 8. License
This project is open-source and available under the MIT License. Anyone can use this foundational framework to build their applications, contribute improvements, and help enhance its capabilities for the community.
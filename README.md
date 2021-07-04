# Ohmyhome REST API
This REST API is an application solution for Ohmyhome Backend Engineer Examination.

The problem was to create a REST API is provide "Property" and "Country" endpoints for users to post properties they want to sell or rent out.

The REST API was written using Go Language and the Data Storage uses MySQL.

Some sample data (Countries and Properties), will be initialized into the database when the application runs.

## Models
Property and Country Models design are as shown below:
```
Country Model:
- ID
- Country
```
```
Property Model:
- ID
- Address
- Country
- Description
- Available
```

## API Endpoints
The REST API Endpoints allows the user to do simple CRUD operations for Property and Country information. The endpoint only accepts request where the "Content-type" is "application/json".

```
Country Endpoints:
- GET Methods
  - Get All Countries 
  - Get Single Country
- POST Method
  - Create New Country
- PUT Method
  - Update Country Information
- DELETE Method
  - Delete Country
```
```
Property Endpoints:
- GET Methods
  - Get All Properties (GET)
  - Get Available Properties (GET)
  - Get Properties In Country (GET)
  - Get Single Property (GET)
- POST Methods
  - Create New Property (POST)
- PUT Method
  - Update Property Information (PUT)
- DELETE Method
  - Delete Property (DELETE)

```

## Libraries
Libraries used are:
- Standard Go Libraries ("fmt", "net/http", etc)
- https://github.com/gorilla/mux v1.8.0 - For Routing
- https://github.com/jinzhu/gorm v1.9.16 - For ORM
- https://github.com/joho/godotenv v1.3.0 - For Reading .env file
- https://github.com/microcosm-cc/bluemonday v1.0.14 - For Sanitization


## Running the Application

1. Clone the project from https://github.com/ksw95/OMH_BEx.git (Unless you already have the files)
2. Create an .env file in the root folder, with these required values:
```
MYSQL_HOSTNAME=<host>
MYSQL_USER=<username>
MYSQL_PASSWORD=<password>
MYSQL_DBNAME=<database>
MYSQL_PORT=<port_no>
```
3. Once project has been cloned and .env file has been created in the root folder, open a terminal and change directory into the project root folder and run the main.go file using the command ```go run main.go```.

## Potential Improvements
- Dockerizing the application
- CI/CD inclusion
- Versioning and Throttling
- Proper User Authentication with API Key

## Testing
Application was manually tested using Postman.
Tested the basic CRUD operations, ensuring sanitzation and validation are working and basic error handling via error messages.

Screenshot examples:

Getting all properties information
![image](https://user-images.githubusercontent.com/73837999/124390104-af9c2200-dd1c-11eb-9e55-7637f7ba8b1a.png)

Posting new property
![image](https://user-images.githubusercontent.com/73837999/124390302-829c3f00-dd1d-11eb-92b9-efc224b4678e.png)

Getting single country
![image](https://user-images.githubusercontent.com/73837999/124390323-9ba4f000-dd1d-11eb-8e8c-2c75d196cab7.png)

Updating country information
![image](https://user-images.githubusercontent.com/73837999/124390338-b11a1a00-dd1d-11eb-8672-d773dabaa293.png)

Deleting country entry
![image](https://user-images.githubusercontent.com/73837999/124390356-c1ca9000-dd1d-11eb-8114-ba6777d9ead0.png)

Testing error handling via error responses
![image](https://user-images.githubusercontent.com/73837999/124390391-f0e10180-dd1d-11eb-92ab-fe528db63ce9.png)

Validation testing
![image](https://user-images.githubusercontent.com/73837999/124390438-37366080-dd1e-11eb-8291-22561cfb3f9e.png)

Sanitization testing
![image](https://user-images.githubusercontent.com/73837999/124390505-89778180-dd1e-11eb-80aa-fb14081a33c2.png)




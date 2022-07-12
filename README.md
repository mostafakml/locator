# locator
a service  try to locate a data-bank location for drones
Mostafakmlaghrabasian4
Mostafakmlaghrabasian4
# project structure

- Clean architecture template with gin framework,
- one custom validation add to validate data that cant be convert to float
- use docker for deploying fast and easy
- adding unit test for controller layer and service layer 

## Running the project

- Make sure you have docker installed.
- Copy `.env.example` to `.env`
- Run `docker build -t locator`
- Run `docker run -p 8080:8080 locator`
- Go to `localhost:8080/ping` to verify if the server works.



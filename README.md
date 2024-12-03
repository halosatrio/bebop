project structure

/bebop
  /config         # Configuration files, e.g., database connection settings
  /models         # Data structures
  /service
    /delivery    # Handles incoming HTTP requests
    /repository  # Database access logic
      /postgres  # PostgreSQL-specific code
    /usecase     # Business logic
  main.go

## TODO:

- [x] auth
  - [x] register
  - [x] login
  - [x] get user
- [x] transactions
  - [x] input cashflow: amount, date, category, notes
  - [x] edit cashflow by id: amount, category, notes
  - [x] delete cashflow by id
  - [x] get cashflow: by date, by month
  - [ ] get outflow, inflow: by month, by category, by quarter
- [x] reports
  - [x] quarterly, annually, by category
- [x] assets
  - [x] get assets
  - [x] post record assets

## Endpoints:

- POST register user
- POST auth - login user
- POST input cashflow
- PUT edit cashflow
- DELETE transaction
- GET cashflow, params: date, month, category (?)
- GET Summary and Reports by date-range (months, quarter)
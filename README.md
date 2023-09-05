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

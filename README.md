# go-snippetbox
  Exploring go web application

  This little practice follows the book, [Let's Go](https://lets-go.alexedwards.net/) by @alexedwards


### Start web application
`make web`

### Check available flags
`make help`

# Table of Contents
1. [Introduction](#introduction)
    - [x] 1.1. Prerequisites
2. [Foundations](#foundations)
    - [x] 2.1. Project setup and creating a module
    - [x] 2.2. Web application basics
    - [x] 2.3. Routing requests
    - [x] 2.4. Customizing HTTP headers
    - [x] 2.5. URL query strings
    - [x] 2.6. Project structure and organization
    - [x] 2.7. HTML templating and inheritance
    - [x] 2.8. Serving static files
    - [x] 2.9. The http.Handler interface
3. [Configuration and error handling](#configuration-and-error-handling)
    - [x] 3.1. Managing configuration settings
    - [x] 3.2. Structured logging
    - [x] 3.3. Dependency injection
    - [x] 3.4. Centralized error handling
    - [x] 3.5. Isolating the application routes
4. [Database-driven responses](#database-driven-responses)
    - [x] 4.1. Setting up MySQL
    - [x] 4.2. Installing a database driver
    - [x] 4.3. Modules and reproducible builds
    - [x] 4.4. Creating a database connection pool
    - [x] 4.5. Designing a database model
    - [x] 4.6. Executing SQL statements
    - [x] 4.7. Single-record SQL queries
    - [x] 4.8. Multiple-record SQL queries
    - [x] 4.9. Transactions and other details
5. [Dynamic HTML templates](#dynamic-html-templates)
    - [x] 5.1. Displaying dynamic data
    - [x] 5.2. Template actions and functions
    - [x] 5.3. Caching templates
    - [x] 5.4. Catching runtime errors
    - [x] 5.5. Common dynamic data
    - [x] 5.6. Custom template functions
6. [Middleware](#middleware)
    - [x] 6.1. How middleware works
    - [x] 6.2. Setting security headers
    - [x] 6.3. Request logging
    - [x] 6.4. Panic recovery
    - [x] 6.5. Composable middleware chains
7. [Advanced routing](#advanced-routing)
    - [x] 7.1. Choosing a router
    - [x] 7.2. Clean URLs and method-based routing
8. [Processing forms](#processing-forms)
    - [x] 8.1. Setting up an HTML form
    - [x] 8.2. Parsing form data
    - [x] 8.3. Validating form data
    - [x] 8.4. Displaying errors and repopulating fields
    - [x] 8.5. Creating validation helpers
    - [x] 8.6. Automatic form parsing
9. [Stateful HTTP](#stateful-http)
    - [x] 9.1. Choosing a session manager
    - [x] 9.2. Setting up the session manager
    - [x] 9.3. Working with session data
10. [Server and security improvements](#server-and-security-improvements)
    - [x] 10.1. The http.Server struct
    - [x] 10.2. The server error log
    - [x] 10.3. Generating a self-signed TLS certificate
    - [x] 10.4. Running a HTTPS server
    - [x] 10.5. Configuring HTTPS settings
    - [x] 10.6. Connection timeouts
11. [User authentication](#user-authentication)
    - [x] 11.1. Routes setup
    - [x] 11.2. Creating a users model
    - [x] 11.3. User signup and password encryption
    - [x] 11.4. User login
    - [x] 11.5. User logout
    - [x] 11.6. User authorization
    - [x] 11.7. CSRF protection
12. [Using request context](#using-request-context)
    - [x] 12.1. How request context works
    - [x] 12.2. Request context for authentication/authorization
13. [File embedding](#file-embedding)
    - [x] 13.1. Embedding static files
    - [x] 13.2. Embedding HTML templates
14. [Testing](#testing)
    - [ ] 14.1. Unit testing and sub-tests
    - [ ] 14.2. Testing HTTP handlers and middleware
    - [ ] 14.3. End-to-end testing
    - [ ] 14.4. Customizing how tests run
    - [ ] 14.5. Mocking dependencies
    - [ ] 14.6. Testing HTML forms
    - [ ] 14.7. Integration testing
    - [ ] 14.8. Profiling test coverage
15. [Conclusion](#conclusion)
16. [Further reading and useful links](#further-reading-and-useful-links)
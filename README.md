# 30DaysGolangChallenge

The tasks listed are designed to be approachable within a 30-minute session, but it's important to recognize that some tasks may take more than 30 minutes, especially if they involve learning new concepts or debugging. The idea is to break each exercise into meaningful chunks that can be tackled day-by-day, even if you need to carry some tasks over to the next day.

To help keep each session within the 30-minute timeframe, I suggest the following strategy:

1. Preparation: Review the task ahead of time. Understanding what you need to accomplish will allow you to dive right in.
2. Break Down Tasks: Sometimes, a task can be broken down into even smaller sub-tasks. For example, "Implementing endpoints" might be broken down into "Setting up the route" and "Writing the handler function."
3. Pacing Yourself: Don’t rush through exercises. If you find yourself hitting a roadblock, note where you left off and pick it back up the next day.
4. Minimal Viable Product (MVP) Approach: Aim to get a basic version working first and then iterate to improve it. For example, for an HTTP server, get it to respond with a basic message before adding complex routing and middleware.


## Week 1: Building a CLI Todo List Manager

### Day 1: Project Setup and Basic CLI Structure

* Initialize a new Go project and install the cobra package.
* Set up a basic CLI application with a root command.


### Day 2: Adding Basic Commands

* Implement the add command to add a new todo item.
* Create a simple Todo struct.


### Day 3: Storing Data

* Implement basic file-based storage (reading and writing JSON).


### Day 4: Listing Todo Items

* Add a list command to display all todo items from the file.


### Day 5: Marking Todos as Completed

* Implement a complete command and update the stored data.


### Day 6: Removing Todo Items

* Add a remove command to delete a todo item by its ID.


### Day 7: Refactoring and Testing

* Refactor the code for better modularity and write unit tests.


## Week 2: Developing a Simple Web-based Blog

### Day 8: Setting up a Web Server

* Initialize a new Go project and set up a basic web server.


### Day 9: Handling Routes

* Implement routing for home and individual post views.


### Day 10: Designing the Blog Post Struct

* Define a BlogPost struct and create in-memory storage.


### Day 11: Serving HTML Templates

* Use the html/template package to render a list of blog posts.


### Day 12: Creating and Viewing Blog Posts

* Implement an endpoint to create new blog posts with basic form handling.


### Day 13: Editing and Deleting Blog Posts

* Add endpoints to edit and delete blog posts.


### Day 14: Database Integration

* Integrate SQLite for persistent blog post storage (focus on simple CRUD operations).


## Week 3: Developing a Simple RESTful API for a Bookstore

### Day 15: Setting up the Project

* Initialize a new project and define API routes.


### Day 16: Designing the Book Struct

* Define a Book struct and create in-memory storage.


### Day 17: Implementing the API Endpoints (Part 1)

* Implement POST and GET endpoints for creating and listing books.


### Day 18: Implementing the API Endpoints (Part 2)

* Implement PUT and DELETE endpoints for updating and deleting books.


### Day 19: Middleware for Logging

* Implement a logging middleware.


### Day 20: Authentication Middleware

* Add basic JWT-based authentication middleware.


### Day 21: Data Validation and Error Handling

* Improve data validation and error handling for all endpoints.

## Week 4: Building a Concurrent Web Scraper

### Day 22: Project Setup and HTTP Requests

* Initialize a new project and write a function to fetch web pages.


### Day 23: Parsing HTML Content

* Use golang.org/x/net/html to parse HTML and extract data.


### Day 24: Structuring and Storing Data

* Define a struct for scraped data and save it to JSON.


### Day 25: Adding Concurrency

* Use goroutines to fetch multiple web pages concurrently.


### Day 26: Rate Limiting and Retrying

* Implement rate limiting and retry logic.


### Day 27: Error Handling

* Improve error handling to ensure the scraper recovers from errors.


### Day 28: Persisting Results

* Save the scraped data to a database.


## Final Project: Integrating Skills

### Day 29: Planning and Setting Up

Plan the mini project, set up the project structure, and create basic routes.
### Day 30: Implementation

Implement key functionalities prioritizing MVP (maybe start with basic CRUD operations).
By segmenting the tasks this way and focusing on minimal viable implementations, you’re more likely to complete each task within 30 minutes. If you find some tasks taking longer, remember it’s perfectly okay to extend into the next session! The key is consistent, structured practice. Happy coding! 🚀
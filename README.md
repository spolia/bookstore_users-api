# bookstore_users-api

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

Users API Service

## Overview

This application handles users.

## Project Structure

This project was written in golang, designed to use MVC architecture and DAO pattern to abstract data persistence.

![alt](assets/users-diagram.png)


## Endpoints

- `GET  /user/:user_id` : Returns the user data.
  
- `POST /users` : Saves a new user. 
  
- `PUT  /user/:user_id` : Updates a user.
  
- `PATCH /user/:user_id` : Updates some user data. 
  
- `DELETE /user/:user_id`: Deletes a user.


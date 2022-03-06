# Golang API

## To run this project follow the bellow steps:

### 1. Setting .env variables
1. Create a `.env` file. 
1. Copy `.env.example` content, paste in your `.env` file and finally set your database variables.

### 2. Setting database

#### To create a database and table for this example run in your postgres SGBD:

* `create database golang;`

* `create table students(id SERIAL, name varchar not null, course varchar not null );`
* `create table teachers(id SERIAL, name varchar not null, subject varchar not null );`

### 3. Running the code

1. Make sure you already have `build-essential` installed on your computer, if not, run in terminal `sudo apt update` and `sudo apt install build-essential`
1. In your terminal within directory, run `make run`
1. Check `specifications` folder to run your requests.

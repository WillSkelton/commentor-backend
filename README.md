# Commentor Backend

## Table Of Contents:
- [Commentor Backend](#commentor-backend)
  - [Table Of Contents:](#table-of-contents)
  - [Description](#description)
  - [Usage](#usage)
    - [Setup](#setup)
    - [Starting the server](#starting-the-server)
  - [Team Members](#team-members)
  - [Tools Used](#tools-used)

## Description
This program is called "Commentor". It will find all the function definitions of all the source code files in a given directory and give the user a clean GUI for adding comments to them. This repository contains the backend for the program.

## Usage
### Setup
The only setup step needed to start the backend is to clone the repository into your gopath (`go/src`) and switch to the branch `develop`. 

### Starting the server
To start the server, open the root directory of the repository in a terminal window and run the command
```
go run main.go
```
If you want to create an executable, you can run
```
go build -o main.exe
./main.exe
```
which will build and run the executable.

## Team Members
- Will Skelton
- Brett Anzlovar
- Lucas Mason
- Odeysiuss Tuon
- Peyton Urquhart

## Tools Used
Written in Golang and tested with goconvey.
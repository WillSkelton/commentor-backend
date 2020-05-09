# Commentor Backend
## Table Of Contents:
- [Commentor Backend](#commentor-backend)
  - [Table Of Contents:](#table-of-contents)
  - [Description](#description)
  - [Presentation and Demonstration:](#presentation-and-demonstration)
  - [Usage](#usage)
    - [Overview](#overview)
    - [Setup](#setup)
      - [Frontend](#frontend)
      - [Starting the Frontend](#starting-the-frontend)
      - [Backend:](#backend)
      - [Starting the Backend](#starting-the-backend)
  - [Team Members](#team-members)
  - [Tools Used](#tools-used)

## Description
This program is called "Commentor". It will find all the function definitions of all the source code files in a given directory and give the user a clean GUI for adding comments to them. This repository contains the frontend for the program.


## Presentation and Demonstration:
Here is the link to our presentation and a short demonstration of the application: https://youtu.be/Az5f4_Hr0hU


## Usage

### Overview
This application is made up of two components: a frontend and a backend which are two separate repositories that must be run in two different terminal windows. The repositories are as listed:
- Frontend: https://github.com/WillSkelton/commentor/tree/develop
- Backend: https://github.com/WillSkelton/commentor-backend/tree/develop

**Both Repositories must be on the `develop` branch**

### Setup
#### Frontend
To start the frontend, clone [this repository](https://github.com/WillSkelton/commentor/tree/develop), switch to the branch `develop`, and run the command

```
git pull origin develop
```

In order to start the frontend, [Yarn](https://yarnpkg.com/en/) must be installed. Once Yarn is installed, open the root directory of the repository in a terminal and run the command

```
yarn install
```

#### Starting the Frontend
This will download all the dependencies needed for the frontend to work. Once those are downloaded, run the command

```
yarn start
```
This will try start on port 3000 of localhost but will pick the next highest available port if 3000 is busy. During development, the frontend will be hosted on a local server but eventually, we will either bundle it together into huge javascript, html, and css files or in our case, a desktop application.

To open the frontend, just open a browser and type `localhost:3000` into the search bar and hit `enter`.

#### Backend:
Now you're half way there. To start the backend, `Golang` must be installed. You can follow the instructions here: https://golang.org/doc/install. Once you do that and configure your Gopath, clone [this repository](https://github.com/WillSkelton/commentor-backend/tree/develop) into your `Gopath`, switch to the branch `develop`, and run the command

```
git pull origin develop
```

#### Starting the Backend
Once you have all the files, you can just use the command
```
go run main.go
```
from the root folder of the repository and it'll start the backend on port `42201`.


## Team Members
- Will Skelton
- Brett Anzlovar
- Lucas Mason
- Odeysiuss Tuon
- Peyton Urquhart

## Tools Used
The frontend is made with React.js and Material UI, tested with Jest, managed with Yarn, and packaged with Electron.js.

The backend is written in Golang.
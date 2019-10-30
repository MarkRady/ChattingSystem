# Welcome to My first Chatting System with Golang and Docker

Highly concurrent app for chatting app using Revel framework and jobs 

### Requirements
    - Docker
    - Docker-machine
    - Docker-compose

### Start App:

   1. Clone repository
   ```bash
        git clone https://github.com/MarkRady/ChattingSystem.git
   ```
   2. run app 
   ```bash
        cd ChattingSystem/Docker
        docker-composer up
   ```
  3. it's take up-to 1mins to getting start after pulling images and build containers
  ####  Go to http://localhost:9000/ and you'll see:
  4) Apis documentation
 [![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/4cc1380527b2fb8309a5)

## Code Layout

The directory structure of a generated Revel application:

    src/instapp/conf/             Configuration directory
    src/instapp/app.conf          Main app configuration file
    src/instapp/routes            Routes definition file

    src/instapp/app/   App sources
        init.go        Interceptor registration
        controllers/   App controllers go here
        models/        Database models and init to database and elasticsearch

## TODO:
1) Implement env file for general configurations
2) figure-out better ways for concurrency 

## Challenges
1) Using Docker and Golang for first time [Done]
2) Requests to be processed concurrently and handling race conditions [Done]


# Welcome to My first Chatting System with Golang, Ruby On Rails and Docker.

Highly concurrent app for chatting app using Revel & RoR framework with sidekiq and Mysql.

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
  4.  Make sure services are running by visiting
  - http://localhost:9000 -> Golang app
  - http://localhost:9200 -> Elasticsearch service
  - http://localhost:3000 -> Ruby on rails app
  - http://localhost:3000/sidekiq -> track Sidekiq queues

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

The directory structure of a generated Ruby on rails application:

    src/ruby-app/config/                         Configuration directory
    src/ruby-app/config/database.yml             Database configurations
    src/ruby-app/config/sidekiq.yml              Workers and queue configurations
    src/ruby-app/config/application.rb           Cache driver & general configurations
    src/ruby-app/config/routes.rb                routes for resources and apis

    src/ruby-app/app/models                      Contain models based on ORM Active record for query database
    src/ruby-app/app/workers                     Contain workers for queues based on Sidekiq
    src/ruby-app/app/controllers/api/v1          Contain controllers for apis


## TODO:
1) Implement env file for general configurations
2) figure-out better ways for concurrency 

## Challenges i faced in this project
1) Using Docker [Done]
2) Using Golang and Revel as web framework [Done]
3) Using Ruby and Ruby on rails as web framework [Done]
4) Using Sidekiq queues driver for job and worker handling and tracking [Done]
2) Requests to be processed concurrently and handling race conditions [Done]



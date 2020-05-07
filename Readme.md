ChessBuddy
==========

Service that counts users which opened current page of website
WebSocket Support. If you open page in different tabs you can see asynchronous page update. 

Quick Start
-----------

Start application using docker-compose:

    docker-compose up

This will install a start the HTTP server inside docker container.
Visit <http://localhost:8080/> to see application results.

You can use `docker-compose dovn -v` to clear visiting history.

---
Start application without docker:
    
    make
    
After build run application:

    ./server


Missing / Planned Features
--------------------------

* testing
* configuration from env variables
* session support for distinguish between users
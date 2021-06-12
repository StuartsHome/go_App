## Go Application
To build container:
- make docker-build

To start container:
- make docker-start

To stop container:
- make docker-start

With service running, query port:
- http://0.0.0.0:1001/



Go application with:
- Concurrency
- Database: MySQL
- Testing: Testify, Mockery
- Resty (Simple HTTP and REST client library)


## Containerisation
- Docker, Docker-compose, Makefile


### :wrench: To Do
- Database 
    - Firebase (NoSQL)
    (MySQL)
- Testing (Testify, Mock)
    - Mock and stub REST server and http calls



## Notes
// you need to download that module and record its version in your go.mod file. 
// The go mod tidy command adds missing module requirements for imported packages
// and removes requirements on modules that aren't used anymore.



## Docker cmds
- docker-compose logs -f (stream)
- docker-compose exec server bash
- 

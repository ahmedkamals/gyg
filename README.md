Prerequisites
=============

- Install git, golang, php7.1, composer, docker, docker-compose

Installation
============

Go inside backend folder, and then run:

    composer install

Copy the middleware folder to be under `$GOPATH/src/gyg/`, so the final paht should be `$GOPATH/src/gyg/middleware/.....`

Install golang dependencies, run:

    go install .
    go build .

Then you should have binary called "middleware", run:

    ./middleware

Go inside "backend" folder and run:

    composer install

Go inside "backend/src" folder and run:

    php -S localhost:8080

Go inside "frontend" folder and open `index.html`

Middleware
==========

- The middleware is configurable "middleware/config/main.json", where the backend configuration be changed,
    also the path to the config file can be overridden during run time by passing the option
     `--config.path Config_Path` which has the default value of `config/main.json`.
- The fetching process is done concurrently with loading rating data.

Backend
========
### Coding
- Used Symfony coding standard.
- Used composer to install dependencies.
- Applied `PSR-0` for loading files.
- Used exception for handling errors.
- Used standard php error codes.

### Modifications
- Changed the query to use GET instead of post to ease the visual tracing,
to ease the caching in case of http caching like Varnish, also search data won't be that big.
- Changed the header to use direct json rendering instead of an attachment, as this might slow the backend when having
concurrent requests, as the first requests get the file handle, and the subsequent requests would have to wait
in the queue to get  the file handle.
- Changed the price to be returned as float.

### Code flow
 - A call is made to a service which accepts a search record.
 - A search record is passed to a data reader inspired from Domain Driven Design "DDD" terminologies.
 - A search record composed from a search data, and a search strategy.
 - A find by city strategy has a criteria which is `EqualToCriteria`.

Todo
----
- Using docker containers.
- Caching processed data at middleware layer.
- Applying code coverage using unit test.
- Using Apache thrift or Google protocol buffer for large outputs.
- Using Value objects as in DDD.
- Optimizing logs.
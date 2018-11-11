# Fibonacci REST API

This API receives a number and returns the fibonacci sequence up to that number. _Please note the sequence starts with 0 in this API._ This API was created as part of a code challenge to dive in and learn more about Golang.

### Available endpoints

- `/api/` Is the home endpoint, returns a list of available endpoints
- `/api/fibonacci/:num` Returns the fibonacci sequence in JSON (ie: `/api/fibonacci/6` returns {[0,1,1,2,3,5]}))

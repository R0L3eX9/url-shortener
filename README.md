# Go Short
A URL Shortener written in Go to learn more about Go and Redis.


## Endpoints
POST - /api/create -> requires the url to be shorten

GET - /app/<shorten_url> -> redirects user to the url linked to the shorten_url


## TODO
* [x] Run Redis on localhost
* [x] Successfully set a value in the Redis db
* [x] Successfully read a value from the Redis db
* [x] Successfully redirect the user to the desired address 
* [x] Check if the generated URL is unique
* [x] Change seed for the random URL
* [ ] Add a frontend

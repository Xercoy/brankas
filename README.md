I wasn't able to complete a large portion of the assessment, most of the time taken was developing successful tests for the POST endpoint. I had a difficult time both manually and automatically sending a POST request with the correct information, therefore the handler is also incomplete. 

To the best of my ability in the design and development of content that was familiar, it took about 4 hours for me to develop this. 

Endpoints:
GET /routes 
POST /routes

I was able to ensure that that the GET request contained an auth token. 

To run, simply do a build and run with a flag to set the auth token (there's additional flags such as the maximim image size in bytes):
```
go build -o foo; ./foo --auth-token="SECRETPASSWORD"
```

The default port is 8049.

To view a full report of the failing/passing tests, run the tests in verbose mode:
```
go test -v
```

This report gives a list of the tests that had passed and the ones that failed.

I tested the GET endpoint such that the response body would be the correct HTML with the auth token provided.

I intended to test the POST endpoint in multiple ways:
- within and outside of the file limit
- a mime type that did and did not conform to the handler
- an accepted upload with the correct and incorrect token
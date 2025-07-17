# Simple Go HTTP Server

This started out as a CodeCrafters.io guided project.

You can visit the original challenge [here](https://app.codecrafters.io/courses/http-server/introduction)

I dropped off that course after the 4th or so stage, and have decided to make this an actual HTTP server package which could be pulled into any project and used to serve HTTP requests.

This won't be a production-grade HTTP server and should *NOT* be used in an actual web-based application, but I would like for its final state to be such that it could be used in that regard.

---

### Current State
- [x] Serve basic HTTP Requests (a successful 200 OK or 404 Not Found)
- [x] Send user-defined response headers
- [x] Automatically calculate and send `Content-Length` response headers
- [x] Concurrent request handling, allow multiple client requests to be executed at once
- [x] Parse out request headers
- [ ] Make use of certain key request headers that alter behaviour
- [x] Allow users to pull give their own request paths and content-returning functions
- [ ] Better error handling to be as resilient as possible
- [x] Logging
- [x] Cleanup project so that it can be used as a library in a basic HTTP application
- [ ] Whatever else I can think of and realistically implement!

---

Thank you for checking out this repo and let me know if you have any suggestions, or if it helped you learn some TCP basics!

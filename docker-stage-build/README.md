# Multi Stage Docker Builds

# To build
`docker build -t nick .`

# To Run
1. `docker run -p 50000:50000 nick`
2. In another terminal: `curl http://localhost:50000/user`
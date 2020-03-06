# Introduction 
A simple REST API that I use to test Azure DevOps, Kubernetes, Docker and other tools/methodologies 

# Build
* cd src
* docker build -t bjd145/whatos:1.4 .

# Run 
* docker run -p 8081:8081 -e API_VERSION=v3 bjd145/whatos:1.4

# Test
* curl http://127.0.0.1:8081/api/os

  {"Time":"Friday, 06-Mar-20 16:12:35 UTC","Host":"4921deb4d042","OSType":"linux","Version":"v3"}
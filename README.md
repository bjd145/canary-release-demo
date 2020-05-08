# Introduction 
  A repo to test Azure DevOps, Kubernetes, Docker and other tools or methodologies 

# Build
* cd src
* docker build -t bjd145/whatos:1.4 .

# Run 
* docker run -p 8081:8081 -e API_VERSION=v3 bjd145/whatos:1.4

# Test
* curl http://127.0.0.1:8081/api/os
  {"Time":"Friday, 06-Mar-20 16:12:35 UTC","Host":"4921deb4d042","OSType":"linux","Version":"v3"}

# Canary Release with Blue/Green Deployment and Helm

## 10% Traffic to V2
* helm upgrade -n whatosapi -i --values .\deploy\values.yaml \
  --set service.greenWeight=10%,service.blueWeight=90%,blueReplicaCount=3,greenReplicaCount=1 \ 
  whatos-release \
  .\deploy

## 50% Traffic to V2
* helm upgrade -n whatosapi -i --values .\deploy\values.yaml \
  --set service.greenWeight=50%,service.blueWeight=50%,blueReplicaCount=2,greenReplicaCount=2 \
  whatos-release \
  .\deploy

## 90% Traffic to V2
* helm upgrade -n whatosapi -i --values .\deploy\values.yaml \
  --set service.greenWeight=90%,service.blueWeight=10%,blueReplicaCount=1,greenReplicaCount=3 \
  whatos-release \
  .\deploy

## 100% Traffic to V2
* helm upgrade -n whatosapi -i --values .\deploy\values.yaml \
  --set service.greenWeight=100%,service.blueWeight=0%,blueReplicaCount=0,greenReplicaCount=4 \
  whatos-release \
  .\deploy
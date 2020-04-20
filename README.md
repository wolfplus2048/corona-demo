# corona-demo

### Installing
clone the app
```
git clone https://github.com/wolfplus2048/corona-demo.git
```
### Running the tests
run server
```
docker-compose -f ./deploy/docker-compose.yml up -d etcd nats
go run -trimpath -gcflags "all=-N -l" main.go
```
run client
open browser => http://localhost:3334/web/

You should see the message: [hello corona]

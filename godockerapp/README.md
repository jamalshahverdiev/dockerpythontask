##### In the current folder just execute the following command to create executable file which will be used inside of the `Dockerfile`:
```bash
$ go build -o gowebapp
```

##### Build and push new image:
```bash
$ docker build -t jamalshahverdiev/gowebapp .
$ docker push jamalshahverdiev/gowebapp:latest
```

##### Result of the API must be like s followwing:
- / - returns back `Success!`
- /ping - returns back `Ok`

#### To run and test the current image use the following commands:
```bash
$ docker container run -d --name goapi -p 8080:8080 jamalshahverdiev/gowebapp:latest
$ curl -XGET http://localhost:8080/
Success!
$ curl -XGET http://localhost:8080/ping
Ok
```

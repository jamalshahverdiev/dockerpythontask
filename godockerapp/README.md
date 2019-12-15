##### In the current folder just execute the following command to create executable file which will be used inside of the `Dockerfile`:
```bash
$ go build -o gowebapp
```

##### Build and push new image:
```bash
$ docker build -t jamalshahverdiev/gowebapp .
$ docker push jamalshahverdiev/gowebapp:latest
```

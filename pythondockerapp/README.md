#### With the following code files we will prepare docker image to push some app with 2 routings `/` and `/ping`

#### From `python:3.6.1-alpine` image `Dockerfile` does prepare environment with `requirements.txt` file and executes `pyapp.py` with 2 routings which will be exposed to the port `8080`.

- / - returns back `Success!`
- /ping - returns back `Ok`

#### Note if you want to push your image like as public to the `https://hub.docker.com` registry server use this command `docker login` with the your own credentials.

#### Execute the following command inside of the current folder to prepare image and push it to the public `https://hub.docker.com` registry server
```bash
$ docker build -t jamalshahverdiev/pyapp .
$ docker push jamalshahverdiev/pyapp:latest
```

#### To run and test the current image use the following commands:
```bash
$ docker container run -d --name pyapi -p 8080:8080 jamalshahverdiev/pyapp:latest
$ curl -XGET http://localhost:8080/
Success!
$ curl -XGET http://localhost:8080/ping
Ok
```


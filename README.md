# mooapi

This is a simple Golang API that I'm using to test and demo. 


## Makefile

This repository includes a [Makefile](https://www.gnu.org/software/make/manual/make.html). I've included it so that all the commands you need to test, build, and run this project are documented.

### fmt
`gofmt -w -s -d .`

[Gofmt](https://go.dev/blog/gofmt) is a tool that automatically formats Go source code. 

### test
`go test -v`

### build
`docker build --pull -t mooapi .`

### run
`docker run --rm -it -p 8080:8080 mooapi`

### all

Will run the fmt, test, and build targets as they are documented above.


## Demo

### Build
To format, test, and build mooapi execute `make`.

```console
$ make
make
=== RUN   TestIndexHandler
--- PASS: TestIndexHandler (0.00s)
=== RUN   TestHelloHandler
--- PASS: TestHelloHandler (0.00s)
=== RUN   TestGreetHandler
=== RUN   TestGreetHandler/NoName
=== RUN   TestGreetHandler/WithName
--- PASS: TestGreetHandler (0.00s)
    --- PASS: TestGreetHandler/NoName (0.00s)
    --- PASS: TestGreetHandler/WithName (0.00s)
PASS
ok      github.com/triangletodd/mooapi  0.002s
[+] Building 5.9s (11/11) FINISHED                                                                                                         docker:default
        => [internal] load build definition from Dockerfile                                                                                                 0.0s
        => => transferring dockerfile: 205B                                                                                                                 0.0s
        => [internal] load metadata for docker.io/library/golang:alpine                                                                                     0.4s
        => [auth] library/golang:pull token for registry-1.docker.io                                                                                        0.0s
        => [internal] load .dockerignore                                                                                                                    0.0s
        => => transferring context: 49B                                                                                                                     0.0s
        => [internal] load build context                                                                                                                    0.0s
        => => transferring context: 15.02kB                                                                                                                 0.0s
        => CACHED [builder 1/4] FROM docker.io/library/golang:alpine@sha256:9bdd5692d39acc3f8d0ea6f81327f87ac6b473dd29a2b6006df362bff48dd1f8                0.0s
        => [builder 2/4] COPY . .                                                                                                                           0.1s
        => [builder 3/4] RUN go test -v                                                                                                                     4.1s
        => [builder 4/4] RUN go build -o /mooapi                                                                                                            0.7s
        => CACHED [stage-1 1/1] COPY --from=builder /mooapi /mooapi                                                                                         0.0s
        => exporting to image                                                                                                                               0.0s
        => => exporting layers                                                                                                                              0.0s
        => => writing image sha256:657d7399197a5f0a429f5df7ea7a61d9319d504158ee16e0757e360eb3a27f96                                                         0.0s
        => => naming to docker.io/library/mooapi
$
```

### Run
To run the docker container execute `make run` after building. Press ctrl+c to quit and delete the container.

```console
$ make run
Starting server at port 8080
```

In another terminal use curl to hit the API.

```console
$ curl localhost:8080
Moo

$ curl  -H "User-Agent: Not Curl 0.0.1" localhost:8080/moo
{"message":"This endpoint is for curl users only!"}

$ curl localhost:8080/moo

                                         ______________
                                        < Hello, curl! >
                                         --------------
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣾⡟⠷⢦⡄⠀⠀⠀⠀/⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡶⢶⡦⣄⡀⠀⠀⠀⢀⣤⠞⠻⠛⠃⠀⠈⠛⢦⡀⠀/⠀⠀⠀⠀⢀⣀⣀⡀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⠉⠒⠬⢍⡉⠉⣹⠃⠀⠀⠀⠀⠀⠀⠀⠀⢹⠒⢶⣶⡾⠛⠛⠁⠀⠙
⠀⠀⠀⠀⠀⠀⠀⢀⣠⠾⠋⠉⠉⠉⠛⠲⠤⢤⣀⣀⣀⡘⣧⡀⠀⠀⠀⠀⠈⣹⠁⠀⠀⠀⠀⠀⠀⠀⢀⠀⠸⡦⠊⠁⠀⠀⠀⠀⢄⡼
⠀⠀⠀⢀⣀⡤⠶⠿⠅⠀⢀⠀⠔⠊⠀⠀⠀⠀⠀⠀⠈⠉⠙⠛⠮⣄⣀⣠⣼⣧⡄⠀⠀⠀⠆⠀⠀⣤⣼⠲⣄⢀⡀⢀⣀⣤⣤⠴⠋⠀
⠀⠀⢀⡟⠁⠀⠀⠀⠀⠀⠀⠀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢘⣧⠀⠀⠀⠀⠀⠀⠉⠀⢰⠃⠈⠉⠉⠉⠁⠀⠀⠀⠀
⠀⠀⣸⠁⠀⠀⡰⠁⠀⠀⠀⠊⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡴⠀⠀⠀⠀⠀⣸⠃⠀⠀⠀⠀⠀⠁⠀⠀⣼⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⢰⣿⠀⠀⣰⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠞⠀⠀⠀⠀⠀⣰⡃⠀⠀⠀⠀⠀⠀⠀⢀⣴⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢠⠇⣿⠀⢠⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠎⠀⠀⠀⠀⠀⠀⢿⢻⡍⢠⡄⠀⠀⢠⣠⠎⣼⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⡸⢰⡏⠀⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡏⠀⠀⠀⠀⠀⠀⠀⠈⠻⣗⣚⣿⣴⣺⠟⢁⡼⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⣇⣸⡇⠀⠸⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠀⠀⠀⢀⡔⠀⠀⠀⠀⠀⠀⠉⠉⠉⠀⢀⡞⠁⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⣇⢸⡇⠀⠀⢿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠇⠀⠀⠀⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡠⠊⠀⠀⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢹⣸⡇⠀⠀⠘⣷⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠞⠁⠀⠀⠀⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠈⡏⢷⠀⠀⠀⠘⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠃⠀⠀⠀⠀⣼⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠘⣼⡀⠀⠀⠀⣼⢷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠈⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢻⡇⠀⠀⢠⢿⠀⠉⠛⢷⠶⠤⢤⣄⣀⣀⣀⣹⠰⠀⠀⠀⠀⣠⠄⠀⠀⠀⠀⠀⣀⡀⠀⢠⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢸⠁⠀⠀⡏⢸⠀⠀⠀⡿⠀⠀⠀⠀⠀⠀⠈⠙⡇⠀⠀⠀⢰⣇⠀⠀⠀⢀⡠⢾⠉⠀⠀⣸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢺⠀⠀⢸⠇⢸⠀⠀⣼⠃⠀⠀⠀⠀⠀⠀⠀⠀⢻⡀⠀⠀⠀⡏⠓⠦⠴⠋⠁⢸⠀⠀⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢸⡄⠀⡾⠀⡞⠀⢰⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣇⠀⠀⠀⡇⠀⠀⠀⠀⠀⣼⠀⠀⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠘⠃⢀⡇⢰⠁⠀⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⢰⠃⠀⠀⠀⠀⠀⢻⠀⠀⢰⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢀⡀⠈⣇⡸⠀⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⢸⠀⠀⠀⠀⠀⠀⢸⡀⠀⡜⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⣼⠀⠀⢸⡇⠀⠠⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⠀⠀⢸⠀⠀⠀⠀⠀⠀⢸⠀⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⢰⡟⠀⠀⠚⡧⣄⣀⣸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⡆⠀⣼⠀⠀⠀⠀⠀⠀⣸⠀⠀⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠻⠤⠤⠖⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⠀⢹⡀⠀⠀⠀⠀⠀⢹⡀⠀⠸⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⠃⠀⠈⣇⠀⠀⠀⠀⠀⠀⣧⡀⠐⠻⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡞⠒⠀⠐⠻⡄⠀⠀⠀⠀⠀⠙⠒⠒⠊⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠳⣤⣠⠤⠴⠃⠀
```


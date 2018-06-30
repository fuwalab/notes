# notes
A simple online notes.

## Requirements (Local Environment)
- [dev-env](https://github.com/fuwalab/dev-env)
    - Simple developing tool using vagrant
- [docker](https://github.com/fuwalab/docker)
    - Docker container, contains Go

## Preparation

1. Clone Repositories / Packages
    ```bash
    $ git clone git@github.com:fuwalab/dev-env.git
    $ cd dev-env
    $ git submodule init
    $ git submodule update
    $ cd docker/data/go
    $ git clone git@github.com:fuwalab/notes.git
    $ cd notes
    $ glide install
    ```

1. Run Vagrant
    ```bash
    $ cd path/to/dev-env
    $ vagrant up
    ```
1. Login to vagrant
    ```bash
    $ vagrant ssh
    ```

1. Login to docker container(on Vagrant)
    ```bash
    $ cd /vagrant/docker/
    $ docker-compose restart
    $ docker exec -it go sh
    ```
    
1. Run go project
    ```
    /go # go run src/notes/*.go

       ____    __
      / __/___/ /  ___
     / _// __/ _ \/ _ \
    /___/\__/_//_/\___/ v3.3.5
    High performance, minimalist Go web framework
    https://echo.labstack.com
    ____________________________________O/_______
                                    O\
    ⇨ http server started on [::]:5000
    ```

1. See the following
    - http://go.dev-env.fuwalab
    
## For the production environment
- Build the go package (in the go container)
    ```
    /go # go build -o bin/notes notes
    /go # ./bin/notes &
    ```
    
## Go web bootstrap
This is a sample repository to bootstrap any new web app I'd wish to write in go. The purpose of this repository is to save a lot of time when trying to write a new webapp and setting up everything from scratch.

There are a lot of common components which I wish I'd not have to write everytime while creating any new project and this repository is going to contain exactly that.

## What this repository does?
* Create a `main.go` which is the entry point of the program. Among other things it, 
    * tries to load config from local.json when the `$ENV` is `LOCAL` via [viper](https://github.com/spf13/viper).
        * loads the config for database, redis, http server and log level.
    * configures [zap](https://github.com/uber-go/zap) as the application's logger and write the logs to stdout.
    * creates connection to postgres using [gorm](https://gorm.io/)
    * creates connection to redis using [go-redis](https://github.com/redis/go-redis)
    * creates an http server with chi's mux as it's handler, runs on port 8080
    * Creates an status handler endpoint which can handle shallow and deep healthchecks for postgres and redis.
* Uses generics, so go 1.18+ is must.
* Creates a justfile, which can be used to directly run the project.
* Creates `.gitignore`.
* Creates gh actions workflows for linter tests and go tests on PR's and push to main.

## How to use it?
Simply download the repository or clone it. Then,
```bash
go mod edit -module github.com/vishal1132/bootstrap-go-web; find . -type f -name '*.go'   -exec sed -i '' -e 's/github.com\/vishal1132\/bootstrap-go-web/<new-module-name>/g' {};
```

## External dependencies it use-> 
```bash
❯ go list -m -f '{{if not .Indirect}}{{.Path}}{{end}}' -mod=mod all | rg -v $(cat go.mod| head -n 1 | awk -F ' ' '{print $2}')
github.com/go-chi/chi/v5
github.com/redis/go-redis/v9
github.com/spf13/viper
go.uber.org/zap
gorm.io/driver/postgres
gorm.io/gorm
```

* Chi-> Chi is a light weight wrapper over net/http and is performant and compatible with the stdlib, so that's the default choice for routing http endpoints. With it's first class support as middleware and only having to implement an interface to response that to json makes it a good candidate for the routing part.

* go-redis-> go redis is a default choice I have seen so far for the redis client communication via go. It's current major version at the time of writing this project is v9.

* viper-> viper is also a standard package to manage configs, since the configs are going to be loaded before the service starts serving any traffic, it doesn't have to be performant just extensible and reliable, which I have found it to be very.

* zap-> Zap is only behind in performance to the zerlog, but offers a lot more features and the performance tradeoff isn't worth to sacrifice the features it provides.

* gorm-> gorm is an accessibility wrapper over the `database/sql` package in the stdlib and again the performance impact isn't drastically different and there are ways to write raw sqls for querying updating or anything complex which would bear cognitive load to do it in the `gorm` way. The other `gorm.io/driver/postgres` package provides the `pgx` driver to connect to the postgres server.

## Configured with

* [go-gorp](https://github.com/go-gorp/gorp): Go Relational Persistence
* [RedisStore](https://github.com/gin-gonic/contrib/tree/master/sessions): Gin middleware for session management with multi-backend support (currently cookie, Redis).
* [Gin Framework](https://gin-gonic.github.io/gin/)

### Installation

```
$ go get github.com/frnd/schedule-hub
```

```
$ cd $GOPATH/src/github.com/frnd/schedule-hub
```

```
$ go get -t -v ./...
```

> Sometimes you need to get this package manually
```
$ go get github.com/bmizerany/assert
```

You will find the **database.sql** in `db/database.sql`

And you can import the postgres database using this command:
```
$ psql -U postgres -h localhost < ./db/database.sql
```

## Running Your Application

Use docker-compose to start required services:

```
$ docker-compose up
```

Start the application with:
```
$ go run *.go
```

## Building Your Application

```
$ go build -v
```

```
$ ./schedule-hub
```

## Testing Your Application

```
$ go test -v ./tests/*
```

## Contribution

You are welcome to contribute to keep it up to date and always improving!

---

## License
(The MIT License)

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

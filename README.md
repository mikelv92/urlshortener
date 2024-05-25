# URL shortening app

To run the app:

```bash
$ docker run -p 27017:27017 mongo
```
```bash
$ go run ./cmd/api
```

To create a short URL:

```bash
$ curl -X POST -d '{"longurl":"https://github.com/mikelv92"}' -H "Content-Type: application/json" \
http://localhost:4000/create

{"LongURL":"https://github.com/mikelv92","ShortURL":"TlklzNUf"}
```

Then visit `http://localhost:4000/TlklzNUf`

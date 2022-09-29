# Passing parameter to postgres not working?

Not sure if I'm doing something wrong, or if there's a genuine problem. Using `sqlc` to generate my database code, using `pgx` under the covers.

Running this up with `docker compose up` results in:

`{"level":"fatal","error":"Cannot encode []store.Fruit into oid 16392 - []store.Fruit must implement Encoder or be converted to a string","time":"2022-09-29T21:55:53Z"}`
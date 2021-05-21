Flint has three http end-points. Flint will preserve the Content-Type 
of the PUT request. The path is the key. The body is the value.

| Method | Path | Body | Description               |
| ------ | ---- | ---- | ------------------------- |
| PUT    | *    | *    | create a key/value        |
| GET    | *    | none | get a value by its key    |
| DELETE | *    | none | delete a value by its key |
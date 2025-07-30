# Flint <img src="https://static.wikia.nocookie.net/minecraft/images/6/67/FlintNew.png" alt="flint" width="36"/>


Flint is a key/value store, where the key is the HTTP request path, and the value is the HTTP request body.  Flint will preserve the Content-Type of the PUT request. Flint has three HTTP endpoints. 

| Method | Path | Body | Description               |
| ------ | ---- | ---- | ------------------------- |
| PUT    | *    | *    | create a key/value        |
| GET    | *    | none | get a value by its key    |
| DELETE | *    | none | delete a value by its key |

## Quick Start


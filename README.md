
# Flint <img src="https://static.wikia.nocookie.net/minecraft/images/6/67/FlintNew.png" alt="flint" width="36"/>


A general purpose config server.

Basic URL Structure:

```
GET http://{host}/conf?app={name}&env={name}&v={version}
```



## Git as Storage

Flint can use a git repository as storage. Here are the basic terminologies.

| Flint Term  | Git Term              | Note                                |
| ----------- | --------------------- | ----------------------------------- |
| application | repository            | configured when starting the server |
| environment | branch                | required                            |
| version     | tag                   | optional, defaults to latest commit |
| properties  | files and directories | files read in as objects            |

In general filenames should not contain any value that would be illegal in a variable name in your programming language. since Flint is not designed for any specific programing language, it follows some generally safe rules: begin with a letter, followed by any number of letters, digits, and underscores. the exact pattern: `^[a-zA-Z_][a-zA-Z0-9_]*[.](?:yaml)$`. Any filename that does not match this regular expression will be ignored.

If you would like the properties of a file to be flattened to the parent directory object, there is a special filename for that `_.yaml`. 

A simple way of getting Flint to ignoring a file or directory would be to begins the name with a `.` dot. However there is also a formal way of ignoring files in the `flint.yaml` file.

Flint only works with one file extension at a time, per application. This is to prevent naming collisions when converting file and directory names to property names, since file systems generally do not allow the same name filename and extension to be included in a single directory. The file extension is defaults to `yaml`, but can be configured in the `flint.yaml` settings file.

A single Flint cluster can be configured to sever multiple applications and has a REST API for managing them, to allow for live changes.

### Example

Here is a hypothetical example.

- app name: `foo` (maps to a repo)
- branch name: `dev`
- tag: `3.2.1`


directory structure:

```
.
|-- http.yaml
|-- db
    |-- _.yaml
    |-- mysql.yaml
    |-- redis.yaml
```

`http.yaml`

```yaml
port: 8080
```

`db/_.yaml`

```yaml
logs: https://log.eample.com
```


`db/mysql.yaml`

```yaml
host: localhost:3306
user: user1
pass: passwd1
```

`db/redis.yaml`

```yaml
host: localhost:6379
```



**request:**

```bash
curl "http://flint.example.com/conf?app=foo&env=dev&v=3.2.1"
```

**response:**

```json
{
    "http": {
        "port": 8080
    },
    "db": {
        "logs": "https://log.example.com",
        "mysql": {
            "host":"localhost:3306",
            "user":"user1",
            "pass":"passwd1"
        },
        "redis": {
            "host":"redis.example.com"
        }
    }
}
```

## configuring settings

There is one filename that is reserved for settings: `flint.yaml`. The file may only exist at the root of an applications directory. The contents of this file will never be included in a response, and is strictly for configuring the flint server itself. The file allows you to point a Flint server to a repository or a directory in a repository, and then manage the application specific behavior from the repository itself.

### Properties

The properties of `flint.yaml`.

| Property  | Value         | Description                                                                 | Example                              |
| --------- | ------------- | --------------------------------------------------------------------------- | ------------------------------------ |
| token     | random string | An optional security measure. required when specified                       | `token: a1b2c3` e.g. `?token=a1b2c3` |
| root      | path          | The base path to begin searching for config files, relative to the app root | `root: conf/`                        |
| ignore    | list of paths | Relative paths to ignore (from flint root), supports wildcards              | `ignore: ["*private.yaml"]`          |
| extension | string        | The extension of the files that will be parsed, default is `yaml`.          | `extension: json`                    |

## Management End-Points


```bash
# list all apps being served.
GET /apps

# add or replace an app.
PUT /apps/{name}

# remove an app.
DELETE /apps/{name}
```
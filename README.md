# Shrls

A simple to deploy URL shortener written in Go.

## Features

  - [x] Admin web interface
  - [-] Single binary contains server and administration cli
  - [x] Any short url becomes a QR code by appending `.qr` to the end of it.
    - [x] Text rendered QR codes available for terminals
  - [x] File Uploads
  - [ ] Curl-able API
  - [x] Automatically strip tracking url parameters from configurable hosts
  - [-] Includes a convenient, feature rich bookmarklet
    - Shorten a URL or save a screenshot of the entire webpage.

## Deploying

Shrls is provided as a docker container to deploy to wherever you would like.
Docker-compose is recommended for new users and a sample
[docker-compose](./docker-compose.yml) file has been provided within the
repository.

## Configuring

Shrls can currently be configured via environment variables and/or a config file existing within one of the following paths.

  - /etc/shrls/config.yaml
  - $HOME/.config/shrls/config.yaml
  - ./config.yaml

The current configuration can be viewed by running `shrls config`. This is a convenient way to get started with a configuration.

By default the configuration looks like this

```yaml
host: localhost
port: 3000
default_redirect: /admin
state:
  bolt:
    path: shrls.db
uploads:
  directory:
    path: uploads
```

If you would like to instead use MongoDB, specify the connection string as en environment variable.
```sh
export SHRLS_MONGO_CONNECTION_STRING="mongodb://username:password@localhost:27017"
```

This can also be saved within the configuration file as

```yaml
state:
  mongodb:
    connection_string: "mongodb://username:password@localhost:27017"
```

Additional environment variables and their descriptions are found below

### SHRLS_HOST

The public facing hostname of the url shortener service.

Used for generating QR codes on the server

### SHRLS_PORT

The port which the url shortener service should run on.

Defaults to `3000`

### SHRLS_MONGO_CONNECTION_STRING

A database connection string to MongoDB.

### SHRLS_DB_PATH

If using the built-in database, a path to where the database should be saved.

Defaults to `shrls.db`

### SHRLS_USERNAME

If set, requires basic auth to access the web admin interface

### SHRLS_PASSWORD

If set, requires basic auth to access the web admin interface

### SHRLS_DEFAULT_REDIRECT

If a short url does not exist, redirect the user to the configured default_redirect. Otherwise 404.

### SHRLS_DEFAULT_REDIRECT_SSL

If the default redirect and QR codes should use `https` in their redirects, set this value to `true`

Defaults to `false`

### SHRLS_UPLOAD_DIRECTORY

Path which contains files uploaded to Shrls.

Defaults to `./uploads`

## Contributing

Development environments are provided by Nix and should come with everything
needed to build Shrls from scratch.

### Architecture Refresh

Shrls is currently in a revamp to migrate from being strictly a server to being
an all-in-one server and admin cli. This design is entirely inspired by how
Hashicorp designs software. While the resultant binary is larger, there is only
one tool which you need to download. This is also forcing components that used
to be within a single namespace to be separated into logical namespaces for the
project, making it much easier to extend in the future.

### Go-Shrls Directory Layout

- Server/

  Interfaces and proto definitions for service and state layers.

- Service/

  User facing APIs. All gRPC methods ending with Request/Response go no deeper.
  Instead the protobuf objects within the request are sent to ServerState and
  the protobuf objects received from ServerState are translated back to
  Response objects.

  Contains implementation of Shrls gRPC service.

- State/

  Implementations of different backend states. Should always accept and respond
  with protobuf.

- Pkg/

  Any additional helpers for the service.

- Cmd/

  Main entrypoints for the application.

- UI/

  The frontend UI for the web server.

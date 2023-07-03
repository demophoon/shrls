# Shrls

A simple to deploy URL shortener written in Go.

## Features

  - Admin web interface
  - Single binary contains server and administration cli
  - Any short url becomes a QR code by appending `.qr` to the end of it.
    - Text rendered QR codes available for terminals
  - Curl-able API
  - Automatically strip tracking url parameters from configurable hosts
  - Includes a convenient, feature rich bookmarklet
    - Shorten a URL or save a screenshot of the entire webpage.
  - File Uploads

## Configuring

Shrls can currently be configured via environment variables and/or a config file existing within one of the following paths.

  - /etc/shrls/config.yaml
  - $HOME/.config/shrls/config.yaml
  - ./config.yaml

The options which can be configured are listed below. You can also configure any of these with the `shrls config` command.

### base_url

The base url of the url shortener service.

### mongo_uri

*Required*

A database connection string to MongoDB.

Defaults to `mongodb://mongo:password@localhost:27017`

### admin_username

If set, requires basic auth to access the web admin interface

### admin_password

If set, requires basic auth to access the web admin interface

### port

The HTTP port of the Shrls service.

Defaults to `3000`

### grpc_port

The gRPC port of the Shrls service.

Defaults to `3001`

### default_redirect

If a short url does not exist, redirect the user to the configured default_redirect. Otherwise 404.

### upload_directory

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

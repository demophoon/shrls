# Shrls design

Shrls is currently in a revamp to migrate from being strictly a server to being
an all-in-one server and admin cli. This design is entirely inspired by how
Hashicorp design's software.

## Go-Shrls Directory Layout

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

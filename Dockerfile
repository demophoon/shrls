FROM golang:alpine AS base
RUN apk update && apk add ca-certificates && update-ca-certificates 2>/dev/null || true
RUN mkdir /new_tmp

FROM nixos/nix:2.24.9 AS builder

RUN nix-channel --add "https://github.com/NixOS/nixpkgs/archive/63dacb46bf939521bdc93981b4cbb7ecb58427a0.tar.gz" nixpkgs && nix-env --install direnv

COPY . /build
WORKDIR /build
RUN direnv allow /build && \
    direnv exec /build make dist

# Final Artifact
FROM scratch
COPY --from=base /new_tmp /tmp
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/shrls /shrls
WORKDIR /config

ENTRYPOINT ["/shrls"]
CMD ["serve"]

EXPOSE 3000

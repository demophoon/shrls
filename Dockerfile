FROM nixos/nix:2.24.9 AS builder

RUN nix-channel --add "https://github.com/NixOS/nixpkgs/archive/63dacb46bf939521bdc93981b4cbb7ecb58427a0.tar.gz" nixpkgs && nix-env --install direnv

COPY . /build
WORKDIR /build
RUN direnv allow /build && \
    direnv exec /build make dist

RUN mkdir /new_tmp

# Final Artifact
FROM scratch
COPY --from=builder /new_tmp /tmp
COPY --from=builder /build/shrls /shrls
WORKDIR /config

ENTRYPOINT ["/shrls"]
CMD ["serve"]

EXPOSE 3000

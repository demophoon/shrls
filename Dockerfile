FROM nixos/nix AS builder

RUN nix-channel --update && nix-env --install direnv

COPY . /build
WORKDIR /build
RUN direnv allow /build && \
    direnv exec /build make dist

# Final Artifact
FROM scratch
COPY --from=builder /build/shrls /shrls
CMD ["/shrls", "serve", "--trace"]

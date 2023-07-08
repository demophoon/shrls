# TODO: Add build steps (with Nix build environment)

FROM scratch

WORKDIR /
COPY ./shrls /shrls

CMD ["/shrls", "serve", "--trace"]

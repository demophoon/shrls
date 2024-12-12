let
  pkgs = import <nixpkgs> {
    config.allowUnfree = true;
  };
in
  pkgs.mkShell {
    buildInputs = with pkgs; [
      # Development Shell
      gotools
      gopls
      go-outline
      gopls
      gopkgs
      godef
      golint
      delve
      asmfmt
      errcheck
      reftools
      golangci-lint
      gomodifytags
      gotags
      impl
      iferr
      delve

      # Build environment
      gnumake42
      go

      # Backend
      buf
      grpcurl

      # Frontend
      nodejs_20
      nodePackages.npm
    ];
    shellHook = ''
    export GOPATH=$(pwd)/.gopath
    export PATH=$GOPATH/bin:$PATH
    '';
  }

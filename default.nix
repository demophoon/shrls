let
  pkgs = import <nixpkgs> {};
in
  pkgs.mkShell {
    buildInputs = with pkgs; [
      # Development Shell
      gotools
      gopls
      go-outline
      gocode
      gopkgs
      gocode-gomod
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

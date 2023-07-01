let
  pkgs = import <nixpkgs> {};
in
  pkgs.mkShell {
    buildInputs = with pkgs; [
      go_1_18
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

      buf
    ];
    shellHook = ''
    export GOPATH=$(pwd)/.gopath
    export PATH=$GOPATH/bin:$PATH
    '';
  }

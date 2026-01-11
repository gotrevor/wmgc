{
  description = "wmgc-secret - keychain helper with caller verification";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = [
            pkgs.go
          ] ++ pkgs.lib.optionals pkgs.stdenv.isDarwin (with pkgs.apple-sdk.frameworks; [
            Security
          ]);
        };
      }
    );
}

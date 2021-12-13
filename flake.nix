{
  description = "IPLD Schema playground";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { system = system; };
      in
      {
        packages.ipldChecker =
          pkgs.buildGoModule rec {
            pname = "ipld-schema-playground";
            version = "0.0.1";

            src = ./.;

            vendorSha256 = pkgs.lib.fakeSha256;

            runVend = true;

            meta = with pkgs.lib; {
              description = "A IPLD Schema playground";
              homepage = "https://github.com/marcopolo/go-ipld-schema-playground";
              license = licenses.mit;
              platforms = platforms.linux ++ platforms.darwin;
            };
          };
        defaultPackage = self.packages.${system}.ipldChecker;
        devShell = pkgs.mkShell {
          buildInputs = [ pkgs.go pkgs.simple-http-server ];
        };
      });
}

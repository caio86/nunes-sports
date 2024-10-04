{
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      systems = [
        "x86_64-linux"
        "aarch64-linux"
      ];
      forAllSystems = function:
        nixpkgs.lib.genAttrs systems
          (system:
            function (nixpkgs.legacyPackages.${system})
          );
      node2nix = pkgs: import ./composition.nix { inherit pkgs; };
    in
    {
      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            nodejs_18
            nodePackages.typescript
            (node2nix pkgs)."@angular/cli-16.1.0"
          ];
        };
      });
    };
}

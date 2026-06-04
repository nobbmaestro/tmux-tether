{
  description = "A simple terminal UI for tmux sessions, written in Go!";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    systems.url = "github:nix-systems/default";
    treefmt-nix.url = "github:numtide/treefmt-nix";
    git-hooks-nix.url = "github:cachix/git-hooks.nix";
  };

  outputs =
    inputs@{
      self,
      flake-parts,
      systems,
      ...
    }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import systems;
      imports = [
        inputs.treefmt-nix.flakeModule
        inputs.git-hooks-nix.flakeModule
      ];

      perSystem =
        {
          config,
          pkgs,
          system,
          ...
        }:
        let
          goMod = builtins.readFile ./go.mod;
          versionMatch = builtins.match ".*go[[:space:]]([0-9]+\\.[0-9]+)(\\.[0-9]+)?.*" goMod;

          goVersion =
            if versionMatch != null then
              builtins.head versionMatch
            else
              throw "Could not extract Go version from go.mod";

          goOverlay = final: prev: {
            go = prev."go_${builtins.replaceStrings [ "." ] [ "_" ] goVersion}";
          };

          tmux-tether = pkgs.callPackage ./default.nix {
            src = ./.;
            gitCommit = self.rev or self.dirtyRev or "dev";
          };
        in
        {
          _module.args.pkgs = import inputs.nixpkgs {
            inherit system;
            overlays = [ goOverlay ];
            config = { };
          };

          packages = {
            default = tmux-tether;
            inherit tmux-tether;
          };

          devShells.default = pkgs.mkShell {
            name = "tmux-tether-dev";

            packages = [
              pkgs.go
              pkgs.gotools
              pkgs.golangci-lint
              pkgs.golines
              pkgs.gnumake
            ];
          };

          treefmt = {
            programs.nixfmt.enable = true;
            programs.gofmt.enable = true;
          };

          checks.build = tmux-tether;

          pre-commit = {
            settings.hooks = {
              treefmt = {
                enable = true;
                package = config.treefmt.build.wrapper;
              };
              golangci-lint.enable = true;
            };
          };

          apps.install-hooks = {
            type = "app";
            program = "${pkgs.writeShellScript "install-hooks" config.pre-commit.installationScript}";
          };
        };

      flake = {
        overlays.default = final: prev: {
          tmux-tether = inputs.self.packages.${final.system}.tmux-tether;
        };
      };
    };
}

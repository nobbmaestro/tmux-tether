{
  description = "A simple terminal UI for tmux sessions, written in Go!";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    systems.url = "github:nix-systems/default";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  outputs =
    inputs@{
      flake-parts,
      systems,
      ...
    }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import systems;
      imports = [
        inputs.treefmt-nix.flakeModule
      ];

      perSystem =
        {
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

          tmux-tether = pkgs.buildGoModule rec {
            pname = "tmux-tether";
            version = "dev";

            gitCommit = inputs.self.rev or inputs.self.dirtyRev or "dev";

            src = inputs.self;

            vendorHash = "sha256-komX1AmHt2NoF1x6xsNa2RFkfVzOXfYEMPhT0zwMxjw=";

            doCheck = true;

            nativeBuildInputs = with pkgs; [
              pkg-config
            ];

            ldflags = [
              "-X main.version=${version}"
              "-X main.commit=${gitCommit}"
            ];

            postInstall = ''
              ln -s $out/bin/tmux-tether $out/bin/tt
            '';

            meta = {
              description = "A simple terminal UI for tmux sessions, written in Go!";
              homepage = "https://github.com/nobbmaestro/tmux-tether";
              license = pkgs.lib.licenses.mit;
              maintainers = [ "nobbmaestro" ];
              platforms = pkgs.lib.platforms.unix;
              mainProgram = "tmux-tether";
            };
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
        };

      flake = {
        overlays.default = final: prev: {
          tmux-tether = inputs.self.packages.${final.system}.tmux-tether;
        };
      };
    };
}

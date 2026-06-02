<h1 align="center">
  Tmux-Tether (tt)
</h1>

<p align="center">
  A simple terminal UI for tmux session management, written in <a href="https://github.com/golang/go">Go</a>!
</p>

<p align="center">
  <a href="https://github.com/nobbmaestro/tmux-tether/releases">
      <img src="https://img.shields.io/github/v/release/nobbmaestro/tmux-tether" alt="GitHub Release" />
  </a>
  <a href="https://github.com/nobbmaestro/tmux-tether/commits/development">
      <img src="https://img.shields.io/github/last-commit/nobbmaestro/tmux-tether/development" alt="GitHub Last Commit" />
  </a>
  <a href="https://github.com/nobbmaestro/nvim-andromeda/compare/0.1.0...development">
      <img src="https://img.shields.io/github/commits-since/nobbmaestro/tmux-tether/0.1.0/development" alt="GitHub Commits Since" />
  </a>
  <a href="https://opensource.org/licenses/MIT">
      <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="MIT License" />
  </a>
</p>

## Table of Contents

- [Installation](#installation)
  - [Nix](#nix)
  - [Build from Source](#build-from-source)
- [Usage](#usage)
- [Customization](#customization)

## Installation

### Nix

```sh
nix build github:nobbmaestro/tmux-tether
# or from a local clone
nix build
```

To integrate into your own flake, add tmux-tether as an input and apply the overlay:

```nix
# flake.nix
{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    tmux-tether.url = "github:nobbmaestro/tmux-tether";
  };

  outputs = { nixpkgs, tmux-tether, ... }: {
    nixosConfigurations.myhost = nixpkgs.lib.nixosSystem {
      system = "x86_64-linux";
      modules = [
        {
          nixpkgs.overlays = [ tmux-tether.overlays.default ];
          environment.systemPackages = [ pkgs.tmux-tether ];
        }
      ];
    };
  };
}
```

The overlay works the same way with [Home Manager](https://github.com/nix-community/home-manager):

```nix
{
  nixpkgs.overlays = [ tmux-tether.overlays.default ];
  home.packages = [ pkgs.tmux-tether ];
}
```

### Build from Source

```sh
git clone git@github.com:nobbmaestro/tmux-tether.git
cd tmux-tether
make install
```

## Usage

Run `tt` to open the session picker. It will scan your configured directories, find session roots, and let you switch to or create the selected tmux session.

```sh
tt
```

Available flags:

| Flag | Description                |
| ---- | -------------------------- |
| `-c` | Print the user config      |
| `-C` | Print the default config   |
| `-d` | Print the config file path |

## Customization

Check out the [configuration docs](docs/config.md).

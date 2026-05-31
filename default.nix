{
  lib,
  buildGoModule,
  pkg-config,
  src,
  gitCommit ? "dev",
}:

buildGoModule rec {
  pname = "tmux-tether";
  version = "0.1.0";

  inherit src;

  vendorHash = "sha256-ABluLEPAIKKcJ1yZrxkPBYJvBd2RUz1NSmtBX/VrZnM=";

  nativeBuildInputs = [ pkg-config ];

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
    license = lib.licenses.mit;
    maintainers = [ "nobbmaestro" ];
    platforms = lib.platforms.unix;
    mainProgram = "tmux-tether";
  };
}

{ lib, buildGoModule }:
buildGoModule {
  pname = "lueur";
  version = "0.0.1";

  src = ./.;

  vendorHash = "sha256-/orMRiTvdIy0wvFfkNJzg6pzBk988Kyt6nfstJsmvVw=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "A gophermap walker for Markdown and HTML";
    homepage = "https://github.com/theobori/lueur";
    license = lib.licenses.mit;
    mainProgram = "lueur";
  };
}

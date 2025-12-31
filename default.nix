{ lib, buildGoModule }:
buildGoModule {
  pname = "lueur";
  version = "0.0.1";

  src = ./.;

  vendorHash = "sha256-x6ZaLaKB5e0bIaZeepjO1dJfjP7T8hcdGeh/8x3arQY=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "Renderer for Gophermap, gph and txt from Markdown and HTML";
    homepage = "https://github.com/theobori/lueur";
    license = lib.licenses.mit;
    mainProgram = "lueur";
  };
}

{ lib, buildGoModule }:
buildGoModule {
  pname = "lueur";
  version = "0.0.1";

  src = ./.;

  vendorHash = "sha256-14SmdbFerWjV/RiDaMC4WvN9eHwZYVpKUPoma58KEjs=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "Gophermap renderer for Markdown and HTML";
    homepage = "https://github.com/theobori/lueur";
    license = lib.licenses.mit;
    mainProgram = "lueur";
  };
}

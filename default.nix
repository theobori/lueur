{ lib, buildGoModule }:
buildGoModule {
  pname = "lueur";
  version = "0.0.1";

  src = ./.;

  vendorHash = null;

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "My project description";
    homepage = "My project homepage";
    license = lib.licenses.mit;
    mainProgram = "lueur";
  };
}

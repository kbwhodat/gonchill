{ lib, ...}:
self: super:
{
  gonchill = super.buildGoModule rec {
    pname = "gonchill";
    version = "1.0.6";

    src = ./.;

    vendorSha256 = null;

    meta = {
      description = "What whatever you want...";
      homepage = "https://github.com/kbwhodat/gonchill";
      license = lib.licenses.mit;
      maintainers = with lib.maintainers; [ kbwhodat ];
    };
  };
}

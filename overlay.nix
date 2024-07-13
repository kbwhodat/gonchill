{ lib, ...}:
self: super:
{
  gonchill = super.buildGoModule rec {
    pname = "gonchill";
    version = "1.0.6";

    src = ./.;

    vendorSha256 = "f11c215ec98a4665603b28b6c064948cf42435b48fe5386020046b75ca242e8b";

    meta = {
      description = "What whatever you want...";
      homepage = "https://github.com/kbwhodat/gonchill";
      license = lib.licenses.mit;
      maintainers = with lib.maintainers; [ kbwhodat ];
    };
  };
}

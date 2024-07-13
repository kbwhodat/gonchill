{ lib, ...}:
self: super:
{

  gonchill = super.buildGoModule rec {
    pname = "gonchill";
    version = "1.0.6";

    src = super.fetchFromGitHub {
      owner = "kbwhodat";
      repo = "gonchill";
      rev = "${version}";
      sha256 = lib.fakesha256;
    };

  };
}

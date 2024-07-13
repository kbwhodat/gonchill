{ lib, ...}:
self: super:
{

  gonchill = super.buildGoModule rec {
    pname = "gonchill";
    version = "1.0.6";

    src = super.fetchFromGitHub {
      owner = "kbwhodat";
      repo = "gonchill";
      rev = "v${version}";
      hash = "sha256-Gjw1dRrgM8D3G7v6WIM2+50r4HmTXvx0Xxme2fH9TlQ=";
    };

    vendorHash = null;

    meta = {
      description = "What whatever you want...";
      homepage = "https://github.com/kbwhodat/gonchill";
      license = lib.licenses.mit;
      maintainers = with lib.maintainers; [ kbwhodat ];
    };
  };
}

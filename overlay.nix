{ lib, ...}:
self: super:
{
  gonchill = super.stdenv.mkDerivation rec {
    pname = "gonchill";
    version = "1.0.6";

    src = ./.;

    buildInputs = [ super.go ];
    sourceRoot = ".";
    phases = [ "unpackPhase" "installPhase" ];

    installPhase = ''
      runHook preInstall

      go build -o gonchill
      mkdir -p $out/bin
      cp gonchill $out/bin/

      runHook postInstall
    '';

    meta = with lib.super.lib; {
      description = "Watch whatever you want...";
      homepage = "https://github.com/kbwhodat/gonchill";
      license = lib.licenses.mit;
      maintainers = [ lib.maintainers.kbwhodat ];
    };
  };
}

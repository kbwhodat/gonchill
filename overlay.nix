self: super: {
  gonchill = super.buildGoModule rec {
    pname = "gonchill";
    version = "1.0.10";

    src = ./.;

    vendorHash = "sha256-Ov++3fC39zPv0CPZvneE0slc5jAewhr917xIggR5jms=";

    meta = with super.lib; {
      description = "Watch whatever you want...";
      homepage = "https://github.com/kbwhodat/gonchill";
      license = licenses.mit;
      maintainers = [ maintainers.kbwhodat ];
    };
  };
}

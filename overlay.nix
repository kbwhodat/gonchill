self: super: {
  gonchill = super.buildGoModule rec {
    pname = "gonchill";
    version = "1.0.10";

    src = ./.;

    vendorHash = "sha256-/9JDYnHfn4do1LZf3jVcGdoJ9W9s3uugCNRa+x+tpyE=";

    meta = with super.lib; {
      description = "Watch whatever you want...";
      homepage = "https://github.com/kbwhodat/gonchill";
      license = licenses.mit;
      maintainers = [ maintainers.kbwhodat ];
    };
  };
}

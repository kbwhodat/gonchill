{ pkgs ? import <nixpkgs> {}}:

pkgs.mkShell {
  packages = with pkgs; [ 
    python313Packages.websockets
    python313Packages.numpy
    python313Packages.aiofiles
    python313Packages.matplotlib
    python313Packages.scipy
    python313Packages.platformdirs
    python313Packages.aiohttp
    python313Packages.jsondiff
    python313Packages.orjson
    python313Packages.selenium
    (python313.withPackages (ps: with ps; let

      seleniuim_driverless = ps.buildPythonPackage rec {
        pname = "selenium-driverless";
        version = "1.9.4";
        format = "setuptools";
        build-system = [ setuptools ];
        src = pkgs.fetchurl {
          url = "https://files.pythonhosted.org/packages/5e/92/3fcf637eebbc334543de61b319c4f00d01526053edf33c2f25aa08f05c13/selenium_driverless-1.9.4.tar.gz";
          sha256 = "151ccf57d399691ec4e943a941a496dbe575d0154a520cc2eca988ebe5d07a76";
        };

        doCheck = false;
      };

      cdp-socket = ps.buildPythonPackage rec {
        pname = "cdp-socket";
        version = "1.9.4";
        format = "setuptools";
        build-system = [ setuptools ];
        src = pkgs.fetchurl {
          url = "https://files.pythonhosted.org/packages/7d/28/58812797e54fb8cf22bff61125e5a7d2763de1a86855549ecc417bdd06d5/cdp-socket-1.2.8.tar.gz";
          sha256 = "d8a3d55883205c7c45c05292cf5ef5a5c74534873e369e258e61213cce15be1a";
        };

        doCheck = false;
      };
    in
      [
      cdp-socket
      seleniuim_driverless
    ]))
  ];
}



# Go And Chill (gonchill)
_________________________________________________________


If you want to watch things just use this. Use a VPN first if your ISP are annoying...

![](images/output.gif)

#### Depedencies
- **mpv** - for the -m option
- **vlc** - for the -v option
- **go**
- [peerflix](https://github.com/mafintosh/peerflix)
- python3
- selenium-profiles

##### Python Dependencies
```python3
pip install -r requirements.txt
```

## How to Install
#### Arch Linux
1. ```yay -S gonchill```

#### From Source
1. ```git clone https://github.com/kbwhodat/gonchill.git```

#### NixOS
Getting this installed requires a few more steps.
1. Python packages need to be installed, [here](https://github.com/kbwhodat/configs/blob/main/nix-config/common/packages/packages.nix#L55-L82) is how you can do it.
2. In your `flake.nix`, `undetected-chromedriver` and `gonchill` overlays need to be used. [Here](https://github.com/kbwhodat/configs/blob/main/nix-config/flake.nix#L23-L36) is how you can make that happen.
3. Then finally, install the `gonchill` [package](https://github.com/kbwhodat/configs/blob/main/nix-config/common/packages/packages.nix#L6), and now you should be good to go.


## How to use
#### Series
```gonchill series -v true detective```

#### Movies
```gonchill movies -v equalizer```

### License 📜
This project is licensed under [GPL-3.0](https://raw.githubusercontent.com/Illumina/licenses/master/gpl-3.0.txt).

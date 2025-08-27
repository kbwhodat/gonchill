
# Go And Chill (gonchill)
_________________________________________________________


If you want to watch things just use this. Use a VPN first if your ISP are annoying...

![Who wants to watch Alien Earth? I do!](images/alien.gif)

#### Depedencies
- **mpv** - for the -m option
- **vlc** - for the -v option
- **go**
- [peerflix](https://github.com/mafintosh/peerflix)
- python3
- selenium-profiles

## How to Install
#### Arch Linux
1. ```yay -S gonchill```

#### From Source
1. ```git clone https://github.com/kbwhodat/gonchill.git```
2. ```pip install -r requirements.txt```
3. ```go build .```
4. ```./gonchill series -m aliens```

#### NixOS
Getting this installed requires a few more steps.
1. Python packages need to be installed [here](https://github.com/kbwhodat/configs/blob/main/nix-config/common/personal/default.nix#L24-L132) is how you can do it.
2. In your `flake.nix`, `undetected-chromedriver` and `gonchill` overlays need to be used. [Here](https://github.com/kbwhodat/configs/blob/main/nix-config/flake.nix#L28-L32) is how you can make that happen.
3. Then finally, install the `gonchill` [package](https://github.com/kbwhodat/configs/blob/197fb70d2b961615078d406ea1fc0e8ad62030a4/nix-config/common/personal/default.nix#L10), and now you should be good to go.
4. `sudo nixos-rebuild switch --flake <dir_to_flake>#machine`
5. Once it's done building you should be able to utilize the `gonchill` cli 


## How to use
#### Series
```gonchill series -v true detective```

#### Movies
```gonchill movies -v equalizer```

### License 📜
This project is licensed under [GPL-3.0](https://raw.githubusercontent.com/Illumina/licenses/master/gpl-3.0.txt).

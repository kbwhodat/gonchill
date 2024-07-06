
# Go And Chill - Your Go-To GO CLI Torrent Streamer 🚀
_________________________________________________________


Welcome to `gonchill`! Dive into the world of instant streaming with the speed of light. Built with ❤️ in Go, `gonchill` lets you stream your favorite content directly via torrents, without the wait. Whether it's movies, series, or shows, `gonchill` brings them to you swiftly, because who likes waiting anyway?

![](images/output.gif)

## Features ✨
- _CLI Magic_: Pure command-line bliss. Simple commands, powerful streaming.
- _Versatile_: Supports a wide range of torrent sources. If it's out there, you can stream it.
- __Open Source_: Peek under the hood; it's all transparent and open for contributions!

#### Depedencies
- **mpv** - for the -m option
- **vlc** - for the -v option
- **go**
- [peerflix](https://github.com/mafintosh/peerflix)
- **chromedriver** - ONLY if you don't want to use the one provided in repo

##### Python Dependencies
```python3
pip -r install requirements.txt
```


## How to Install
##### Arch Linux
```yay -S gonchill```

#### From Source
1. ```git clone https://github.com/kbwhodat/gonchill.git```
2a. ```go run . movies -v avengers``` or ```go run . series -v billions``` - For vlc
2b. ```go run . movies -m avengers``` or ```go run . series -m billions``` - For mpv


## How to use
#### Series
```gonchill series -v true detective```

#### Movies
```gonchill movies -v equalizer```

### License 📜
This project is licensed under [GPL-3.0](https://raw.githubusercontent.com/Illumina/licenses/master/gpl-3.0.txt).

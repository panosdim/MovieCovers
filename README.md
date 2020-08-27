# AutoTag

Application that scans a folder for MP4 and MKV movie files and automatic attach cover art in the files.
In order to find the cover art it uses the [TMDB API](https://developers.themoviedb.org/3/getting-started/introduction)

## Usage

AutoTag can be run manually or installed an run periodically in a systemd compatible system.
To run it manually use

```
$ autotag -p "/mnt/Movies" -e "^Test[0-9]"
```

AutoTag accepts the following two parameters:

```
-e string
        Regex to exclude from scan
-p string
        Path to scan for movies (required)
```

## Installation

In order to install use the install.sh script.

## Libraries Used

-   [GO TMDB](https://github.com/ryanbradynd05/go-tmdb)
-   [Java properties scanner](https://github.com/magiconair/properties)

# marauders

<div align="center">
<a href="https://github.com/9g10f/marauders"><img src="https://img.shields.io/github/stars/9g10f/marauders?style=flat&logo=github&color=orange" alt="Stars"></a>
<a href="https://github.com/9g10f/marauders/network/members"><img src="https://img.shields.io/github/forks/9g10f/marauders?style=flat&logo=github&color=orange" alt="Forks"></a>
<a href="https://codeberg.org/9g10f/marauders"><img src="https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fcodeberg.org%2Fapi%2Fv1%2Frepos%2F9g10f%2Fmarauders&query=%24.stars_count&label=stars&logo=codeberg" alt="Stars"></a>
<a href="https://codeberg.org/9g10f/marauders/forks"><img src="https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fcodeberg.org%2Fapi%2Fv1%2Frepos%2F9g10f%2Fmarauders&query=%24.forks_count&label=forks&logo=codeberg" alt="Forks"></a>
<a href="https://codeberg.org/9g10f/marauders/tags"><img src="https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fcodeberg.org%2Fapi%2Fv1%2Frepos%2F9g10f%2Fmarauders%2Ftags&query=%24.0.name&label=tag&color=white" alt="Tag"></a>
<a href="LICENSE"><img src="https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fapi.github.com%2Frepos%2F9g10f%2Fmarauders%2Flicense&query=%24.license.spdx_id&label=license&color=white" alt="License"></a>
<br>
<img src="https://img.shields.io/badge/platform-Linux%20%7C%20Windows%20%7C%20FreeBSD%20%7C%20macOS-green" alt="Platforms">
</div>

<br>

<div align="center">
<a href="https://codeberg.org/9g10f/marauders"> <img src="https://img.shields.io/badge/Codeberg%20mirror-black?logo=codeberg&style=for-the-badge" alt="Codeberg mirror"> </a>
</div>

---

marauders is the CLI server for Marauder Launcher.

## Installation
You can install `marauders` through Go if you have it installed:
```
go install codeberg.org/9g10f/marauders@latest
```

## Usage
To run `marauders` normally you can do:
```
marauders run --port <port> --s-path <path-to-s-directory> --log-file <path-to-log-file> --log-level <log-level>
```
That will run `marauders` and all the logs will be printed. If you want to run `marauders` in `daemon` mode you can run it with `start` instead:
```
marauders start --port <port> --s-path <path-to-s-directory> --log-file <path-to-log-file> --log-level <log-level>
```
That way you can continue using your shell after starting `marauders` and later check the logs in your `log-file`.

## How to make install scripts
The script file's name is it's id, it doesn't need to be a number, nothing will break if it's id is a text, but that's the standard. In the script itself you have 6 very simple commands: `set`, `download`, `unzip`, `rm`, `rsynca` and `mv`. You can use comments by starting the line with `#`.

### set
You use `set` to define script variables. You can define a variable like this:
```
set <variable-name> = <variable-value>
```
Script variables can either be global or local. If you define a variable inside a version (after a `@`) it is local and is only set for the specific game version, else it is global and is set for every version. Local variables overwrite global variables.

### download
You use `download` to download file(s). It supports both regular HTTP(S) requests and BitTorrent magnets:
```
download <url>/<magnet>
```
`download` writes all files in the game's install directory (`$path`).

### unzip
`unzip` uses `7z` to unzip all files from a zipped file:
```
unzip <path>
```
It unzips everything to it's own directory.

### rm
The `rm` command behaves much like the regular Unix `rm`:
```
rm <path>
```
It takes no options or flags and by default removes files, empty and filled directories.

### rsynca
You use `rsynca` much like you would use `rsync -a`, but due to specific `marauderctl` syntax you separate the 2 paths with `|`:
```
rsynca <path1>|<path2>
```
It's useful for applying game patches, you move a whole `Fix Repair` directory to the main game directory replacing every existing file.

### mv
The `mv` command behaves much like the regular Unix `mv`, but due to specific `marauderctl` syntax you separate the two paths with `|`:
```
mv <path1>|<path2>
```
It moves files and directories from `path1` to `path2`. Can also be used to rename files/directories.

## For the developers
If you wish to include a donation/merch button in your game's page open an issue with a link to your donation/merch page and we'll add it.

## License
marauders is free and licensed under the [GNU General Public License v3.0](LICENSE).
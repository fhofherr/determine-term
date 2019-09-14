# `detterm`

`detterm` is a simple tool that tries to determine if a shell is
being executed from within a terminal emulator.

## Installation

```sh
make
sudo make install
```


## Usage

`detterm` starts with its own process id and walks the tree of
parent processes looking for one of the passed terminal emulators. If it
finds such a terminal emulator it writes the found emulators name to
standard out.

```sh
detterm alacritty konsole
```

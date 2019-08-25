# determine-term

`determine-term` is a simple tool that tries to determine if a shell is
being executed from within a terminal emulator.

## Usage

`determine-term` starts with its own process id and walks the tree of
parent processes looking for one of the passed terminal emulators. If it
finds such a terminal emulator it writes the found emulators name to
standard out.

```sh
determine-term alacritty konsole
```

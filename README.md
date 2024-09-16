# tlsmtools - TL File Splitting and Merging Tools

English | [简体中文](README_zh_CN.md)

`tlsmtools` is a simple command-line utility designed to split and merge TL files.

## Features

- **Split**: Split a `tl` file into `orig` and `alt` parts.
- **Merge**: Merge `orig`, `alt`, and `text` files into a single `tl` file.

## Installation

Ensure that you have Go installed on your system. Currentlt only tested on Go version `1.23.1`.

```bash
git clone https://github.com/MZWNET/tlsmtools.git
cd tlsmtools
go build
```

## Usage

### Split a File

```bash
tlsmtools split -f [file path] [-d [output directory]]
```

- `-f` or `--file`: Path to the file you want to split (required).
- `-d` or `--outputDir`: Directory to save the split files (default is `.`).

### Merge Files

```bash
tlsmtools merge -orig [orig file path] -alt [alt file path] -text [text file path] [-o [output file path]]
```

- `-orig`: Path to the `orig` file (required).
- `-alt`: Path to the `alt` file (required).
- `-text`: Path to the `text` file (required).
- `-o` or `--output`: Output file name and path (default is `output.tl`).

## License

This project is licensed under [The Unlicense](http://unlicense.org/).

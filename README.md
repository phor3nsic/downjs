
<div align="center">

```

	██████╗  ██████╗ ██╗    ██╗███╗   ██╗     ██╗███████╗
	██╔══██╗██╔═══██╗██║    ██║████╗  ██║     ██║██╔════╝
	██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║     ██║███████╗
	██║  ██║██║   ██║██║███╗██║██║╚██╗██║██   ██║╚════██║
	██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║╚█████╔╝███████║
	╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝ ╚════╝ ╚══════╝
```
</div>
<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/deeplooklabs/downjs?color=56BEB8">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/deeplooklabs/downjs?color=56BEB8">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/deeplooklabs/downjs?color=56BEB8">

  <img alt="License" src="https://img.shields.io/github/license/deeplooklabs/downjs?color=56BEB8">
</p>

## ##
## Intro
**downjs** is a versatile Go script designed to simplify the process of downloading JavaScript files from various URLs and perform checks for the existence of associated .map files. This functionality is invaluable for developers and analysts who want to streamline the process of collecting JavaScript files for further analysis and debugging.

[@xen00rw](https://github.com/xen00rw)

It's your idea man! ;)

## Features

- **Effortless JavaScript Downloads:** With downjs, you can effortlessly download JavaScript files from a list of URLs without the need for manual downloads or complex scripts.

- **Map File Analysis:** The script also checks for the existence of associated .map files, which are essential for debugging and analyzing minified JavaScript code.

- **Flexible Configuration:** Customize the script to suit your needs by adjusting the list of URLs and file destinations.

- **Concurrent Downloads:** downjs leverages Go's concurrency capabilities, allowing you to download multiple files simultaneously, making the process faster and more efficient.

## Install

```bash
go install github.com/deeplooklabs/downjs@latest
```

## Usage

1. Create a text file named list.txt containing a list of URLs pointing to the JavaScript files you want to download. Each URL should be on a separate line.
2. Execute the script using the following command:

```bash

cat jslist.txt | downjs
```
3. Downjs will download the specified JavaScript files and check for the existence of .map files, saving them for future analysis in the same directory.
4. Access the downloaded files and .map files in the designated output folder for further analysis and debugging.

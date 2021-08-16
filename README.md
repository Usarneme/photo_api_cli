# photo_api_cli

A command line application for accessing photo gallery data via JSON API.

---

### Requirements

1. A terminal program such as iTerm on Mac, cmd prompt or GitBash on Windows, and Terminal on Linux
2. Permissions to compile and/or run applications on your machine
3. The Go programming language needed to build the executable

Visit [https://golang.org/doc/install](https://golang.org/doc/install) to get and install Go on your machine

---

### Creating the Executable & Usage

1. Clone this repository with `git clone https://github.com/Usarneme/photo_api_cli`
2. Enter project repository's main program entrypoint directory with `cd photo_api_cli/main/`
3. Build the executable with `go build`
4. Run the executable that was just created
4a. If on a Mac or POSIX-compliant Linux machine, run with the command `./main`
4b. If on Windows, run with the command `main.exe`
5. Run the command with the help flag `./main --help` to see these usage instructions:

```
[usage] - You can call this CLI application with an optional album number.
[usage] - Either provide a number: ./main 18
[usage] - or, with no arguments to see the first set of albums: ./main
```

---

### GoLang Installation Instructions

* See the installation instructions for Windows, Mac, or Linux at [https://golang.org/doc/install](https://golang.org/doc/install)
* Essentially, for Windows and Mac you can download the installer from the above link and run it. For Linux there is likely a package that can be installed with your package manager of choice, for example: `sudo apt-get install golang`. Try searching aptitude with `apt-cache search golang`
* Confirm Go is installed with `go version`

---

### Copyright &copy; 2021 Tom/Usarneme

Please open a pull request if any bugs are found!

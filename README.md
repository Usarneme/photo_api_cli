# photo_api_cli

A command line application for accessing photo gallery data via JSON API.

---

### Requirements

1. A terminal program such as iTerm on Mac, cmd prompt or GitBash on Windows, and Terminal on Linux
2. Permissions to compile and/or run applications on your machine
3. (Optional) The Go language needed to build your own binary. main/main is an executable compiled on a recent Mac so if you also have a Mac it should run without needing recompile.

Visit [https://golang.org/doc/install](https://golang.org/doc/install) to get and install Go on your machine

---

### Installation Instructions

1. Clone this repository with `git clone https://github.com/Usarneme/photo_api_cli`
2. Enter the new directory with `cd photo_api_cli`
3. Enter the main program entrypoint directory with `cd main`
4. Compile an executable with `go build`; this will generate a file called _main_

---

### Usage Instructions

- If on a Mac or POSIX-compliant Linux machine, run the compiled binary from within the main/ directory by typing `./main` in your terminal of choice
- Utilize `./main --help` to see usage instructions:

```
[usage] - You can call this CLI application with an optional album number.
[usage] - Either provide a number: ./main 18
[usage] - or, with no arguments to see the first set of albums: ./main
```

---

### Copyright &copy; 2021 Tom/Usarneme

Please open a pull request if any bugs are found!

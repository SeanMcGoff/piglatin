# piglatin
Here lies the CLI portion of the piglatin project :pig:

# Installation
At the moment, there is no other option but to build the project from source (sorry), but overall the process should be pretty painless.

You first need to download the source code. This can either be done by clicking that big green button that says "Code" and downloading the compressed code, or by simply typing:

```git clone https://github.com/SeanMcGoff/piglatin.git```

in your terminal of choice (must have git installed, [install Git here if not](https://git-scm.com/downloads)).

Next, enter the ```cmd/piglatin``` directory of the project and type

```make```

into your terminal to build the project (must have Go installed, [install Go hereif not](https://golang.org/dl/))

After that, the executable will be in the __bin__ folder. From there, you can move the executable wherever you want (i.e. /usr/bin if you are on a *nix system)

# Usage

```bash
piglatin -t "English text goes here"
``` 

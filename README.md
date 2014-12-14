# dotfile-manager

This is my attempt to write a basic package manager to let me install my 
dotfiles from various locations on GitHub. I chose to write a package manager
instead of keeping them all in one place such that I would be able to decouple
as much as possible.

### INSTALLING THE TOOL

To install *dotfile-manager* you have two primary options. The first is manually
download and build the source code, while the other is relying upon `go get`
instead.

First, `go get`: `go get github.com/crockeo/dotfile-manager`. Doing this will
put an executable into your `$GOPATH/bin` folder. Alternatively you can install
it through manually building the source code:

```bash
# Making the directory to clone into.
$ mkdir -p $GOPATH/src/github.com/crockeo

# Moving to the directory.
$ cd $GOPATH/src/github.com/crockeo

# Getting the code
$ git clone http://github.com/crockeo/dotfile-manager

# Moving into the directory.
$ cd dotfile-manager

# Building
$ go build

# Moving the executable into the GOPATH's bin.
$ mv ./dotfile-manager $GOPATH/bin
```

### INSTALLING PACKAGES

Assuming that `dotfile-manager` is in the `$PATH`, you can simply run:

```bash
$ dotfile-manager install <package/name>
```

The package name should be in the form of `author/repository`. For instance, my
repository for [vim configuration](http://github.com/crockeo/vimconfig) which is
hosted at `http://github.com/crockeo/vimconfig` would be installed like so:

```bash
$ dotfile-manager install crockeo/vimconfig
```

### LICENSE

Check out the LICENSE folder for specific licensing information. (Hint: It's
licensed under the MIT License.)

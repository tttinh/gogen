# Go Generator

Gogen is a CLI tool that will create scaffold structure for your gRPC application.

Install the cobra generator with the command `go install github.com/tinhtt/gogen/gogen@latest`.

Go will automatically install it in your `$GOPATH/bin` directory which should be in your $PATH.

Once installed you should have the `gogen` command available. Confirm by typing `gogen` at a command line.

## Supported Operation

### gogen init

The `gogen init [app]` command will create your initial application code
for you. It will populate your program with
the right structure so you can immediately start working with.

#### Initializing a module

**If you already have a module, skip this step.**

If you want to initialize a new Go module:

1.  Create a new directory
2.  `cd` into that directory
3.  run `go mod init <MODNAME>`

e.g.

```
cd $HOME/code
mkdir myapp
cd myapp
go mod init github.com/tinhtt/abc
```

#### Initializing a application

From within a Go module run `gogen init`. This will create a new bare-bones project
for you to edit.

You should be able to run your new application immediately. Try it with
`go run main.go`.

e.g.

```
cd $HOME/code/myapp
gogen init
go run main.go
```

The basic structure of this application should be similar to the following:

You now have a basic application up and running. Next step is to edit the files and customize them for your application.

Have fun!

<h1 align="center">
  <br />
  <img src="https://github.com/jeffotoni/goworkshopdevops/blob/master/godevops.png" alt="logo" width="700" />
  <br />
  <br />
  <br />
</h1>

# Golang for DevOps and Developers

This is a guide in Golang for devops. The guide aims to present an overview of the language from the basic to the intermediate. In this guide we will cover all the concepts that we believe are important to anyone who starts or even those who already program in Golang can use this guide as a reference and help in some way dozens or even hundreds of golang users.
There are thousands of references today when it comes to Golang, let's start at the beginning and we could not stop talking about [Golang Tour](https://tour.golang.org).

This site is one of the most important, it is here that we find and we take a lot of our doubts [Blog Golang](https://blog.golang.org/) it is simply fantastic.

Here is where we will find all the documentation [Doc Golang](https://golang.org/doc/) on this page is complete with everything we need in Golang, here is the faq [Faq Golang](https://golang.org/doc/faq) here we will find several cool things.
And this page is where we found all the specifications of Golang, it is very complete [Spec Golang](https://golang.org/ref/spec)
and here the package [Package](https://golang.org/src/).
## Contents

- [Overview](#overview)
- [Installation](#installation)
  - [Introduction](#introduction-installation)
    - [Linux](#linux)
    - [Test your installation](#test-your-installation)
    - [$GOPATH](#gopath)
    - [Workspace](#workspace)
    - [Outside GOPATH](#outside-gopath)
    - [Func Main](#func-main)
- [Clean Architecture](#clean-architecture)
  - [Introduction Clean Architecture](#introduction-clean-architecture)
    - [directory organization](#directory-organization)
- [Go commands](#go-commands)
  - [Go commands introduction](#go-commands-introduction)
  - [GC and GCCGO](#gc-and-gccgo)2
  - [go run](#go-run) 
  - [go build](#go-build)
    - [Build modes](#gobuildmodes)
    - [Go and C](#go-and-c)
  - [go install](#go-install)
  - [go test](#go-test)
  - [go clean](#goclean)
  - [go get](#goget)
  - [go tool](#gotool)
  - [go doc](#godoc)
  - [Go mod](#gomod)
    - [GO111MODULE](#)
    - [go mod init](#gomodinit)
    - [go mod verify](#gomodverify)
    - [go mod download](#gomoddownload)
    - [go mod vendor](#gomodvendor)
    - [go mod graph](#gomodgraph)
- [C-style](#cstyle)
  - [println](#println)
    - [Variables](#variables)
    - [Scopo](#scopo)
- [Constants](#constants)
  - [Control structures](#controlstructures)
    - [if/else](#ifelse)
    - [for](#for)
    - [range](#range)
- [Functions](#functions)
  - [introduction](#)
    - [return multiple values](#returnmulti) 
    - [Variadic Functions](#variadicfunc) 
    - [functions as a parameter](#funcparameter) 
    - [Closures](#closures)
    - [Recursion](#recursion)
    - [Asynchronous Functions](#asynchromous)
- [Array](#array)
  - [introduction](#)
    - [Declaring Array](#)
    - [Initializing Array](#)
    - [Accessing Array Elements](#)
    - [Array can not be resized](#)
- [Slice](#slice)
  - [introduction](#)
    - [Slices - make](#)
    - [Slices - append](#)
    - [Slices - range](#)
    - [Slices of Slices](#)
    - [Nil Slices](#)
    - [Slice len () and cap ()](#)
- [Struct](#struct)
  - [introduction](#)
    - [Allocation New](#allocationnew)
    - [Initializing a struct](#structinit)
    - [Pointer to a struct](#structpointer)
    - [Type Struct Json](#structjson)
    - [Anonymus Struct](#structanonymus)
    - [Struct Reflect](#structreflect)
- [Erros](#erros)
  - [introduction](#)
    - [How Error Control Works](#)
    - [Log.fatal](#)
    - [Log.panic](#)
    - [panic](#panic)
    - [Custom Errors](#customerrors)
    - [Interface Errors](#)
    - [Errors.New](#)
    - [fmt.Errorf](#)
- [My first packages](#package)
  - [introduction](#)
    - [fmt](#fmt)
    - [log](#log)
    - [time](#time)
    - [encoding/json](#encodingjson)
    - [os](#os)
    - [io/ioutil](#ioutil)
    - [bufio](#bufio)
    - [flag](#flag)
    - [strconv](#strconv)
    - [strings](#strings)
    - [net/http](#nethttp)
    - [errors](#packageerros)
    - [testing](#testing)
- [Date and Time](#)
  - [introduction](#)
    - [Date](#)
    - [Duration](#)
    - [Ticker](#)
    - [Sleep](#)
    - [Team](#)
    - [Parse](#)
    - [Format](#)
- [Pointer](#pointer)
  - [introduction](#)
    - [How Pointers in Go Works](#howpointers)
    - [When using pointers](#whenusingpointers)
- [Methods](#methods)
  - [introduction](#)
    - [Methods x Func](#methodsfunc)
    - [Interface](#interface)
- [Defer](#defer)
- [Interfaces](#interfaces)
  - [Introduction](#introducaointerface)
    - [OOP on Golang](#oopgolang)
    - [Empty interface](#emptyinterface)
    - [Methods and Interfaces](#methodsandinterfaces)
    - [Conversion vs. Assertion](#conversionvsassertion)
    - [Sets Methods](#setmethods)
- [Maps](#)
  - [introduction](#)
    - [Make](#)
    - [Type Int, string, interfaces](#)
    - [Struct](#)
    - [Mutant](#)
    - [Not Sync](#)
- [Json](#Json)
  - [introduction](#)
    - [Json marshal](#jsonmarshal)
    - [Json Unmarshal](#jsonunmarshal)
    - [json Encode](#jsonencode)
    - [Json Decode](#jsondecode)
- [Parse Json](#Json)
  - [introduction](#)
    - [Toml](#jsontoml)
    - [Yaml](#jsonyaml)
    - [Gcfg](#jsongcfg)
- [Tests / Golang](#)
  - [Introduction](#)
    - [Examples Test](#)
    - [go test](#)
    - [Mocks](#)
- [net/http Server](#)
  - [introduction](#)
    - [http.NewServeMux](#)
    - [http.HandlerFunc](#)
    - [http.Handle](#)
    - [http.Handler](#)
    - [http.Server](#)
    - [next.ServeHTTP](#)
    - [ListenAndServe](#)
    - [ListenAndServeTLS](#)
    - [Server.Shutdown](#)
    - [Middleware](#)
- [net/http Server Pages](#)
  - [introduction](#)
    - [http.FileServer](#)
    - [http.NotFound](#)
    - [Disable http.FileServer](#)
    - [http.Dir](#)
    - [http.StripPrefix](#)     
- [net/http Client](#)
  - [introduction](#)
    - [http.Transport](#)
    - [http.Client](#)
    - [http.Get](#)
    - [http,Post](#)
    - [http.NewRequest](#)
    - [Context.WithCancel](#)
- [Encryption](#)
  - [introduction](#)
    - [Uuid](#)
    - [Blowfish](#)
    - [sha256](#)
    - [base64](#)
    - [Md5](#)
    - [Jwt](#)
- [Socket](#)
  - [introduction](#)
    - [TCP](#)
    - [UPD](#)
    - [UNIX](#)
- [Websocket](#)
  - [introduction](#)
    - [WS SERVER](#)
- [Database/Sql](#) 
  - [introduction](#)
    - [Creating Connection](#)
    - [Creating singleton connection](#)
    - [Creating singleton thread safe connection](#)
    - [Create Table](#)
    - [Insert Table](#)
    - [Update Table](#)
    - [Select Table](#)
    - [Delete Table](#)
    - [Transaction](#)
    - [Errors](#)
- [Database relational](#)
  - [introduction](#)
    - [Postgres](#)
    - [Mysql](#)
    - [SqlLite3](#)
- [Database noSql](#)
  - [introduction](#)
    - [Mongo](#)
    - [InfluxDb](#)
    - [DynamoDB](#)
    - [BoltDB](#)
    - [Redis](#)
- [Files](#)
  - [introduction](#)
    - [Reading and Writing](#)
    - [Stdin](#)
    - [Stdout](#)
    - [Stderr](#)
    - [Reading and Writing to Files](#)
    - [GetEnv](#)
    - [IsDir](#)
    - [IsFile](#)
- [Sync.Maps](#)
  - [introduction](#)
    - [Sync.Maps](#)
    - [Load](#)
    - [Store](#)
    - [LoadOrStore](#)
    - [Delete](#)
- [Goroutine](#)
  - [introduction](#)
    - [Channel](#)
    - [Parallelism](#)
    - [GOMAXPROCS](#)
    - [WaitGroup](#)
    - [Race Conditions](#)
    - [Mutex](#)
    - [Atomicity](#)
    - [1 for N](#)
    - [N to 1](#)
    - [Fan Out / Fan In](#)
    - [Worker](#)
    - [Pipeline](#)
    - [Deadlock](#)
    - [Context](#)
    - [Ticker](#)
    - [Singleton Connect Thread Safe](#)
- [Pprof](#)
  - [Introduction](#)
    - [pprof basics](#)
    - [cpuprofile](#)
    - [memprofile](#)
    - [go tool pprof](#)
- [Documentation/API](#)
  - [Introduction](#)
    - [go doc](#)
    - [Documenting API with Swager](#)
    - [Documenting API with blueprint](#)
- [Lambdas/Serverless](#)
  - [Introduction](#)
    - [gofn](#)
    - [Running thousands of docker run](#)
    - [Dcoker build in our Role](#)
    - [Aws Lambda invoke golang](#)
    - [Aws Lambda RDS golang](#)
    - [Aws Lambda EC2 golang](#)
    - [Aws Lambda API-Gateway golang](#)
    - [Aws Lambda S3 golang](#)
    - [Aws Lambda creating and using awscli](#)
    - [Aws Lambda Deploy and using awscli](#)
    - [Aws SAM](#)
- [API REST](#)
  - [introduction](#)
    - [Clean Code](#)
    - [Types of architecture (hexagonal, layered)](#)
    - [Server Standalone](#)
    - [Server with nginx](#)
    - [Docker](#)
    - [Monitoring](#)
    - [Microservices](#)
    - [databases](#)     
     
### Overview
---
Go is a powerful language when it comes to competition and high performance, with a clean and efficient architecture Go grow growing year after year and each day their communities get bigger.
Some paradigms have been shattered to make it a high performance language where competition is one of its strengths. Go facilitates the creation of programs that take full advantage of multicore and networked machines, while the new type system allows the construction of flexible and modular programs.
It is a fast, statically compiled language that looks like a dynamically interpreted language. This characteristic becomes Golang a unique language as the subject is web.

### Installation
---

#### Introduction

We will download the file, unpack it and install it in /usr/local/go, if we have golang already installed in the machine we will have to remove the existing one to leave our installation as unique.
Let's create our directory in our workspace and test to see if everything went well

#### Linux

```bash
$ sudo rm -rf /usr/local/go
$ wget https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

#### $GOPATH

$GOPATH is the golang in your $HOME, this is necessary for your projects to use pkg and build properly. This was mandatory for all versions before version 1.11. The cool thing is that from now on we will not have to create projects in $GOPATH, we can create in any other directory that is not in $GOPATH.

Here is the link to the versioning proposal [Proposal: Versioned Go Modules](https://go.googlesource.com/proposal/+/master/design/24301-versioned-go.md/) or [Go 1.11 Modules](https://github.com/golang/go/wiki/Modules/)

We'll detail how to work with **go mod**, it was one of the best experiences I had for versioning projects using Golang.

Let's set up our environment to run Go. Add **/usr/local/go/bin** to the PATH environment variable. You can do this by adding this line to your **/etc/profile** (for a system-wide installation) or **$HOME/.profile**. 

```bash
$ export PATH=$PATH:/usr/local/go/bin
```

**Note**: changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile. 

```bash
$ echo "export GOPATH=$HOME/go" >> $HOME/.profile
$ echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.profile
$ echo "export PATH=$PATH:$GOPATH/bin" >> $HOME/.profile
```

#### Test your installation

Let's run go version to see if everything is correct.

```bash
$ go version
go version go1.11.4 linux/amd64
```

Check that Go is installed correctly by setting up a workspace and building a simple program, as follows. 

Create your **workspace** directory, $HOME/go. (If you'd like to use a different directory, you will need to set the $GOPATH environment variable.)

Next, make the directory src/hello inside your workspace, and in that directory create a file named hello.go that looks like:

#### Workspace

Workspace is our place of work, where we will organize our directories with our projects. As shown above, until **Go version 1.11** we were forced to do everything under the Workspace. $GOPATH Down Projects.

**Example hello**

```bash
$ export GOPATH=$HOME/go
$ mkdir $HOME/go
$ mkdir $HOME/go/src
$ mkdir $HOME/go/src/hello
$ vim $HOME/go/src/hello/hello.go
```
```bash
$GOPATH/
  |-src
    |-hello
      |-hello.go
```

**Example Project**

```bash
$ export GOPATH=$HOME/go
$ mkdir $HOME/go/src/project1
$ mkdir $HOME/go/src/project1/my-pkg
$ mkdir $HOME/go/src/project1/my-cmd
$ mkdir $HOME/go/src/project1/my-vendor
$ mkdir $HOME/go/src/project1/my-logs
$ mkdir $HOME/go/src/project1/my-models
$ mkdir $HOME/go/src/project1/my-repo
$ mkdir $HOME/go/src/project1/my-handler
```

```bash
$GOPATH/
  |-src
    |-github.com/user/project1/
      	|-cmd (of project1)
        	|-main.go
      	|-vendor
      	|-logs
      	|-models
      	|-repo
      	|-handler
    |-github.com/user/project2/
      ....
      ....
```

The $GOPATH environment variable tells the Go tool where your workspace is located. 

```go
$ go get github.com/user/project1
```
The **go get** command fetches source repositories from the internet and places them in your workspace.
Package paths matter to the Go tool. Using "github.com/..." means the tool knows how to fetch your repository. 

In the scenario above everything would have to stay in our **$GOPATH** so that our projects worked correctly.

#### Outside $GOPATH

Now we can do our projects without being in $GOPATH, we can, for example, do it in any directory.

**Project outside GOPATH**

```bash
$ export GOPATH=$HOME/go
$ mkdir $HOME/2019/project1
$ mkdir $HOME/2019/project1/my-pkg
$ mkdir $HOME/2019/project1/my-cmd
$ mkdir $HOME/2019/project1/my-logs
$ mkdir $HOME/2019/project1/my-models
$ mkdir $HOME/2019/project1/my-repo
$ mkdir $HOME/2019/project1/my-handler
```
```bash
$HOME/
  |-2019
    |-github.com/user/project1/
      |-cmd
        |-main.go
      |-vendor
      |-logs
      |-models
      |-repo
      |-handler
```

We can put our project in any directory now.

```bash
$HOME/
  |-any-directory
    |-github.com/user/project1/
      |-cmd
        |-main.go
      |-vendor
      |-logs
      |-models
      |-repo
      |-handler
```

For the above scenario, we will have to use **go mod** in our project so that all external packages can work correctly, in this way we will be able to manage them correctly and version.
More information can be found here: [Wiki Go Modules](https://github.com/golang/go/wiki/Modules)

```go
$ go mod init github.com/user/project1
```
**Note**: 
When we use go mod in $GOPATH we will have to enable using GO111MODULE=on, so that it can work within the $GOPATH structure.
So our program can compile successfully.

```go
$ GO111MODULE=on go run cmd/main.go
$ GO111MODULE=on go build -o project1 cmd/main.go
```

#### Func Main

```go
package main

import "fmt"

func main() {
	fmt.Printf("hello, Gophers\n")
}

```

Then **build** it with the **go tool**: 

```go
$ cd $HOME/go/src/hello
$ go build
```

The command above will build an executable named hello in the directory alongside your source code. Execute it to see the greeting: 

```go
$ ./hello
hello, Gophers
```

Check also the command **run** it with the go: 

```go
$ go run hello.go
hello, Gophers
```

If you see the "hello, Gophers" message then your Go installation **is working**.

You can run **go install** to install the binary into your workspace's **bin** directory or **go clean -i** to remove it.

```go
$ pwd
$ $HOME/go/src/hello
$ cd $HOME/go/src/hello
$ go install
$ ls -lhs $HOME/go/bin
-rwxrwxr-x 1 user user 2,9M nov  8 03:11 hello
$ go clean -i 
$ ls -lhs $HOME/go/bin
```
### Clean Architecture
---

#### Introduction Clean Architecture

The term came about when the book "Clean Architecture: A Craftsman's Guide to Software Structure and Design" was published in 2017 by Robert C. Martin that aims to present what he calls universal rules that are applied in the current architectures you can improve drastically during the life of any software system.
Are they:

- Independent of Frameworks.
- Testable.
- Independent of UI.
- Independent of Database.
- Independent of any external agency.

We divide our architecture into 4 layers:

- Entities
- Use cases
- Controller
- Framework & Driver

The structure would look like below as clean architecture

```bash
$HOME/
  |-any-directory
    |-github.com/user/project/
      |-cmd
      |-models
      |-repository
      |-service
      |-vendor
      |-Makefile
```
We can read an article that becomes popular on the internet [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

The cool thing about this theme is that you can apply directly to our projects using Golang.

We can organize our directories as per Clean Architecture.
There are some organizational models that are already very common in projects in Golang, which can help us to organize our projects in a healthy way.
In this link you will find a good approach on [Standards Project](https://github.com/golang-standards/project-layout)

#### Directory organization

In Golang we have many ways of organizing our projects, everything will depend on the objective of the project, and so your organization will be according to each need.

We can not use just one model or just one standard for everything. In Golang we make systems of different types that have totally different needs and realities. Golang is a general-purpose language, which means it can be applied to thousands of totally different and diverse business models.

**Some types:**
- Systems to serve web front-end
- Systems to serve web back-end
- Microservices
- Functions as a service
- Creating APIs using rEST in General
- Creating APIs using GraphQL in General
- Creating APIs using SOAP in General
- Creation of protocols in General
- Creating and Using Sockets in General
- Embedded systems
- Systems for automation
- Packages to be consumed (libs)
- Creation of new programming languages,
- Monitoring systems
- Games
- Integration with UI of various types
- Webassembly 
- Systems for terminals

It is a multitude of different ways of organizing our structure in Golagn.
What we must do is always to use good practices and this is independent of the type of project and its purpose.

The most common scenarios that use Clean Architecture are projects aimed at microservices and the web.

This article shows us one of the ways to organize our codes in Golang.[Organizing Go code 2012](https://blog.golang.org/organizing-go-code) and [slide 2014 presentation](https://talks.golang.org/2014/organizeio.slide).
Very cool this video, worth checking:  [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0).

Below I have described some directories that are used very frequently in **Web projects in General**.

```bash
$HOME/
  |-any-directory
    |-github.com/user/project/
      |-config
      |-cmd
        |-main.go
      |-logs
      |-models
      |-repo
      |-handler/controller
      |-middleware
      |-pkg
      |-mocks
      |-model
      |-views/template
      |-loggs
      |-vendor
      |-Makefile
```

**Web Design with Front End**

```bash
$HOME/
  |-any-directory
    |-github.com/user/project2/
      |-config
      |-data
      |-screenshots
      |-script
      |-static
      |-vendor
      |-main.go
      |-Makefile
```

### Go commands
---
#### Go commands introduction

In golang we have an arsenal to help us when it comes to compiling, testing, documenting, managing Profiling etc.

```bash
bug         start a bug report
build       compile packages and dependencies
clean       remove object files and cached files
doc         show documentation for package or symbol
env         print Go environment information
fix         update packages to use new APIs
fmt         gofmt (reformat) package sources
generate    generate Go files by processing source
get         download and install packages and dependencies
install     compile and install packages and dependencies
list        list packages or modules
mod         module maintenance
run         compile and run Go program
test        test packages
tool        run specified go tool
version     print Go version
vet         report likely mistakes in packages
```
 
Use "go help <command>" for more information about a command. 

#### GC and GCCGO
 
The Go language has always been defined by a spec, not an implementation. The Go team has written two different compilers that implement that [spec](https://golang.org/ref/spec): [gc and gccgo](https://golang.org/doc/install/gccgo).

 - Gc is the original compiler, and the go tool uses it by default.
 - Gccgo is a different implementation with a different focus.
 - Compared to gc, gccgo is slower to compile code but supports more powerful optimizations, so a CPU-bound program built by gccgo will usually run faster.
 - The gc compiler supports only the most popular processors: x86 (32-bit and 64-bit) and ARM.
 - Gccgo, however, supports all the processors that GCC supports.
 - Not all those processors have been thoroughly tested for gccgo, but many have, including x86 (32-bit and 64-bit), SPARC, MIPS, PowerPC and even Alpha.
 - Gccgo has also been tested on operating systems that the gc compiler does not support, notably Solaris.

if you install the go command from a standard Go release, it already supports gccgo via the -compiler option: go build -compiler gccgo myprog.go 

#### go bug

Bug opens the default browser and starts a new bug report. The report includes useful system information.
In this link: [Issue Go](https://github.com/golang/go/issues)

```go
$ go bug
```

#### go build

Build compiles the packages named by the import paths, along with their dependencies, but it does not install the results. 

When compiling packages, build ignores files that end in '_test.go'.

The -o flag, only allowed when compiling a single package, forces build to write the resulting executable or object to the named output file, instead of the default behavior described in the last two paragraphs.

The -i flag installs the packages that are dependencies of the target.

```go
$ go build [-o output] [-i] [build flags] [packages]
```

The build flags are shared by the build, clean, get, install, list, run, and test commands:

```bash
- a
	force rebuilding of packages that are already up-to-date.
	
-race
	enable data race detection.
	Supported only on linux/amd64, freebsd/amd64, darwin/amd64 and windows/amd64.
	
-work
	print the name of the temporary work directory and
	do not delete it when exiting.
	
-buildmode mode
	build mode to use. See 'go help buildmode' for more.
	
-compiler name
	name of compiler to use, as in runtime.Compiler (gccgo or gc).
	
-gccgoflags '[pattern=]arg list'
	arguments to pass on each gccgo compiler/linker invocation.

-gcflags '[pattern=]arg list'
	arguments to pass on each go tool compile invocation.

-ldflags '[pattern=]arg list'
	arguments to pass on each go tool link invocation.

-linkshared
	link against shared libraries previously created with
	-buildmode=shared.

-mod mode
	module download mode to use: readonly or vendor.
	See 'go help modules' for more.

-tags 'tag list'
	a space-separated list of build tags to consider satisfied during the
	build. For more information about build tags, see the description of
	build constraints in the documentation for the go/build package.
 
```

**Some examples**

Compiling to run on AWS platform, lambda.

```go
$ GOOS=linux GOARCH=amd64 go build -o lambda lambda.go
```

Compiling for WebAssembly

```go
$ go build GOARCH=wasm GOOS=js go build -o test.wasm hello.go
```

Compiling and generating an .o file, it generates the assembly

```go
$ GOOS=linux GOARCH=amd64 go tool compile -S hello.go 
```

```go
$ go tool compile -S hello.go > hello.S
```

```go
$ go build -gcflags -S hello.go
```

The objdump tool shows references to the line numbers of the original code for what is executing.

```go
$ go tool objdump hello > ref-assembly
```

This is in the code for when we compile it to differentiate and compile only those that contain these tags.

```go
// +build dev
```

```go
$ go build -tags dev -o hello hello.go
``

```go
$ go help buildmode
```

The 'go build' and 'go install' commands take a -buildmode argument which
indicates which kind of object file is to be built. Currently supported values
are:

-buildmode=archive
	Build the listed non-main packages into .a files. Packages named
	main are ignored.

-buildmode=c-archive
	Build the listed main package, plus all packages it imports,
	into a C archive file. The only callable symbols will be those
	functions exported using a cgo //export comment. Requires
	exactly one main package to be listed.

-buildmode=c-shared
	Build the listed main package, plus all packages it imports,
	into a C shared library. The only callable symbols will
	be those functions exported using a cgo //export comment.
	Requires exactly one main package to be listed.

-buildmode=default
	Listed main packages are built into executables and listed
	non-main packages are built into .a files (the default
	behavior).

-buildmode=shared
	Combine all the listed non-main packages into a single shared
	library that will be used when building with the -linkshared
	option. Packages named main are ignored.

-buildmode=exe
	Build the listed main packages and everything they import into
	executables. Packages not named main are ignored.

-buildmode=pie
	Build the listed main packages and everything they import into
	position independent executables (PIE). Packages not named
	main are ignored.

-buildmode=plugin
	Build the listed main packages, plus all packages that they
	import, into a Go plugin. Packages not named main are ignored.

As of Go 1.8, there is a new Go plug-in system. This feature allows programmers to create loosely coupled modular programs using packages compiled as shared object libraries that can be loaded and dynamically linked at run time.

You can check here some comments: [Go Plugin](https://golang.org/pkg/plugin/)

```go
$ go build -buildmode=plugin -o goplugin.go method-plugin.go
```

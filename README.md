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
    - [Types of systems in Go](#types-of-systems-in-go)
    - [Small example API Web](#Small-example-API-Web)
    - [Small example API Web Front-end](#Small-example-API-Web-Front-end)
- [Go commands](#go-commands)
  - [Go commands introduction](#go-commands-introduction)
  - [GC and GCCGO](#gc-and-gccgo)
  - [go run](#go-run) 
  - [go build](#go-build)
    - [Some Examples](#some-examples)
    - [Build modes](#buildmode)
    - [Go Plugin](#go-plugin)
  - [go install](#go-install)
  - [go test](#go-test)
  - [go clean](#go-clean)
  - [go get](#go-get)
  - [go tool](#go-tool)
  - [go doc](#go-doc)
  - [Go mod](#go-mod)
    - [GO111MODULE](#go111module)
    - [go mod init](#gomodinit)
    - [edit go.mod](#edit-go.mod-from-tools-or-scripts)
    - [download modules to local cache](#download-modules-to-local-cache )
    - [go mod verify](#go-mod-verify)
    - [go mod vendor](#go-mod-vendor)
    - [go mod why](#go-mod-why)
- [Environment variables](#environment-variables)
- [Comments and C-style](#comments-and-c-style)
- [Writing on the screen with print](#Writing-on-the-screen-with-print)
- [Introduction Golang](#introduction-golang)
  - [Golang language](#Golang-language)
    - [Keywords](#Keywords)
    - [Operators and punctuation](#Operators-and-punctuation)
    - [Rune literals](#Rune-literals)
    - [String literals](#String-literals)
    - [Constants](#Constants)
    - [Iota](#iota)
    - [Variables](#Variables)
   - [Types](#Types)
     - [Numeric Types](#numeric-type)
     - [String Type](#string-type)
     - [Array Type](#string-type)
     - [Slice Type](#string-type)
     - [Struct Type](#string-type)
     - [Pointer Type](#string-type)
     - [Function Type](#string-type)
     - [Interface Type](#string-type)
     - [Map Type](#string-type)
     - [Channel Type](#string-type)
     - [Boolean Type](#string-type)
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

The term came about when the book __"Clean Architecture: A Craftsman's Guide to Software Structure and Design"__ was published in 2017 by Robert C. Martin that aims to present what he calls universal rules that are applied in the current architectures you can improve drastically during the life of any software system.
Are they:

```bash
- Independent of Frameworks.
- Testable.
- Independent of UI.
- Independent of Database.
- Independent of any external agency.
```

We divide our architecture into 4 layers:

```bash
- Entities
- Use cases
- Controller
- Framework & Driver
```

The structure would look like below as **Clean Architecture**

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

#### Types of systems in Go

```bash
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
```

It is a multitude of different ways of organizing our structure in Golagn.
What we must do is always to use good practices and this is independent of the type of project and its purpose.

The most common scenarios that use __Clean Architecture__ are projects aimed at microservices and the web.

This article shows us one of the ways to organize our codes in Golang.[Organizing Go code 2012](https://blog.golang.org/organizing-go-code) and [slide 2014 presentation](https://talks.golang.org/2014/organizeio.slide).
Very cool this video, worth checking:  [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0).

Below I have described some directories that are used very frequently in **Web projects in General**.

#### Small example API Web

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

#### Small example Clean Architecture

```bash
$HOME/
  |-any-directory
    |-github.com/user/project2/
      |-config
      |-data
      |-pkg
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

```bash
 - Gc is the original compiler, and the go tool uses it by default.
 - Gccgo is a different implementation with a different focus.
 - Compared to gc, gccgo is slower to compile code but supports more powerful optimizations, so a CPU-bound program built by gccgo will usually run faster.
 - The gc compiler supports only the most popular processors: x86 (32-bit and 64-bit) and ARM.
 - Gccgo, however, supports all the processors that GCC supports.
 - Not all those processors have been thoroughly tested for gccgo, but many have, including x86 (32-bit and 64-bit), SPARC, MIPS, PowerPC and even Alpha.
 - Gccgo has also been tested on operating systems that the gc compiler does not support, notably Solaris.
```

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

#### Some Examples

Compiling to run on AWS platform, lambda.
package main

```go
// @jeffotoni
// aws lambda
import (
	"github.com/aws/aws-lambda-go/lambda"
)

type MyResponse struct {
	Body string `json:"body"`
}

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(event MyEvent) (*MyResponse, error) {

	return &MyResponse{Body: event.Name}, nil
}

func main() {

	lambda.Start(HandleRequest)
}
```

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

When compiling just inform the tag you put in your code

```go
$ go build -tags dev -o hello hello.go
```

Passing parameter, passing values to variables in code

```go
package main

import "fmt"

var (
	version string
	date    string
)

func main() {
	fmt.Printf("version=%s, date=%s", version, date)
}
```

```go
$ go build -ldflags "-X main.version=0.0.1 -X main.date=2019-01-18" version.go
```

Com1piling and forcing a rebuilding and for linux platforms, disabling cgo and saying that it is static

```go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o hello hello.go
```

You can omit debugging information by passing the '-w' flag to the linker and you can omit the symbol table by passing '-s'

```go
package main

import "fmt"

func main() {
	fmt.Println("vim-go, hello.")
}
```

Compilation with ldflags

```go
go build -ldflags="-s -w" hello.go
```

output

```bash
$ ls -lh 
-rwxrwxr-x 1 root root 1,3M jan 18 12:42 hello
-rw-rw-r-- 1 root root   75 jan 17 12:04 hello.go
```

Normal compilation

```go
go build -o hello hello.go
```
output

```bash
$ ls -lh 
-rwxrwxr-x 1 root root 1,9M jan 18 12:42 hello
-rw-rw-r-- 1 root root   75 jan 17 12:04 hello.go
```

#### Buildmode

```go
$ go help buildmode
```

The 'go build' and 'go install' commands take a -buildmode argument which
indicates which kind of object file is to be built. Currently supported values
are:

```bash
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
```

#### Go Plugin

As of **Go 1.8**, there is a new Go plug-in system. This feature allows programmers to create loosely coupled modular programs using packages compiled as shared object libraries that can be loaded and dynamically linked at run time.

You can check here some comments: [Go Plugin](https://golang.org/pkg/plugin/)

Check out the example below (lower.go):

```go
// Go in action
// @jeffotoni
// 2019-01-18

package main

import (
	"strings"
)

type lower string

func (t lower) MustLower(name string) string {

	return strings.ToLower(name)
}
```

```go
$ go build -buildmode=plugin -o lower.so lower.go
```

Check out the example below (main.go):

```go
// Go in action
// @jeffotoni
// 2019-01-18

package main

import (
	"fmt"
	"os"
	"plugin"
)

var (
	mod = "./tolower.so"
)

type Tolower interface {
	MustLower(name string) string
}

func main() {
	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	l, err := plug.Lookup("Tolower")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var lower Tolower
	lower, ok := l.(Tolower)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// 4. use the module
	fmt.Println(lower.MustLower("@JEFFOTONI/LAMBDAMAN"))
}
```
```go
$ go run main.go
```
output

```bash
@jeffotoni/lambdaman
```


Let's discuss Exploring shared objects in Go in the next topics, it's an interesting subject to understand how to create libs using Go but using **-buildmode = shared-linksshared.**

It would be something like this:

```go
$ go install -buildmode=shared -linkshared github.com/user/dso/lib
```

Let's create a topic just to deal with objects and shareds in Go.

### Go Install
---

Install packages and dependencies

Usage:

```go
go install [-i] [build flags] [packages]
```
Install compiles and installs the packages named by the import paths.

The -i flag installs the dependencies of the named packages as well.

For more about the build flags, see 'go help build'. For more about specifying packages, see 'go help packages'.


### Go Test
---

Test packages

Usage:

```go
go test [build/test flags] [packages] [build/test flags & test binary flags]
```

Go **test** automates testing the packages named by the import paths. It prints a summary of the test results in the format: 

```bash
ok   archive/tar   0.011s
FAIL archive/zip   0.022s
ok   compress/gzip 0.033s
```
 followed by detailed output for each failed package.

'Go test' recompiles each package along with any files with names matching the file pattern "*_test.go". These additional files can contain test functions, benchmark functions, and example functions. See 'go help testfunc' for more. Each listed package causes the execution of a separate test binary. Files whose names begin with "_" (including "_test.go") or "." are ignored.

Test files that declare a package with the suffix "_test" will be compiled as a separate package, and then linked and run with the main test binary.

The go tool will ignore a directory named "testdata", making it available to hold ancillary data needed by the tests. 

As part of building a test binary, go test runs go vet on the package and its test source files to identify significant problems. If go vet finds any problems, go test reports those and does not run the test binary. Only a high-confidence subset of the default go vet checks are used. That subset is: 'atomic', 'bool', 'buildtags', 'nilfunc', and 'printf'. You can see the documentation for these and other vet tests via "go doc cmd/vet". To disable the running of go vet, use the -vet=off flag. 

 All test output and summary lines are printed to the go command's standard output, even if the test printed them to its own standard error. (The go command's standard error is reserved for printing errors building the tests.)

 Go test runs in two different modes:

The first, called local directory mode, occurs when go test is invoked with no package arguments (for example, 'go test' or 'go test -v'). In this mode, go test compiles the package sources and tests found in the current directory and then runs the resulting test binary. In this mode, caching (discussed below) is disabled. After the package test finishes, go test prints a summary line showing the test status ('ok' or 'FAIL'), package name, and elapsed time.

The second, called package list mode, occurs when go test is invoked with explicit package arguments (for example 'go test math', 'go test ./...', and even 'go test .'). In this mode, go test compiles and tests each of the packages listed on the command line. If a package test passes, go test prints only the final 'ok' summary line. If a package test fails, go test prints the full test output. If invoked with the -bench or -v flag, go test prints the full output even for passing package tests, in order to display the requested benchmark results or verbose logging. 

In addition to the build flags, the flags handled by 'go test' itself are: 

```bash
-args
    Pass the remainder of the command line (everything after -args)
    to the test binary, uninterpreted and unchanged.
    Because this flag consumes the remainder of the command line,
    the package list (if present) must appear before this flag.

-c
    Compile the test binary to pkg.test but do not run it
    (where pkg is the last element of the package's import path).
    The file name can be changed with the -o flag.

-exec xprog
    Run the test binary using xprog. The behavior is the same as
    in 'go run'. See 'go help run' for details.

-i
    Install packages that are dependencies of the test.
    Do not run the test.

-json
    Convert test output to JSON suitable for automated processing.
    See 'go doc test2json' for the encoding details.

-o file
    Compile the test binary to the named file.
    The test still runs (unless -c or -i is specified).
```
For more about build flags, see 'go help build'. For more about specifying packages, see 'go help packages'. 

### Go Clean
---

Lorem Ipsum es simplemente el texto de relleno de las imprentas y archivos de texto. Lorem Ipsum ha sido el texto de relleno estándar de las industrias desde el año 1500, cuando un impresor (N. del T. persona que se dedica a la imprenta) desconocido usó una galería de textos y los mezcló de tal manera que logró hacer un libro de textos especimen. No sólo sobrevivió 500 años, sino que tambien ingresó como texto de relleno en documentos electrónicos, quedando esencialmente igual al original. Fue popularizado en los 60s con la creación de las hojas "Letraset", las cuales contenian pasajes de Lorem Ipsum, y más recientemente con software de autoedición, como por ejemplo Aldus PageMaker, el cual incluye versiones de Lorem Ipsum.

### Go Get
---

The 'go get' command changes behavior depending on whether the go command is running in module-aware mode or legacy **GOPATH** mode. This help text, accessible as 'go help module-get' even in legacy **GOPATH mode**, describes **'go get'** as it operates in module-aware mode.

```bash
Usage: go get [-d] [-m] [-u] [-v] [-insecure] [build flags] [packages]
```

Get downloads the packages named by the import paths, along with their dependencies. It then installs the named packages, like 'go install'.

The -d flag instructs get to stop after downloading the packages; that is, it instructs get not to install the packages.

The -f flag, valid only when -u is set, forces get -u not to verify that each package has been checked out from the source control repository implied by its import path. This can be useful if the source is a local fork of the original.

The -fix flag instructs get to run the fix tool on the downloaded packages before resolving dependencies or building the code.

The -insecure flag permits fetching from repositories and resolving custom domains using insecure schemes such as HTTP. Use with caution.

The -t flag instructs get to also download the packages required to build the tests for the specified packages.

The -u flag instructs get to use the network to update the named packages and their dependencies. By default, get uses the network to check out missing packages but does not use it to look for updates to existing packages.

The -v flag enables verbose progress and debug output. 

When checking out a new package, get creates the target directory GOPATH/src/<import-path>. If the GOPATH contains multiple entries, get uses the first one. For more details see: 'go help gopath'.

When checking out or updating a package, get looks for a branch or tag that matches the locally installed version of Go. The most important rule is that if the local installation is running version "go1", get searches for a branch or tag named "go1". If no such version exists it retrieves the default branch of the package.

When go get checks out or updates a Git repository, it also updates any git submodules referenced by the repository.

Get never checks out or updates code stored in vendor directories. 

### Go tool
---

Usage:

```bash
go tool [-n] command [args...]
```

Tool runs the go tool command identified by the arguments. With no arguments it prints the list of known tools.

The -n flag causes tool to print the command that would be executed but not execute it.

For more about each tool command, see [Go CMD](https://golang.org/cmd/)

### Go doc
---

Usage:

```bash
go doc [-u] [-c] [package|[package.]symbol[.methodOrField]]
```

Doc prints the documentation comments associated with the item identified by its arguments (a package, const, func, type, var, method, or struct field) followed by a one-line summary of each of the first-level items "under" that item (package-level declarations for a package, methods for a type, etc.).

Doc accepts zero, one, or two arguments.

Given no arguments, that is, when run as 

```bash
go doc
```

it prints the package documentation for the package in the current directory. If the package is a command (package main), the exported symbols of the package are elided from the presentation unless the -cmd flag is provided.

When run with one argument, the argument is treated as a Go-syntax-like representation of the item to be documented. What the argument selects depends on what is installed in GOROOT and GOPATH, as well as the form of the argument, which is schematically one of these: 

```bash
go doc <pkg>
go doc <sym>[.<methodOrField>]
go doc [<pkg>.]<sym>[.<methodOrField>]
go doc [<pkg>.][<sym>.]<methodOrField>
```

The first item in this list matched by the argument is the one whose documentation is printed. (See the examples below.) However, if the argument starts with a capital letter it is assumed to identify a symbol or method in the current directory.

For packages, the order of scanning is determined lexically in breadth-first order. That is, the package presented is the one that matches the search and is nearest the root and lexically first at its level of the hierarchy. The GOROOT tree is always scanned in its entirety before GOPATH.

If there is no package specified or matched, the package in the current directory is selected, so "go doc Foo" shows the documentation for symbol Foo in the current package. 

### Go mod
---

A module is a collection of related Go packages. Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules. Modules replace the old GOPATH-based approach to specifying which source files are used in a given build. 

Usage:

```bash
go mod <command> [arguments]
```
A module is defined by a tree of Go source files with a **go.mod** file in the tree's root directory. The directory containing the go.mod file is called the module root. Typically the module root will also correspond to a source code repository root (but in general it need not). The module is the set of all Go packages in the module root and its subdirectories, but excluding subtrees with their own go.mod files.

The "module path" is the import path prefix corresponding to the module root. The go.mod file defines the module path and lists the specific versions of other modules that should be used when resolving imports during a build, by giving their module paths and versions.

For example, this go.mod declares that the directory containing it is the root of the module with path example.com/m, and it also declares that the module depends on specific versions of golang.org/x/text and gopkg.in/yaml.v2: 

```bash
go mod init github.com/user/gomyproject

require (
	golang.org/x/text v0.3.0
	gopkg.in/yaml.v2 v2.1.0
)
```
The go.mod file can also specify replacements and excluded versions that only apply when building the module directly; they are ignored when the module is incorporated into a larger build. For more about the go.mod file, see 'go help go.mod'.

To start a new module, simply create a go.mod file in the root of the module's directory tree, containing only a module statement. The 'go mod init' command can be used to do this: 

```bash
go mod init github.com/user/gomyproject
```
In a project already using an existing dependency management tool like **godep, glide, or dep, 'go mod init'** will also add require statements matching the existing configuration.

Once the go.mod file exists, no additional steps are required: go commands like **'go build'**, **'go test'**, or even **'go list'** will automatically add new dependencies as needed to satisfy imports.

The commands are: 

```bash
download    download modules to local cache
edit        edit go.mod from tools or scripts
graph       print module requirement graph
init        initialize new module in current directory
tidy        add missing and remove unused modules
vendor      make vendored copy of dependencies
verify      verify dependencies have expected content
why         explain why packages or modules are needed
```
Use "go help mod <command>" for more information about a command.

#### GO111MODULE

Go 1.11 includes preliminary support for Go modules, including a new module-aware 'go get' command. We intend to keep revising this support, while preserving compatibility, until it can be declared official (no longer preliminary), and then at a later point we may remove support for work in GOPATH and the old 'go get' command.

The quickest way to take advantage of the new Go 1.11 module support is to check out your repository into a directory outside GOPATH/src, create a go.mod file (described in the next section) there, and run go commands from within that file tree.

For more fine-grained control, the module support in Go 1.11 respects a temporary environment variable, GO111MODULE, which can be set to one of three string values: off, on, or auto (the default). If GO111MODULE=off, then the go command never uses the new module support. Instead it looks in vendor directories and GOPATH to find dependencies; we now refer to this as "GOPATH mode." If GO111MODULE=on, then the go command requires the use of modules, never consulting GOPATH. We refer to this as the command being module-aware or running in "module-aware mode". If GO111MODULE=auto or is unset, then the go command enables or disables module support based on the current directory. Module support is enabled only when the current directory is outside GOPATH/src and itself contains a go.mod file or is below a directory containing a go.mod file.

In module-aware mode, GOPATH no longer defines the meaning of imports during a build, but it still stores downloaded dependencies (in GOPATH/pkg/mod) and installed commands (in GOPATH/bin, unless GOBIN is set).

#### Go Init

Initialize new module in current directory

Usage:

```bash
go mod init [module]
```

Init initializes and writes a new **go.mod** to the current directory, in effect creating a new module rooted at the current directory. The file go.mod must not already exist. If possible, init will guess the module path from import comments (see 'go help importpath') or from version control configuration. To override this guess, supply the module path as an argument. 


```bash
go mod init github.com/user/gomyproject2

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/didip/tollbooth v4.0.0+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	golang.org/x/crypto v0.0.0-20190103213133-ff983b9c42bc
	golang.org/x/time v0.0.0-20181108054448-85acf8d2951c // indirect
)
```

#### Edit go.mod from tools or scripts 

Usage:

```bash
go mod edit [editing flags] [go.mod]
```

#### Download modules to local cache 

Usage:

```bash
go mod download [-json] [modules]
```

Download downloads the named modules, which can be module patterns selecting dependencies of the main module or module queries of the form path@version. With no arguments, download applies to all dependencies of the main module.

The go command will automatically download modules as needed during ordinary execution. The "go mod download" command is useful mainly for pre-filling the local cache or to compute the answers for a Go module proxy. 

By default, download reports errors to standard error but is otherwise silent. The -json flag causes download to print a sequence of JSON objects to standard output, describing each downloaded module (or failure), corresponding to this Go struct: 

```go
type Module struct {
    Path     string // module path
    Version  string // module version
    Error    string // error loading module
    Info     string // absolute path to cached .info file
    GoMod    string // absolute path to cached .mod file
    Zip      string // absolute path to cached .zip file
    Dir      string // absolute path to cached source root directory
    Sum      string // checksum for path, version (as in go.sum)
    GoModSum string // checksum for go.mod (as in go.sum)
}
```
#### Go mod Verify

Verify dependencies have expected content.

Usage:

```bash
go mod verify
```

Verify checks that the dependencies of the current module, which are stored in a local downloaded source cache, have not been modified since being downloaded. If all the modules are unmodified, verify prints "all modules verified." Otherwise it reports which modules have been changed and causes 'go mod' to exit with a non-zero status. 

#### Go mod Vendor

Make vendored copy of dependencies

Usage:

```bash
go mod vendor [-v]
```

Vendor resets the main module's vendor directory to include all packages needed to build and test all the main module's packages. It does not include test code for vendored packages.

The -v flag causes vendor to print the names of vendored modules and packages to standard error.

Explain why packages or modules are needed

#### Go mod why

Usage:

```bash
go mod why [-m] [-vendor] packages...
```

Why shows a shortest path in the import graph from the main module to each of the listed packages. If the -m flag is given, why treats the arguments as a list of modules and finds a path to any package in each of the modules.

By default, why queries the graph of packages matched by "go list all", which includes tests for reachable packages. The -vendor flag causes why to exclude tests of dependencies.

The output is a sequence of stanzas, one for each package or module name on the command line, separated by blank lines. Each stanza begins with a comment line "# package" or "# module" giving the target package or module. Subsequent lines give a path through the import graph, one package per line. If the package or module is not referenced from the main module, the stanza will display a single parenthesized note indicating that fact.

For example:

```bash
$ go mod why golang.org/x/text/language golang.org/x/text/encoding
 golang.org/x/text/language
rsc.io/quote
rsc.io/sampler
golang.org/x/text/language

golang.org/x/text/encoding
(main module does not need package golang.org/x/text/encoding)
```

### Environment variables
---
The go command, and the tools it invokes, examine a few different environment variables. For many of these, you can see the default value of on your system by running 'go env NAME', where NAME is the name of the variable.

General-purpose environment variables: 

```bash
GCCGO
	The gccgo command to run for 'go build -compiler=gccgo'.
GOARCH
	The architecture, or processor, for which to compile code.
	Examples are amd64, 386, arm, ppc64.
GOBIN
	The directory where 'go install' will install a command.
GOCACHE
	The directory where the go command will store cached
	information for reuse in future builds.
GOFLAGS
	A space-separated list of -flag=value settings to apply
	to go commands by default, when the given flag is known by
	the current command. Flags listed on the command-line
	are applied after this list and therefore override it.
GOOS
	The operating system for which to compile code.
	Examples are linux, darwin, windows, netbsd.
GOPATH
	For more details see: 'go help gopath'.
GOPROXY
	URL of Go module proxy. See 'go help goproxy'.
GORACE
	Options for the race detector.
	See https://golang.org/doc/articles/race_detector.html.
GOROOT
	The root of the go tree.
GOTMPDIR
	The directory where the go command will write
	temporary source files, packages, and binaries.
```

Each entry in the GOFLAGS list must be a standalone flag. Because the entries are space-separated, flag values must not contain spaces. In some cases, you can provide multiple flag values instead: for example, to set '-ldflags=-s -w' you can use 'GOFLAGS=-ldflags=-s -ldflags=-w'.

Architecture-specific environment variables:

```bash
GOARM
	For GOARCH=arm, the ARM architecture for which to compile.
	Valid values are 5, 6, 7.
GO386
	For GOARCH=386, the floating point instruction set.
	Valid values are 387, sse2.
GOMIPS
	For GOARCH=mips{,le}, whether to use floating point instructions.
	Valid values are hardfloat (default), softfloat.
GOMIPS64
	For GOARCH=mips64{,le}, whether to use floating point instructions.
	Valid values are hardfloat (default), softfloat.
```

### Comments and C-style

Go provides C-style /* */ block comments and C++-style // line comments. Line comments are the norm; block comments appear mostly as package comments, but are useful within an expression or to disable large swaths of code.

```go
// Go in action
// @jeffotoni
// 2019-01-16

/**
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// define flags
var (
	flagRegion   = flag.String("region", "", "example: us-east-1")
	flagInstance = flag.String("instance", "", "example: i-05bef719c14d68d10")
)
*/
```

### Writing on the screen with print

Let's learn how to send data to screen which is actually **stdout** standard output we will see more ahead with details on **stdout** and **stdin**.

Let's know **print, println and fmt.Println**

Current implementations provide several built-in functions useful during bootstrapping. These functions are documented for completeness but are not guaranteed to stay in the language. They do not return a result. 

Implementation restriction: **print** and **println** need not accept arbitrary argument types, but printing of boolean, numeric, and string types must be supported. 

**println is an built-in function** (into the runtime) which may eventually be removed, while the **fmt package** is in the standard library, which will persist.


```bash
Function   Behavior

print      prints all arguments; formatting of arguments is implementation-specific
println    like print but prints spaces between arguments and a newline at the end
```

using print:
```go
// test print
package main

func main() {
   print("debugging my system with print")
}
```

Output:
```bash
debugging my system with print
```

using println:
```go
// test println
package main

func main() {
   println("debugging my system with println")
}
```

Output:
```bash
debugging my system with println
```


using fmt.Println:
```go
package main

import "fmt"

func main() {
   fmt.Println("debugging my system with fmt.Println")
}
```

Output:
```bash
debugging my system with fmt.Println
```

The goal of starting and running the print, println or fmt.Println command is to help us with the tests we will be performing from now on at every step of our Go learning. 


### Introduction golang
---

Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. 
Programs are constructed from packages, whose properties allow efficient management of dependencies.

The grammar is compact and regular, allowing for easy analysis by automatic tools such as integrated development environments.

#### Golang language
---

##### Keywords

The following keywords are reserved and may not be used as identifiers. 

```bash
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

##### Operators and punctuation

The following character sequences represent operators (including assignment operators) and punctuation: 

```bash
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```

##### Rune literals

A rune literal represents a rune constant, an integer value identifying a Unicode code point. A rune literal is expressed as one or more characters enclosed in single quotes, as in 'x' or '\n'. Within the quotes, any character may appear except newline and unescaped single quote. A single quoted character represents the Unicode value of the character itself, while multi-character sequences beginning with a backslash encode values in various formats.

The simplest form represents the single character within the quotes; since Go source text is Unicode characters encoded in UTF-8, multiple UTF-8-encoded bytes may represent a single integer value. For instance, the literal 'a' holds a single byte representing a literal a, Unicode U+0061, value 0x61, while 'ä' holds two bytes (0xc3 0xa4) representing a literal a-dieresis, U+00E4, value 0xe4.

Several backslash escapes allow arbitrary values to be encoded as ASCII text. There are four ways to represent the integer value as a numeric constant: \x followed by exactly two hexadecimal digits; \u followed by exactly four hexadecimal digits; \U followed by exactly eight hexadecimal digits, and a plain backslash \ followed by exactly three octal digits. In each case the value of the literal is the value represented by the digits in the corresponding base.

Although these representations all result in an integer, they have different valid ranges. Octal escapes must represent a value between 0 and 255 inclusive. Hexadecimal escapes satisfy this condition by construction. The escapes \u and \U represent Unicode code points so within them some values are illegal, in particular those above 0x10FFFF and surrogate halves.

After a backslash, certain single-character escapes represent special values:

```bash
\a   U+0007 alert or bell
\b   U+0008 backspace
\f   U+000C form feed
\n   U+000A line feed or newline
\r   U+000D carriage return
\t   U+0009 horizontal tab
\v   U+000b vertical tab
\\   U+005c backslash
\'   U+0027 single quote  (valid escape only within rune literals)
\"   U+0022 double quote  (valid escape only within string literals)
```

##### String literals

 A string literal represents a string constant obtained from concatenating a sequence of characters. There are two forms: raw string literals and interpreted string literals.

Raw string literals are character sequences between back quotes, as in `foo`. Within the quotes, any character may appear except back quote. The value of a raw string literal is the string composed of the uninterpreted (implicitly UTF-8-encoded) characters between the quotes; in particular, backslashes have no special meaning and the string may contain newlines. Carriage return characters ('\r') inside raw string literals are discarded from the raw string value. 

Interpreted string literals are character sequences between double quotes, as in "bar". Within the quotes, any character may appear except newline and unescaped double quote. The text between the quotes forms the value of the literal, with backslash escapes interpreted as they are in rune literals (except that \' is illegal and \" is legal), with the same restrictions. The three-digit octal (\nnn) and two-digit hexadecimal (\xnn) escapes represent individual bytes of the resulting string; all other escapes represent the (possibly multi-byte) UTF-8 encoding of individual characters. Thus inside a string literal \377 and \xFF represent a single byte of value 0xFF=255, while ÿ, \u00FF, \U000000FF and \xc3\xbf represent the two bytes 0xc3 0xbf of the UTF-8 encoding of character U+00FF.

```bash
`abc`                // same as "abc"
`\n
\n`                  // same as "\\n\n\\n"
"\n"
"\""                 // same as `"`
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // illegal: surrogate half
"\U00110000"         // illegal: invalid Unicode code point
```

These examples all represent the same string: 

```bash
"日本語"                                 // UTF-8 input text
`日本語`                                 // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
```
##### Constants

There are boolean constants, rune constants, integer constants, floating-point constants, complex constants, and string constants. Rune, integer, floating-point, and complex constants are collectively called numeric constants.

A constant value is represented by a rune, integer, floating-point, imaginary, or string literal, an identifier denoting a constant, a constant expression, a conversion with a result that is a constant, or the result value of some built-in functions such as unsafe.Sizeof applied to any value, cap or len applied to some expressions, real and imag applied to a complex constant and complex applied to numeric constants. The boolean truth values are represented by the predeclared constants true and false. The predeclared identifier iota denotes an integer constant.

In general, complex constants are a form of constant expression and are discussed in that section. 

Example: 

```go
package main

const (
	PATH        = "/myhome/app"
	KB     int  = 1 << (10 * iota) // 1Kb
	MB     int  = 1 << (10 * iota) // 1MB
	GB     int  = 1 << (10 * iota) // 1G
	MALE   bool = true
	FEMALE      = true
	DOLLAR      = 3.99
)

// EXPRESSIONS PERMITTED
const a = 2 + 3.0        // a == 5.0   (untyped floating-point constant)
const b = 15 / 4         // b == 3     (untyped integer constant)
const c = 15 / 4.0       // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3 / 2  // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3 / 2. // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 << 3.0       // d == 8     (untyped integer constant)
const e = 1.0 << 3       // e == 8     (untyped integer constant)
//const f = int32(1) << 33 // illegal    (constant 8589934592 overflows int32)
//const g = float64(2) >> 1 // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" > "bar" // h == true  (untyped boolean constant)
const j = true          // j == true  (untyped boolean constant)
const k = 'w' + 1       // k == 'x'   (untyped rune constant)
const l = "hi"          // l == "hi"  (untyped string constant)
const m = string(k)     // m == "x"   (type string)

const Σ = 1 - 0.707i     //            (untyped complex constant)
const Δ = Σ + 2.0e-4     //            (untyped complex constant)
const Φ = iota*1i - 1/1i //            (untyped complex constant)
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays // this constant is not exported
)

const Pi float64 = 3.14159265358979323846
const zero = 0.0 // untyped floating-point constant
const (
	size int64 = 1024
	eof        = -1 // untyped integer constant
)
const xa, xb, xc = 3, 4, "foo" // a = 3, b = 4, c = "foo", untyped integer and string constants
const xu, xv float32 = 0, 3    // u = 0.0, v = 3.0

func main() {

	println("###############")
	println(Sunday)
	println(Monday)
	println(Tuesday)
	println(Wednesday)
	println(Thursday)
	println(Friday)
	println(Partyday)
	println(numberOfDays)
	println(PATH)
	println(KB)
	println(MB)
	println(GB)
	println(MALE)
	println(FEMALE)
	println(DOLLAR)
	println("###############")
	println("")
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(h)
	println(j)
	println(k)
	println(l)
	println(m)
	println(Σ)
	println(Δ)
	println(Φ)
	println("###############")
}
```

Output:
```bash
###############
0
1
2
3
4
5
6
7
/myhome/app
1024
1048576
1073741824
true
true
+3.990000e+000
###############

+5.000000e+000
3
+3.750000e+000
8
8
true
true
120
hi
x
(+1.000000e+000-7.070000e-001i)
(+1.000200e+000-7.070000e-001i)
(+0.000000e+000+1.000000e+000i)
###############
```

##### Iota

Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants. Its value is the index of the respective ConstSpec in that constant declaration, starting at zero. It can be used to construct a set of related constants:

Example:

```bash
const (
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)

const (
	a = 1 << iota  // a == 1  (iota == 0)
	b = 1 << iota  // b == 2  (iota == 1)
	c = 3          // c == 3  (iota == 2, unused)
	d = 1 << iota  // d == 8  (iota == 3)
)

const (
	u         = iota * 42  // u == 0     (untyped integer constant)
	v float64 = iota * 42  // v == 42.0  (float64 constant)
	w         = iota * 42  // w == 84    (untyped integer constant)
)

const x = iota  // x == 0
const y = iota  // y == 0
```
##### Variable

 A variable is a storage location for holding a value. The set of permissible values is determined by the variable's type.
 
  A variable declaration or, for function parameters and results, the signature of a function declaration or function literal reserves storage for a named variable. Calling the built-in function new or taking the address of a composite literal allocates storage for a variable at run time. Such an anonymous variable is referred to via a (possibly implicit) pointer indirection.

Structured variables of array, slice, and struct types have elements and fields that may be addressed individually. Each such element acts like a variable.

The static type (or just type) of a variable is the type given in its declaration, the type provided in the new call or composite literal, or the type of an element of a structured variable. Variables of interface type also have a distinct dynamic type, which is the concrete type of the value assigned to the variable at run time (unless the value is the predeclared identifier nil, which has no type). The dynamic type may vary during execution but values stored in interface variables are always assignable to the static type of the variable. 

Example:

```go
package main

import "fmt"

type T struct{ y int }

var x interface{} // x is nil and has static type interface{}
var v *T          // v has value nil, static type *T

type tv T

var (
	Float32    float32
	Float64    float64
	Boolean    bool
	Int        int
	String     = "@jeffotoni"
	Byte       = []byte("string here")
	Uint8      uint8
	Rune       rune
	Complex128 complex128
)

func main() {

	x = 42 // x has value 42 and dynamic type int
	fmt.Println(x)

	x = v // x has value (*T)(nil) and dynamic type *T
	fmt.Println(x)

	x = T{y: 2}
	fmt.Println(x)

	vx := tv{y: 10}
	fmt.Println(vx)
}
```

Output:
```bash
42
<nil>
{2}
{10}
```


##### Types
---

A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a type name, if it has one, or specified using a type literal, which composes a type from existing types. 

The language predeclares certain type names. Others are introduced with type declarations. Composite types—array, struct, pointer, function, interface, slice, map, and channel types—may be constructed using type literals.

Each type T has an underlying type: If T is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is T itself. Otherwise, T's underlying type is the underlying type of the type to which T refers in its type declaration. 

```bash
type (
    A1 = string
    A2 = A1
)

type (
    B1 string
    B2 B1
    B3 []B1
    B4 B3
)
```

The underlying type of string, A1, A2, B1, and B2 is string. The underlying type of []B1, B3, and B4 is []B1. 

##### Numeric Types

A numeric type represents sets of integer or floating-point values. The predeclared architecture-independent numeric types are: 

```bash
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```
The value of an n-bit integer is n bits wide and represented using two's complement arithmetic.

There is also a set of predeclared numeric types with implementation-specific sizes:

```bash
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

To avoid portability issues all numeric types are defined types and thus distinct except byte, which is an alias for uint8, and rune, which is an alias for int32. Conversions are required when different numeric types are mixed in an expression or assignment. For instance, int32 and int are not the same type even though they may have the same size on a particular architecture. 

##### String types

A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is string; it is a defined type.

The length of a string s (its size in bytes) can be discovered using the built-in function len. The length is a compile-time constant if the string is a constant. A string's bytes can be accessed by integer indices 0 through len(s)-1. It is illegal to take the address of such an element; if s[i] is the i'th byte of a string, &s[i] is invalid. 

```go
package main

import "fmt"

type S string

var (
	String = "@jeffotoni"
)

func main() {
	var text string
	var str S

	mypicture := "@Photograph-jeffotoni"
	str = "@workshop-devOpsBh"
	text = "@jeffotoni-golang"

	fmt.Println(str)
	fmt.Println(String)
	fmt.Println(text)
	fmt.Println(mypicture)

	// example string
	s := "日本語"
	fmt.Printf("Glyph:             %q\n", s)
	fmt.Printf("UTF-8:             [% x]\n", []byte(s))
	fmt.Printf("Unicode codepoint: %U\n", []rune(s))
}
```
Output:

```go
@workshop-devOpsBh
@jeffotoni
@jeffotoni-golang
@Photograph-jeffotoni
Glyph:             "日本語"
UTF-8:             [e6 97 a5 e6 9c ac e8 aa 9e]
Unicode codepoint: [U+65E5 U+672C U+8A9E]
```

#### Array types

An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length and is never negative.

```bash
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
```

The length is part of the array's type; it must evaluate to a non-negative constant representable by a value of type int. The length of array a can be discovered using the built-in function len. The elements can be addressed by integer indices 0 through len(a)-1. Array types are always one-dimensional but may be composed to form multi-dimensional types.

```
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays. 

##### Declaring an Array in Golang

```go
var a[n]T
```

Example:

```go
package main

import "fmt"

func main() {

	// var a []string // wrong

	// An array of 10 integers
	var a1 [10]int
	a1[0] = 10
	fmt.Println(a1)

	// An array of 3 strings
	var a2 [3]string
	a2[0] = "Jeff"
	a2[1] = "Lambda"
	fmt.Println(a2)
	fmt.Println(a2[0], a2[1])

	// An array of 3 strings
	var a3 [5]string
	fmt.Println(a3)
}
```

Output:
```bash
[10 0 0 0 0 0 0 0 0 0]
[Jeff Lambda ]
Jeff Lambda
[    ]
```

##### Initializing an array using an array literal

You can declare and initialize an array at the same time like this -

Example:

```go

package main

import "fmt"

func main() {

	// Declaring and initializing an array at the same time
	var a = [5]int{12, 24, 64, 55, 99}
	fmt.Println(a)

	// Declaring and initializing an array type string
	var a2 = [3]string{"lambda", "serverless", "Go"}
	fmt.Println(a2)

	// Short hand declaration
	a3 := [2]string{"2019", "Golang"}
	fmt.Println(a3)
}
```

Output:
```bash
[12 24 64 55 99]
[lambda serverless Go]
[2019 Golang]
```
##### Golang compiler infer the length of the array

An array has a fixed size. Even using [...] to declare values without determining size, the Golang compiler will increment positions, but you will not be able to make changes like adding or changing non-existent features.

Example:
```go
package main

import "fmt"

func main() {

	// Letting Go compiler infer the length of the array
	a := [...]string{"C", "C++", "B", "Fortran", "Lisp", "Pascal", "Assembly"}
	fmt.Println(a)

	// Letting Go compiler infer the length of the array
	a2 := [...]string{}
	// a2[0] = "" // error invalid array index 0
	fmt.Println(a2)
	// a2 = append(a2, "Cloud") // error must be slice
	//
	a3 := [...]string{"@go_br", "@awsbrasil"}
	a3[0] = "Golang is life!"
	fmt.Println(a3)
}
```


Output:
```bash
[C C++ B Fortran Lisp Pascal Assembly]
[]
[Golang is life! @awsbrasil]
```

If we do this a2: = [...] string {} we will not be able to add elements in this array anymore.

##### Array Dimensional

We can also make dimensional arrays row and column [row] [column], and can also use this way [column].

Example:
```go
package main

import "fmt"

func main() {

	a := [2][2]int{
		{33, 51},
		{72, 92}, // This trailing comma is mandatory
	}

	fmt.Println(a)
	fmt.Println(a[0][0])
	fmt.Println(a[0][1])
	fmt.Println(a[1][0])
	fmt.Println(a[1][1])

	a3 := [...][3]int{
		{10, 11, 20},
		{12, 14, 22},
		{30, 15, 24},
		{32, 16, 26},
		{34, 17, 28},
		{36, 18, 30},
	}
	fmt.Println(a3)
}
```

Output:
```bash
[[33 51] [72 92]]
33
51
72
92
[[10 11 20] [12 14 22] [30 15 24] [32 16 26] [34 17 28] [36 18 30]]
```

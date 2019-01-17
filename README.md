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
- [Clean structure](#clean-structure)
  - [directory organization](#)
    - [src](#)
    - [pkg](#)
    - [cmd](#)
    - [mocks](#)
    - [repo](#)
    - [model](#)
    - [views/template](#)
    - [handler](#)
    - [middleware](#)
    - [config](#)
    - [loggs](#)
- [Go commands](#gocommands)
  - [go run](#gorun) 
  - [go build](#gobuild)
      - [Build modes](#gobuildmodes)
      - [Go and C](#goandc)
  - [go install](#goinstall)
  - [go test](#gotest)
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

Create your **workspace** directory, $HOME/go. (If you'd like to use a different directory, you will need to set the GOPATH environment variable.)

Next, make the directory src/hello inside your workspace, and in that directory create a file named hello.go that looks like:

#### Workspace

Workspace is our place of work, where we will organize our directories with our projects. As shown above, until **Go version 1.11** we were forced to do everything under the Workspace. $GOPATH Down Projects.

**Example hello**

```bash
$ mkdir $HOME/go
$ mkdir $HOME/go/src
$ mkdir $HOME/go/src/hello
$ vim $HOME/go/src/hello/hello.go
```

**Example Project**

```bash
$ mkdir $HOME/go/src/project
$ mkdir $HOME/go/src/project/my-pkg
$ mkdir $HOME/go/src/project/my-cmd
$ mkdir $HOME/go/src/project/my-logs
$ mkdir $HOME/go/src/project/my-models
```
In the scenario above everything would have to stay in our $ GOPATH so that our projects worked correctly.

#### Outside $GOPATH
Now we can make our projects out of $GOPATH, we can for example do so.


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

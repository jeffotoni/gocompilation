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
     - [Linux](#linux)
     - [$GOPATH](#gopath)
     - [Go Workspace](#goworkspace)
   - [Clean structure](#cleanstructure)
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
    - [return multiple values](#returnmulti) 
    - [Variadic Functions](#variadicfunc) 
    - [functions as a parameter](#funcparameter) 
    - [Closures](#closures)
    - [Recursion](#recursion)
    - [Asynchronous Functions](#asynchromous)
  - [Array](#array)
    - [Declaring Array](#)
    - [Initializing Array](#)
    - [Accessing Array Elements](#)
    - [Array can not be resized](#)
  - [Slice](#slice)
    - [Slices - make](#)
    - [Slices - append](#)
    - [Slices - range](#)
    - [Slices of Slices](#)
    - [Nil Slices](#)
    - [Slice len () and cap ()](#)
  - [Struct](#struct)
    - [Allocation New](#allocationnew)
    - [Initializing a struct](#structinit)
    - [Pointer to a struct](#structpointer)
    - [Type Struct Json](#structjson)
    - [Anonymus Struct](#structanonymus)
    - [Struct Reflect](#structreflect)
  - [Erros](#erros)
     - [How Error Control Works](#)
     - [Log.fatal](#)
     - [Log.panic](#)
     - [panic](#panic)
     - [Custom Errors](#customerrors)
     - [Interface Errors](#)
     - [Errors.New](#)
     - [fmt.Errorf](#)
  - [My first packages](#package)
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
  - [Pointer](#pointer)
    - [How Pointers in Go Works](#howpointers)
    - [When using pointers](#whenusingpointers)
  - [Methods](#methods)
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
   - [Maps]
     - [Make](#)
     - [Type Int, string, interfaces](#)
     - [Struct](#)
     - [Mutant](#)
     - [Not Sync](#)
   - [Json](#Json)
     - [Json marshal](#jsonmarshal)
     - [Json Unmarshal](#jsonunmarshal)
     - [json Encode](#jsonencode)
     - [Json Decode](#jsondecode)
   - [Parse Json](#Json)
     - [Toml](#jsontoml)
     - [Yaml](#jsonyaml)
     - [Gcfg](#jsongcfg)
   - [net/http Server](#)
     - [http.NewServeMux](#)
     - [http.HandlerFunc](#)
     - [http.Handle](#)
     - [http.Handler](#)
     - [http.Server](#)
     - [http.FileServer](#)
     - [next.ServeHTTP](#)
     - [http.NotFound](#)
     - [Disable http.FileServer](#)
        ├── main.go
        └── static
          ├── css
          │   ├── index.html
          │   └── main.css
          ├── img
          │   ├── index.html
          │   └── logo.png
          ├── index.html
          └── robots.txt
     - [http.Dir](#)
     - [http.StripPrefix](#)
     - [ListenAndServe](#)
     - [Middleware](#)
   - [net/http Client](#)
     - [http.Transport](#)
     - [http.Client](#)
     - [http.Get](#)
     - [http,Post](#)
     - [http.NewRequest](#)
   

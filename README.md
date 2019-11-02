# SweetHomeBot core package

This package is needed for develop a SweetHomeBot GoLang plugin. 

These plugins have certain peculiarities. 
To ensure compatibility of older installations and plugins, this package use a separate versioning. 
While importing this package use following snippet:

    import core "github.com/sweethomebot/go/core/v1"
    
Also is it important, that this package is located in `/sweethomebot/go/` on your compile device. (Add this path to **$GOPATH**)

The main SweetHomBot software and all plugin I provide are compiled with Go version 1.12.4 on a Raspberry Pi 2.

## subpackages

### core
Implement the communication framework, which allow plugins to interact with another. 
Very common instructions has a shortcut, listed in **FunctionShortcuts.go**

### hhcType 
Basic types for **Hardware Host Communication**

### ui
Structure of the UI elements

## pkg
This folder contains the compiled version of this repository, which is also used in the main program.
(outdated)

## Install
	go get github.com/sweethomebot/go/...

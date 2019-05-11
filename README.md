# SweetHomeBot core package

This package is needed for develop a SweetHomeBot plugin.

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
go 1.12.4 @ raspberry pi 3+  

## Install
	go get github.com/sweethomebot/go/...

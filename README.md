# GoLang
This is the repository of my journey with GoLang
Day - 01
Today I got to know that Golang is not a complete Object oriented Progrmming language.
Go language won't support classes and objects but supports interfaces and structs.
This is a statically typed programming langusge.
this is similar to the programming language C++
It have automatic garbage collection.
This is a compiled language.
Go lang is faster than c++
This is most popular language in companies like google, netflix and other big tech giants.
The popular applications that use Golang are docker, kubernetes and Terraform are the some of the popular ones.

Syntac of Golang
package <package name>
import <packages>

func main (){
statements
}

Example:
package main
import "fmt"

func main() {
  Print("Hello world")
}

How to code
first we need to download Go
https://golang.org/dl/
After installing it we can check by using the command in the terminal
go version 
we should see the version else go is not installed properly.

After this we need to have a code editor and VS code is the popular one.
Install and setup vscode from here
https://code.visualstudio.com/

After installing go to the extensions and install GO at Google team

Then open the command palette and run Go: Install/Update Tools

Now run the command below
go mod init example.com/hello

Now our vscode is set now start coding by creating a file with the extension .go and do code.

Comment lines:
there are 2 types
1) Single line
   The single line comments are start with //  
2) Multiline
   Multiline comments are start with /* and end with */

   These comment lines are very useful for making code more readable and also we can skip a set of code so that it will be easier for us to debug.

Variables:
Variables are named memory locations to store the memory.
In Go again we can declare variables in 2 different ways.
1) By using var keyword.
   syntax: var <variable_name> <type> = <value>
   Example: var name string="lavakumar"
2) By using the :=
   syntax: variable_name := value
   when we do like this that the compiker will infer the type of the variable by the value.

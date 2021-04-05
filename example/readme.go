package main

import (
  "github.com/n-arms/cbreakable"
)

func main(){
  /*
cbrk.Cbreak(true)
defer cbrk.Cbreak(false)
cbrk.Stdout.Print("Allow this program to continue? (y/n)")
char := cbrk.Stdin.Get()
cbrk.Stdout.Println(char)
if char.Equals("n") {
  return
}
//stdout.Println("guess what, it continued!")
cbrk.Stdout.Println("guess what, it continued!")
*/
cbrk.Cbreak(true)
defer cbrk.Cbreak(false)
cbrk.Stdout.Println("What is your password?")
//if cbrk.CharString(cbrk.Stdin.Getln()) == "incorrect" {
if cbrk.CharString(cbrk.Stdin.Getln()) == "incorrect\n" {
  cbrk.Stdout.Println("thats right!")
}else{
  cbrk.Stdout.Println("thats wrong :(")
}
}

package main

import (
  "github.com/n-arms/cbreakable"
)

func render_menu(options []string, cursor int) {
    out := cbrk.Stdout;
    for i := 0; i < len(options); i++ {
        if i == cursor {
            out.Print(cbrk.Red);
            out.Print(options[i]);
            out.Println(cbrk.Reset);
        } else {
            out.Println(options[i]);
        }
    }
}

func main(){
    cbrk.Cbreak(true)
    defer cbrk.Cbreak(false)

    options := []string{"this one", "or this one", "what about the other one", "don't pick this one"};
    cursor := 0;

    for {
        render_menu(options, cursor);
        char := cbrk.Stdin.Get();
        if char.Equals(cbrk.Up) || char.Equals(cbrk.Left) {
            if cursor == 0 {
                cursor = len(options) - 1;
            } else {
                cursor -= 1;
            }
        } else if char.Equals(cbrk.Down) || char.Equals(cbrk.Right) {
            cursor = (cursor + 1) % len(options);
        }
        for i := 0; i < len(options); i++ {
            cbrk.Stdout.Clearln();
        }
    }
}

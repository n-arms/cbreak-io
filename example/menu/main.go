package main

import (
  "github.com/n-arms/cbreakable"
  "strings"
  "io/ioutil"
  "log"
  "os"
)

func render_menu(options []string, cursor int, search string) int {
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
    out.Print("> ");
    out.Print(cbrk.Cyan);
    out.Print(search);
    out.Println(cbrk.Reset);
    return len(options) + 1;
}

func clear(lines int) {
    for i := 0; i < lines; i++ {
        cbrk.Stdout.Clearln();
    }
}

func run_search(list []string, key string) []string {
    new_list := []string{};
    for _, elem := range(list) {
        if strings.Contains(elem, key) {
            new_list = append(new_list, elem);
        }
    }
    return new_list;
}

func files(path string) []string {
    files, err := ioutil.ReadDir(path);
    if err != nil {
        log.Fatal(err);
    }
    names := make([]string, len(files));
    for i, file := range(files) {
        names[i] = file.Name();
    }
    return names;
}

func clamp(val int, min int, max int) int {
    if val > max {
        return max;
    } else if val < min {
        return min;
    } else {
        return val;
    }
}

const (
    del = "127"
    enter = "10"
)

func main(){
    cbrk.Cbreak(true)
    defer cbrk.Cbreak(false)

    path := os.Args[1];
    elems := files(path);
    current := []string{};
    cursor := 0;
    key := "";
    lines := 0;

    for {
        current = run_search(elems, key);
        cursor = clamp(cursor, 0, len(current) - 1);
        lines = render_menu(current, cursor, key);
        char := cbrk.Stdin.Get();

        if char.Equals(cbrk.Up) || char.Equals(cbrk.Left) {
            cursor -= 1;
        } else if char.Equals(cbrk.Down) || char.Equals(cbrk.Right) {
            cursor += 1;
        } else if char.Repr() == del {
            key = key[:len(key) - 1];
        } else if char.Repr() == enter {
            break;
        } else {
            key += char.String();
        }

        clear(lines);
    }

    clear(lines);
    cbrk.Stdout.Println(current[cursor]);
}

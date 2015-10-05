package lmod

import (
    "testing"
    "fmt"
)

func TestLoadedModules(t *testing.T) {
    modlist := LoadedModules()

    for _,element := range modlist {
        a := fmt.Sprint(element, "\n")
        fmt.Printf(a)
    }
}

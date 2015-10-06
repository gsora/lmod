package lmod

import (
    "testing"
    "fmt"
)

func TestLoadedModules(t *testing.T) {
    modlist, err := LoadedModules()

    if err != nil {
        t.Errorf(err.Error())
    }

    for _,element := range modlist {
        a := fmt.Sprint(element, "\n")
        fmt.Printf(a)
    }
}

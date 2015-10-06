package lmod

import (
    "testing"
    "fmt"
)

func TestIsModuleLoaded(t *testing.T) {
    // the intel module is always loaded
    testModule := "i915"

    loaded, err := IsModuleLoaded(testModule)

    if err != nil {
        t.Errorf(err.Error())
    }

    if loaded {
        fmt.Printf("i195 is loaded\n")
    } else {
        fmt.Printf("i915 is not loaded\n")
    }
}

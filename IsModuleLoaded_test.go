package lmod

import (
    "testing"
    "fmt"
)

func TestIsModuleLoaded(t *testing.T) {
    // the intel module is always loaded
    testModule := "i915"

    if IsModuleLoaded(testModule) {
        fmt.Printf("i195 is loaded\n")
    } else {
        fmt.Printf("i915 is not loaded\n")
    }
}

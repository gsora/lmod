package lmod

import (
    "testing"
    "fmt"
)

func TestModuleOptions(t *testing.T) {

    // needs root powers

    testModule := "i915"

    modOptions, err := ModuleOptions(testModule)

    if err != nil {
        t.Errorf(err.Error())
    }

    fmt.Println(modOptions)
}

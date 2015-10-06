package lmod

import (
    "testing"
    "fmt"
)

func TestModuleHasParameter(t *testing.T) {
    // needs root powers
    testModule := "i915"
    testParam := "modeset"

    loaded, paramValue, err := ModuleHasParameter(testModule, testParam)

    if err != nil {
        t.Errorf(err.Error())
    }

    fmt.Println("loaded: ", loaded, " parameter value: ", paramValue)
}

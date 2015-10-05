package lmod

import (
    "io/ioutil"
    "os"
)

func LoadedModules() ([]string) {

    var modules []string

    flist, err := ioutil.ReadDir("/sys/module")

    if err != nil {
        os.Exit(1)
    }

    for _,element := range flist {
        modules = append(modules, element.Name())
    }

    return modules
}

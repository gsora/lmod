package lmod

import (
    "io/ioutil"
    "fmt"
)

func LoadedModules() ([]string, error) {

    var modules []string

    flist, err := ioutil.ReadDir("/sys/module")

    if err != nil {
        return modules, fmt.Errorf("lmod: cannot access loaded module list") 
    }

    for _,element := range flist {
        modules = append(modules, element.Name())
    }

    return modules, nil
}

package lmod

import (
    "fmt"
    "io/ioutil"
)

func ModuleOptions(module string) (map[string]string, error) {
    basePath := "/sys/module/"
    modOptions := make(map[string]string)

    loaded, err := IsModuleLoaded(module)

    if err != nil {
        return modOptions, err
    }

    if loaded == true {
        modPath := fmt.Sprint(basePath + module + "/parameters/")
        plist, err := ioutil.ReadDir(modPath)

        if err != nil {
            return modOptions, fmt.Errorf("lmod: cannot open module parameters path")
        }

        for _, element := range plist {
            modParamContent, err := ioutil.ReadFile(fmt.Sprint(modPath + element.Name()))

            if err != nil {
                return modOptions, fmt.Errorf("lmod: cannot open '%s'module parameter for reading", module)
            }

            modOptions[element.Name()] = string(modParamContent[:])
        }

        return modOptions, nil
    } else {
        return modOptions, fmt.Errorf("lmod: module %s not found", module)
    }
}

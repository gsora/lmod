package lmod

import (
    "io/ioutil"
    "fmt"
)

func ModuleHasParameter(module string, parameter string) (bool, string, error) {
    loaded, err := IsModuleLoaded(module)

    if loaded == true {
        modOptions, err := ModuleOptions(module)

        if err != nil {
            return false, "", err
        }

        _, prs := modOptions[parameter]

        if prs != false {
            paramPath := "/sys/module/" + module + "/parameters/" + parameter
            modParamContent, err := ioutil.ReadFile(paramPath)

            if err != nil {
                return false, "", err
            }

            return true, string(modParamContent[:]), nil
        } else {
            return false, "", fmt.Errorf("lmod: parameter '%s' for module '%s' not found", parameter, module)
        }
    } else {
        return false, "", err
    }
}

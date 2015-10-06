package lmod

func IsModuleLoaded(moduleName string) (bool, error) {
    mods, err := LoadedModules()

    if err != nil {
        return false, err
    }

    for _,element := range mods {
        if element == moduleName {
            return true, nil
        }
    }
    return false, nil
}

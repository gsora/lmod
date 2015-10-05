package lmod

func IsModuleLoaded(moduleName string) (bool) {
    for _,element := range LoadedModules() {
        if element == moduleName {
            return true
        }
    }
    
    return false
}

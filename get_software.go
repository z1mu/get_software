package main

import (
    "fmt"
    "golang.org/x/sys/windows/registry"
)

func main() {
    key, exists, _ := registry.CreateKey(registry.LOCAL_MACHINE, `SOFTWARE\Wow6432Node\Microsoft\Windows\CurrentVersion\Uninstall`, registry.ALL_ACCESS)
    defer key.Close()
    if exists {
        println(`[+] x64 regedit path is vaild`)
        keys, _ := key.ReadSubKeyNames(0)
	    for _, key_subkey := range keys {
            p := `SOFTWARE\Wow6432Node\Microsoft\Windows\CurrentVersion\Uninstall\` + key_subkey
	        k, e, _ := registry.CreateKey(registry.LOCAL_MACHINE, p, registry.ALL_ACCESS)
	        if e {
	        	displayName, _, _  := k.GetStringValue(`DisplayName`)
                if len(displayName) != 0 {
                    displayVersion, _, _  := k.GetStringValue(`DisplayVersion`)
                    fmt.Printf("[+] displayName: %s , CurrentVersion : %s \n",displayName, displayVersion)
                }
	        }
    		
	    }
    } else {
        println(`[-] x64 regedit path is invaild`)
    }

    key, exists, _ = registry.CreateKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`, registry.ALL_ACCESS)
    defer key.Close()
    if exists {
        println(`[+] x86 regedit path is vaild`)
        keys, _ := key.ReadSubKeyNames(0)
        for _, key_subkey := range keys {
            p := `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\` + key_subkey
            k, e, _ := registry.CreateKey(registry.LOCAL_MACHINE, p, registry.ALL_ACCESS)
            if e {
                displayName, _, _  := k.GetStringValue(`DisplayName`)
                if len(displayName) != 0 {
                    displayVersion, _, _  := k.GetStringValue(`DisplayVersion`)
                    fmt.Printf("[+] displayName: %s , CurrentVersion : %s \n",displayName, displayVersion)
                }
            }
            
        }
    } else {
        println(`[-] x86 regedit path is invaild`)
    }
}
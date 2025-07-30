package main

import (
    "fmt"
    "os"

    "github.com/hashicorp/packer-plugin-sdk/plugin"
    "github.com/hashicorp/packer-plugin-sdk/version"
)

func main() {
    pps := plugin.NewSet()
    pps.RegisterPostProcessor(plugin.DEFAULT_NAME, new(PostProcessor))
    pps.SetVersion(version.InitializePluginVersion("2.1.0", ""))
    err := pps.Run()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}

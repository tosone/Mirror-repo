package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/viper"
)

// Version version
var Version = "no provided"

// BuildStamp BuildStamp
var BuildStamp = "no provided"

// GitHash GitHash
var GitHash = "no provided"

// Setting ..
func Setting(version, buildStamp, gitHash string) {
	Version = version
	BuildStamp = buildStamp
	GitHash = gitHash
}

// Initialize ..
func Initialize() {
	fmt.Printf("%s %s %s/%s\n", viper.GetString("Setting.Name"), Version, runtime.GOOS, runtime.GOARCH)
	fmt.Printf("BuildDate: %s\n", BuildStamp)
	fmt.Printf("BuildHash: %s\n", GitHash)
}

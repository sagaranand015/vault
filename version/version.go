package version

import (
	"bytes"
	"fmt"
)

var (
	// The git commit that was compiled. This will be filled in by the compiler.
	GitCommit   string
	GitDescribe string

	Version           = "unknown"
	VersionPrerelease = "unknown"
)

// VersionInfo
type VersionInfo struct {
	Revision          string
	Version           string
	VersionPrerelease string
}

func GetVersion() *VersionInfo {
	ver := Version
	rel := VersionPrerelease
	if GitDescribe != "" {
		ver = GitDescribe
	}
	if GitDescribe == "" && rel == "" && VersionPrerelease != "" {
		rel = "dev"
	}

	return &VersionInfo{
		Revision:          GitCommit,
		Version:           ver,
		VersionPrerelease: rel,
	}
}

func (c *VersionInfo) String() string {
	var versionString bytes.Buffer

	if Version == "unknown" && VersionPrerelease == "unknown" {
		fmt.Fprintf(&versionString, "Vault (version unknown)")
	}
	fmt.Fprintf(&versionString, "Vault v%s", c.Version)
	if c.VersionPrerelease != "" {
		fmt.Fprintf(&versionString, "-%s", c.VersionPrerelease)

		if c.Revision != "" {
			fmt.Fprintf(&versionString, " (%s)", c.Revision)
		}
	}

	return versionString.String()
}

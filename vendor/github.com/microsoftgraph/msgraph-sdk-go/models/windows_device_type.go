package models
import (
    "errors"
)
// Contains properties for Windows device type.
type WindowsDeviceType int

const (
    // No flags set.
    NONE_WINDOWSDEVICETYPE WindowsDeviceType = iota
    // Whether or not the Desktop Windows device type is supported.
    DESKTOP_WINDOWSDEVICETYPE
    // Whether or not the Mobile Windows device type is supported.
    MOBILE_WINDOWSDEVICETYPE
    // Whether or not the Holographic Windows device type is supported.
    HOLOGRAPHIC_WINDOWSDEVICETYPE
    // Whether or not the Team Windows device type is supported.
    TEAM_WINDOWSDEVICETYPE
)

func (i WindowsDeviceType) String() string {
    return []string{"none", "desktop", "mobile", "holographic", "team"}[i]
}
func ParseWindowsDeviceType(v string) (any, error) {
    result := NONE_WINDOWSDEVICETYPE
    switch v {
        case "none":
            result = NONE_WINDOWSDEVICETYPE
        case "desktop":
            result = DESKTOP_WINDOWSDEVICETYPE
        case "mobile":
            result = MOBILE_WINDOWSDEVICETYPE
        case "holographic":
            result = HOLOGRAPHIC_WINDOWSDEVICETYPE
        case "team":
            result = TEAM_WINDOWSDEVICETYPE
        default:
            return 0, errors.New("Unknown WindowsDeviceType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsDeviceType(values []WindowsDeviceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

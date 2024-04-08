package models
import (
    "errors"
)
// Contains properties for Windows architecture.
type WindowsArchitecture int

const (
    // No flags set.
    NONE_WINDOWSARCHITECTURE WindowsArchitecture = iota
    // Whether or not the X86 Windows architecture type is supported.
    X86_WINDOWSARCHITECTURE
    // Whether or not the X64 Windows architecture type is supported.
    X64_WINDOWSARCHITECTURE
    // Whether or not the Arm Windows architecture type is supported.
    ARM_WINDOWSARCHITECTURE
    // Whether or not the Neutral Windows architecture type is supported.
    NEUTRAL_WINDOWSARCHITECTURE
)

func (i WindowsArchitecture) String() string {
    return []string{"none", "x86", "x64", "arm", "neutral"}[i]
}
func ParseWindowsArchitecture(v string) (any, error) {
    result := NONE_WINDOWSARCHITECTURE
    switch v {
        case "none":
            result = NONE_WINDOWSARCHITECTURE
        case "x86":
            result = X86_WINDOWSARCHITECTURE
        case "x64":
            result = X64_WINDOWSARCHITECTURE
        case "arm":
            result = ARM_WINDOWSARCHITECTURE
        case "neutral":
            result = NEUTRAL_WINDOWSARCHITECTURE
        default:
            return 0, errors.New("Unknown WindowsArchitecture value: " + v)
    }
    return &result, nil
}
func SerializeWindowsArchitecture(values []WindowsArchitecture) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

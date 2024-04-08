package models
import (
    "errors"
)
// Generic visibility state.
type VisibilitySetting int

const (
    // Not configured.
    NOTCONFIGURED_VISIBILITYSETTING VisibilitySetting = iota
    // Hide.
    HIDE_VISIBILITYSETTING
    // Show.
    SHOW_VISIBILITYSETTING
)

func (i VisibilitySetting) String() string {
    return []string{"notConfigured", "hide", "show"}[i]
}
func ParseVisibilitySetting(v string) (any, error) {
    result := NOTCONFIGURED_VISIBILITYSETTING
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_VISIBILITYSETTING
        case "hide":
            result = HIDE_VISIBILITYSETTING
        case "show":
            result = SHOW_VISIBILITYSETTING
        default:
            return 0, errors.New("Unknown VisibilitySetting value: " + v)
    }
    return &result, nil
}
func SerializeVisibilitySetting(values []VisibilitySetting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

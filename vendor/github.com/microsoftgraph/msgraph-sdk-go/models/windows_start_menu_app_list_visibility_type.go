package models
import (
    "errors"
)
// Type of start menu app list visibility.
type WindowsStartMenuAppListVisibilityType int

const (
    // User defined. Default value.
    USERDEFINED_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE WindowsStartMenuAppListVisibilityType = iota
    // Collapse the app list on the start menu.
    COLLAPSE_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
    // Removes the app list entirely from the start menu.
    REMOVE_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
    // Disables the corresponding toggle (Collapse or Remove) in the Settings app.
    DISABLESETTINGSAPP_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
)

func (i WindowsStartMenuAppListVisibilityType) String() string {
    return []string{"userDefined", "collapse", "remove", "disableSettingsApp"}[i]
}
func ParseWindowsStartMenuAppListVisibilityType(v string) (any, error) {
    result := USERDEFINED_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
    switch v {
        case "userDefined":
            result = USERDEFINED_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
        case "collapse":
            result = COLLAPSE_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
        case "remove":
            result = REMOVE_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
        case "disableSettingsApp":
            result = DISABLESETTINGSAPP_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
        default:
            return 0, errors.New("Unknown WindowsStartMenuAppListVisibilityType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsStartMenuAppListVisibilityType(values []WindowsStartMenuAppListVisibilityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

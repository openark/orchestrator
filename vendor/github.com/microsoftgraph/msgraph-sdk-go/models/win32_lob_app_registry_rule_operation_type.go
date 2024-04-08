package models
import (
    "errors"
)
// Contains all supported registry data detection type.
type Win32LobAppRegistryRuleOperationType int

const (
    // Not configured.
    NOTCONFIGURED_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE Win32LobAppRegistryRuleOperationType = iota
    // The specified registry key or value exists.
    EXISTS_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
    // The specified registry key or value does not exist.
    DOESNOTEXIST_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
    // String value type.
    STRING_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
    // Integer value type.
    INTEGER_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
    // Version value type.
    VERSION_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
)

func (i Win32LobAppRegistryRuleOperationType) String() string {
    return []string{"notConfigured", "exists", "doesNotExist", "string", "integer", "version"}[i]
}
func ParseWin32LobAppRegistryRuleOperationType(v string) (any, error) {
    result := NOTCONFIGURED_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
        case "exists":
            result = EXISTS_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
        case "doesNotExist":
            result = DOESNOTEXIST_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
        case "string":
            result = STRING_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
        case "integer":
            result = INTEGER_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
        case "version":
            result = VERSION_WIN32LOBAPPREGISTRYRULEOPERATIONTYPE
        default:
            return 0, errors.New("Unknown Win32LobAppRegistryRuleOperationType value: " + v)
    }
    return &result, nil
}
func SerializeWin32LobAppRegistryRuleOperationType(values []Win32LobAppRegistryRuleOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

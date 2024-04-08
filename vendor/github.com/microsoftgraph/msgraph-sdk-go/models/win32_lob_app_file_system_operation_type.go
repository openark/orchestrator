package models
import (
    "errors"
)
// Contains all supported file system detection type.
type Win32LobAppFileSystemOperationType int

const (
    // Not configured.
    NOTCONFIGURED_WIN32LOBAPPFILESYSTEMOPERATIONTYPE Win32LobAppFileSystemOperationType = iota
    // Whether the specified file or folder exists.
    EXISTS_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
    // Last modified date.
    MODIFIEDDATE_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
    // Created date.
    CREATEDDATE_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
    // Version value type.
    VERSION_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
    // Size detection type.
    SIZEINMB_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
)

func (i Win32LobAppFileSystemOperationType) String() string {
    return []string{"notConfigured", "exists", "modifiedDate", "createdDate", "version", "sizeInMB"}[i]
}
func ParseWin32LobAppFileSystemOperationType(v string) (any, error) {
    result := NOTCONFIGURED_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
        case "exists":
            result = EXISTS_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
        case "modifiedDate":
            result = MODIFIEDDATE_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
        case "createdDate":
            result = CREATEDDATE_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
        case "version":
            result = VERSION_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
        case "sizeInMB":
            result = SIZEINMB_WIN32LOBAPPFILESYSTEMOPERATIONTYPE
        default:
            return 0, errors.New("Unknown Win32LobAppFileSystemOperationType value: " + v)
    }
    return &result, nil
}
func SerializeWin32LobAppFileSystemOperationType(values []Win32LobAppFileSystemOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

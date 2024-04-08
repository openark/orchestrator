package models
import (
    "errors"
)
// Storage locations where managed apps can potentially store their data
type ManagedAppDataStorageLocation int

const (
    // OneDrive for business
    ONEDRIVEFORBUSINESS_MANAGEDAPPDATASTORAGELOCATION ManagedAppDataStorageLocation = iota
    // SharePoint
    SHAREPOINT_MANAGEDAPPDATASTORAGELOCATION
    // Box
    BOX_MANAGEDAPPDATASTORAGELOCATION
    // Local storage on the device
    LOCALSTORAGE_MANAGEDAPPDATASTORAGELOCATION
)

func (i ManagedAppDataStorageLocation) String() string {
    return []string{"oneDriveForBusiness", "sharePoint", "box", "localStorage"}[i]
}
func ParseManagedAppDataStorageLocation(v string) (any, error) {
    result := ONEDRIVEFORBUSINESS_MANAGEDAPPDATASTORAGELOCATION
    switch v {
        case "oneDriveForBusiness":
            result = ONEDRIVEFORBUSINESS_MANAGEDAPPDATASTORAGELOCATION
        case "sharePoint":
            result = SHAREPOINT_MANAGEDAPPDATASTORAGELOCATION
        case "box":
            result = BOX_MANAGEDAPPDATASTORAGELOCATION
        case "localStorage":
            result = LOCALSTORAGE_MANAGEDAPPDATASTORAGELOCATION
        default:
            return 0, errors.New("Unknown ManagedAppDataStorageLocation value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppDataStorageLocation(values []ManagedAppDataStorageLocation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

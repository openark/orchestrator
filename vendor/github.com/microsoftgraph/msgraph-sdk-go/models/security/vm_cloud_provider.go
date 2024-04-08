package security
import (
    "errors"
)
// 
type VmCloudProvider int

const (
    UNKNOWN_VMCLOUDPROVIDER VmCloudProvider = iota
    AZURE_VMCLOUDPROVIDER
    UNKNOWNFUTUREVALUE_VMCLOUDPROVIDER
)

func (i VmCloudProvider) String() string {
    return []string{"unknown", "azure", "unknownFutureValue"}[i]
}
func ParseVmCloudProvider(v string) (any, error) {
    result := UNKNOWN_VMCLOUDPROVIDER
    switch v {
        case "unknown":
            result = UNKNOWN_VMCLOUDPROVIDER
        case "azure":
            result = AZURE_VMCLOUDPROVIDER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_VMCLOUDPROVIDER
        default:
            return 0, errors.New("Unknown VmCloudProvider value: " + v)
    }
    return &result, nil
}
func SerializeVmCloudProvider(values []VmCloudProvider) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

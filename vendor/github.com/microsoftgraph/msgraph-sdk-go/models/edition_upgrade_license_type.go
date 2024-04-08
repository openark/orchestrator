package models
import (
    "errors"
)
// Edition Upgrade License type
type EditionUpgradeLicenseType int

const (
    // Product Key Type
    PRODUCTKEY_EDITIONUPGRADELICENSETYPE EditionUpgradeLicenseType = iota
    // License File Type
    LICENSEFILE_EDITIONUPGRADELICENSETYPE
)

func (i EditionUpgradeLicenseType) String() string {
    return []string{"productKey", "licenseFile"}[i]
}
func ParseEditionUpgradeLicenseType(v string) (any, error) {
    result := PRODUCTKEY_EDITIONUPGRADELICENSETYPE
    switch v {
        case "productKey":
            result = PRODUCTKEY_EDITIONUPGRADELICENSETYPE
        case "licenseFile":
            result = LICENSEFILE_EDITIONUPGRADELICENSETYPE
        default:
            return 0, errors.New("Unknown EditionUpgradeLicenseType value: " + v)
    }
    return &result, nil
}
func SerializeEditionUpgradeLicenseType(values []EditionUpgradeLicenseType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Possible values for internet site security level.
type InternetSiteSecurityLevel int

const (
    // User Defined, default value, no intent.
    USERDEFINED_INTERNETSITESECURITYLEVEL InternetSiteSecurityLevel = iota
    // Medium.
    MEDIUM_INTERNETSITESECURITYLEVEL
    // Medium-High.
    MEDIUMHIGH_INTERNETSITESECURITYLEVEL
    // High.
    HIGH_INTERNETSITESECURITYLEVEL
)

func (i InternetSiteSecurityLevel) String() string {
    return []string{"userDefined", "medium", "mediumHigh", "high"}[i]
}
func ParseInternetSiteSecurityLevel(v string) (any, error) {
    result := USERDEFINED_INTERNETSITESECURITYLEVEL
    switch v {
        case "userDefined":
            result = USERDEFINED_INTERNETSITESECURITYLEVEL
        case "medium":
            result = MEDIUM_INTERNETSITESECURITYLEVEL
        case "mediumHigh":
            result = MEDIUMHIGH_INTERNETSITESECURITYLEVEL
        case "high":
            result = HIGH_INTERNETSITESECURITYLEVEL
        default:
            return 0, errors.New("Unknown InternetSiteSecurityLevel value: " + v)
    }
    return &result, nil
}
func SerializeInternetSiteSecurityLevel(values []InternetSiteSecurityLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

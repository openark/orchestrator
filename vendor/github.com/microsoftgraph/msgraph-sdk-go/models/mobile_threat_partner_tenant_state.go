package models
import (
    "errors"
)
// Partner state of this tenant.
type MobileThreatPartnerTenantState int

const (
    // Partner is unavailable.
    UNAVAILABLE_MOBILETHREATPARTNERTENANTSTATE MobileThreatPartnerTenantState = iota
    // Partner is available.
    AVAILABLE_MOBILETHREATPARTNERTENANTSTATE
    // Partner is enabled.
    ENABLED_MOBILETHREATPARTNERTENANTSTATE
    // Partner is unresponsive.
    UNRESPONSIVE_MOBILETHREATPARTNERTENANTSTATE
)

func (i MobileThreatPartnerTenantState) String() string {
    return []string{"unavailable", "available", "enabled", "unresponsive"}[i]
}
func ParseMobileThreatPartnerTenantState(v string) (any, error) {
    result := UNAVAILABLE_MOBILETHREATPARTNERTENANTSTATE
    switch v {
        case "unavailable":
            result = UNAVAILABLE_MOBILETHREATPARTNERTENANTSTATE
        case "available":
            result = AVAILABLE_MOBILETHREATPARTNERTENANTSTATE
        case "enabled":
            result = ENABLED_MOBILETHREATPARTNERTENANTSTATE
        case "unresponsive":
            result = UNRESPONSIVE_MOBILETHREATPARTNERTENANTSTATE
        default:
            return 0, errors.New("Unknown MobileThreatPartnerTenantState value: " + v)
    }
    return &result, nil
}
func SerializeMobileThreatPartnerTenantState(values []MobileThreatPartnerTenantState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

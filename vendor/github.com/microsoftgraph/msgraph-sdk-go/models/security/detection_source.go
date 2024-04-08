package security
import (
    "errors"
)
// 
type DetectionSource int

const (
    UNKNOWN_DETECTIONSOURCE DetectionSource = iota
    MICROSOFTDEFENDERFORENDPOINT_DETECTIONSOURCE
    ANTIVIRUS_DETECTIONSOURCE
    SMARTSCREEN_DETECTIONSOURCE
    CUSTOMTI_DETECTIONSOURCE
    MICROSOFTDEFENDERFOROFFICE365_DETECTIONSOURCE
    AUTOMATEDINVESTIGATION_DETECTIONSOURCE
    MICROSOFTTHREATEXPERTS_DETECTIONSOURCE
    CUSTOMDETECTION_DETECTIONSOURCE
    MICROSOFTDEFENDERFORIDENTITY_DETECTIONSOURCE
    CLOUDAPPSECURITY_DETECTIONSOURCE
    MICROSOFT365DEFENDER_DETECTIONSOURCE
    AZUREADIDENTITYPROTECTION_DETECTIONSOURCE
    MANUAL_DETECTIONSOURCE
    MICROSOFTDATALOSSPREVENTION_DETECTIONSOURCE
    APPGOVERNANCEPOLICY_DETECTIONSOURCE
    APPGOVERNANCEDETECTION_DETECTIONSOURCE
    UNKNOWNFUTUREVALUE_DETECTIONSOURCE
)

func (i DetectionSource) String() string {
    return []string{"unknown", "microsoftDefenderForEndpoint", "antivirus", "smartScreen", "customTi", "microsoftDefenderForOffice365", "automatedInvestigation", "microsoftThreatExperts", "customDetection", "microsoftDefenderForIdentity", "cloudAppSecurity", "microsoft365Defender", "azureAdIdentityProtection", "manual", "microsoftDataLossPrevention", "appGovernancePolicy", "appGovernanceDetection", "unknownFutureValue"}[i]
}
func ParseDetectionSource(v string) (any, error) {
    result := UNKNOWN_DETECTIONSOURCE
    switch v {
        case "unknown":
            result = UNKNOWN_DETECTIONSOURCE
        case "microsoftDefenderForEndpoint":
            result = MICROSOFTDEFENDERFORENDPOINT_DETECTIONSOURCE
        case "antivirus":
            result = ANTIVIRUS_DETECTIONSOURCE
        case "smartScreen":
            result = SMARTSCREEN_DETECTIONSOURCE
        case "customTi":
            result = CUSTOMTI_DETECTIONSOURCE
        case "microsoftDefenderForOffice365":
            result = MICROSOFTDEFENDERFOROFFICE365_DETECTIONSOURCE
        case "automatedInvestigation":
            result = AUTOMATEDINVESTIGATION_DETECTIONSOURCE
        case "microsoftThreatExperts":
            result = MICROSOFTTHREATEXPERTS_DETECTIONSOURCE
        case "customDetection":
            result = CUSTOMDETECTION_DETECTIONSOURCE
        case "microsoftDefenderForIdentity":
            result = MICROSOFTDEFENDERFORIDENTITY_DETECTIONSOURCE
        case "cloudAppSecurity":
            result = CLOUDAPPSECURITY_DETECTIONSOURCE
        case "microsoft365Defender":
            result = MICROSOFT365DEFENDER_DETECTIONSOURCE
        case "azureAdIdentityProtection":
            result = AZUREADIDENTITYPROTECTION_DETECTIONSOURCE
        case "manual":
            result = MANUAL_DETECTIONSOURCE
        case "microsoftDataLossPrevention":
            result = MICROSOFTDATALOSSPREVENTION_DETECTIONSOURCE
        case "appGovernancePolicy":
            result = APPGOVERNANCEPOLICY_DETECTIONSOURCE
        case "appGovernanceDetection":
            result = APPGOVERNANCEDETECTION_DETECTIONSOURCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DETECTIONSOURCE
        default:
            return 0, errors.New("Unknown DetectionSource value: " + v)
    }
    return &result, nil
}
func SerializeDetectionSource(values []DetectionSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package security
import (
    "errors"
)
// 
type ServiceSource int

const (
    UNKNOWN_SERVICESOURCE ServiceSource = iota
    MICROSOFTDEFENDERFORENDPOINT_SERVICESOURCE
    MICROSOFTDEFENDERFORIDENTITY_SERVICESOURCE
    MICROSOFTDEFENDERFORCLOUDAPPS_SERVICESOURCE
    MICROSOFTDEFENDERFOROFFICE365_SERVICESOURCE
    MICROSOFT365DEFENDER_SERVICESOURCE
    AZUREADIDENTITYPROTECTION_SERVICESOURCE
    MICROSOFTAPPGOVERNANCE_SERVICESOURCE
    DATALOSSPREVENTION_SERVICESOURCE
    UNKNOWNFUTUREVALUE_SERVICESOURCE
)

func (i ServiceSource) String() string {
    return []string{"unknown", "microsoftDefenderForEndpoint", "microsoftDefenderForIdentity", "microsoftDefenderForCloudApps", "microsoftDefenderForOffice365", "microsoft365Defender", "azureAdIdentityProtection", "microsoftAppGovernance", "dataLossPrevention", "unknownFutureValue"}[i]
}
func ParseServiceSource(v string) (any, error) {
    result := UNKNOWN_SERVICESOURCE
    switch v {
        case "unknown":
            result = UNKNOWN_SERVICESOURCE
        case "microsoftDefenderForEndpoint":
            result = MICROSOFTDEFENDERFORENDPOINT_SERVICESOURCE
        case "microsoftDefenderForIdentity":
            result = MICROSOFTDEFENDERFORIDENTITY_SERVICESOURCE
        case "microsoftDefenderForCloudApps":
            result = MICROSOFTDEFENDERFORCLOUDAPPS_SERVICESOURCE
        case "microsoftDefenderForOffice365":
            result = MICROSOFTDEFENDERFOROFFICE365_SERVICESOURCE
        case "microsoft365Defender":
            result = MICROSOFT365DEFENDER_SERVICESOURCE
        case "azureAdIdentityProtection":
            result = AZUREADIDENTITYPROTECTION_SERVICESOURCE
        case "microsoftAppGovernance":
            result = MICROSOFTAPPGOVERNANCE_SERVICESOURCE
        case "dataLossPrevention":
            result = DATALOSSPREVENTION_SERVICESOURCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SERVICESOURCE
        default:
            return 0, errors.New("Unknown ServiceSource value: " + v)
    }
    return &result, nil
}
func SerializeServiceSource(values []ServiceSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

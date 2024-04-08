package security
import (
    "errors"
)
// 
type DataSourceScopes int

const (
    NONE_DATASOURCESCOPES DataSourceScopes = iota
    ALLTENANTMAILBOXES_DATASOURCESCOPES
    ALLTENANTSITES_DATASOURCESCOPES
    ALLCASECUSTODIANS_DATASOURCESCOPES
    ALLCASENONCUSTODIALDATASOURCES_DATASOURCESCOPES
    UNKNOWNFUTUREVALUE_DATASOURCESCOPES
)

func (i DataSourceScopes) String() string {
    return []string{"none", "allTenantMailboxes", "allTenantSites", "allCaseCustodians", "allCaseNoncustodialDataSources", "unknownFutureValue"}[i]
}
func ParseDataSourceScopes(v string) (any, error) {
    result := NONE_DATASOURCESCOPES
    switch v {
        case "none":
            result = NONE_DATASOURCESCOPES
        case "allTenantMailboxes":
            result = ALLTENANTMAILBOXES_DATASOURCESCOPES
        case "allTenantSites":
            result = ALLTENANTSITES_DATASOURCESCOPES
        case "allCaseCustodians":
            result = ALLCASECUSTODIANS_DATASOURCESCOPES
        case "allCaseNoncustodialDataSources":
            result = ALLCASENONCUSTODIALDATASOURCES_DATASOURCESCOPES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DATASOURCESCOPES
        default:
            return 0, errors.New("Unknown DataSourceScopes value: " + v)
    }
    return &result, nil
}
func SerializeDataSourceScopes(values []DataSourceScopes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

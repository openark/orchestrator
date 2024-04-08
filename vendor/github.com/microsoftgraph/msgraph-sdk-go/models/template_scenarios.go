package models
import (
    "errors"
)
// 
type TemplateScenarios int

const (
    NEW_TEMPLATESCENARIOS TemplateScenarios = iota
    SECUREFOUNDATION_TEMPLATESCENARIOS
    ZEROTRUST_TEMPLATESCENARIOS
    REMOTEWORK_TEMPLATESCENARIOS
    PROTECTADMINS_TEMPLATESCENARIOS
    EMERGINGTHREATS_TEMPLATESCENARIOS
    UNKNOWNFUTUREVALUE_TEMPLATESCENARIOS
)

func (i TemplateScenarios) String() string {
    return []string{"new", "secureFoundation", "zeroTrust", "remoteWork", "protectAdmins", "emergingThreats", "unknownFutureValue"}[i]
}
func ParseTemplateScenarios(v string) (any, error) {
    result := NEW_TEMPLATESCENARIOS
    switch v {
        case "new":
            result = NEW_TEMPLATESCENARIOS
        case "secureFoundation":
            result = SECUREFOUNDATION_TEMPLATESCENARIOS
        case "zeroTrust":
            result = ZEROTRUST_TEMPLATESCENARIOS
        case "remoteWork":
            result = REMOTEWORK_TEMPLATESCENARIOS
        case "protectAdmins":
            result = PROTECTADMINS_TEMPLATESCENARIOS
        case "emergingThreats":
            result = EMERGINGTHREATS_TEMPLATESCENARIOS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEMPLATESCENARIOS
        default:
            return 0, errors.New("Unknown TemplateScenarios value: " + v)
    }
    return &result, nil
}
func SerializeTemplateScenarios(values []TemplateScenarios) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

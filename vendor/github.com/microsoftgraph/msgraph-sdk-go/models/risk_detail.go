package models
import (
    "errors"
)
// 
type RiskDetail int

const (
    NONE_RISKDETAIL RiskDetail = iota
    ADMINGENERATEDTEMPORARYPASSWORD_RISKDETAIL
    USERPERFORMEDSECUREDPASSWORDCHANGE_RISKDETAIL
    USERPERFORMEDSECUREDPASSWORDRESET_RISKDETAIL
    ADMINCONFIRMEDSIGNINSAFE_RISKDETAIL
    AICONFIRMEDSIGNINSAFE_RISKDETAIL
    USERPASSEDMFADRIVENBYRISKBASEDPOLICY_RISKDETAIL
    ADMINDISMISSEDALLRISKFORUSER_RISKDETAIL
    ADMINCONFIRMEDSIGNINCOMPROMISED_RISKDETAIL
    HIDDEN_RISKDETAIL
    ADMINCONFIRMEDUSERCOMPROMISED_RISKDETAIL
    UNKNOWNFUTUREVALUE_RISKDETAIL
    ADMINCONFIRMEDSERVICEPRINCIPALCOMPROMISED_RISKDETAIL
    ADMINDISMISSEDALLRISKFORSERVICEPRINCIPAL_RISKDETAIL
    M365DADMINDISMISSEDDETECTION_RISKDETAIL
)

func (i RiskDetail) String() string {
    return []string{"none", "adminGeneratedTemporaryPassword", "userPerformedSecuredPasswordChange", "userPerformedSecuredPasswordReset", "adminConfirmedSigninSafe", "aiConfirmedSigninSafe", "userPassedMFADrivenByRiskBasedPolicy", "adminDismissedAllRiskForUser", "adminConfirmedSigninCompromised", "hidden", "adminConfirmedUserCompromised", "unknownFutureValue", "adminConfirmedServicePrincipalCompromised", "adminDismissedAllRiskForServicePrincipal", "m365DAdminDismissedDetection"}[i]
}
func ParseRiskDetail(v string) (any, error) {
    result := NONE_RISKDETAIL
    switch v {
        case "none":
            result = NONE_RISKDETAIL
        case "adminGeneratedTemporaryPassword":
            result = ADMINGENERATEDTEMPORARYPASSWORD_RISKDETAIL
        case "userPerformedSecuredPasswordChange":
            result = USERPERFORMEDSECUREDPASSWORDCHANGE_RISKDETAIL
        case "userPerformedSecuredPasswordReset":
            result = USERPERFORMEDSECUREDPASSWORDRESET_RISKDETAIL
        case "adminConfirmedSigninSafe":
            result = ADMINCONFIRMEDSIGNINSAFE_RISKDETAIL
        case "aiConfirmedSigninSafe":
            result = AICONFIRMEDSIGNINSAFE_RISKDETAIL
        case "userPassedMFADrivenByRiskBasedPolicy":
            result = USERPASSEDMFADRIVENBYRISKBASEDPOLICY_RISKDETAIL
        case "adminDismissedAllRiskForUser":
            result = ADMINDISMISSEDALLRISKFORUSER_RISKDETAIL
        case "adminConfirmedSigninCompromised":
            result = ADMINCONFIRMEDSIGNINCOMPROMISED_RISKDETAIL
        case "hidden":
            result = HIDDEN_RISKDETAIL
        case "adminConfirmedUserCompromised":
            result = ADMINCONFIRMEDUSERCOMPROMISED_RISKDETAIL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RISKDETAIL
        case "adminConfirmedServicePrincipalCompromised":
            result = ADMINCONFIRMEDSERVICEPRINCIPALCOMPROMISED_RISKDETAIL
        case "adminDismissedAllRiskForServicePrincipal":
            result = ADMINDISMISSEDALLRISKFORSERVICEPRINCIPAL_RISKDETAIL
        case "m365DAdminDismissedDetection":
            result = M365DADMINDISMISSEDDETECTION_RISKDETAIL
        default:
            return 0, errors.New("Unknown RiskDetail value: " + v)
    }
    return &result, nil
}
func SerializeRiskDetail(values []RiskDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

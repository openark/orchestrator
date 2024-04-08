package security
import (
    "errors"
)
// 
type CaseAction int

const (
    CONTENTEXPORT_CASEACTION CaseAction = iota
    APPLYTAGS_CASEACTION
    CONVERTTOPDF_CASEACTION
    INDEX_CASEACTION
    ESTIMATESTATISTICS_CASEACTION
    ADDTOREVIEWSET_CASEACTION
    HOLDUPDATE_CASEACTION
    UNKNOWNFUTUREVALUE_CASEACTION
    PURGEDATA_CASEACTION
)

func (i CaseAction) String() string {
    return []string{"contentExport", "applyTags", "convertToPdf", "index", "estimateStatistics", "addToReviewSet", "holdUpdate", "unknownFutureValue", "purgeData"}[i]
}
func ParseCaseAction(v string) (any, error) {
    result := CONTENTEXPORT_CASEACTION
    switch v {
        case "contentExport":
            result = CONTENTEXPORT_CASEACTION
        case "applyTags":
            result = APPLYTAGS_CASEACTION
        case "convertToPdf":
            result = CONVERTTOPDF_CASEACTION
        case "index":
            result = INDEX_CASEACTION
        case "estimateStatistics":
            result = ESTIMATESTATISTICS_CASEACTION
        case "addToReviewSet":
            result = ADDTOREVIEWSET_CASEACTION
        case "holdUpdate":
            result = HOLDUPDATE_CASEACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CASEACTION
        case "purgeData":
            result = PURGEDATA_CASEACTION
        default:
            return 0, errors.New("Unknown CaseAction value: " + v)
    }
    return &result, nil
}
func SerializeCaseAction(values []CaseAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

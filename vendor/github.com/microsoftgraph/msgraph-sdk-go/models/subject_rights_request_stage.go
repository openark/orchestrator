package models
import (
    "errors"
)
// 
type SubjectRightsRequestStage int

const (
    CONTENTRETRIEVAL_SUBJECTRIGHTSREQUESTSTAGE SubjectRightsRequestStage = iota
    CONTENTREVIEW_SUBJECTRIGHTSREQUESTSTAGE
    GENERATEREPORT_SUBJECTRIGHTSREQUESTSTAGE
    CONTENTDELETION_SUBJECTRIGHTSREQUESTSTAGE
    CASERESOLVED_SUBJECTRIGHTSREQUESTSTAGE
    CONTENTESTIMATE_SUBJECTRIGHTSREQUESTSTAGE
    UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTSTAGE
)

func (i SubjectRightsRequestStage) String() string {
    return []string{"contentRetrieval", "contentReview", "generateReport", "contentDeletion", "caseResolved", "contentEstimate", "unknownFutureValue"}[i]
}
func ParseSubjectRightsRequestStage(v string) (any, error) {
    result := CONTENTRETRIEVAL_SUBJECTRIGHTSREQUESTSTAGE
    switch v {
        case "contentRetrieval":
            result = CONTENTRETRIEVAL_SUBJECTRIGHTSREQUESTSTAGE
        case "contentReview":
            result = CONTENTREVIEW_SUBJECTRIGHTSREQUESTSTAGE
        case "generateReport":
            result = GENERATEREPORT_SUBJECTRIGHTSREQUESTSTAGE
        case "contentDeletion":
            result = CONTENTDELETION_SUBJECTRIGHTSREQUESTSTAGE
        case "caseResolved":
            result = CASERESOLVED_SUBJECTRIGHTSREQUESTSTAGE
        case "contentEstimate":
            result = CONTENTESTIMATE_SUBJECTRIGHTSREQUESTSTAGE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTSTAGE
        default:
            return 0, errors.New("Unknown SubjectRightsRequestStage value: " + v)
    }
    return &result, nil
}
func SerializeSubjectRightsRequestStage(values []SubjectRightsRequestStage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

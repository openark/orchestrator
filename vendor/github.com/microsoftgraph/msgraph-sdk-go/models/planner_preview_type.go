package models
import (
    "errors"
)
// 
type PlannerPreviewType int

const (
    AUTOMATIC_PLANNERPREVIEWTYPE PlannerPreviewType = iota
    NOPREVIEW_PLANNERPREVIEWTYPE
    CHECKLIST_PLANNERPREVIEWTYPE
    DESCRIPTION_PLANNERPREVIEWTYPE
    REFERENCE_PLANNERPREVIEWTYPE
)

func (i PlannerPreviewType) String() string {
    return []string{"automatic", "noPreview", "checklist", "description", "reference"}[i]
}
func ParsePlannerPreviewType(v string) (any, error) {
    result := AUTOMATIC_PLANNERPREVIEWTYPE
    switch v {
        case "automatic":
            result = AUTOMATIC_PLANNERPREVIEWTYPE
        case "noPreview":
            result = NOPREVIEW_PLANNERPREVIEWTYPE
        case "checklist":
            result = CHECKLIST_PLANNERPREVIEWTYPE
        case "description":
            result = DESCRIPTION_PLANNERPREVIEWTYPE
        case "reference":
            result = REFERENCE_PLANNERPREVIEWTYPE
        default:
            return 0, errors.New("Unknown PlannerPreviewType value: " + v)
    }
    return &result, nil
}
func SerializePlannerPreviewType(values []PlannerPreviewType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

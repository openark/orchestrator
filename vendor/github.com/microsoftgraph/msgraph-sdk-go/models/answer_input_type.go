package models
import (
    "errors"
)
// 
type AnswerInputType int

const (
    TEXT_ANSWERINPUTTYPE AnswerInputType = iota
    RADIOBUTTON_ANSWERINPUTTYPE
    UNKNOWNFUTUREVALUE_ANSWERINPUTTYPE
)

func (i AnswerInputType) String() string {
    return []string{"text", "radioButton", "unknownFutureValue"}[i]
}
func ParseAnswerInputType(v string) (any, error) {
    result := TEXT_ANSWERINPUTTYPE
    switch v {
        case "text":
            result = TEXT_ANSWERINPUTTYPE
        case "radioButton":
            result = RADIOBUTTON_ANSWERINPUTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ANSWERINPUTTYPE
        default:
            return 0, errors.New("Unknown AnswerInputType value: " + v)
    }
    return &result, nil
}
func SerializeAnswerInputType(values []AnswerInputType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

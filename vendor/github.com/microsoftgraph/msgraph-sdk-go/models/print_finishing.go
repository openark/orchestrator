package models
import (
    "errors"
)
// 
type PrintFinishing int

const (
    NONE_PRINTFINISHING PrintFinishing = iota
    STAPLE_PRINTFINISHING
    PUNCH_PRINTFINISHING
    COVER_PRINTFINISHING
    BIND_PRINTFINISHING
    SADDLESTITCH_PRINTFINISHING
    STITCHEDGE_PRINTFINISHING
    STAPLETOPLEFT_PRINTFINISHING
    STAPLEBOTTOMLEFT_PRINTFINISHING
    STAPLETOPRIGHT_PRINTFINISHING
    STAPLEBOTTOMRIGHT_PRINTFINISHING
    STITCHLEFTEDGE_PRINTFINISHING
    STITCHTOPEDGE_PRINTFINISHING
    STITCHRIGHTEDGE_PRINTFINISHING
    STITCHBOTTOMEDGE_PRINTFINISHING
    STAPLEDUALLEFT_PRINTFINISHING
    STAPLEDUALTOP_PRINTFINISHING
    STAPLEDUALRIGHT_PRINTFINISHING
    STAPLEDUALBOTTOM_PRINTFINISHING
    UNKNOWNFUTUREVALUE_PRINTFINISHING
)

func (i PrintFinishing) String() string {
    return []string{"none", "staple", "punch", "cover", "bind", "saddleStitch", "stitchEdge", "stapleTopLeft", "stapleBottomLeft", "stapleTopRight", "stapleBottomRight", "stitchLeftEdge", "stitchTopEdge", "stitchRightEdge", "stitchBottomEdge", "stapleDualLeft", "stapleDualTop", "stapleDualRight", "stapleDualBottom", "unknownFutureValue"}[i]
}
func ParsePrintFinishing(v string) (any, error) {
    result := NONE_PRINTFINISHING
    switch v {
        case "none":
            result = NONE_PRINTFINISHING
        case "staple":
            result = STAPLE_PRINTFINISHING
        case "punch":
            result = PUNCH_PRINTFINISHING
        case "cover":
            result = COVER_PRINTFINISHING
        case "bind":
            result = BIND_PRINTFINISHING
        case "saddleStitch":
            result = SADDLESTITCH_PRINTFINISHING
        case "stitchEdge":
            result = STITCHEDGE_PRINTFINISHING
        case "stapleTopLeft":
            result = STAPLETOPLEFT_PRINTFINISHING
        case "stapleBottomLeft":
            result = STAPLEBOTTOMLEFT_PRINTFINISHING
        case "stapleTopRight":
            result = STAPLETOPRIGHT_PRINTFINISHING
        case "stapleBottomRight":
            result = STAPLEBOTTOMRIGHT_PRINTFINISHING
        case "stitchLeftEdge":
            result = STITCHLEFTEDGE_PRINTFINISHING
        case "stitchTopEdge":
            result = STITCHTOPEDGE_PRINTFINISHING
        case "stitchRightEdge":
            result = STITCHRIGHTEDGE_PRINTFINISHING
        case "stitchBottomEdge":
            result = STITCHBOTTOMEDGE_PRINTFINISHING
        case "stapleDualLeft":
            result = STAPLEDUALLEFT_PRINTFINISHING
        case "stapleDualTop":
            result = STAPLEDUALTOP_PRINTFINISHING
        case "stapleDualRight":
            result = STAPLEDUALRIGHT_PRINTFINISHING
        case "stapleDualBottom":
            result = STAPLEDUALBOTTOM_PRINTFINISHING
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PRINTFINISHING
        default:
            return 0, errors.New("Unknown PrintFinishing value: " + v)
    }
    return &result, nil
}
func SerializePrintFinishing(values []PrintFinishing) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

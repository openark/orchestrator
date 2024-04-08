package security
import (
    "errors"
)
// 
type DataSourceContainerStatus int

const (
    ACTIVE_DATASOURCECONTAINERSTATUS DataSourceContainerStatus = iota
    RELEASED_DATASOURCECONTAINERSTATUS
    UNKNOWNFUTUREVALUE_DATASOURCECONTAINERSTATUS
)

func (i DataSourceContainerStatus) String() string {
    return []string{"active", "released", "unknownFutureValue"}[i]
}
func ParseDataSourceContainerStatus(v string) (any, error) {
    result := ACTIVE_DATASOURCECONTAINERSTATUS
    switch v {
        case "active":
            result = ACTIVE_DATASOURCECONTAINERSTATUS
        case "released":
            result = RELEASED_DATASOURCECONTAINERSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DATASOURCECONTAINERSTATUS
        default:
            return 0, errors.New("Unknown DataSourceContainerStatus value: " + v)
    }
    return &result, nil
}
func SerializeDataSourceContainerStatus(values []DataSourceContainerStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

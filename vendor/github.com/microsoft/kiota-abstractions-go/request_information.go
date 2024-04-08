package abstractions

import (
	"context"
	"errors"
	"time"

	"reflect"
	"strconv"
	"strings"

	u "net/url"

	"github.com/google/uuid"
	s "github.com/microsoft/kiota-abstractions-go/serialization"
	t "github.com/yosida95/uritemplate/v3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// RequestInformation represents an abstract HTTP request.
type RequestInformation struct {
	// The HTTP method of the request.
	Method HttpMethod
	uri    *u.URL
	// The Request Headers.
	Headers *RequestHeaders
	// The Query Parameters of the request.
	QueryParameters map[string]string
	// The Request Body.
	Content []byte
	// The path parameters to use for the URL template when generating the URI.
	PathParameters map[string]string
	// The Url template for the current request.
	UrlTemplate string
	options     map[string]RequestOption
}

const raw_url_key = "request-raw-url"

// NewRequestInformation creates a new RequestInformation object with default values.
func NewRequestInformation() *RequestInformation {
	return &RequestInformation{
		Headers:         NewRequestHeaders(),
		QueryParameters: make(map[string]string),
		options:         make(map[string]RequestOption),
		PathParameters:  make(map[string]string),
	}
}

// GetUri returns the URI of the request.
func (request *RequestInformation) GetUri() (*u.URL, error) {
	if request.uri != nil {
		return request.uri, nil
	} else if request.UrlTemplate == "" {
		return nil, errors.New("uri cannot be empty")
	} else if request.PathParameters == nil {
		return nil, errors.New("uri template parameters cannot be nil")
	} else if request.QueryParameters == nil {
		return nil, errors.New("uri query parameters cannot be nil")
	} else if request.PathParameters[raw_url_key] != "" {
		uri, err := u.Parse(request.PathParameters[raw_url_key])
		if err != nil {
			return nil, err
		}
		request.SetUri(*uri)
		return request.uri, nil
	} else {
		_, baseurlExists := request.PathParameters["baseurl"]
		if !baseurlExists && strings.Contains(strings.ToLower(request.UrlTemplate), "{+baseurl}") {
			return nil, errors.New("pathParameters must contain a value for \"baseurl\" for the url to be built")
		}

		uriTemplate, err := t.New(request.UrlTemplate)
		if err != nil {
			return nil, err
		}
		values := t.Values{}
		varNames := uriTemplate.Varnames()
		normalizedNames := make(map[string]string)
		for _, varName := range varNames {
			normalizedNames[strings.ToLower(varName)] = varName
		}
		for key, value := range request.PathParameters {
			addParameterWithOriginalName(key, value, normalizedNames, values)
		}
		for key, value := range request.QueryParameters {
			addParameterWithOriginalName(key, value, normalizedNames, values)
		}
		url, err := uriTemplate.Expand(values)
		if err != nil {
			return nil, err
		}
		uri, err := u.Parse(url)
		return uri, err
	}
}

// addParameterWithOriginalName adds the URI template parameter to the template using the right casing, because of go conventions, casing might have changed for the generated property
func addParameterWithOriginalName(key string, value string, normalizedNames map[string]string, values t.Values) {
	lowercaseKey := strings.ToLower(key)
	if paramName, ok := normalizedNames[lowercaseKey]; ok {
		values.Set(paramName, t.String(value))
	} else {
		values.Set(key, t.String(value))
	}
}

// SetUri updates the URI for the request from a raw URL.
func (request *RequestInformation) SetUri(url u.URL) {
	request.uri = &url
	for k := range request.PathParameters {
		delete(request.PathParameters, k)
	}
	for k := range request.QueryParameters {
		delete(request.QueryParameters, k)
	}
}

// AddRequestOptions adds an option to the request to be read by the middleware infrastructure.
func (request *RequestInformation) AddRequestOptions(options []RequestOption) {
	if options == nil {
		return
	}
	if request.options == nil {
		request.options = make(map[string]RequestOption, len(options))
	}
	for _, option := range options {
		request.options[option.GetKey().Key] = option
	}
}

// GetRequestOptions returns the options for this request. Options are unique by type. If an option of the same type is added twice, the last one wins.
func (request *RequestInformation) GetRequestOptions() []RequestOption {
	if request.options == nil {
		return []RequestOption{}
	}
	result := make([]RequestOption, len(request.options))
	idx := 0
	for _, option := range request.options {
		result[idx] = option
		idx++
	}
	return result
}

const contentTypeHeader = "Content-Type"
const binaryContentType = "application/octet-steam"

// SetStreamContent sets the request body to a binary stream.
func (request *RequestInformation) SetStreamContent(content []byte) {
	request.Content = content
	if request.Headers != nil {
		request.Headers.Add(contentTypeHeader, binaryContentType)
	}
}

func (request *RequestInformation) setContentAndContentType(writer s.SerializationWriter, contentType string) error {
	content, err := writer.GetSerializedContent()
	if err != nil {
		return err
	} else if content == nil {
		return errors.New("content cannot be nil")
	}
	request.Content = content
	if request.Headers != nil {
		request.Headers.Add(contentTypeHeader, contentType)
	}
	return nil
}

func (request *RequestInformation) getSerializationWriter(requestAdapter RequestAdapter, contentType string, items ...interface{}) (s.SerializationWriter, error) {
	if contentType == "" {
		return nil, errors.New("content type cannot be empty")
	} else if requestAdapter == nil {
		return nil, errors.New("requestAdapter cannot be nil")
	} else if len(items) == 0 {
		return nil, errors.New("items cannot be nil or empty")
	}
	factory := requestAdapter.GetSerializationWriterFactory()
	if factory == nil {
		return nil, errors.New("factory cannot be nil")
	}
	writer, err := factory.GetSerializationWriter(contentType)
	if err != nil {
		return nil, err
	} else if writer == nil {
		return nil, errors.New("writer cannot be nil")
	} else {
		return writer, nil
	}
}

func (r *RequestInformation) setRequestType(result interface{}, span trace.Span) {
	if result != nil {
		span.SetAttributes(attribute.String("com.microsoft.kiota.request.type", reflect.TypeOf(result).String()))
	}
}

const observabilityTracerName = "github.com/microsoft/kiota-abstractions-go"

// SetContentFromParsable sets the request body from a model with the specified content type.
func (request *RequestInformation) SetContentFromParsable(ctx context.Context, requestAdapter RequestAdapter, contentType string, item s.Parsable) error {
	_, span := otel.GetTracerProvider().Tracer(observabilityTracerName).Start(ctx, "SetContentFromParsable")
	defer span.End()

	writer, err := request.getSerializationWriter(requestAdapter, contentType, item)
	if err != nil {
		span.RecordError(err)
		return err
	}
	defer writer.Close()
	request.setRequestType(item, span)
	err = writer.WriteObjectValue("", item)
	if err != nil {
		span.RecordError(err)
		return err
	}
	err = request.setContentAndContentType(writer, contentType)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}

// SetContentFromParsableCollection sets the request body from a model with the specified content type.
func (request *RequestInformation) SetContentFromParsableCollection(ctx context.Context, requestAdapter RequestAdapter, contentType string, items []s.Parsable) error {
	_, span := otel.GetTracerProvider().Tracer(observabilityTracerName).Start(ctx, "SetContentFromParsableCollection")
	defer span.End()

	writer, err := request.getSerializationWriter(requestAdapter, contentType, items)
	if err != nil {
		span.RecordError(err)
		return err
	}
	defer writer.Close()
	if len(items) > 0 {
		request.setRequestType(items[0], span)
	}
	err = writer.WriteCollectionOfObjectValues("", items)
	if err != nil {
		span.RecordError(err)
		return err
	}
	err = request.setContentAndContentType(writer, contentType)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}

// SetContentFromScalar sets the request body from a scalar value with the specified content type.
func (request *RequestInformation) SetContentFromScalar(ctx context.Context, requestAdapter RequestAdapter, contentType string, item interface{}) error {
	_, span := otel.GetTracerProvider().Tracer(observabilityTracerName).Start(ctx, "SetContentFromScalar")
	defer span.End()
	writer, err := request.getSerializationWriter(requestAdapter, contentType, item)
	if err != nil {
		span.RecordError(err)
		return err
	}
	defer writer.Close()
	request.setRequestType(item, span)

	if sv, ok := item.(*string); ok {
		err = writer.WriteStringValue("", sv)
	} else if bv, ok := item.(*bool); ok {
		err = writer.WriteBoolValue("", bv)
	} else if byv, ok := item.(*byte); ok {
		err = writer.WriteByteValue("", byv)
	} else if i8v, ok := item.(*int8); ok {
		err = writer.WriteInt8Value("", i8v)
	} else if i32v, ok := item.(*int32); ok {
		err = writer.WriteInt32Value("", i32v)
	} else if i64v, ok := item.(*int64); ok {
		err = writer.WriteInt64Value("", i64v)
	} else if f32v, ok := item.(*float32); ok {
		err = writer.WriteFloat32Value("", f32v)
	} else if f64v, ok := item.(*float64); ok {
		err = writer.WriteFloat64Value("", f64v)
	} else if uv, ok := item.(*uuid.UUID); ok {
		err = writer.WriteUUIDValue("", uv)
	} else if tv, ok := item.(*time.Time); ok {
		err = writer.WriteTimeValue("", tv)
	} else if dv, ok := item.(*s.ISODuration); ok {
		err = writer.WriteISODurationValue("", dv)
	} else if tov, ok := item.(*s.TimeOnly); ok {
		err = writer.WriteTimeOnlyValue("", tov)
	} else if dov, ok := item.(*s.DateOnly); ok {
		err = writer.WriteDateOnlyValue("", dov)
	}
	if err != nil {
		span.RecordError(err)
		return err
	}
	err = request.setContentAndContentType(writer, contentType)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}

// SetContentFromScalarCollection sets the request body from a scalar value with the specified content type.
func (request *RequestInformation) SetContentFromScalarCollection(ctx context.Context, requestAdapter RequestAdapter, contentType string, items []interface{}) error {
	_, span := otel.GetTracerProvider().Tracer(observabilityTracerName).Start(ctx, "SetContentFromScalarCollection")
	defer span.End()
	writer, err := request.getSerializationWriter(requestAdapter, contentType, items...)
	if err != nil {
		span.RecordError(err)
		return err
	}
	defer writer.Close()
	if len(items) > 0 {
		value := items[0]
		request.setRequestType(value, span)
		if _, ok := value.(string); ok {
			sc := make([]string, len(items))
			for i, v := range items {
				if sv, ok := v.(string); ok {
					sc[i] = sv
				}
			}
			err = writer.WriteCollectionOfStringValues("", sc)
		} else if _, ok := value.(bool); ok {
			bc := make([]bool, len(items))
			for i, v := range items {
				if sv, ok := v.(bool); ok {
					bc[i] = sv
				}
			}
			err = writer.WriteCollectionOfBoolValues("", bc)
		} else if _, ok := value.(byte); ok {
			byc := make([]byte, len(items))
			for i, v := range items {
				if sv, ok := v.(byte); ok {
					byc[i] = sv
				}
			}
			err = writer.WriteCollectionOfByteValues("", byc)
		} else if _, ok := value.(int8); ok {
			i8c := make([]int8, len(items))
			for i, v := range items {
				if sv, ok := v.(int8); ok {
					i8c[i] = sv
				}
			}
			err = writer.WriteCollectionOfInt8Values("", i8c)
		} else if _, ok := value.(int32); ok {
			i32c := make([]int32, len(items))
			for i, v := range items {
				if sv, ok := v.(int32); ok {
					i32c[i] = sv
				}
			}
			err = writer.WriteCollectionOfInt32Values("", i32c)
		} else if _, ok := value.(int64); ok {
			i64c := make([]int64, len(items))
			for i, v := range items {
				if sv, ok := v.(int64); ok {
					i64c[i] = sv
				}
			}
			err = writer.WriteCollectionOfInt64Values("", i64c)
		} else if _, ok := value.(float32); ok {
			f32c := make([]float32, len(items))
			for i, v := range items {
				if sv, ok := v.(float32); ok {
					f32c[i] = sv
				}
			}
			err = writer.WriteCollectionOfFloat32Values("", f32c)
		} else if _, ok := value.(float64); ok {
			f64c := make([]float64, len(items))
			for i, v := range items {
				if sv, ok := v.(float64); ok {
					f64c[i] = sv
				}
			}
			err = writer.WriteCollectionOfFloat64Values("", f64c)
		} else if _, ok := value.(uuid.UUID); ok {
			uc := make([]uuid.UUID, len(items))
			for i, v := range items {
				if sv, ok := v.(uuid.UUID); ok {
					uc[i] = sv
				}
			}
			err = writer.WriteCollectionOfUUIDValues("", uc)
		} else if _, ok := value.(time.Time); ok {
			tc := make([]time.Time, len(items))
			for i, v := range items {
				if sv, ok := v.(time.Time); ok {
					tc[i] = sv
				}
			}
			err = writer.WriteCollectionOfTimeValues("", tc)
		} else if _, ok := value.(s.ISODuration); ok {
			dc := make([]s.ISODuration, len(items))
			for i, v := range items {
				if sv, ok := v.(s.ISODuration); ok {
					dc[i] = sv
				}
			}
			err = writer.WriteCollectionOfISODurationValues("", dc)
		} else if _, ok := value.(s.TimeOnly); ok {
			toc := make([]s.TimeOnly, len(items))
			for i, v := range items {
				if sv, ok := v.(s.TimeOnly); ok {
					toc[i] = sv
				}
			}
			err = writer.WriteCollectionOfTimeOnlyValues("", toc)
		} else if _, ok := value.(s.DateOnly); ok {
			doc := make([]s.DateOnly, len(items))
			for i, v := range items {
				if sv, ok := v.(s.DateOnly); ok {
					doc[i] = sv
				}
			}
			err = writer.WriteCollectionOfDateOnlyValues("", doc)
		} else if _, ok := value.(byte); ok {
			ba := make([]byte, len(items))
			for i, v := range items {
				if sv, ok := v.(byte); ok {
					ba[i] = sv
				}
			}
			err = writer.WriteByteArrayValue("", ba)
		}
	}
	if err != nil {
		span.RecordError(err)
		return err
	}
	err = request.setContentAndContentType(writer, contentType)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}

// AddQueryParameters adds the query parameters to the request by reading the properties from the provided object.
func (request *RequestInformation) AddQueryParameters(source interface{}) {
	if source == nil || request == nil {
		return
	}
	valOfP := reflect.ValueOf(source)
	fields := reflect.TypeOf(source)
	numOfFields := fields.NumField()
	for i := 0; i < numOfFields; i++ {
		field := fields.Field(i)
		fieldName := field.Name
		fieldValue := valOfP.Field(i)
		tagValue := field.Tag.Get("uriparametername")
		if tagValue != "" {
			fieldName = tagValue
		}
		value := fieldValue.Interface()
		if value == nil {
			continue
		}
		str, ok := value.(*string)
		if ok && str != nil && *str != "" {
			request.QueryParameters[fieldName] = *str
		}
		bl, ok := value.(*bool)
		if ok && bl != nil {
			request.QueryParameters[fieldName] = strconv.FormatBool(*bl)
		}
		it, ok := value.(*int32)
		if ok && it != nil {
			request.QueryParameters[fieldName] = strconv.FormatInt(int64(*it), 10)
		}
		arr, ok := value.([]string)
		if ok && len(arr) > 0 {
			request.QueryParameters[fieldName] = strings.Join(arr, ",")
		}
	}
}

package msgraphgocore

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type batchResponse struct {
	responses     []BatchItem
	indexResponse map[string]BatchItem
	isIndexed     bool
}

// GetResponses returns a slice of BatchItem to the user
func (br *batchResponse) GetResponses() []BatchItem {
	return br.responses
}

// SetResponses adds a slice of BatchItem to the response
func (br *batchResponse) SetResponses(responses []BatchItem) {
	br.responses = responses
}

// GetResponseById returns a response payload as a batch item
func (br *batchResponse) GetResponseById(itemId string) BatchItem {
	if !br.isIndexed {

		for _, resp := range br.GetResponses() {
			br.indexResponse[*(resp.GetId())] = resp
		}

		br.isIndexed = true
	}

	return br.indexResponse[itemId]
}

// CreateBatchResponseDiscriminator creates a new instance of the appropriate class based on discriminator value
func CreateBatchResponseDiscriminator(serialization.ParseNode) (serialization.Parsable, error) {
	res := batchResponse{
		indexResponse: make(map[string]BatchItem),
		isIndexed:     false,
	}
	return &res, nil
}

// BatchResponse instance of batch request result payload
type BatchResponse interface {
	serialization.Parsable
	GetResponses() []BatchItem
	SetResponses(responses []BatchItem)
	GetResponseById(itemId string) BatchItem
}

// Serialize serializes information the current object
func (br *batchResponse) Serialize(serialization.SerializationWriter) error {
	panic("batch responses are not serializable")
}

// GetFieldDeserializers the deserialization information for the current model
func (br *batchResponse) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	res := make(map[string]func(serialization.ParseNode) error)
	res["responses"] = func(n serialization.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateBatchRequestItemDiscriminator)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]BatchItem, len(val))
			for i, v := range val {
				res[i] = v.(BatchItem)
			}
			br.SetResponses(res)
		}
		return nil
	}
	return res
}

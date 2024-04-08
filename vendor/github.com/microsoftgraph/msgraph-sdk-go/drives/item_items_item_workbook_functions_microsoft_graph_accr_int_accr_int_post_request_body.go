package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The basis property
    basis iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The calcMethod property
    calcMethod iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The firstInterest property
    firstInterest iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The frequency property
    frequency iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The issue property
    issue iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The par property
    par iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The rate property
    rate iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The settlement property
    settlement iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBasis gets the basis property value. The basis property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) GetBasis()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.basis
}
// GetCalcMethod gets the calcMethod property value. The calcMethod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) GetCalcMethod()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.calcMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["basis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBasis(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["calcMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalcMethod(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["firstInterest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstInterest(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["frequency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrequency(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["issue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssue(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["par"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPar(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["rate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRate(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["settlement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettlement(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetFirstInterest gets the firstInterest property value. The firstInterest property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) GetFirstInterest()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.firstInterest
}
// GetFrequency gets the frequency property value. The frequency property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) GetFrequency()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.frequency
}
// GetIssue gets the issue property value. The issue property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) GetIssue()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.issue
}
// GetPar gets the par property value. The par property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) GetPar()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.par
}
// GetRate gets the rate property value. The rate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) GetRate()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.rate
}
// GetSettlement gets the settlement property value. The settlement property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) GetSettlement()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.settlement
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("basis", m.GetBasis())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("calcMethod", m.GetCalcMethod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("firstInterest", m.GetFirstInterest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("frequency", m.GetFrequency())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("issue", m.GetIssue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("par", m.GetPar())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rate", m.GetRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("settlement", m.GetSettlement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBasis sets the basis property value. The basis property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) SetBasis(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.basis = value
}
// SetCalcMethod sets the calcMethod property value. The calcMethod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) SetCalcMethod(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.calcMethod = value
}
// SetFirstInterest sets the firstInterest property value. The firstInterest property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) SetFirstInterest(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.firstInterest = value
}
// SetFrequency sets the frequency property value. The frequency property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) SetFrequency(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.frequency = value
}
// SetIssue sets the issue property value. The issue property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) SetIssue(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.issue = value
}
// SetPar sets the par property value. The par property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) SetPar(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.par = value
}
// SetRate sets the rate property value. The rate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) SetRate(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.rate = value
}
// SetSettlement sets the settlement property value. The settlement property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntPostRequestBody) SetSettlement(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.settlement = value
}

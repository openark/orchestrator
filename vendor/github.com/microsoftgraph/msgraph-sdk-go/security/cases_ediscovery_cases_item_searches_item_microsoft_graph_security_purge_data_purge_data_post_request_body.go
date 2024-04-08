package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody 
type CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The purgeAreas property
    purgeAreas *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PurgeAreas
    // The purgeType property
    purgeType *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PurgeType
}
// NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody instantiates a new CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody()(*CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) {
    m := &CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["purgeAreas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ParsePurgeAreas)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurgeAreas(val.(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PurgeAreas))
        }
        return nil
    }
    res["purgeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ParsePurgeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurgeType(val.(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PurgeType))
        }
        return nil
    }
    return res
}
// GetPurgeAreas gets the purgeAreas property value. The purgeAreas property
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) GetPurgeAreas()(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PurgeAreas) {
    return m.purgeAreas
}
// GetPurgeType gets the purgeType property value. The purgeType property
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) GetPurgeType()(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PurgeType) {
    return m.purgeType
}
// Serialize serializes information the current object
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPurgeAreas() != nil {
        cast := (*m.GetPurgeAreas()).String()
        err := writer.WriteStringValue("purgeAreas", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPurgeType() != nil {
        cast := (*m.GetPurgeType()).String()
        err := writer.WriteStringValue("purgeType", &cast)
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
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPurgeAreas sets the purgeAreas property value. The purgeAreas property
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) SetPurgeAreas(value *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PurgeAreas)() {
    m.purgeAreas = value
}
// SetPurgeType sets the purgeType property value. The purgeType property
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) SetPurgeType(value *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PurgeType)() {
    m.purgeType = value
}

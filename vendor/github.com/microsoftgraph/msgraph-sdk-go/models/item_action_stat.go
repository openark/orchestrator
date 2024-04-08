package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionStat 
type ItemActionStat struct {
    // The number of times the action took place. Read-only.
    actionCount *int32
    // The number of distinct actors that performed the action. Read-only.
    actorCount *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
}
// NewItemActionStat instantiates a new itemActionStat and sets the default values.
func NewItemActionStat()(*ItemActionStat) {
    m := &ItemActionStat{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionStatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionStatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionStat(), nil
}
// GetActionCount gets the actionCount property value. The number of times the action took place. Read-only.
func (m *ItemActionStat) GetActionCount()(*int32) {
    return m.actionCount
}
// GetActorCount gets the actorCount property value. The number of distinct actors that performed the action. Read-only.
func (m *ItemActionStat) GetActorCount()(*int32) {
    return m.actorCount
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionStat) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActionStat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionCount(val)
        }
        return nil
    }
    res["actorCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorCount(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ItemActionStat) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ItemActionStat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("actionCount", m.GetActionCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("actorCount", m.GetActorCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetActionCount sets the actionCount property value. The number of times the action took place. Read-only.
func (m *ItemActionStat) SetActionCount(value *int32)() {
    m.actionCount = value
}
// SetActorCount sets the actorCount property value. The number of distinct actors that performed the action. Read-only.
func (m *ItemActionStat) SetActorCount(value *int32)() {
    m.actorCount = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionStat) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ItemActionStat) SetOdataType(value *string)() {
    m.odataType = value
}

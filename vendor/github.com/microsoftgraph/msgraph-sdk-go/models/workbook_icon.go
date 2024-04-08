package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookIcon 
type WorkbookIcon struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Represents the index of the icon in the given set.
    index *int32
    // The OdataType property
    odataType *string
    // Represents the set that the icon is part of. The possible values are: Invalid, ThreeArrows, ThreeArrowsGray, ThreeFlags, ThreeTrafficLights1, ThreeTrafficLights2, ThreeSigns, ThreeSymbols, ThreeSymbols2, FourArrows, FourArrowsGray, FourRedToBlack, FourRating, FourTrafficLights, FiveArrows, FiveArrowsGray, FiveRating, FiveQuarters, ThreeStars, ThreeTriangles, FiveBoxes.
    set *string
}
// NewWorkbookIcon instantiates a new workbookIcon and sets the default values.
func NewWorkbookIcon()(*WorkbookIcon) {
    m := &WorkbookIcon{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkbookIconFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookIconFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookIcon(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookIcon) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookIcon) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
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
    res["set"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSet(val)
        }
        return nil
    }
    return res
}
// GetIndex gets the index property value. Represents the index of the icon in the given set.
func (m *WorkbookIcon) GetIndex()(*int32) {
    return m.index
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkbookIcon) GetOdataType()(*string) {
    return m.odataType
}
// GetSet gets the set property value. Represents the set that the icon is part of. The possible values are: Invalid, ThreeArrows, ThreeArrowsGray, ThreeFlags, ThreeTrafficLights1, ThreeTrafficLights2, ThreeSigns, ThreeSymbols, ThreeSymbols2, FourArrows, FourArrowsGray, FourRedToBlack, FourRating, FourTrafficLights, FiveArrows, FiveArrowsGray, FiveRating, FiveQuarters, ThreeStars, ThreeTriangles, FiveBoxes.
func (m *WorkbookIcon) GetSet()(*string) {
    return m.set
}
// Serialize serializes information the current object
func (m *WorkbookIcon) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("index", m.GetIndex())
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
        err := writer.WriteStringValue("set", m.GetSet())
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
func (m *WorkbookIcon) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIndex sets the index property value. Represents the index of the icon in the given set.
func (m *WorkbookIcon) SetIndex(value *int32)() {
    m.index = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkbookIcon) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSet sets the set property value. Represents the set that the icon is part of. The possible values are: Invalid, ThreeArrows, ThreeArrowsGray, ThreeFlags, ThreeTrafficLights1, ThreeTrafficLights2, ThreeSigns, ThreeSymbols, ThreeSymbols2, FourArrows, FourArrowsGray, FourRedToBlack, FourRating, FourTrafficLights, FiveArrows, FiveArrowsGray, FiveRating, FiveQuarters, ThreeStars, ThreeTriangles, FiveBoxes.
func (m *WorkbookIcon) SetSet(value *string)() {
    m.set = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TargetManager 
type TargetManager struct {
    SubjectSet
    // Manager level, between 1 and 4. The direct manager is 1.
    managerLevel *int32
}
// NewTargetManager instantiates a new TargetManager and sets the default values.
func NewTargetManager()(*TargetManager) {
    m := &TargetManager{
        SubjectSet: *NewSubjectSet(),
    }
    odataTypeValue := "#microsoft.graph.targetManager"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTargetManagerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTargetManagerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTargetManager(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TargetManager) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubjectSet.GetFieldDeserializers()
    res["managerLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagerLevel(val)
        }
        return nil
    }
    return res
}
// GetManagerLevel gets the managerLevel property value. Manager level, between 1 and 4. The direct manager is 1.
func (m *TargetManager) GetManagerLevel()(*int32) {
    return m.managerLevel
}
// Serialize serializes information the current object
func (m *TargetManager) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubjectSet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("managerLevel", m.GetManagerLevel())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetManagerLevel sets the managerLevel property value. Manager level, between 1 and 4. The direct manager is 1.
func (m *TargetManager) SetManagerLevel(value *int32)() {
    m.managerLevel = value
}

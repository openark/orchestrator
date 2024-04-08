package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AgreementFileProperties 
type AgreementFileProperties struct {
    Entity
    // The date time representing when the file was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Localized display name of the policy file of an agreement. The localized display name is shown to end users who view the agreement.
    displayName *string
    // Data that represents the terms of use PDF document. Read-only.
    fileData AgreementFileDataable
    // Name of the agreement file (for example, TOU.pdf). Read-only.
    fileName *string
    // If none of the languages matches the client preference, indicates whether this is the default agreement file . If none of the files are marked as default, the first one is treated as the default. Read-only.
    isDefault *bool
    // Indicates whether the agreement file is a major version update. Major version updates invalidate the agreement's acceptances on the corresponding language.
    isMajorVersion *bool
    // The language of the agreement file in the format 'languagecode2-country/regioncode2'. 'languagecode2' is a lowercase two-letter code derived from ISO 639-1, while 'country/regioncode2' is derived from ISO 3166 and usually consists of two uppercase letters, or a BCP-47 language tag. For example, U.S. English is en-US. Read-only.
    language *string
}
// NewAgreementFileProperties instantiates a new agreementFileProperties and sets the default values.
func NewAgreementFileProperties()(*AgreementFileProperties) {
    m := &AgreementFileProperties{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAgreementFilePropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgreementFilePropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.agreementFile":
                        return NewAgreementFile(), nil
                    case "#microsoft.graph.agreementFileLocalization":
                        return NewAgreementFileLocalization(), nil
                    case "#microsoft.graph.agreementFileVersion":
                        return NewAgreementFileVersion(), nil
                }
            }
        }
    }
    return NewAgreementFileProperties(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date time representing when the file was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AgreementFileProperties) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. Localized display name of the policy file of an agreement. The localized display name is shown to end users who view the agreement.
func (m *AgreementFileProperties) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgreementFileProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["fileData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAgreementFileDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileData(val.(AgreementFileDataable))
        }
        return nil
    }
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isMajorVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMajorVersion(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    return res
}
// GetFileData gets the fileData property value. Data that represents the terms of use PDF document. Read-only.
func (m *AgreementFileProperties) GetFileData()(AgreementFileDataable) {
    return m.fileData
}
// GetFileName gets the fileName property value. Name of the agreement file (for example, TOU.pdf). Read-only.
func (m *AgreementFileProperties) GetFileName()(*string) {
    return m.fileName
}
// GetIsDefault gets the isDefault property value. If none of the languages matches the client preference, indicates whether this is the default agreement file . If none of the files are marked as default, the first one is treated as the default. Read-only.
func (m *AgreementFileProperties) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetIsMajorVersion gets the isMajorVersion property value. Indicates whether the agreement file is a major version update. Major version updates invalidate the agreement's acceptances on the corresponding language.
func (m *AgreementFileProperties) GetIsMajorVersion()(*bool) {
    return m.isMajorVersion
}
// GetLanguage gets the language property value. The language of the agreement file in the format 'languagecode2-country/regioncode2'. 'languagecode2' is a lowercase two-letter code derived from ISO 639-1, while 'country/regioncode2' is derived from ISO 3166 and usually consists of two uppercase letters, or a BCP-47 language tag. For example, U.S. English is en-US. Read-only.
func (m *AgreementFileProperties) GetLanguage()(*string) {
    return m.language
}
// Serialize serializes information the current object
func (m *AgreementFileProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("fileData", m.GetFileData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMajorVersion", m.GetIsMajorVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date time representing when the file was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AgreementFileProperties) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. Localized display name of the policy file of an agreement. The localized display name is shown to end users who view the agreement.
func (m *AgreementFileProperties) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFileData sets the fileData property value. Data that represents the terms of use PDF document. Read-only.
func (m *AgreementFileProperties) SetFileData(value AgreementFileDataable)() {
    m.fileData = value
}
// SetFileName sets the fileName property value. Name of the agreement file (for example, TOU.pdf). Read-only.
func (m *AgreementFileProperties) SetFileName(value *string)() {
    m.fileName = value
}
// SetIsDefault sets the isDefault property value. If none of the languages matches the client preference, indicates whether this is the default agreement file . If none of the files are marked as default, the first one is treated as the default. Read-only.
func (m *AgreementFileProperties) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetIsMajorVersion sets the isMajorVersion property value. Indicates whether the agreement file is a major version update. Major version updates invalidate the agreement's acceptances on the corresponding language.
func (m *AgreementFileProperties) SetIsMajorVersion(value *bool)() {
    m.isMajorVersion = value
}
// SetLanguage sets the language property value. The language of the agreement file in the format 'languagecode2-country/regioncode2'. 'languagecode2' is a lowercase two-letter code derived from ISO 639-1, while 'country/regioncode2' is derived from ISO 3166 and usually consists of two uppercase letters, or a BCP-47 language tag. For example, U.S. English is en-US. Read-only.
func (m *AgreementFileProperties) SetLanguage(value *string)() {
    m.language = value
}

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileApp an abstract class containing the base properties for Intune mobile apps.
type MobileApp struct {
    Entity
    // The list of group assignments for this mobile app.
    assignments []MobileAppAssignmentable
    // The list of categories for this app.
    categories []MobileAppCategoryable
    // The date and time the app was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description of the app.
    description *string
    // The developer of the app.
    developer *string
    // The admin provided or imported title of the app.
    displayName *string
    // The more information Url.
    informationUrl *string
    // The value indicating whether the app is marked as featured by the admin.
    isFeatured *bool
    // The large icon, to be displayed in the app details and used for upload of the icon.
    largeIcon MimeContentable
    // The date and time the app was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Notes for the app.
    notes *string
    // The owner of the app.
    owner *string
    // The privacy statement Url.
    privacyInformationUrl *string
    // The publisher of the app.
    publisher *string
    // Indicates the publishing state of an app.
    publishingState *MobileAppPublishingState
}
// NewMobileApp instantiates a new mobileApp and sets the default values.
func NewMobileApp()(*MobileApp) {
    m := &MobileApp{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.androidLobApp":
                        return NewAndroidLobApp(), nil
                    case "#microsoft.graph.androidStoreApp":
                        return NewAndroidStoreApp(), nil
                    case "#microsoft.graph.iosiPadOSWebClip":
                        return NewIosiPadOSWebClip(), nil
                    case "#microsoft.graph.iosLobApp":
                        return NewIosLobApp(), nil
                    case "#microsoft.graph.iosStoreApp":
                        return NewIosStoreApp(), nil
                    case "#microsoft.graph.iosVppApp":
                        return NewIosVppApp(), nil
                    case "#microsoft.graph.macOSLobApp":
                        return NewMacOSLobApp(), nil
                    case "#microsoft.graph.macOSMicrosoftEdgeApp":
                        return NewMacOSMicrosoftEdgeApp(), nil
                    case "#microsoft.graph.macOSOfficeSuiteApp":
                        return NewMacOSOfficeSuiteApp(), nil
                    case "#microsoft.graph.managedAndroidLobApp":
                        return NewManagedAndroidLobApp(), nil
                    case "#microsoft.graph.managedAndroidStoreApp":
                        return NewManagedAndroidStoreApp(), nil
                    case "#microsoft.graph.managedApp":
                        return NewManagedApp(), nil
                    case "#microsoft.graph.managedIOSLobApp":
                        return NewManagedIOSLobApp(), nil
                    case "#microsoft.graph.managedIOSStoreApp":
                        return NewManagedIOSStoreApp(), nil
                    case "#microsoft.graph.managedMobileLobApp":
                        return NewManagedMobileLobApp(), nil
                    case "#microsoft.graph.microsoftStoreForBusinessApp":
                        return NewMicrosoftStoreForBusinessApp(), nil
                    case "#microsoft.graph.mobileLobApp":
                        return NewMobileLobApp(), nil
                    case "#microsoft.graph.webApp":
                        return NewWebApp(), nil
                    case "#microsoft.graph.win32LobApp":
                        return NewWin32LobApp(), nil
                    case "#microsoft.graph.windowsMicrosoftEdgeApp":
                        return NewWindowsMicrosoftEdgeApp(), nil
                    case "#microsoft.graph.windowsMobileMSI":
                        return NewWindowsMobileMSI(), nil
                    case "#microsoft.graph.windowsUniversalAppX":
                        return NewWindowsUniversalAppX(), nil
                    case "#microsoft.graph.windowsWebApp":
                        return NewWindowsWebApp(), nil
                }
            }
        }
    }
    return NewMobileApp(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments for this mobile app.
func (m *MobileApp) GetAssignments()([]MobileAppAssignmentable) {
    return m.assignments
}
// GetCategories gets the categories property value. The list of categories for this app.
func (m *MobileApp) GetCategories()([]MobileAppCategoryable) {
    return m.categories
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the app was created.
func (m *MobileApp) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description of the app.
func (m *MobileApp) GetDescription()(*string) {
    return m.description
}
// GetDeveloper gets the developer property value. The developer of the app.
func (m *MobileApp) GetDeveloper()(*string) {
    return m.developer
}
// GetDisplayName gets the displayName property value. The admin provided or imported title of the app.
func (m *MobileApp) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppCategoryable)
            }
            m.SetCategories(res)
        }
        return nil
    }
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["developer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeveloper(val)
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
    res["informationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInformationUrl(val)
        }
        return nil
    }
    res["isFeatured"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFeatured(val)
        }
        return nil
    }
    res["largeIcon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLargeIcon(val.(MimeContentable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["privacyInformationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyInformationUrl(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["publishingState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppPublishingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingState(val.(*MobileAppPublishingState))
        }
        return nil
    }
    return res
}
// GetInformationUrl gets the informationUrl property value. The more information Url.
func (m *MobileApp) GetInformationUrl()(*string) {
    return m.informationUrl
}
// GetIsFeatured gets the isFeatured property value. The value indicating whether the app is marked as featured by the admin.
func (m *MobileApp) GetIsFeatured()(*bool) {
    return m.isFeatured
}
// GetLargeIcon gets the largeIcon property value. The large icon, to be displayed in the app details and used for upload of the icon.
func (m *MobileApp) GetLargeIcon()(MimeContentable) {
    return m.largeIcon
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the app was last modified.
func (m *MobileApp) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNotes gets the notes property value. Notes for the app.
func (m *MobileApp) GetNotes()(*string) {
    return m.notes
}
// GetOwner gets the owner property value. The owner of the app.
func (m *MobileApp) GetOwner()(*string) {
    return m.owner
}
// GetPrivacyInformationUrl gets the privacyInformationUrl property value. The privacy statement Url.
func (m *MobileApp) GetPrivacyInformationUrl()(*string) {
    return m.privacyInformationUrl
}
// GetPublisher gets the publisher property value. The publisher of the app.
func (m *MobileApp) GetPublisher()(*string) {
    return m.publisher
}
// GetPublishingState gets the publishingState property value. Indicates the publishing state of an app.
func (m *MobileApp) GetPublishingState()(*MobileAppPublishingState) {
    return m.publishingState
}
// Serialize serializes information the current object
func (m *MobileApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("developer", m.GetDeveloper())
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
        err = writer.WriteStringValue("informationUrl", m.GetInformationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFeatured", m.GetIsFeatured())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("largeIcon", m.GetLargeIcon())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("privacyInformationUrl", m.GetPrivacyInformationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    if m.GetPublishingState() != nil {
        cast := (*m.GetPublishingState()).String()
        err = writer.WriteStringValue("publishingState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments for this mobile app.
func (m *MobileApp) SetAssignments(value []MobileAppAssignmentable)() {
    m.assignments = value
}
// SetCategories sets the categories property value. The list of categories for this app.
func (m *MobileApp) SetCategories(value []MobileAppCategoryable)() {
    m.categories = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the app was created.
func (m *MobileApp) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description of the app.
func (m *MobileApp) SetDescription(value *string)() {
    m.description = value
}
// SetDeveloper sets the developer property value. The developer of the app.
func (m *MobileApp) SetDeveloper(value *string)() {
    m.developer = value
}
// SetDisplayName sets the displayName property value. The admin provided or imported title of the app.
func (m *MobileApp) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetInformationUrl sets the informationUrl property value. The more information Url.
func (m *MobileApp) SetInformationUrl(value *string)() {
    m.informationUrl = value
}
// SetIsFeatured sets the isFeatured property value. The value indicating whether the app is marked as featured by the admin.
func (m *MobileApp) SetIsFeatured(value *bool)() {
    m.isFeatured = value
}
// SetLargeIcon sets the largeIcon property value. The large icon, to be displayed in the app details and used for upload of the icon.
func (m *MobileApp) SetLargeIcon(value MimeContentable)() {
    m.largeIcon = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the app was last modified.
func (m *MobileApp) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNotes sets the notes property value. Notes for the app.
func (m *MobileApp) SetNotes(value *string)() {
    m.notes = value
}
// SetOwner sets the owner property value. The owner of the app.
func (m *MobileApp) SetOwner(value *string)() {
    m.owner = value
}
// SetPrivacyInformationUrl sets the privacyInformationUrl property value. The privacy statement Url.
func (m *MobileApp) SetPrivacyInformationUrl(value *string)() {
    m.privacyInformationUrl = value
}
// SetPublisher sets the publisher property value. The publisher of the app.
func (m *MobileApp) SetPublisher(value *string)() {
    m.publisher = value
}
// SetPublishingState sets the publishingState property value. Indicates the publishing state of an app.
func (m *MobileApp) SetPublishingState(value *MobileAppPublishingState)() {
    m.publishingState = value
}

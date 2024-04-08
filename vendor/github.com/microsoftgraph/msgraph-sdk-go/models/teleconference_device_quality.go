package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeleconferenceDeviceQuality 
type TeleconferenceDeviceQuality struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A unique identifier for all  the participant calls in a conference or a unique identifier for two participant calls in P2P call. This needs to be copied over from Microsoft.Graph.Call.CallChainId.
    callChainId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // A geo-region where the service is deployed, such as ProdNoam.
    cloudServiceDeploymentEnvironment *string
    // A unique deployment identifier assigned by Azure.
    cloudServiceDeploymentId *string
    // The Azure deployed cloud service instance name, such as FrontEnd_IN_3.
    cloudServiceInstanceName *string
    // The Azure deployed cloud service name, such as contoso.cloudapp.net.
    cloudServiceName *string
    // Any additional description, such as VTC Bldg 30/21.
    deviceDescription *string
    // The user media agent name, such as Cisco SX80.
    deviceName *string
    // A unique identifier for a specific media leg of a participant in a conference.  One participant can have multiple media leg identifiers if retargeting happens. CVI partner assigns this value.
    mediaLegId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The list of media qualities in a media session (call), such as audio quality, video quality, and/or screen sharing quality.
    mediaQualityList []TeleconferenceDeviceMediaQualityable
    // The OdataType property
    odataType *string
    // A unique identifier for a specific participant in a conference. The CVI partner needs to copy over Call.MyParticipantId to this property.
    participantId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewTeleconferenceDeviceQuality instantiates a new teleconferenceDeviceQuality and sets the default values.
func NewTeleconferenceDeviceQuality()(*TeleconferenceDeviceQuality) {
    m := &TeleconferenceDeviceQuality{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeleconferenceDeviceQualityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeleconferenceDeviceQualityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeleconferenceDeviceQuality(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeleconferenceDeviceQuality) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCallChainId gets the callChainId property value. A unique identifier for all  the participant calls in a conference or a unique identifier for two participant calls in P2P call. This needs to be copied over from Microsoft.Graph.Call.CallChainId.
func (m *TeleconferenceDeviceQuality) GetCallChainId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.callChainId
}
// GetCloudServiceDeploymentEnvironment gets the cloudServiceDeploymentEnvironment property value. A geo-region where the service is deployed, such as ProdNoam.
func (m *TeleconferenceDeviceQuality) GetCloudServiceDeploymentEnvironment()(*string) {
    return m.cloudServiceDeploymentEnvironment
}
// GetCloudServiceDeploymentId gets the cloudServiceDeploymentId property value. A unique deployment identifier assigned by Azure.
func (m *TeleconferenceDeviceQuality) GetCloudServiceDeploymentId()(*string) {
    return m.cloudServiceDeploymentId
}
// GetCloudServiceInstanceName gets the cloudServiceInstanceName property value. The Azure deployed cloud service instance name, such as FrontEnd_IN_3.
func (m *TeleconferenceDeviceQuality) GetCloudServiceInstanceName()(*string) {
    return m.cloudServiceInstanceName
}
// GetCloudServiceName gets the cloudServiceName property value. The Azure deployed cloud service name, such as contoso.cloudapp.net.
func (m *TeleconferenceDeviceQuality) GetCloudServiceName()(*string) {
    return m.cloudServiceName
}
// GetDeviceDescription gets the deviceDescription property value. Any additional description, such as VTC Bldg 30/21.
func (m *TeleconferenceDeviceQuality) GetDeviceDescription()(*string) {
    return m.deviceDescription
}
// GetDeviceName gets the deviceName property value. The user media agent name, such as Cisco SX80.
func (m *TeleconferenceDeviceQuality) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeleconferenceDeviceQuality) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["callChainId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallChainId(val)
        }
        return nil
    }
    res["cloudServiceDeploymentEnvironment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudServiceDeploymentEnvironment(val)
        }
        return nil
    }
    res["cloudServiceDeploymentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudServiceDeploymentId(val)
        }
        return nil
    }
    res["cloudServiceInstanceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudServiceInstanceName(val)
        }
        return nil
    }
    res["cloudServiceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudServiceName(val)
        }
        return nil
    }
    res["deviceDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDescription(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["mediaLegId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaLegId(val)
        }
        return nil
    }
    res["mediaQualityList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeleconferenceDeviceMediaQualityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeleconferenceDeviceMediaQualityable, len(val))
            for i, v := range val {
                res[i] = v.(TeleconferenceDeviceMediaQualityable)
            }
            m.SetMediaQualityList(res)
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
    res["participantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipantId(val)
        }
        return nil
    }
    return res
}
// GetMediaLegId gets the mediaLegId property value. A unique identifier for a specific media leg of a participant in a conference.  One participant can have multiple media leg identifiers if retargeting happens. CVI partner assigns this value.
func (m *TeleconferenceDeviceQuality) GetMediaLegId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.mediaLegId
}
// GetMediaQualityList gets the mediaQualityList property value. The list of media qualities in a media session (call), such as audio quality, video quality, and/or screen sharing quality.
func (m *TeleconferenceDeviceQuality) GetMediaQualityList()([]TeleconferenceDeviceMediaQualityable) {
    return m.mediaQualityList
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeleconferenceDeviceQuality) GetOdataType()(*string) {
    return m.odataType
}
// GetParticipantId gets the participantId property value. A unique identifier for a specific participant in a conference. The CVI partner needs to copy over Call.MyParticipantId to this property.
func (m *TeleconferenceDeviceQuality) GetParticipantId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.participantId
}
// Serialize serializes information the current object
func (m *TeleconferenceDeviceQuality) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteUUIDValue("callChainId", m.GetCallChainId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudServiceDeploymentEnvironment", m.GetCloudServiceDeploymentEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudServiceDeploymentId", m.GetCloudServiceDeploymentId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudServiceInstanceName", m.GetCloudServiceInstanceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudServiceName", m.GetCloudServiceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceDescription", m.GetDeviceDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("mediaLegId", m.GetMediaLegId())
        if err != nil {
            return err
        }
    }
    if m.GetMediaQualityList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMediaQualityList()))
        for i, v := range m.GetMediaQualityList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("mediaQualityList", cast)
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
        err := writer.WriteUUIDValue("participantId", m.GetParticipantId())
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
func (m *TeleconferenceDeviceQuality) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCallChainId sets the callChainId property value. A unique identifier for all  the participant calls in a conference or a unique identifier for two participant calls in P2P call. This needs to be copied over from Microsoft.Graph.Call.CallChainId.
func (m *TeleconferenceDeviceQuality) SetCallChainId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.callChainId = value
}
// SetCloudServiceDeploymentEnvironment sets the cloudServiceDeploymentEnvironment property value. A geo-region where the service is deployed, such as ProdNoam.
func (m *TeleconferenceDeviceQuality) SetCloudServiceDeploymentEnvironment(value *string)() {
    m.cloudServiceDeploymentEnvironment = value
}
// SetCloudServiceDeploymentId sets the cloudServiceDeploymentId property value. A unique deployment identifier assigned by Azure.
func (m *TeleconferenceDeviceQuality) SetCloudServiceDeploymentId(value *string)() {
    m.cloudServiceDeploymentId = value
}
// SetCloudServiceInstanceName sets the cloudServiceInstanceName property value. The Azure deployed cloud service instance name, such as FrontEnd_IN_3.
func (m *TeleconferenceDeviceQuality) SetCloudServiceInstanceName(value *string)() {
    m.cloudServiceInstanceName = value
}
// SetCloudServiceName sets the cloudServiceName property value. The Azure deployed cloud service name, such as contoso.cloudapp.net.
func (m *TeleconferenceDeviceQuality) SetCloudServiceName(value *string)() {
    m.cloudServiceName = value
}
// SetDeviceDescription sets the deviceDescription property value. Any additional description, such as VTC Bldg 30/21.
func (m *TeleconferenceDeviceQuality) SetDeviceDescription(value *string)() {
    m.deviceDescription = value
}
// SetDeviceName sets the deviceName property value. The user media agent name, such as Cisco SX80.
func (m *TeleconferenceDeviceQuality) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetMediaLegId sets the mediaLegId property value. A unique identifier for a specific media leg of a participant in a conference.  One participant can have multiple media leg identifiers if retargeting happens. CVI partner assigns this value.
func (m *TeleconferenceDeviceQuality) SetMediaLegId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.mediaLegId = value
}
// SetMediaQualityList sets the mediaQualityList property value. The list of media qualities in a media session (call), such as audio quality, video quality, and/or screen sharing quality.
func (m *TeleconferenceDeviceQuality) SetMediaQualityList(value []TeleconferenceDeviceMediaQualityable)() {
    m.mediaQualityList = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeleconferenceDeviceQuality) SetOdataType(value *string)() {
    m.odataType = value
}
// SetParticipantId sets the participantId property value. A unique identifier for a specific participant in a conference. The CVI partner needs to copy over Call.MyParticipantId to this property.
func (m *TeleconferenceDeviceQuality) SetParticipantId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.participantId = value
}

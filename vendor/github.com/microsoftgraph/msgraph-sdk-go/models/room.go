package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Room 
type Room struct {
    Place
    // Specifies the name of the audio device in the room.
    audioDeviceName *string
    // Type of room. Possible values are standard, and reserved.
    bookingType *BookingType
    // Specifies the building name or building number that the room is in.
    building *string
    // Specifies the capacity of the room.
    capacity *int32
    // Specifies the name of the display device in the room.
    displayDeviceName *string
    // Email address of the room.
    emailAddress *string
    // Specifies a descriptive label for the floor, for example, P.
    floorLabel *string
    // Specifies the floor number that the room is on.
    floorNumber *int32
    // Specifies whether the room is wheelchair accessible.
    isWheelChairAccessible *bool
    // Specifies a descriptive label for the room, for example, a number or name.
    label *string
    // Specifies a nickname for the room, for example, 'conf room'.
    nickname *string
    // Specifies additional features of the room, for example, details like the type of view or furniture type.
    tags []string
    // Specifies the name of the video device in the room.
    videoDeviceName *string
}
// NewRoom instantiates a new Room and sets the default values.
func NewRoom()(*Room) {
    m := &Room{
        Place: *NewPlace(),
    }
    odataTypeValue := "#microsoft.graph.room"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRoomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoom(), nil
}
// GetAudioDeviceName gets the audioDeviceName property value. Specifies the name of the audio device in the room.
func (m *Room) GetAudioDeviceName()(*string) {
    return m.audioDeviceName
}
// GetBookingType gets the bookingType property value. Type of room. Possible values are standard, and reserved.
func (m *Room) GetBookingType()(*BookingType) {
    return m.bookingType
}
// GetBuilding gets the building property value. Specifies the building name or building number that the room is in.
func (m *Room) GetBuilding()(*string) {
    return m.building
}
// GetCapacity gets the capacity property value. Specifies the capacity of the room.
func (m *Room) GetCapacity()(*int32) {
    return m.capacity
}
// GetDisplayDeviceName gets the displayDeviceName property value. Specifies the name of the display device in the room.
func (m *Room) GetDisplayDeviceName()(*string) {
    return m.displayDeviceName
}
// GetEmailAddress gets the emailAddress property value. Email address of the room.
func (m *Room) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Room) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Place.GetFieldDeserializers()
    res["audioDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioDeviceName(val)
        }
        return nil
    }
    res["bookingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBookingType(val.(*BookingType))
        }
        return nil
    }
    res["building"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuilding(val)
        }
        return nil
    }
    res["capacity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacity(val)
        }
        return nil
    }
    res["displayDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayDeviceName(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["floorLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloorLabel(val)
        }
        return nil
    }
    res["floorNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloorNumber(val)
        }
        return nil
    }
    res["isWheelChairAccessible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsWheelChairAccessible(val)
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["nickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNickname(val)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    res["videoDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoDeviceName(val)
        }
        return nil
    }
    return res
}
// GetFloorLabel gets the floorLabel property value. Specifies a descriptive label for the floor, for example, P.
func (m *Room) GetFloorLabel()(*string) {
    return m.floorLabel
}
// GetFloorNumber gets the floorNumber property value. Specifies the floor number that the room is on.
func (m *Room) GetFloorNumber()(*int32) {
    return m.floorNumber
}
// GetIsWheelChairAccessible gets the isWheelChairAccessible property value. Specifies whether the room is wheelchair accessible.
func (m *Room) GetIsWheelChairAccessible()(*bool) {
    return m.isWheelChairAccessible
}
// GetLabel gets the label property value. Specifies a descriptive label for the room, for example, a number or name.
func (m *Room) GetLabel()(*string) {
    return m.label
}
// GetNickname gets the nickname property value. Specifies a nickname for the room, for example, 'conf room'.
func (m *Room) GetNickname()(*string) {
    return m.nickname
}
// GetTags gets the tags property value. Specifies additional features of the room, for example, details like the type of view or furniture type.
func (m *Room) GetTags()([]string) {
    return m.tags
}
// GetVideoDeviceName gets the videoDeviceName property value. Specifies the name of the video device in the room.
func (m *Room) GetVideoDeviceName()(*string) {
    return m.videoDeviceName
}
// Serialize serializes information the current object
func (m *Room) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Place.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("audioDeviceName", m.GetAudioDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetBookingType() != nil {
        cast := (*m.GetBookingType()).String()
        err = writer.WriteStringValue("bookingType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("building", m.GetBuilding())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("capacity", m.GetCapacity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayDeviceName", m.GetDisplayDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("floorLabel", m.GetFloorLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("floorNumber", m.GetFloorNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isWheelChairAccessible", m.GetIsWheelChairAccessible())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nickname", m.GetNickname())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("videoDeviceName", m.GetVideoDeviceName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudioDeviceName sets the audioDeviceName property value. Specifies the name of the audio device in the room.
func (m *Room) SetAudioDeviceName(value *string)() {
    m.audioDeviceName = value
}
// SetBookingType sets the bookingType property value. Type of room. Possible values are standard, and reserved.
func (m *Room) SetBookingType(value *BookingType)() {
    m.bookingType = value
}
// SetBuilding sets the building property value. Specifies the building name or building number that the room is in.
func (m *Room) SetBuilding(value *string)() {
    m.building = value
}
// SetCapacity sets the capacity property value. Specifies the capacity of the room.
func (m *Room) SetCapacity(value *int32)() {
    m.capacity = value
}
// SetDisplayDeviceName sets the displayDeviceName property value. Specifies the name of the display device in the room.
func (m *Room) SetDisplayDeviceName(value *string)() {
    m.displayDeviceName = value
}
// SetEmailAddress sets the emailAddress property value. Email address of the room.
func (m *Room) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// SetFloorLabel sets the floorLabel property value. Specifies a descriptive label for the floor, for example, P.
func (m *Room) SetFloorLabel(value *string)() {
    m.floorLabel = value
}
// SetFloorNumber sets the floorNumber property value. Specifies the floor number that the room is on.
func (m *Room) SetFloorNumber(value *int32)() {
    m.floorNumber = value
}
// SetIsWheelChairAccessible sets the isWheelChairAccessible property value. Specifies whether the room is wheelchair accessible.
func (m *Room) SetIsWheelChairAccessible(value *bool)() {
    m.isWheelChairAccessible = value
}
// SetLabel sets the label property value. Specifies a descriptive label for the room, for example, a number or name.
func (m *Room) SetLabel(value *string)() {
    m.label = value
}
// SetNickname sets the nickname property value. Specifies a nickname for the room, for example, 'conf room'.
func (m *Room) SetNickname(value *string)() {
    m.nickname = value
}
// SetTags sets the tags property value. Specifies additional features of the room, for example, details like the type of view or furniture type.
func (m *Room) SetTags(value []string)() {
    m.tags = value
}
// SetVideoDeviceName sets the videoDeviceName property value. Specifies the name of the video device in the room.
func (m *Room) SetVideoDeviceName(value *string)() {
    m.videoDeviceName = value
}

package shakearound

import (
    "github.com/chanxuehong/wechat/mp"
)

//指定的设备信息
type Devices struct {
    DeviceId int `json:"device_id"`   //设备编号
    Uuid string `json:"uuid"`           //UUID
    Major int `json:"major"`          //major
    Minor int `json:"minor"`          //minor
    Status int `json:"status"`          //激活状态，0：未激活，1：已激活（但不活跃），2：活跃
    PoiId int `json:"poi_id"`         //设备关联的门店ID
    Comment string `json:"comment"`     //设备的备注信息
    PageIds string `json:"page_ids"`    //与此设备关联的页面ID列表
}

//申请设备ID
//  quantity:       申请的设备ID的数量，单次新增设备超过500个，需走人工审核流程
//  applyReason:    申请理由，不超过100个字
//  comment:        备注，不超过15个汉字或30个英文字母
//  poiId:          设备关联的门店ID
func (clt Client) ApplyDeviceId(quantity int, applyReason, comment, poiId string) (deviceses *[]Devices, applyId int, err error) {
    var request = struct {
        Quantity   int `json:"quantity"`
        Apply_reason string `json:"apply_reason"`
        Comment string `json:"comment"`
        PoiId int `json:"poi_id"`
    }{
        Quantity:   quantity,
        Apply_reason: applyReason,
        Comment: comment,
        PoiId: poiId,
    }

    var result struct {
        mp.Error
        DeviceIdentifiers []Devices `json:"device_identifiers"`
        AuditStatus int `json:"audit_status"`
        AuditComment string `json:"audit_comment"`
        ApplyId int `json:"apply_id"`
    }

    incompleteURL := "https://api.weixin.qq.com/shakearound/device/applyid?access_token="
    if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
        return
    }

    if result.ErrCode != mp.ErrCodeOK {
        err = &result.Error
        return
    }

    if len(result.DeviceIdentifiers) == 0 {
        err = mp.Error{
            ErrCode: result.AuditStatus,
            ErrMsg: result.AuditComment,
        }
    }

    deviceses = &result.DeviceIdentifiers
    applyId = result.ApplyId
    return
}

//编辑设备信息（使用device_id）
//  deviceId:   设备编号
//  comment:    设备的备注信息，不超过15个汉字或30个英文字母。
func (clt Client) UpdateDeviceByDeviceId(deviceId int, comment string) (err error) {
    err = clt.UpdateDevice(deviceId, "", 0, 0, comment)
    return
}

//编辑设备信息（使用UUID、major、minor）
//  uuid:   uuid
//  major:  major
//  minor: minor
//  comment:    设备的备注信息，不超过15个汉字或30个英文字母。
func (clt Client) UpdateDeviceByUuid(uuid string, major, minor int, comment string) (err error) {
    err = clt.UpdateDevice(0, uuid, major, minor, comment)
    return
}

//编辑设备信息
func (clt Client) UpdateDevice(deviceId int, uuid string, major, minor int, comment string)(err error){
    type DeviceIdentifier struct{
        DeviceId int `json:""device_id,omitempty`   //设备编号
        Uuid string `json:"uuid,omitempty"`         //UUID
        Major int `json:"major"`                    //major
        Minor int `json:"minor"`                    //minor
    }
    var deviceIdentifier = DeviceIdentifier{
        DeviceId: deviceId,
        Uuid: uuid,
        Major: major,
        Minor: minor,
    }

    var request = struct {
        DeviceIdentifier   DeviceIdentifier `json:"device_identifier"`
        Comment string `json:"comment"`
    }{
        DeviceIdentifier: deviceIdentifier,
        Comment: comment,
    }

    var result struct {
        mp.Error
    }

    incompleteURL := "https://api.weixin.qq.com/shakearound/device/update?access_token="
    if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
        return
    }

    if result.ErrCode != mp.ErrCodeOK {
        err = &result.Error
        return
    }
    return
}

//配置设备与门店的关联关系（使用device_id）
//  deviceId:   设备编号
//  comment:    设备关联的门店ID
func (clt Client) DeviceBindLocationByDeviceId(deviceId, poiId int) (err error) {
    err = clt.DeviceBindLocation(deviceId, "", 0, 0, poiId)
    return
}

//配置设备与门店的关联关系（使用UUID、major、minor）
//  uuid:   uuid
//  major:  major
//  minor: minor
//  comment:    设备关联的门店ID
func (clt Client) DeviceBindLocationByUuid(uuid string, major, minor, poiId int) (err error) {
    err = clt.DeviceBindLocation(0, uuid, major, minor, poiId)
    return
}

//配置设备与门店的关联关系
func (clt Client) DeviceBindLocation(deviceId int, uuid string, major, minor, poiId int) (err error) {
    type DeviceIdentifier struct{
        DeviceId int `json:""device_id,omitempty`   //设备编号
        Uuid string `json:"uuid,omitempty"`         //UUID
        Major int `json:"major"`                    //major
        Minor int `json:"minor"`                    //minor
    }
    var deviceIdentifier = DeviceIdentifier{
        DeviceId: deviceId,
        Uuid: uuid,
        Major: major,
        Minor: minor,
    }

    var request = struct {
        DeviceIdentifier   DeviceIdentifier `json:"device_identifier"`
        PoiId string `json:"poi_id"`
    }{
        DeviceIdentifier: deviceIdentifier,
        PoiId: poiId,
    }

    var result struct {
        mp.Error
    }

    incompleteURL := "https://api.weixin.qq.com/shakearound/device/bindlocation?access_token="
    if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
        return
    }

    if result.ErrCode != mp.ErrCodeOK {
        err = &result.Error
        return
    }
    return
}

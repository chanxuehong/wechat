package shakearound

import (
    "github.com/chanxuehong/wechat/mp"
)


type Devices struct {
    DeviceId int `json:"device_id"`
    Uuid string `json:"uuid"`
    Major int `json:"major"`
    Minor int `json:"minor"`
    Status int `json:"status"`
    PoiId int `json:"poi_id"`
    Comment string `json:"comment"`
    PageIds string `json:"page_ids"`
}

func (clt Client) ApplyDeviceId(quantity int, applyReason, comment string, poiId int) (deviceses *[]Devices, applyId int, err error) {
    var request = struct {
        Quantity   int `json:"quantity"`
        Apply_reason string `json:"apply_reason"`
        Comment string `json:"comment"`
        PoiId int `json:"poi_id,omitempty"`
    }{
        Quantity:   quantity,
        Apply_reason: applyReason,
        Comment: comment,
        PoiId: poiId,
    }

    var result struct {
        mp.Error
		Data struct {
			DeviceIdentifiers []Devices `json:"device_identifiers"`
			AuditStatus int `json:"audit_status"`
			AuditComment string `json:"audit_comment"`
			ApplyId int `json:"apply_id"`
		} `json:"data"`
    }

    incompleteURL := "https://api.weixin.qq.com/shakearound/device/applyid?access_token="
    if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
        return
    }

    if result.ErrCode != mp.ErrCodeOK {
        err = &result.Error
        return
    }

    if len(result.Data.DeviceIdentifiers) == 0 {
        err = &(mp.Error{
            ErrCode: result.Data.AuditStatus,
            ErrMsg: result.Data.AuditComment,
        })
    }

    deviceses = &result.Data.DeviceIdentifiers
    applyId = result.Data.ApplyId
    return
}


func (clt Client) UpdateDeviceByDeviceId(deviceId int, comment string) (err error) {
    err = clt.UpdateDevice(deviceId, "", 0, 0, comment)
    return
}


func (clt Client) UpdateDeviceByUuid(uuid string, major, minor int, comment string) (err error) {
    err = clt.UpdateDevice(0, uuid, major, minor, comment)
    return
}

func (clt Client) UpdateDevice(deviceId int, uuid string, major, minor int, comment string)(err error){
    type deviceIdentifier struct {
        DeviceId int `json:""device_id,omitempty`
        Uuid string `json:"uuid,omitempty"`         //UUID
        Major int `json:"major"`                    //major
        Minor int `json:"minor"`                    //minor
    }
    var request = struct {
        DeviceIdentifier struct{
            DeviceId int `json:""device_id,omitempty`
            Uuid string `json:"uuid,omitempty"`         //UUID
            Major int `json:"major"`                    //major
            Minor int `json:"minor"`                    //minor
        } `json:"device_identifier"`
        Comment string `json:"comment"`
    }{
        DeviceIdentifier: deviceIdentifier{
            DeviceId: deviceId,
            Uuid: uuid,
            Major: major,
            Minor: minor,
        },
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


func (clt Client) DeviceBindLocationByDeviceId(deviceId, poiId int) (err error) {
    err = clt.DeviceBindLocation(deviceId, "", 0, 0, poiId)
    return
}


func (clt Client) DeviceBindLocationByUuid(uuid string, major, minor, poiId int) (err error) {
    err = clt.DeviceBindLocation(0, uuid, major, minor, poiId)
    return
}


func (clt Client) DeviceBindLocation(deviceId int, uuid string, major, minor, poiId int) (err error) {
    type deviceIdentifier struct {
        DeviceId int `json:""device_id,omitempty`
        Uuid string `json:"uuid,omitempty"`         //UUID
        Major int `json:"major"`                    //major
        Minor int `json:"minor"`                    //minor
    }
    var request = struct {
        DeviceIdentifier struct{
            DeviceId int `json:""device_id,omitempty`
            Uuid string `json:"uuid,omitempty"`         //UUID
            Major int `json:"major"`                    //major
            Minor int `json:"minor"`                    //minor
        } `json:"device_identifier"`
        PoiId int `json:"poi_id"`
    }{
        DeviceIdentifier: deviceIdentifier{
            DeviceId: deviceId,
            Uuid: uuid,
            Major: major,
            Minor: minor,
        },
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


func (clt Client)SeachDeviceByDeviceId(deviceId int)(deviceses *[]Devices, totalCount int, err error){
    type deviceIdentifier struct {
        DeviceId int `json:""device_id,omitempty`
    }
    var request = struct {
        DeviceIdentifier struct{
            DeviceId int `json:""device_id,omitempty`
        } `json:"device_identifier"`
    }{
        DeviceIdentifier: deviceIdentifier{
            DeviceId: deviceId,
        },
    }
    return clt.SeachDevice(request)
}


func (clt Client)SeachDeviceByUuid(uuid string, major, minor int)(deviceses *[]Devices, totalCount int, err error){
    type deviceIdentifier struct {
        Uuid string `json:"uuid"`         //UUID
        Major int `json:"major"`                    //major
        Minor int `json:"minor"`                    //minor
    }
    var request = struct {
        DeviceIdentifier struct{
            Uuid string `json:"uuid"`         //UUID
            Major int `json:"major"`                    //major
            Minor int `json:"minor"`                    //minor
        } `json:"device_identifier"`
    }{
        DeviceIdentifier: deviceIdentifier{
            Uuid: uuid,
            Major: major,
            Minor: minor,
        },
    }
    return clt.SeachDevice(request)
}


func (clt Client)SeachDeviceByCount(begin, count, applyId int)(deviceses *[]Devices, totalCount int, err error){
    var request = struct {
        ApplyId int `json:"apply_id,omitempty"`
        Begin int `json:"begin"`
        Count int `json:count`
    }{
        ApplyId:applyId,
        Begin:begin,
        Count:count,
    }
    return clt.SeachDevice(request)
}


func (clt Client)SeachDevice(v interface{}) (deviceses *[]Devices, totalCount int, err error) {
    var result struct {
        mp.Error
        Data struct{
            Devices []Devices `json:"devices"`
            TotalCount int `json:"total_count"`
        } `json:"data"`
    }

    incompleteURL := "https://api.weixin.qq.com/shakearound/device/search?access_token="
    if err = clt.PostJSON(incompleteURL, &v, &result); err != nil {
        return
    }

    if result.ErrCode != mp.ErrCodeOK {
        err = &result.Error
        return
    }

    deviceses = &result.Data.Devices
    totalCount = result.Data.TotalCount
    return
}


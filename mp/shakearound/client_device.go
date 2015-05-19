// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com) Harry Rong(harrykobe@gmail.com)
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

//  申请设备ID
//  quantity:       申请的设备ID的数量，单次新增设备超过500个，需走人工审核流程
//  applyReason:    申请理由，不超过100个字
//  comment:        备注，不超过15个汉字或30个英文字母
//  poiId:          设备关联的门店ID
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

//  编辑设备信息
//  deviceId:   设备编号
//  comment:    设备的备注信息，不超过15个汉字或30个英文字母
func (clt Client) UpdateDeviceByDeviceId(deviceId int, comment string) (err error) {
    err = clt.UpdateDevice(deviceId, "", 0, 0, comment)
    return
}

//  编辑设备信
//  uuid:       UUID
//  major:      major
//  minor:      minor
//  comment:    设备的备注信息，不超过15个汉字或30个英文字母
func (clt Client) UpdateDeviceByUuid(uuid string, major, minor int, comment string) (err error) {
    err = clt.UpdateDevice(0, uuid, major, minor, comment)
    return
}

//  编辑设备信
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

//  配置设备与门店的关联关系
//  deviceId:   设备编号
//  poiId:      设备关联的门店ID
func (clt Client) DeviceBindLocationByDeviceId(deviceId, poiId int) (err error) {
    err = clt.DeviceBindLocation(deviceId, "", 0, 0, poiId)
    return
}

//  配置设备与门店的关联关系
//  uuid:       UUID
//  major:      major
//  minor:      minor
//  poiId:      设备关联的门店ID
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

//  查询一个设备
//  deviceId:   设备编号
func (clt Client)SeachDeviceByDeviceId(deviceId int)(devices *Devices, totalCount int, err error){
    type DeviceIdentifier struct {
        DeviceId int `json:"device_id"`
    }
    var request = struct {
        DeviceIdentifier []DeviceIdentifier  `json:"device_identifiers"`
    }{
        DeviceIdentifier: []DeviceIdentifier{
            DeviceIdentifier{
                DeviceId: deviceId,
            },
        },
    }
    deviceses, totalCount, err := clt.SeachDevice(request)
    if err != nil{
        return
    }
    devices = &(*deviceses)[0]
    return
}

//  查询一个设备
//  uuid:       UUID
//  major:      major
//  minor:      minor
func (clt Client)SeachDeviceByUuid(uuid string, major, minor int)(devices *Devices, totalCount int, err error){
    type deviceIdentifier struct {
        Uuid string `json:"uuid"`                   //UUID
        Major int `json:"major"`                    //major
        Minor int `json:"minor"`                    //minor
    }
    var request = struct {
        DeviceIdentifier []deviceIdentifier  `json:"device_identifiers"`
    }{
        DeviceIdentifier: []deviceIdentifier{
            deviceIdentifier{
                Uuid: uuid,
                Major: major,
                Minor: minor,
            },
        },
    }
    deviceses, totalCount, err := clt.SeachDevice(request)
    if err != nil{
        return
    }
    devices = &(*deviceses)[0]
    return
}

//  查询设备列表
//  deviceses   设备列表
func (clt Client)SeachDeviceByDevices(deviceses *[]Devices)(devices *[]Devices, totalCount int, err error){
    var request = struct {
        DeviceIdentifier *[]Devices  `json:"device_identifiers"`
    }{
        DeviceIdentifier: deviceses,
    }
    return clt.SeachDevice(request)
}

//  查询设备列表
//  begin:          设备列表的起始索引值
//  count:          待查询的设备个数
//  applyId:        批次ID，申请设备ID时所返回的批次ID
func (clt Client)SeachDeviceByCount(begin, count int, applyIds ...int)(deviceses *[]Devices, totalCount int, err error){
    var applyId int
    if len(applyIds) > 0 {
        applyId = applyIds[0]
    }
    var request = struct {
        ApplyId int `json:"apply_id,omitempty"`
        Begin int `json:"begin"`
        Count int `json:"count"`
    }{
        ApplyId:applyId,
        Begin:begin,
        Count:count,
    }
    return clt.SeachDevice(request)
}

//  查询设备列表
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

func (clt Client)DeviceBindPageByDeviceId(deviceId int, pageIds []int, append ...bool)(err error){
    var deviceIdentifier = struct{
        DeviceId int `json:"device_id"`
    }{
        DeviceId: deviceId,
    }
    var appendNum int = 0
    if len(append) > 0{
        if append[0] == true{
            appendNum = 1
        }
    }
    return clt.DeviceBindPage(deviceIdentifier, pageIds, 1, appendNum)
}


func (clt Client)DeviceBindPageByUuid(uuid string, major, minor int, pageIds []int, append ...bool)(err error){
    var deviceIdentifier = struct{
        Uuid string `json:"uuid"`                   //UUID
        Major int `json:"major"`                    //major
        Minor int `json:"minor"`                    //minor
    }{
        Uuid: uuid,
        Major: major,
        Minor: minor,
    }
    var appendNum int = 0
    if len(append) > 0{
        if append[0] == true{
            appendNum = 1
        }
    }
    return clt.DeviceBindPage(deviceIdentifier, pageIds, 1, appendNum)
}


func (clt Client)DeviceBindPage(v interface{}, pageIds []int, bind int, append int)(err error){
    var request = struct{
        DeviceIdentifier interface{}    `json:"device_identifier"`
        PageIds []int `json:"page_ids"`
        Bind int `json:"bind"`
        Append int  `json:"append"`
    }{
        DeviceIdentifier: v,
        PageIds: pageIds,
        Bind: bind,
        Append: append,
    }

    var result mp.Error
    incompleteURL := "https://api.weixin.qq.com/shakearound/device/bindpage?access_token="
    if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
        return
    }

    if result.ErrCode != mp.ErrCodeOK {
        err = &result.Error
        return
    }
    return
}
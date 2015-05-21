// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com) Harry Rong(harrykobe@gmail.com)
package shakearound

import (
    "github.com/chanxuehong/wechat/mp"
)

type DeviceBase struct {
    DeviceId int `json:"device_id"`     //设备编号
    Uuid string `json:"uuid"`           //UUID、major、minor
    Major int `json:"major"`
    Minor int `json:"minor"`
}

type Device struct {
    DeviceBase                          //设备基础信息
    Status int `json:"status"`          //激活状态
    PoiId int `json:"poi_id"`           //设备关联的门店ID
    Comment string `json:"comment"`     //设备的备注信息
    PageIds string `json:"page_ids"`    //与此设备关联的页面ID列表，用逗号隔开
}

//  申请设备ID
//  quantity:       申请的设备ID的数量，单次新增设备超过500个，需走人工审核流程
//  applyReason:    申请理由，不超过100个字
//  comment:        备注，不超过15个汉字或30个英文字母
//  poiId:          设备关联的门店ID
func (clt Client) ApplyDeviceId(quantity int, applyReason, comment string, poiIds ...int) (devices *[]Device, applyId int, err error) {
    var poiId int
    if len(poiIds) > 0  {
        poiId = poiIds[0]
    }
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
			DeviceIdentifiers []Device `json:"device_identifiers"`
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

    devices = &result.Data.DeviceIdentifiers
    applyId = result.Data.ApplyId
    return
}

//  编辑设备信息
//  deviceId:   设备编号
//  comment:    设备的备注信息，不超过15个汉字或30个英文字母
func (clt Client) UpdateDeviceByDeviceId(deviceId int, comment string) (err error) {
    var deviceBase = DeviceBase{
        DeviceId: deviceId,
    }
    err = clt.UpdateDevice(&deviceBase, comment)
    return
}

//  编辑设备信息
//  uuid:       UUID
//  major:      major
//  minor:      minor
//  comment:    设备的备注信息，不超过15个汉字或30个英文字母
func (clt Client) UpdateDeviceByUuid(uuid string, major, minor int, comment string) (err error) {
    var deviceBase = DeviceBase{
        Uuid: uuid,
        Major: major,
        Minor: minor,
    }
    err = clt.UpdateDevice(&deviceBase, comment)
    return
}

//  编辑设备信息
func (clt Client) UpdateDevice(deviceBase *DeviceBase, comment string)(err error){
    var request = struct {
        DeviceIdentifier *DeviceBase `json:"device_identifier"`
        Comment string `json:"comment"`
    }{
        DeviceIdentifier: deviceBase,
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
    var deviceBase = DeviceBase{
        DeviceId: deviceId,
    }
    err = clt.DeviceBindLocation(&deviceBase, poiId)
    return
}

//  配置设备与门店的关联关系
//  uuid:       UUID
//  major:      major
//  minor:      minor
//  poiId:      设备关联的门店ID
func (clt Client) DeviceBindLocationByUuid(uuid string, major, minor, poiId int) (err error) {
    var deviceBase = DeviceBase{
        Uuid: uuid,
        Major: major,
        Minor: minor,
    }
    err = clt.DeviceBindLocation(&deviceBase, poiId)
    return
}

//  配置设备与门店的关联关系
func (clt Client) DeviceBindLocation(deviceBase *DeviceBase, poiId int) (err error) {
    var request = struct {
        DeviceIdentifier *DeviceBase `json:"device_identifier"`
        PoiId int `json:"poi_id"`
    }{
        DeviceIdentifier: deviceBase,
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
func (clt Client)SearchDeviceByDeviceId(deviceId int)(device *Device, totalCount int, err error){
    var request = struct {
        DeviceIdentifier []DeviceBase  `json:"device_identifiers"`
    }{
        DeviceIdentifier: []DeviceBase{
            DeviceBase{
                DeviceId: deviceId,
            },
        },
    }
    devices, totalCount, err := clt.searchDevice(request)
    if err != nil{
        return
    }
    device = &(*devices)[0]
    return
}

//  查询一个设备
//  uuid:       UUID
//  major:      major
//  minor:      minor
func (clt Client)SearchDeviceByUuid(uuid string, major, minor int)(device *Device, totalCount int, err error){
    var request = struct {
        DeviceIdentifier []DeviceBase  `json:"device_identifiers"`
    }{
        DeviceIdentifier: []DeviceBase{
            DeviceBase{
                Uuid: uuid,
                Major: major,
                Minor: minor,
            },
        },
    }
    devices, totalCount, err := clt.searchDevice(request)
    if err != nil{
        return
    }
    device = &(*devices)[0]
    return
}

//  查询设备列表
//  deviceses   设备列表
func (clt Client)SearchDeviceByDevices(devicesIn *[]DeviceBase)(devices *[]Device, totalCount int, err error){
    var request = struct {
        DeviceIdentifier *[]DeviceBase  `json:"device_identifiers"`
    }{
        DeviceIdentifier: devicesIn,
    }
    return clt.searchDevice(request)
}

//  查询设备列表
//  begin:          设备列表的起始索引值
//  count:          待查询的设备个数
//  applyId:        批次ID，申请设备ID时所返回的批次ID
func (clt Client)SearchDeviceByCount(begin, count int, applyIds ...int)(devices *[]Device, totalCount int, err error){
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
    return clt.searchDevice(request)
}

//  查询设备列表
func (clt Client)searchDevice(v interface{}) (devices *[]Device, totalCount int, err error) {
    var result struct {
        mp.Error
        Data struct{
            Devices []Device `json:"devices"`
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

    devices = &result.Data.Devices
    totalCount = result.Data.TotalCount
    return
}




//  配置设备与页面的关联关系
//  绑定页面，使用device_id
func (clt Client)DeviceBindPageByDeviceId(deviceId int, pageIds []int, append ...bool)(err error){
    var deviceIdentifier = DeviceBase{
        DeviceId: deviceId,
    }
    var appendNum int = 0
    if len(append) > 0{
        if append[0] == true{
            appendNum = 1
        }
    }
    return clt.deviceBindPage(deviceIdentifier, pageIds, 1, appendNum)
}

//  配置设备与页面的关联关系
//  绑定页面，使用uuid
func (clt Client)DeviceBindPageByUuid(uuid string, major, minor int, pageIds []int, append ...bool)(err error){
    var deviceIdentifier = DeviceBase{
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
    return clt.deviceBindPage(deviceIdentifier, pageIds, 1, appendNum)
}

//  配置设备与页面的关联关系
//  绑定页面，使用Devices
func (clt Client)DeviceBindPageByDevices(device *DeviceBase, pageIds []int, append ...bool)(err error){
    var appendNum int = 0
    if len(append) > 0{
        if append[0] == true{
            appendNum = 1
        }
    }
    return clt.deviceBindPage(*device, pageIds, 1, appendNum)
}

//  配置设备与页面的关联关系
//  解绑页面，使用device_id
func (clt Client)DeviceUnbindPageByDeviceId(deviceId int, pageIds []int)(err error){
    var deviceIdentifier = DeviceBase{
        DeviceId: deviceId,
    }
    return clt.deviceBindPage(deviceIdentifier, pageIds, 0, 0)
}

//  配置设备与页面的关联关系
//  解绑页面，使用uuid
func (clt Client)DeviceUnbindPageByUuid(uuid string, major, minor int, pageIds []int)(err error){
    var deviceIdentifier = DeviceBase{
        Uuid: uuid,
        Major: major,
        Minor: minor,
    }
    return clt.deviceBindPage(deviceIdentifier, pageIds, 0, 0)
}

//  配置设备与页面的关联关系
//  解绑页面，使用Devices
func (clt Client)DeviceUnbindPageByDevices(device *DeviceBase, pageIds []int)(err error){
    return clt.deviceBindPage(*device, pageIds, 0, 0)
}

//  配置设备与页面的关联关系
func (clt Client)deviceBindPage(v interface{}, pageIds []int, bind int, append int)(err error){
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
        err = &result
        return
    }
    return
}
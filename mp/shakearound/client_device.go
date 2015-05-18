package shakearound

import (
    "github.com/chanxuehong/wechat/mp"
)

//ָ�����豸��Ϣ
type Devices struct {
    DeviceId int `json:"device_id"`   //�豸���
    Uuid string `json:"uuid"`           //UUID
    Major int `json:"major"`          //major
    Minor int `json:"minor"`          //minor
    Status int `json:"status"`          //����״̬��0��δ���1���Ѽ��������Ծ����2����Ծ
    PoiId int `json:"poi_id"`         //�豸�������ŵ�ID
    Comment string `json:"comment"`     //�豸�ı�ע��Ϣ
    PageIds string `json:"page_ids"`    //����豸������ҳ��ID�б�
}

//�����豸ID
//  quantity:       ������豸ID�����������������豸����500���������˹��������
//  applyReason:    �������ɣ�������100����
//  comment:        ��ע��������15�����ֻ�30��Ӣ����ĸ
//  poiId:          �豸�������ŵ�ID
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

//�༭�豸��Ϣ��ʹ��device_id��
//  deviceId:   �豸���
//  comment:    �豸�ı�ע��Ϣ��������15�����ֻ�30��Ӣ����ĸ��
func (clt Client) UpdateDeviceByDeviceId(deviceId int, comment string) (err error) {
    err = clt.UpdateDevice(deviceId, "", 0, 0, comment)
    return
}

//�༭�豸��Ϣ��ʹ��UUID��major��minor��
//  uuid:   uuid
//  major:  major
//  minor: minor
//  comment:    �豸�ı�ע��Ϣ��������15�����ֻ�30��Ӣ����ĸ��
func (clt Client) UpdateDeviceByUuid(uuid string, major, minor int, comment string) (err error) {
    err = clt.UpdateDevice(0, uuid, major, minor, comment)
    return
}

//�༭�豸��Ϣ
func (clt Client) UpdateDevice(deviceId int, uuid string, major, minor int, comment string)(err error){
    type DeviceIdentifier struct{
        DeviceId int `json:""device_id,omitempty`   //�豸���
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

//�����豸���ŵ�Ĺ�����ϵ��ʹ��device_id��
//  deviceId:   �豸���
//  comment:    �豸�������ŵ�ID
func (clt Client) DeviceBindLocationByDeviceId(deviceId, poiId int) (err error) {
    err = clt.DeviceBindLocation(deviceId, "", 0, 0, poiId)
    return
}

//�����豸���ŵ�Ĺ�����ϵ��ʹ��UUID��major��minor��
//  uuid:   uuid
//  major:  major
//  minor: minor
//  comment:    �豸�������ŵ�ID
func (clt Client) DeviceBindLocationByUuid(uuid string, major, minor, poiId int) (err error) {
    err = clt.DeviceBindLocation(0, uuid, major, minor, poiId)
    return
}

//�����豸���ŵ�Ĺ�����ϵ
func (clt Client) DeviceBindLocation(deviceId int, uuid string, major, minor, poiId int) (err error) {
    type DeviceIdentifier struct{
        DeviceId int `json:""device_id,omitempty`   //�豸���
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

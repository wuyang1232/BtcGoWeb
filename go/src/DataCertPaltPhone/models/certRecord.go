package models

import (
	"bytes"
	"encoding/gob"
)

/*
该结构体用于定义链上数据保存的信息
*/
type CertRecord struct {
	CertID   []byte //认证id，本质是一个md5值
	CertIdHex string
	CertHash []byte //认证文件hash值，本质是sha256值
	CertHashHex string
	CertName string //认证人的名称
	Phone    string //联系方式
	CertCard string //身份证号
	FileName string //认证文件的名称
	FileSize int64  //文件的大小
	CertTime int64  //认证的时间
	CertTimeFormat string
}

/*
序列化操作
*/

func (c CertRecord) Serialize() ([]byte, error) {
	buff := new(bytes.Buffer)//缓冲区
	err := gob.NewEncoder(buff).Encode(c)//encoder编码器
	return buff.Bytes(), err
}

/*
反序列化操作
*/
func DeserializeCertRecord(data []byte) (*CertRecord, error) {
	var certRecord *CertRecord
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&certRecord)//decoder解码
	return certRecord, err
}

package utils

import (
	"bytes"
	"encoding/binary"
)

/*
将一个int64转换为[]byte
*/
func Int64ToByte(num int64) ([]byte,error) {
	//Buffer：缓冲区。 增益
	buff := new(bytes.Buffer)//通过new实例化一个缓冲区
	//buff.Write() 通过一系列的write方法向缓冲区写入数据
	//buff.Bytes() 通过Bytes方法从缓冲区中获取数据
	/*
	两种排列方式：
	   大端位序排列：binary.BigEndian
	   小端位序排列：binary.LittleEndian
	*/
	err := binary.Write(buff,binary.BigEndian,num)
	if err != nil {
		return nil, err
	}
	//从缓冲区获取数据
	return buff.Bytes(),nil
}
/*
将字符串转换为【】byte
*/
func StringToBytes(data string) []byte {
	return []byte(data)
}

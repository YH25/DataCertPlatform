package utils

import (


	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

/*
对一个字符串数据进行MD5哈希计算
*/
func MD5Hashstring(data string) string {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(data))
	bytes := hashMd5.Sum(nil)
	return hex.EncodeToString(bytes)

}

/*
io:input output 输入和输出
*/
func MD5HashReader(reader io.Reader)(string, error)  {
	md5Hash := md5.New()
	readerBytes, err :=ioutil.ReadAll(reader)
	//fmt.Println("读取到的文件:",readerBytes)
	if err != nil {
		return "", err
	}
	md5Hash.Write(readerBytes)
	hashBytes := md5Hash.Sum(nil)
	return  hex.EncodeToString(hashBytes),nil
}


/*
读取io流当中的数据，并对数据进行hash计算，但会sha256 hash值
*/
func SHA256HashReader(reader io.Reader)(string,error)  {
	sha256Hash := sha256.New()
	readerBytes,err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	sha256Hash.Write(readerBytes)
	hashBytes := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}
/*
对区块数据进行SHA256哈希计算
*/
func SHA256HashBlock(bs []byte) []byte {
	//2、对转换后的[]byte字节切片输入write方法
	sha256Hash := sha256.New()
	sha256Hash.Write(bs)
	hash := sha256Hash.Sum(nil)
	return hash
}
package blockchain

import (
	"DataCertPlatform/utils"
	"bytes"
	"time"
)

/*
定义区块结构体，用于表示区块
*/
type Block struct {
	Height    int64  //区块的高度，第几个区块
	TimeStamp int64  //区块产生的时间戳
	PervHash  []byte //前一个区块的hash
	Data      []byte //数据字段
	Hash      []byte //当前区块的hash值
	Version   string //版本号
	Nonce      int64  //区块对应的nonce值
}
/*
创建一个新区快
*/
func NewBlock(height int64, pervHash []byte, data []byte)Block {
	block := Block{
		Height: height + 1,
		TimeStamp: time.Now().Unix(),
		PervHash: pervHash,
		Data: data,
		Version: "0x01",
	}
	//找nonce值，通过工作量证明算法计算寻找
	pow := NewPow(block)
	nonce := pow.Run()
	block.Nonce = nonce


	//1、将block结构体数据转换为[]byte类型
	heightBytes, _ := utils.Int64ToByte(block.Height)
	timeStampBytes, _ := utils.Int64ToByte(block.TimeStamp)
	versionBytes := utils.StringToBytes(block.Version)

    nonceBytes, _ := utils.Int64ToByte(block.Nonce)
	var blockBytes []byte
	//bytes.join拼接
	bytes.Join([][]byte{
	heightBytes,
	timeStampBytes,
	block.PervHash,
	block.Data,
	versionBytes,
	nonceBytes,
	},[]byte{})


		//调用hash计算，对区块进行sha256哈希值计算
	block.Hash = utils.SHA256HashBlock(blockBytes)

	//挖矿竞争，获得记账权
	

	return block


}

/*
创建创世区块
*/
func CerateGenesisBlock() Block {
	gennesisBlock := NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
	return gennesisBlock
}
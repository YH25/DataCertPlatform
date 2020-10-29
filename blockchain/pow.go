package blockchain

import (
	"DataCertPlatform/utils"
	"bytes"
	"math/big"
)

/*
工作量证明算法结构体
*/
type ProofOfwork struct {
	Target *big.Int
	Block Block //要找的nonce值对应的区块

}


/*
实例化一个Pow算法的实例
*/
func NewPow(block Block) ProofOfwork {
	t := big.NewInt(1)
	t = t.Lsh(t,255)
	pow := ProofOfwork{
		Target: t,
		Block: block,
	}

	return pow
}
/*
run方法用于寻找合适的nonce值
*/
func (p ProofOfwork) Run() int64{
	var nonce int64
	nonce = 0

	for {//不知道什么时候结束，使用无限循环

		block := p.Block
		heightBytes, _ := utils.Int64ToByte(block.Height)
		timeBytes, _ := utils.Int64ToByte(block.TimeStamp)
		versionBytes := utils.StringToBytes(block.Version)

		nonceBytes, _ := utils.Int64ToByte(nonce)
		var blockBytes []byte
		//已有区块信息和尝试的nonce值的拼接信息
		bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.PervHash,
			block.Data,
			versionBytes,
			nonceBytes,
		}, []byte{})

		//区块和尝试的nonce值拼接后得到的hash值
		blockHash := utils.SHA256HashBlock(blockBytes)

		target := p.Target     //目标值
		var hashBig *big.Int   //声明定义
		hashBig = new(big.Int) //分配内存空间，为变量分配地址
		hashBig = hashBig.SetBytes(blockHash)
		if hashBig.Cmp(target) == -1 {
			//停止寻找
			break
		}
		nonce++  //自增，继续寻找
	}
	//将找到的符合规则的nonce返回
	return nonce
}
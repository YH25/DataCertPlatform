package blockchain

import "math/big"

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

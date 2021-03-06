package blockchain

import (
	"DataCertPaltPhone/utils"
	"bytes"
	"math/big"
)

const DIFFICULTY = 10

/*
工作量证明算法结构体
*/
type ProofOfWork struct {
	Target *big.Int //系统的目标值
	Block  Block    //要找的nonce值对应的区块
}

//实例化一个pow算法的实例
func NewPow(block Block) ProofOfWork {
	t := big.NewInt(1)
	t = t.Lsh(t, 255-DIFFICULTY)
	pow := ProofOfWork{
		Target: t,
		Block:  block,
	}
	return pow
}

/*
run方法用于寻找合适的nonce值
*/
func (p ProofOfWork) Run() ([]byte, int64) {
	var nonce int64
	nonce = 0
	var blockHash []byte
	for { //使用无限循坏
		heigthBytes, _ := utils.Int64ToByte(p.Block.Height)
		timeBytes, _ := utils.Int64ToByte(p.Block.TimeStamp)
		versionBytes := utils.StringToByte(p.Block.Version)
		nonceBytes, _ := utils.Int64ToByte(nonce)
		//已有的区块信息和尝试的nonce值的拼接信息
		blockBytes := bytes.Join([][]byte{
			heigthBytes,
			timeBytes,
			p.Block.PerviousHash,
			p.Block.Data,
			versionBytes,
			nonceBytes,
		}, []byte{})
		//区块和尝试的nonce值拼接后得到的hash值
		blockHash = utils.Sha256HashBlock(blockBytes)
		target := p.Target     //目标值
		var hashBig *big.Int   //声明和定义
		hashBig = new(big.Int) //分配内存和地址，为变量分配地址
		hashBig = hashBig.SetBytes(blockHash)

		if hashBig.Cmp(target) == -1 {
			//停止寻找
			break
		}
		nonce++ //自增，继续寻找
		//fmt.Println("尝试的nonce值：", nonce)
	}
	return blockHash, nonce
}

package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type block struct {
	ver        int    //版本号
	prev_block string //父区块的哈希值
	mekl_root  string //该区块merkle树的哈希值
	time       string //时间戳
	bits       int    // 难度

}

func main() {

	prev_block := "000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506"

	mrkl_root := "000000000002d01c1fccc21636b607dfd930d31d01c3a62104612a1719011250"

	time := "1293623863"

	bits := 4

	block := block{1, prev_block, mrkl_root, time, bits}

	nc := getNonce(block)

	fmt.Println(nc)

}

func getNonce(blc block) int {

	nonce := 0
	//
	for !mining(getHashValueStr(nonce, blc), blc.bits) {

		fmt.Println(getHashValueStr(nonce, blc))

		nonce++

	}

	fmt.Println(getHashValueStr(nonce, blc))
	fmt.Println("出块成功！")

	return nonce

}

func getHashValueStr(nonce int, blc block) string {

	hv := strconv.Itoa(blc.ver) + blc.time + blc.prev_block +
		blc.mekl_root + strconv.Itoa(nonce)

	first := sha256.New()

	first.Write([]byte(hv))

	return fmt.Sprintf("%x", first.Sum(nil))

}

func mining(hashStr string, diff int) bool {

	var i int

	for i = 0; i < len(hashStr); i++ {

		if hashStr[i] != '0' {
			break
		}
	}
	return i >= diff

}

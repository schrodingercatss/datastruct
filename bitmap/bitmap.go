package bitmap

import (
	"bytes"
)

type BitMap []uint64

const (
	AddressBitsPerWord uint8  = 6  // uint64占64bit，也就是2^6
	WordsPerSize       uint64 = 64 // 单字为64bit
)

// 创建指定大小的BitSet
func NewBitSet(nbits int) *BitMap {
	wordLen := (nbits - 1) >> AddressBitsPerWord // 求出需要多少字来保存n bits数据
	temp := BitMap(make([]uint64, wordLen + 1, wordLen + 1))
	return &temp
}

// 将BitSet的某一位置为1
func (bitMap *BitMap) Set(bitIndex uint64) {
	wordIndex := bitMap.wordIndex(bitIndex)
	bitMap.expandTo(wordIndex)
	(*bitMap)[wordIndex] |= uint64(0x01) << (bitIndex % WordsPerSize)
}
// 将BitSet的某一位置为0
func (bitMap *BitMap) Clear(bitIndex uint64) {
	wordIndex := bitMap.wordIndex(bitIndex)
	bitMap.expandTo(wordIndex)
	(*bitMap)[wordIndex] &^= uint64(0x01) << (bitIndex % WordsPerSize)
}

// 获取BitSet指定位上的值
func (bitMap *BitMap) Get(bitIndex uint64) bool{
	wordIndex := bitMap.wordIndex(bitIndex)
	return (wordIndex < len(*bitMap)) && ((*bitMap)[wordIndex] & (uint64(0x01) << (bitIndex % WordsPerSize)) != 0)
}

// 以二进制的方式打印BitSet
func (bitMap *BitMap) String() string{
	var temp uint64
	strAppend := &bytes.Buffer{}
	for i := 0; i < len(*bitMap); i++ {
		temp = (*bitMap)[i]
		for j := 0; j < 64; j++ {
			if temp & (uint64(0x01) << uint64(j)) != 0 {
				strAppend.WriteString("1");
			} else {
				strAppend.WriteString("0");
			}
		}
	}
	return strAppend.String()
}


//定位位置
func (bitMap *BitMap) wordIndex(bitIndex uint64) int {
	return int(bitIndex >> AddressBitsPerWord)
}

//扩容:每次扩容两倍
func (bitMap *BitMap) expandTo(wordIndex int) {
	wordsRequired := wordIndex + 1
	if len(*bitMap) < wordsRequired {
		if wordsRequired < 2 * len(*bitMap) {
			wordsRequired = 2 * len(*bitMap)
		}
		newBitMap := make([]uint64, wordsRequired, wordsRequired)
		copy(newBitMap, *bitMap)
		*bitMap = newBitMap
	}
}



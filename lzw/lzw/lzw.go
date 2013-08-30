package lzw

import (
	"lzw/tst"
	"fmt"
)
const R int16 = 128
func Compress(str string) {
	var root *tst.TSTNode = nil
	var code int16 = -1
	//init the TSTree with the ASCII value
	var i int16
	for i = 0; i < R; i++ {
		str := string(i)
		str_slice := str[:]
		root = tst.InsertChar(root, str_slice, 0, i)
	}

	code = R
	slice := str[:]//待处理的文本
	for true {
		lcsLen := tst.GetLCSLength(root, slice, 0)//返回TST中与待处理文本slice匹配的最长串的长度
		sliceTmp := slice[:lcsLen]//在TST中找到的最长串
		lcsCode := tst.GetCode(root, sliceTmp, 0)//返回这个最长串对应的编码
		if lcsCode == -1 {
			panic("lcsCode == -1 : invalid character exists\n")
		}
		//fmt.Printf("%s : %x\n",sliceTmp, lcsCode)//输出压缩码
		fmt.Printf("%d ", lcsCode);//输出压缩码
		if lcsLen < len(slice) {
			sliceTmp := slice[:lcsLen + 1]//最长匹配串+贪婪一个字符，插入TST中
			code += 1
			root = tst.InsertChar(root, sliceTmp, 0, code)
			//fmt.Printf("%s : %x {insert} %d\n", sliceTmp,  code, code);
		} else {
			fmt.Printf("%d", R)//R值作为压缩文件终止符
			break //最后剩下的待处理文本出现在了TST中，处理完毕，退出
		}
		slice = slice[lcsLen:]
	}
	fmt.Println()
}


func  Expand(codeInput []int16) {
	st := make(map[int16]string)
	var i int16
	for i = 0; i < R; i++ {
		st[i] = string(i)
	}

	var code int16 = R 

	var codeTmp int16 = codeInput[0]
	var val string = st[codeTmp]
	var codeInputLen int16 = int16(len(codeInput))
//	fmt.Println("codeTmp:", codeTmp)
	for i = 1; i < codeInputLen; i++ {
		code++
		fmt.Printf("%s", val)
		codeTmp = codeInput[i]
		if codeTmp == R {
			break
		}
//		fmt.Println("codeTmp:", codeTmp)
		if code == codeTmp {
			val = val + string(val[0])
			st[code] = val
		} else {
			s := st[codeTmp]
			val = val + string(s[0])
			st[code] = val
			val = s
		}
	}
	fmt.Println()
}

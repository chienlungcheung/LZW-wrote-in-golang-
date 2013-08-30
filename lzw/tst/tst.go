package tst

import (
	
)

type TSTNode struct {
	left, middle, right *TSTNode
	ch byte
	code int16
}

//insert a node into the TSTree
func InsertChar(root *TSTNode, str string, index int, code int16) *TSTNode {
	len_str := len(str)

	if index == len_str {
		return root
	}

	if root == nil {
		root = new(TSTNode)
		root.ch = str[index]
		root.code = -1
		if index == len_str - 1 {
			root.code = code
		} else {
			root.middle = InsertChar(root.middle, str, index + 1, code)
		}
	} else {
		if str[index] < root.ch {
			root.left = InsertChar(root.left, str, index, code)
		} else if str[index] > root.ch {
			root.right = InsertChar(root.right, str, index, code)
		} else {
			if index == len_str - 1 {
				root.code = code
			} else {
				root.middle = InsertChar(root.middle, str, index + 1, code)
			}
		}
	}
	return root
}

//if str exists, return its code; or, return -1
func GetCode(root *TSTNode, str string, index int) int16 {
	len_str := len(str)

	if index == len_str {
		return -1
	}

	if root == nil {
		return -1
	} else {
		if str[index] < root.ch {
			return GetCode(root.left, str, index)
		} else if str[index] > root.ch {
			return GetCode(root.right, str, index)
		} else {
			if index == len_str - 1 {
				return root.code
			} else {
				return GetCode(root.middle, str, index + 1)
			}
		}
	}
}

//get length of the longest comman string
func GetLCSLength(root *TSTNode, str string, index int) int {
	len_str := len(str)

	if index == len_str {
		return index
	}

	if root == nil {
		return index
	} else {
		if str[index] < root.ch {
			return GetLCSLength(root.left, str, index)
		} else if str[index] > root.ch {
			return GetLCSLength(root.right, str, index)
		} else {
			return GetLCSLength(root.middle, str, index + 1)
		}
	}
}

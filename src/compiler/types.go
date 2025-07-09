package compiler

import "github.com/llir/llvm/ir/types"

func I32() types.Type {
	return types.I32
}

func I1() types.Type {
	return types.I1
}

func Void() types.Type {
	return types.Void
}

func I8() types.Type {
	return types.I8
}

func I64() types.Type {
	return types.I64
}

func Float() types.Type {
	return types.Float
}

func Double() types.Type {
	return types.Double
}

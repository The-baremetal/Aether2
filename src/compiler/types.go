package compiler

import "tinygo.org/x/go-llvm"

func I32(context llvm.Context) llvm.Type {
	return context.Int32Type()
}

func I1(context llvm.Context) llvm.Type {
	return context.Int1Type()
}

func Void(context llvm.Context) llvm.Type {
	return context.VoidType()
}

func I8(context llvm.Context) llvm.Type {
	return context.Int8Type()
}

func I64(context llvm.Context) llvm.Type {
	return context.Int64Type()
}

func Float(context llvm.Context) llvm.Type {
	return context.FloatType()
}

func Double(context llvm.Context) llvm.Type {
	return context.DoubleType()
}

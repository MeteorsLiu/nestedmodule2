package mpfr

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const X__GMPXXUSECXX11 = 0
const X__GMPZULILIMBS = 1
/**************** Function objects ****************/

type X__gmpUnaryPlus struct {
	Eval func(MpzPtr, MpzSrcptr)
}

type X__gmpUnaryMinus struct {
	Eval func(MpzPtr, MpzSrcptr)
}

type X__gmpUnaryCom struct {
	Eval func(MpzPtr, MpzSrcptr)
}

type X__gmpBinaryPlus struct {
	Eval func(MpzPtr, MpzSrcptr, MpzSrcptr)
}

type X__gmpBinaryMinus struct {
	Eval func(MpzPtr, MpzSrcptr, MpzSrcptr)
}

type X__gmpBinaryLshift struct {
	Eval func(MpzPtr, MpzSrcptr, MpBitcntT)
}

type X__gmpBinaryRshift struct {
	Eval func(MpzPtr, MpzSrcptr, MpBitcntT)
}

type X__gmpBinaryMultiplies struct {
	Eval func(MpzPtr, MpzSrcptr, MpzSrcptr)
}

type X__gmpBinaryDivides struct {
	Eval func(MpzPtr, MpzSrcptr, MpzSrcptr)
}

type X__gmpBinaryModulus struct {
	Eval func(MpzPtr, MpzSrcptr, MpzSrcptr)
}

type X__gmpBinaryAnd struct {
	Eval func(MpzPtr, MpzSrcptr, MpzSrcptr)
}

type X__gmpBinaryIor struct {
	Eval func(MpzPtr, MpzSrcptr, MpzSrcptr)
}

type X__gmpBinaryXor struct {
	Eval func(MpzPtr, MpzSrcptr, MpzSrcptr)
}

type X__gmpCmpFunction struct {
	Eval func(MpzSrcptr, MpzSrcptr) c.Int
}

type X__gmpBinaryEqual struct {
	Eval func(MpzSrcptr, MpzSrcptr) c.Int
}

type X__gmpBinaryLess struct {
	Eval func(MpzSrcptr, MpzSrcptr) c.Int
}

type X__gmpBinaryGreater struct {
	 c.Int
}

type X__gmpUnaryIncrement struct {
	Eval func(MpzPtr)
}

type X__gmpUnaryDecrement struct {
	Eval func(MpzPtr)
}

type X__gmpAbsFunction struct {
	Eval func(MpzPtr, MpzSrcptr)
}

type X__gmpTruncFunction struct {
	Eval func(MpfPtr, MpfSrcptr)
}

type X__gmpFloorFunction struct {
	Eval func(MpfPtr, MpfSrcptr)
}

type X__gmpCeilFunction struct {
	Eval func(MpfPtr, MpfSrcptr)
}

type X__gmpSqrtFunction struct {
	Eval func(MpzPtr, MpzSrcptr)
}

type X__gmpHypotFunction struct {
	Eval func(MpfPtr, MpfSrcptr, MpfSrcptr)
}

type X__gmpSgnFunction struct {
	Eval func(MpzSrcptr) c.Int
}

type X__gmpGcdFunction struct {
	Eval func(MpzPtr, MpzSrcptr, MpzSrcptr)
}

type X__gmpLcmFunction struct {
	Eval func(MpzPtr, MpzSrcptr, MpzSrcptr)
}

type X__gmpRandFunction struct {
	Eval func(MpzPtr, GmpRandstateT, MpBitcntT)
}

type X__gmpFacFunction struct {
	Eval func(MpzPtr, c.Ulong)
}

type X__gmpPrimorialFunction struct {
	Eval func(MpzPtr, c.Ulong)
}

type X__gmpFibFunction struct {
	Eval func(MpzPtr, c.Ulong)
}

type X__gmpAllocCstring struct {
	Str                *int8
	X__gmpAllocCstring func(*int8) c.Int
}

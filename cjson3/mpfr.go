package mpfr

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const MPFRVERSIONMAJOR = 4
const MPFRVERSIONMINOR = 2
const MPFRVERSIONPATCHLEVEL = 1
const MPFRVERSIONSTRING = "4.2.1"
const MPFRFLAGSUNDERFLOW = 1
const MPFRFLAGSOVERFLOW = 2
const MPFRFLAGSNAN = 4
const MPFRFLAGSINEXACT = 8
const MPFRFLAGSERANGE = 16
const MPFRFLAGSDIVBY0 = 32
const X_MPFRPRECFORMAT = 3
const MPFRPRECMIN = 1
const MPFRUSEC99FEATURE = 1

type MpfrVoid [0]byte
type MpfrInt c.Int
type MpfrUint c.Uint
type MpfrLong c.Long
type MpfrUlong c.Ulong
type MpfrSizeT uintptr
type MpfrFlagsT c.Uint
type MpfrRndT c.Int

const (
	MpfrRndTMPFRRNDN  MpfrRndT = 0
	MpfrRndTMPFRRNDZ  MpfrRndT = 1
	MpfrRndTMPFRRNDU  MpfrRndT = 2
	MpfrRndTMPFRRNDD  MpfrRndT = 3
	MpfrRndTMPFRRNDA  MpfrRndT = 4
	MpfrRndTMPFRRNDF  MpfrRndT = 5
	MpfrRndTMPFRRNDNA MpfrRndT = -1
)

type MpfrPrecT c.Long
type MpfrUprecT c.Ulong
type MpfrSignT c.Int
type MpfrExpT c.Long
type MpfrUexpT c.Ulong

type X__mpfrStruct struct {
	X_mpfrPrec MpfrPrecT
	X_mpfrSign MpfrSignT
	X_mpfrExp  MpfrExpT
	X_mpfrD    *MpLimbT
}
type MpfrT [1]X__mpfrStruct
type MpfrPtr *X__mpfrStruct
type MpfrSrcptr *X__mpfrStruct
type MpfrKindT c.Int

const (
	MpfrKindTMPFRNANKIND     MpfrKindT = 0
	MpfrKindTMPFRINFKIND     MpfrKindT = 1
	MpfrKindTMPFRZEROKIND    MpfrKindT = 2
	MpfrKindTMPFRREGULARKIND MpfrKindT = 3
)

type MpfrFreeCacheT c.Int

const (
	MpfrFreeCacheTMPFRFREELOCALCACHE  MpfrFreeCacheT = 1
	MpfrFreeCacheTMPFRFREEGLOBALCACHE MpfrFreeCacheT = 2
)
//go:linkname MpfrGetVersion C.mpfr_get_version
func MpfrGetVersion() *int8
//go:linkname MpfrGetPatches C.mpfr_get_patches
func MpfrGetPatches() *int8
//go:linkname MpfrBuildoptTlsP C.mpfr_buildopt_tls_p
func MpfrBuildoptTlsP() c.Int
//go:linkname MpfrBuildoptFloat128P C.mpfr_buildopt_float128_p
func MpfrBuildoptFloat128P() c.Int
//go:linkname MpfrBuildoptDecimalP C.mpfr_buildopt_decimal_p
func MpfrBuildoptDecimalP() c.Int
//go:linkname MpfrBuildoptGmpinternalsP C.mpfr_buildopt_gmpinternals_p
func MpfrBuildoptGmpinternalsP() c.Int
//go:linkname MpfrBuildoptSharedcacheP C.mpfr_buildopt_sharedcache_p
func MpfrBuildoptSharedcacheP() c.Int
//go:linkname MpfrBuildoptTuneCase C.mpfr_buildopt_tune_case
func MpfrBuildoptTuneCase() *int8
//go:linkname MpfrGetEmin C.mpfr_get_emin
func MpfrGetEmin() MpfrExpT
// llgo:link MpfrExpT.MpfrSetEmin C.mpfr_set_emin
func (recv_ MpfrExpT) MpfrSetEmin() c.Int {
	return 0
}
//go:linkname MpfrGetEminMin C.mpfr_get_emin_min
func MpfrGetEminMin() MpfrExpT
//go:linkname MpfrGetEminMax C.mpfr_get_emin_max
func MpfrGetEminMax() MpfrExpT
//go:linkname MpfrGetEmax C.mpfr_get_emax
func MpfrGetEmax() MpfrExpT
// llgo:link MpfrExpT.MpfrSetEmax C.mpfr_set_emax
func (recv_ MpfrExpT) MpfrSetEmax() c.Int {
	return 0
}
//go:linkname MpfrGetEmaxMin C.mpfr_get_emax_min
func MpfrGetEmaxMin() MpfrExpT
//go:linkname MpfrGetEmaxMax C.mpfr_get_emax_max
func MpfrGetEmaxMax() MpfrExpT
// llgo:link MpfrRndT.MpfrSetDefaultRoundingMode C.mpfr_set_default_rounding_mode
func (recv_ MpfrRndT) MpfrSetDefaultRoundingMode() {
}
//go:linkname MpfrGetDefaultRoundingMode C.mpfr_get_default_rounding_mode
func MpfrGetDefaultRoundingMode() MpfrRndT
// llgo:link MpfrRndT.MpfrPrintRndMode C.mpfr_print_rnd_mode
func (recv_ MpfrRndT) MpfrPrintRndMode() *int8 {
	return nil
}
//go:linkname MpfrClearFlags C.mpfr_clear_flags
func MpfrClearFlags()
//go:linkname MpfrClearUnderflow C.mpfr_clear_underflow
func MpfrClearUnderflow()
//go:linkname MpfrClearOverflow C.mpfr_clear_overflow
func MpfrClearOverflow()
//go:linkname MpfrClearDivby0 C.mpfr_clear_divby0
func MpfrClearDivby0()
//go:linkname MpfrClearNanflag C.mpfr_clear_nanflag
func MpfrClearNanflag()
//go:linkname MpfrClearInexflag C.mpfr_clear_inexflag
func MpfrClearInexflag()
//go:linkname MpfrClearErangeflag C.mpfr_clear_erangeflag
func MpfrClearErangeflag()
//go:linkname MpfrSetUnderflow C.mpfr_set_underflow
func MpfrSetUnderflow()
//go:linkname MpfrSetOverflow C.mpfr_set_overflow
func MpfrSetOverflow()
//go:linkname MpfrSetDivby0 C.mpfr_set_divby0
func MpfrSetDivby0()
//go:linkname MpfrSetNanflag C.mpfr_set_nanflag
func MpfrSetNanflag()
//go:linkname MpfrSetInexflag C.mpfr_set_inexflag
func MpfrSetInexflag()
//go:linkname MpfrSetErangeflag C.mpfr_set_erangeflag
func MpfrSetErangeflag()
//go:linkname MpfrUnderflowP C.mpfr_underflow_p
func MpfrUnderflowP() c.Int
//go:linkname MpfrOverflowP C.mpfr_overflow_p
func MpfrOverflowP() c.Int
//go:linkname MpfrDivby0P C.mpfr_divby0_p
func MpfrDivby0P() c.Int
//go:linkname MpfrNanflagP C.mpfr_nanflag_p
func MpfrNanflagP() c.Int
//go:linkname MpfrInexflagP C.mpfr_inexflag_p
func MpfrInexflagP() c.Int
//go:linkname MpfrErangeflagP C.mpfr_erangeflag_p
func MpfrErangeflagP() c.Int
// llgo:link MpfrFlagsT.MpfrFlagsClear C.mpfr_flags_clear
func (recv_ MpfrFlagsT) MpfrFlagsClear() {
}
// llgo:link MpfrFlagsT.MpfrFlagsSet C.mpfr_flags_set
func (recv_ MpfrFlagsT) MpfrFlagsSet() {
}
// llgo:link MpfrFlagsT.MpfrFlagsTest C.mpfr_flags_test
func (recv_ MpfrFlagsT) MpfrFlagsTest() MpfrFlagsT {
	return 0
}
//go:linkname MpfrFlagsSave C.mpfr_flags_save
func MpfrFlagsSave() MpfrFlagsT
// llgo:link MpfrFlagsT.MpfrFlagsRestore C.mpfr_flags_restore
func (recv_ MpfrFlagsT) MpfrFlagsRestore(MpfrFlagsT) {
}
//go:linkname MpfrCheckRange C.mpfr_check_range
func MpfrCheckRange(MpfrPtr, c.Int, MpfrRndT) c.Int
//go:linkname MpfrInit2 C.mpfr_init2
func MpfrInit2(MpfrPtr, MpfrPrecT)
//go:linkname MpfrInit C.mpfr_init
func MpfrInit(MpfrPtr)
//go:linkname MpfrClear C.mpfr_clear
func MpfrClear(MpfrPtr)
// llgo:link MpfrPrecT.MpfrInits2 C.mpfr_inits2
func (recv_ MpfrPrecT) MpfrInits2(__llgo_arg_0 MpfrPtr, __llgo_va_list ...interface{}) {
}
//go:linkname MpfrInits C.mpfr_inits
func MpfrInits(__llgo_arg_0 MpfrPtr, __llgo_va_list ...interface{})
//go:linkname MpfrClears C.mpfr_clears
func MpfrClears(__llgo_arg_0 MpfrPtr, __llgo_va_list ...interface{})
//go:linkname MpfrPrecRound C.mpfr_prec_round
func MpfrPrecRound(MpfrPtr, MpfrPrecT, MpfrRndT) c.Int
//go:linkname MpfrCanRound C.mpfr_can_round
func MpfrCanRound(MpfrSrcptr, MpfrExpT, MpfrRndT, MpfrRndT, MpfrPrecT) c.Int
//go:linkname MpfrMinPrec C.mpfr_min_prec
func MpfrMinPrec(MpfrSrcptr) MpfrPrecT
//go:linkname MpfrGetExp C.mpfr_get_exp
func MpfrGetExp(MpfrSrcptr) MpfrExpT
//go:linkname MpfrSetExp C.mpfr_set_exp
func MpfrSetExp(MpfrPtr, MpfrExpT) c.Int
//go:linkname MpfrGetPrec C.mpfr_get_prec
func MpfrGetPrec(MpfrSrcptr) MpfrPrecT
//go:linkname MpfrSetPrec C.mpfr_set_prec
func MpfrSetPrec(MpfrPtr, MpfrPrecT)
//go:linkname MpfrSetPrecRaw C.mpfr_set_prec_raw
func MpfrSetPrecRaw(MpfrPtr, MpfrPrecT)
// llgo:link MpfrPrecT.MpfrSetDefaultPrec C.mpfr_set_default_prec
func (recv_ MpfrPrecT) MpfrSetDefaultPrec() {
}
//go:linkname MpfrGetDefaultPrec C.mpfr_get_default_prec
func MpfrGetDefaultPrec() MpfrPrecT
//go:linkname MpfrSetD C.mpfr_set_d
func MpfrSetD(MpfrPtr, float64, MpfrRndT) c.Int
//go:linkname MpfrSetFlt C.mpfr_set_flt
func MpfrSetFlt(MpfrPtr, float32, MpfrRndT) c.Int
//go:linkname MpfrSetLd C.mpfr_set_ld
func MpfrSetLd(MpfrPtr, float64, MpfrRndT) c.Int
//go:linkname MpfrSetZ C.mpfr_set_z
func MpfrSetZ(MpfrPtr, MpzSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSetZ2exp C.mpfr_set_z_2exp
func MpfrSetZ2exp(MpfrPtr, MpzSrcptr, MpfrExpT, MpfrRndT) c.Int
//go:linkname MpfrSetNan C.mpfr_set_nan
func MpfrSetNan(MpfrPtr)
//go:linkname MpfrSetInf C.mpfr_set_inf
func MpfrSetInf(MpfrPtr, c.Int)
//go:linkname MpfrSetZero C.mpfr_set_zero
func MpfrSetZero(MpfrPtr, c.Int)
//go:linkname MpfrSetF C.mpfr_set_f
func MpfrSetF(MpfrPtr, MpfSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCmpF C.mpfr_cmp_f
func MpfrCmpF(MpfrSrcptr, MpfSrcptr) c.Int
//go:linkname MpfrGetF C.mpfr_get_f
func MpfrGetF(MpfPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSetSi C.mpfr_set_si
func MpfrSetSi(MpfrPtr, c.Long, MpfrRndT) c.Int
//go:linkname MpfrSetUi C.mpfr_set_ui
func MpfrSetUi(MpfrPtr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrSetSi2exp C.mpfr_set_si_2exp
func MpfrSetSi2exp(MpfrPtr, c.Long, MpfrExpT, MpfrRndT) c.Int
//go:linkname MpfrSetUi2exp C.mpfr_set_ui_2exp
func MpfrSetUi2exp(MpfrPtr, c.Ulong, MpfrExpT, MpfrRndT) c.Int
//go:linkname MpfrSetQ C.mpfr_set_q
func MpfrSetQ(MpfrPtr, MpqSrcptr, MpfrRndT) c.Int
//go:linkname MpfrMulQ C.mpfr_mul_q
func MpfrMulQ(MpfrPtr, MpfrSrcptr, MpqSrcptr, MpfrRndT) c.Int
//go:linkname MpfrDivQ C.mpfr_div_q
func MpfrDivQ(MpfrPtr, MpfrSrcptr, MpqSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAddQ C.mpfr_add_q
func MpfrAddQ(MpfrPtr, MpfrSrcptr, MpqSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSubQ C.mpfr_sub_q
func MpfrSubQ(MpfrPtr, MpfrSrcptr, MpqSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCmpQ C.mpfr_cmp_q
func MpfrCmpQ(MpfrSrcptr, MpqSrcptr) c.Int
//go:linkname MpfrGetQ C.mpfr_get_q
func MpfrGetQ(q MpqPtr, f MpfrSrcptr)
//go:linkname MpfrSetStr C.mpfr_set_str
func MpfrSetStr(MpfrPtr, *int8, c.Int, MpfrRndT) c.Int
//go:linkname MpfrInitSetStr C.mpfr_init_set_str
func MpfrInitSetStr(MpfrPtr, *int8, c.Int, MpfrRndT) c.Int
//go:linkname MpfrSet4 C.mpfr_set4
func MpfrSet4(MpfrPtr, MpfrSrcptr, MpfrRndT, c.Int) c.Int
//go:linkname MpfrAbs C.mpfr_abs
func MpfrAbs(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSet C.mpfr_set
func MpfrSet(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrNeg C.mpfr_neg
func MpfrNeg(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSignbit C.mpfr_signbit
func MpfrSignbit(MpfrSrcptr) c.Int
//go:linkname MpfrSetsign C.mpfr_setsign
func MpfrSetsign(MpfrPtr, MpfrSrcptr, c.Int, MpfrRndT) c.Int
//go:linkname MpfrCopysign C.mpfr_copysign
func MpfrCopysign(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrGetZ2exp C.mpfr_get_z_2exp
func MpfrGetZ2exp(MpzPtr, MpfrSrcptr) MpfrExpT
//go:linkname MpfrGetFlt C.mpfr_get_flt
func MpfrGetFlt(MpfrSrcptr, MpfrRndT) float32
//go:linkname MpfrGetD C.mpfr_get_d
func MpfrGetD(MpfrSrcptr, MpfrRndT) float64
//go:linkname MpfrGetLd C.mpfr_get_ld
func MpfrGetLd(MpfrSrcptr, MpfrRndT) float64
//go:linkname MpfrGetD1 C.mpfr_get_d1
func MpfrGetD1(MpfrSrcptr) float64
//go:linkname MpfrGetD2exp C.mpfr_get_d_2exp
func MpfrGetD2exp(*c.Long, MpfrSrcptr, MpfrRndT) float64
//go:linkname MpfrGetLd2exp C.mpfr_get_ld_2exp
func MpfrGetLd2exp(*c.Long, MpfrSrcptr, MpfrRndT) float64
// llgo:link (*MpfrExpT).MpfrFrexp C.mpfr_frexp
func (recv_ *MpfrExpT) MpfrFrexp(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int {
	return 0
}
//go:linkname MpfrGetSi C.mpfr_get_si
func MpfrGetSi(MpfrSrcptr, MpfrRndT) c.Long
//go:linkname MpfrGetUi C.mpfr_get_ui
func MpfrGetUi(MpfrSrcptr, MpfrRndT) c.Ulong
//go:linkname MpfrGetStrNdigits C.mpfr_get_str_ndigits
func MpfrGetStrNdigits(c.Int, MpfrPrecT) uintptr
//go:linkname MpfrGetStr C.mpfr_get_str
func MpfrGetStr(*int8, *MpfrExpT, c.Int, uintptr, MpfrSrcptr, MpfrRndT) *int8
//go:linkname MpfrGetZ C.mpfr_get_z
func MpfrGetZ(z MpzPtr, f MpfrSrcptr, __llgo_arg_2 MpfrRndT) c.Int
//go:linkname MpfrFreeStr C.mpfr_free_str
func MpfrFreeStr(*int8)
//go:linkname MpfrUrandom C.mpfr_urandom
func MpfrUrandom(MpfrPtr, GmpRandstateT, MpfrRndT) c.Int
//go:linkname MpfrGrandom C.mpfr_grandom
func MpfrGrandom(MpfrPtr, MpfrPtr, GmpRandstateT, MpfrRndT) c.Int
//go:linkname MpfrNrandom C.mpfr_nrandom
func MpfrNrandom(MpfrPtr, GmpRandstateT, MpfrRndT) c.Int
//go:linkname MpfrErandom C.mpfr_erandom
func MpfrErandom(MpfrPtr, GmpRandstateT, MpfrRndT) c.Int
//go:linkname MpfrUrandomb C.mpfr_urandomb
func MpfrUrandomb(MpfrPtr, GmpRandstateT) c.Int
//go:linkname MpfrNextabove C.mpfr_nextabove
func MpfrNextabove(MpfrPtr)
//go:linkname MpfrNextbelow C.mpfr_nextbelow
func MpfrNextbelow(MpfrPtr)
//go:linkname MpfrNexttoward C.mpfr_nexttoward
func MpfrNexttoward(MpfrPtr, MpfrSrcptr)
//go:linkname MpfrPrintf C.mpfr_printf
func MpfrPrintf(__llgo_arg_0 *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname MpfrAsprintf C.mpfr_asprintf
func MpfrAsprintf(__llgo_arg_0 **int8, __llgo_arg_1 *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname MpfrSprintf C.mpfr_sprintf
func MpfrSprintf(__llgo_arg_0 *int8, __llgo_arg_1 *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname MpfrSnprintf C.mpfr_snprintf
func MpfrSnprintf(__llgo_arg_0 *int8, __llgo_arg_1 uintptr, __llgo_arg_2 *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname MpfrPow C.mpfr_pow
func MpfrPow(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrPowr C.mpfr_powr
func MpfrPowr(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrPowSi C.mpfr_pow_si
func MpfrPowSi(MpfrPtr, MpfrSrcptr, c.Long, MpfrRndT) c.Int
//go:linkname MpfrCompoundSi C.mpfr_compound_si
func MpfrCompoundSi(MpfrPtr, MpfrSrcptr, c.Long, MpfrRndT) c.Int
//go:linkname MpfrPowUi C.mpfr_pow_ui
func MpfrPowUi(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrUiPowUi C.mpfr_ui_pow_ui
func MpfrUiPowUi(MpfrPtr, c.Ulong, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrUiPow C.mpfr_ui_pow
func MpfrUiPow(MpfrPtr, c.Ulong, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrPowZ C.mpfr_pow_z
func MpfrPowZ(MpfrPtr, MpfrSrcptr, MpzSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSqrt C.mpfr_sqrt
func MpfrSqrt(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSqrtUi C.mpfr_sqrt_ui
func MpfrSqrtUi(MpfrPtr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrRecSqrt C.mpfr_rec_sqrt
func MpfrRecSqrt(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAdd C.mpfr_add
func MpfrAdd(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSub C.mpfr_sub
func MpfrSub(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrMul C.mpfr_mul
func MpfrMul(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrDiv C.mpfr_div
func MpfrDiv(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAddUi C.mpfr_add_ui
func MpfrAddUi(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrSubUi C.mpfr_sub_ui
func MpfrSubUi(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrUiSub C.mpfr_ui_sub
func MpfrUiSub(MpfrPtr, c.Ulong, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrMulUi C.mpfr_mul_ui
func MpfrMulUi(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrDivUi C.mpfr_div_ui
func MpfrDivUi(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrUiDiv C.mpfr_ui_div
func MpfrUiDiv(MpfrPtr, c.Ulong, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAddSi C.mpfr_add_si
func MpfrAddSi(MpfrPtr, MpfrSrcptr, c.Long, MpfrRndT) c.Int
//go:linkname MpfrSubSi C.mpfr_sub_si
func MpfrSubSi(MpfrPtr, MpfrSrcptr, c.Long, MpfrRndT) c.Int
//go:linkname MpfrSiSub C.mpfr_si_sub
func MpfrSiSub(MpfrPtr, c.Long, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrMulSi C.mpfr_mul_si
func MpfrMulSi(MpfrPtr, MpfrSrcptr, c.Long, MpfrRndT) c.Int
//go:linkname MpfrDivSi C.mpfr_div_si
func MpfrDivSi(MpfrPtr, MpfrSrcptr, c.Long, MpfrRndT) c.Int
//go:linkname MpfrSiDiv C.mpfr_si_div
func MpfrSiDiv(MpfrPtr, c.Long, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAddD C.mpfr_add_d
func MpfrAddD(MpfrPtr, MpfrSrcptr, float64, MpfrRndT) c.Int
//go:linkname MpfrSubD C.mpfr_sub_d
func MpfrSubD(MpfrPtr, MpfrSrcptr, float64, MpfrRndT) c.Int
//go:linkname MpfrDSub C.mpfr_d_sub
func MpfrDSub(MpfrPtr, float64, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrMulD C.mpfr_mul_d
func MpfrMulD(MpfrPtr, MpfrSrcptr, float64, MpfrRndT) c.Int
//go:linkname MpfrDivD C.mpfr_div_d
func MpfrDivD(MpfrPtr, MpfrSrcptr, float64, MpfrRndT) c.Int
//go:linkname MpfrDDiv C.mpfr_d_div
func MpfrDDiv(MpfrPtr, float64, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSqr C.mpfr_sqr
func MpfrSqr(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrConstPi C.mpfr_const_pi
func MpfrConstPi(MpfrPtr, MpfrRndT) c.Int
//go:linkname MpfrConstLog2 C.mpfr_const_log2
func MpfrConstLog2(MpfrPtr, MpfrRndT) c.Int
//go:linkname MpfrConstEuler C.mpfr_const_euler
func MpfrConstEuler(MpfrPtr, MpfrRndT) c.Int
//go:linkname MpfrConstCatalan C.mpfr_const_catalan
func MpfrConstCatalan(MpfrPtr, MpfrRndT) c.Int
//go:linkname MpfrAgm C.mpfr_agm
func MpfrAgm(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrLog C.mpfr_log
func MpfrLog(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrLog2 C.mpfr_log2
func MpfrLog2(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrLog10 C.mpfr_log10
func MpfrLog10(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrLog1p C.mpfr_log1p
func MpfrLog1p(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrLog2p1 C.mpfr_log2p1
func MpfrLog2p1(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrLog10p1 C.mpfr_log10p1
func MpfrLog10p1(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrLogUi C.mpfr_log_ui
func MpfrLogUi(MpfrPtr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrExp C.mpfr_exp
func MpfrExp(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrExp2 C.mpfr_exp2
func MpfrExp2(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrExp10 C.mpfr_exp10
func MpfrExp10(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrExpm1 C.mpfr_expm1
func MpfrExpm1(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrExp2m1 C.mpfr_exp2m1
func MpfrExp2m1(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrExp10m1 C.mpfr_exp10m1
func MpfrExp10m1(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrEint C.mpfr_eint
func MpfrEint(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrLi2 C.mpfr_li2
func MpfrLi2(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCmp C.mpfr_cmp
func MpfrCmp(MpfrSrcptr, MpfrSrcptr) c.Int
//go:linkname MpfrCmp3 C.mpfr_cmp3
func MpfrCmp3(MpfrSrcptr, MpfrSrcptr, c.Int) c.Int
//go:linkname MpfrCmpD C.mpfr_cmp_d
func MpfrCmpD(MpfrSrcptr, float64) c.Int
//go:linkname MpfrCmpLd C.mpfr_cmp_ld
func MpfrCmpLd(MpfrSrcptr, float64) c.Int
//go:linkname MpfrCmpUi C.mpfr_cmp_ui
func MpfrCmpUi(MpfrSrcptr, c.Ulong) c.Int
//go:linkname MpfrCmpSi C.mpfr_cmp_si
func MpfrCmpSi(MpfrSrcptr, c.Long) c.Int
//go:linkname MpfrCmpUi2exp C.mpfr_cmp_ui_2exp
func MpfrCmpUi2exp(MpfrSrcptr, c.Ulong, MpfrExpT) c.Int
//go:linkname MpfrCmpSi2exp C.mpfr_cmp_si_2exp
func MpfrCmpSi2exp(MpfrSrcptr, c.Long, MpfrExpT) c.Int
//go:linkname MpfrCmpabs C.mpfr_cmpabs
func MpfrCmpabs(MpfrSrcptr, MpfrSrcptr) c.Int
//go:linkname MpfrCmpabsUi C.mpfr_cmpabs_ui
func MpfrCmpabsUi(MpfrSrcptr, c.Ulong) c.Int
//go:linkname MpfrReldiff C.mpfr_reldiff
func MpfrReldiff(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT)
//go:linkname MpfrEq C.mpfr_eq
func MpfrEq(MpfrSrcptr, MpfrSrcptr, c.Ulong) c.Int
//go:linkname MpfrSgn C.mpfr_sgn
func MpfrSgn(MpfrSrcptr) c.Int
//go:linkname MpfrMul2exp C.mpfr_mul_2exp
func MpfrMul2exp(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrDiv2exp C.mpfr_div_2exp
func MpfrDiv2exp(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrMul2ui C.mpfr_mul_2ui
func MpfrMul2ui(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrDiv2ui C.mpfr_div_2ui
func MpfrDiv2ui(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrMul2si C.mpfr_mul_2si
func MpfrMul2si(MpfrPtr, MpfrSrcptr, c.Long, MpfrRndT) c.Int
//go:linkname MpfrDiv2si C.mpfr_div_2si
func MpfrDiv2si(MpfrPtr, MpfrSrcptr, c.Long, MpfrRndT) c.Int
//go:linkname MpfrRint C.mpfr_rint
func MpfrRint(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrRoundeven C.mpfr_roundeven
func MpfrRoundeven(MpfrPtr, MpfrSrcptr) c.Int
//go:linkname MpfrRound C.mpfr_round
func MpfrRound(MpfrPtr, MpfrSrcptr) c.Int
//go:linkname MpfrTrunc C.mpfr_trunc
func MpfrTrunc(MpfrPtr, MpfrSrcptr) c.Int
//go:linkname MpfrCeil C.mpfr_ceil
func MpfrCeil(MpfrPtr, MpfrSrcptr) c.Int
//go:linkname MpfrFloor C.mpfr_floor
func MpfrFloor(MpfrPtr, MpfrSrcptr) c.Int
//go:linkname MpfrRintRoundeven C.mpfr_rint_roundeven
func MpfrRintRoundeven(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrRintRound C.mpfr_rint_round
func MpfrRintRound(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrRintTrunc C.mpfr_rint_trunc
func MpfrRintTrunc(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrRintCeil C.mpfr_rint_ceil
func MpfrRintCeil(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrRintFloor C.mpfr_rint_floor
func MpfrRintFloor(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFrac C.mpfr_frac
func MpfrFrac(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrModf C.mpfr_modf
func MpfrModf(MpfrPtr, MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrRemquo C.mpfr_remquo
func MpfrRemquo(MpfrPtr, *c.Long, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrRemainder C.mpfr_remainder
func MpfrRemainder(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFmod C.mpfr_fmod
func MpfrFmod(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFmodUi C.mpfr_fmod_ui
func MpfrFmodUi(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrFmodquo C.mpfr_fmodquo
func MpfrFmodquo(MpfrPtr, *c.Long, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFitsUlongP C.mpfr_fits_ulong_p
func MpfrFitsUlongP(MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFitsSlongP C.mpfr_fits_slong_p
func MpfrFitsSlongP(MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFitsUintP C.mpfr_fits_uint_p
func MpfrFitsUintP(MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFitsSintP C.mpfr_fits_sint_p
func MpfrFitsSintP(MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFitsUshortP C.mpfr_fits_ushort_p
func MpfrFitsUshortP(MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFitsSshortP C.mpfr_fits_sshort_p
func MpfrFitsSshortP(MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFitsUintmaxP C.mpfr_fits_uintmax_p
func MpfrFitsUintmaxP(MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFitsIntmaxP C.mpfr_fits_intmax_p
func MpfrFitsIntmaxP(MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrExtract C.mpfr_extract
func MpfrExtract(MpzPtr, MpfrSrcptr, c.Uint)
//go:linkname MpfrSwap C.mpfr_swap
func MpfrSwap(MpfrPtr, MpfrPtr)
//go:linkname MpfrDump C.mpfr_dump
func MpfrDump(MpfrSrcptr)
//go:linkname MpfrNanP C.mpfr_nan_p
func MpfrNanP(MpfrSrcptr) c.Int
//go:linkname MpfrInfP C.mpfr_inf_p
func MpfrInfP(MpfrSrcptr) c.Int
//go:linkname MpfrNumberP C.mpfr_number_p
func MpfrNumberP(MpfrSrcptr) c.Int
//go:linkname MpfrIntegerP C.mpfr_integer_p
func MpfrIntegerP(MpfrSrcptr) c.Int
//go:linkname MpfrZeroP C.mpfr_zero_p
func MpfrZeroP(MpfrSrcptr) c.Int
//go:linkname MpfrRegularP C.mpfr_regular_p
func MpfrRegularP(MpfrSrcptr) c.Int
//go:linkname MpfrGreaterP C.mpfr_greater_p
func MpfrGreaterP(MpfrSrcptr, MpfrSrcptr) c.Int
//go:linkname MpfrGreaterequalP C.mpfr_greaterequal_p
func MpfrGreaterequalP(MpfrSrcptr, MpfrSrcptr) c.Int
//go:linkname MpfrLessP C.mpfr_less_p
func MpfrLessP(MpfrSrcptr, MpfrSrcptr) c.Int
//go:linkname MpfrLessequalP C.mpfr_lessequal_p
func MpfrLessequalP(MpfrSrcptr, MpfrSrcptr) c.Int
//go:linkname MpfrLessgreaterP C.mpfr_lessgreater_p
func MpfrLessgreaterP(MpfrSrcptr, MpfrSrcptr) c.Int
//go:linkname MpfrEqualP C.mpfr_equal_p
func MpfrEqualP(MpfrSrcptr, MpfrSrcptr) c.Int
//go:linkname MpfrUnorderedP C.mpfr_unordered_p
func MpfrUnorderedP(MpfrSrcptr, MpfrSrcptr) c.Int
//go:linkname MpfrAtanh C.mpfr_atanh
func MpfrAtanh(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAcosh C.mpfr_acosh
func MpfrAcosh(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAsinh C.mpfr_asinh
func MpfrAsinh(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCosh C.mpfr_cosh
func MpfrCosh(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSinh C.mpfr_sinh
func MpfrSinh(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrTanh C.mpfr_tanh
func MpfrTanh(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSinhCosh C.mpfr_sinh_cosh
func MpfrSinhCosh(MpfrPtr, MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSech C.mpfr_sech
func MpfrSech(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCsch C.mpfr_csch
func MpfrCsch(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCoth C.mpfr_coth
func MpfrCoth(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAcos C.mpfr_acos
func MpfrAcos(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAsin C.mpfr_asin
func MpfrAsin(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAtan C.mpfr_atan
func MpfrAtan(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSin C.mpfr_sin
func MpfrSin(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSinCos C.mpfr_sin_cos
func MpfrSinCos(MpfrPtr, MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCos C.mpfr_cos
func MpfrCos(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrTan C.mpfr_tan
func MpfrTan(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAtan2 C.mpfr_atan2
func MpfrAtan2(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSec C.mpfr_sec
func MpfrSec(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCsc C.mpfr_csc
func MpfrCsc(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCot C.mpfr_cot
func MpfrCot(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSinu C.mpfr_sinu
func MpfrSinu(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrCosu C.mpfr_cosu
func MpfrCosu(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrTanu C.mpfr_tanu
func MpfrTanu(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrAcosu C.mpfr_acosu
func MpfrAcosu(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrAsinu C.mpfr_asinu
func MpfrAsinu(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrAtanu C.mpfr_atanu
func MpfrAtanu(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrAtan2u C.mpfr_atan2u
func MpfrAtan2u(MpfrPtr, MpfrSrcptr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrAcospi C.mpfr_acospi
func MpfrAcospi(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAsinpi C.mpfr_asinpi
func MpfrAsinpi(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAtanpi C.mpfr_atanpi
func MpfrAtanpi(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAtan2pi C.mpfr_atan2pi
func MpfrAtan2pi(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSinpi C.mpfr_sinpi
func MpfrSinpi(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCospi C.mpfr_cospi
func MpfrCospi(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrTanpi C.mpfr_tanpi
func MpfrTanpi(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrHypot C.mpfr_hypot
func MpfrHypot(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrErf C.mpfr_erf
func MpfrErf(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrErfc C.mpfr_erfc
func MpfrErfc(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCbrt C.mpfr_cbrt
func MpfrCbrt(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrRoot C.mpfr_root
func MpfrRoot(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrRootnUi C.mpfr_rootn_ui
func MpfrRootnUi(MpfrPtr, MpfrSrcptr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrRootnSi C.mpfr_rootn_si
func MpfrRootnSi(MpfrPtr, MpfrSrcptr, c.Long, MpfrRndT) c.Int
//go:linkname MpfrGamma C.mpfr_gamma
func MpfrGamma(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrGammaInc C.mpfr_gamma_inc
func MpfrGammaInc(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrBeta C.mpfr_beta
func MpfrBeta(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrLngamma C.mpfr_lngamma
func MpfrLngamma(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrLgamma C.mpfr_lgamma
func MpfrLgamma(MpfrPtr, *c.Int, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrDigamma C.mpfr_digamma
func MpfrDigamma(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrZeta C.mpfr_zeta
func MpfrZeta(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrZetaUi C.mpfr_zeta_ui
func MpfrZetaUi(MpfrPtr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrFacUi C.mpfr_fac_ui
func MpfrFacUi(MpfrPtr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrJ0 C.mpfr_j0
func MpfrJ0(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrJ1 C.mpfr_j1
func MpfrJ1(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrJn C.mpfr_jn
func MpfrJn(MpfrPtr, c.Long, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrY0 C.mpfr_y0
func MpfrY0(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrY1 C.mpfr_y1
func MpfrY1(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrYn C.mpfr_yn
func MpfrYn(MpfrPtr, c.Long, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAi C.mpfr_ai
func MpfrAi(MpfrPtr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrMin C.mpfr_min
func MpfrMin(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrMax C.mpfr_max
func MpfrMax(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrDim C.mpfr_dim
func MpfrDim(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrMulZ C.mpfr_mul_z
func MpfrMulZ(MpfrPtr, MpfrSrcptr, MpzSrcptr, MpfrRndT) c.Int
//go:linkname MpfrDivZ C.mpfr_div_z
func MpfrDivZ(MpfrPtr, MpfrSrcptr, MpzSrcptr, MpfrRndT) c.Int
//go:linkname MpfrAddZ C.mpfr_add_z
func MpfrAddZ(MpfrPtr, MpfrSrcptr, MpzSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSubZ C.mpfr_sub_z
func MpfrSubZ(MpfrPtr, MpfrSrcptr, MpzSrcptr, MpfrRndT) c.Int
//go:linkname MpfrZSub C.mpfr_z_sub
func MpfrZSub(MpfrPtr, MpzSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrCmpZ C.mpfr_cmp_z
func MpfrCmpZ(MpfrSrcptr, MpzSrcptr) c.Int
//go:linkname MpfrFma C.mpfr_fma
func MpfrFma(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFms C.mpfr_fms
func MpfrFms(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFmma C.mpfr_fmma
func MpfrFmma(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrFmms C.mpfr_fmms
func MpfrFmms(MpfrPtr, MpfrSrcptr, MpfrSrcptr, MpfrSrcptr, MpfrSrcptr, MpfrRndT) c.Int
//go:linkname MpfrSum C.mpfr_sum
func MpfrSum(MpfrPtr, *MpfrPtr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrDot C.mpfr_dot
func MpfrDot(MpfrPtr, *MpfrPtr, *MpfrPtr, c.Ulong, MpfrRndT) c.Int
//go:linkname MpfrFreeCache C.mpfr_free_cache
func MpfrFreeCache()
// llgo:link MpfrFreeCacheT.MpfrFreeCache2 C.mpfr_free_cache2
func (recv_ MpfrFreeCacheT) MpfrFreeCache2() {
}
//go:linkname MpfrFreePool C.mpfr_free_pool
func MpfrFreePool()
//go:linkname MpfrMpMemoryCleanup C.mpfr_mp_memory_cleanup
func MpfrMpMemoryCleanup() c.Int
//go:linkname MpfrSubnormalize C.mpfr_subnormalize
func MpfrSubnormalize(MpfrPtr, c.Int, MpfrRndT) c.Int
//go:linkname MpfrStrtofr C.mpfr_strtofr
func MpfrStrtofr(MpfrPtr, *int8, **int8, c.Int, MpfrRndT) c.Int
//go:linkname MpfrRoundNearestAwayBegin C.mpfr_round_nearest_away_begin
func MpfrRoundNearestAwayBegin(MpfrPtr)
//go:linkname MpfrRoundNearestAwayEnd C.mpfr_round_nearest_away_end
func MpfrRoundNearestAwayEnd(MpfrPtr, c.Int) c.Int
// llgo:link MpfrPrecT.MpfrCustomGetSize C.mpfr_custom_get_size
func (recv_ MpfrPrecT) MpfrCustomGetSize() uintptr {
	return 0
}
//go:linkname MpfrCustomInit C.mpfr_custom_init
func MpfrCustomInit(unsafe.Pointer, MpfrPrecT)
//go:linkname MpfrCustomGetSignificand C.mpfr_custom_get_significand
func MpfrCustomGetSignificand(MpfrSrcptr) unsafe.Pointer
//go:linkname MpfrCustomGetExp C.mpfr_custom_get_exp
func MpfrCustomGetExp(MpfrSrcptr) MpfrExpT
//go:linkname MpfrCustomMove C.mpfr_custom_move
func MpfrCustomMove(MpfrPtr, unsafe.Pointer)
//go:linkname MpfrCustomInitSet C.mpfr_custom_init_set
func MpfrCustomInitSet(MpfrPtr, c.Int, MpfrExpT, MpfrPrecT, unsafe.Pointer)
//go:linkname MpfrCustomGetKind C.mpfr_custom_get_kind
func MpfrCustomGetKind(MpfrSrcptr) c.Int
//go:linkname MpfrTotalOrderP C.mpfr_total_order_p
func MpfrTotalOrderP(MpfrSrcptr, MpfrSrcptr) c.Int

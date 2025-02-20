package mpfr

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const X__GMPHAVEHOSTCPUFAMILYPower = 0
const X__GMPHAVEHOSTCPUFAMILYPowerpc = 0
const GMPLIMBBITS = 64
const GMPNAILBITS = 0
const X__GNUMP = 6
const X__GMPLIBGMPDLL = 0
const X__GMPMPSIZETINT = 0
const X__GMPINLINEPROTOTYPES = 1
const X__GMPCC = "gcc"
const X__GMPCFLAGS = " -m64 -O3"
const X__GNUMPVERSION = 6
const X__GNUMPVERSIONMINOR = 3
const X__GNUMPVERSIONPATCHLEVEL = 0

type MpLimbT c.Ulong
type MpLimbSignedT c.Long
type MpBitcntT c.Ulong

type X__mpzStruct struct {
	X_mpAlloc c.Int
	X_mpSize  c.Int
	X_mpD     *MpLimbT
}
type MPINT X__mpzStruct
type MpzT [1]X__mpzStruct
type MpPtr *MpLimbT
type MpSrcptr *MpLimbT
type MpSizeT c.Long
type MpExpT c.Long

type X__mpqStruct struct {
	X_mpNum X__mpzStruct
	X_mpDen X__mpzStruct
}
type MPRAT X__mpqStruct
type MpqT [1]X__mpqStruct

type X__mpfStruct struct {
	X_mpPrec c.Int
	X_mpSize c.Int
	X_mpExp  MpExpT
	X_mpD    *MpLimbT
}
type MpfT [1]X__mpfStruct
type GmpRandalgT c.Int

const (
	GmpRandalgTGMPRANDALGDEFAULT GmpRandalgT = 0
	GmpRandalgTGMPRANDALGLC      GmpRandalgT = 0
)

type X__gmpRandstateStruct struct {
	X_mpSeed    MpzT
	X_mpAlg     GmpRandalgT
	X_mpAlgdata struct {
		X_mpLc unsafe.Pointer
	}
}
type GmpRandstateT [1]X__gmpRandstateStruct
type MpzSrcptr *X__mpzStruct
type MpzPtr *X__mpzStruct
type MpfSrcptr *X__mpfStruct
type MpfPtr *X__mpfStruct
type MpqSrcptr *X__mpqStruct
type MpqPtr *X__mpqStruct
type GmpRandstatePtr *X__gmpRandstateStruct
type GmpRandstateSrcptr *X__gmpRandstateStruct
//go:linkname X__GmpSetMemoryFunctions C.__gmp_set_memory_functions
func X__GmpSetMemoryFunctions(func(uintptr) unsafe.Pointer, func(unsafe.Pointer, uintptr, uintptr) unsafe.Pointer, func(unsafe.Pointer, uintptr))
//go:linkname X__GmpGetMemoryFunctions C.__gmp_get_memory_functions
func X__GmpGetMemoryFunctions(func(uintptr) unsafe.Pointer, func(unsafe.Pointer, uintptr, uintptr) unsafe.Pointer, func(unsafe.Pointer, uintptr))
//go:linkname X__GmpRandinit C.__gmp_randinit
func X__GmpRandinit(__llgo_arg_0 GmpRandstatePtr, __llgo_arg_1 GmpRandalgT, __llgo_va_list ...interface{})
//go:linkname X__GmpRandinitDefault C.__gmp_randinit_default
func X__GmpRandinitDefault(GmpRandstatePtr)
//go:linkname X__GmpRandinitLc2exp C.__gmp_randinit_lc_2exp
func X__GmpRandinitLc2exp(GmpRandstatePtr, MpzSrcptr, c.Ulong, MpBitcntT)
//go:linkname X__GmpRandinitLc2expSize C.__gmp_randinit_lc_2exp_size
func X__GmpRandinitLc2expSize(GmpRandstatePtr, MpBitcntT) c.Int
//go:linkname X__GmpRandinitMt C.__gmp_randinit_mt
func X__GmpRandinitMt(GmpRandstatePtr)
//go:linkname X__GmpRandinitSet C.__gmp_randinit_set
func X__GmpRandinitSet(GmpRandstatePtr, GmpRandstateSrcptr)
//go:linkname X__GmpRandseed C.__gmp_randseed
func X__GmpRandseed(GmpRandstatePtr, MpzSrcptr)
//go:linkname X__GmpRandseedUi C.__gmp_randseed_ui
func X__GmpRandseedUi(GmpRandstatePtr, c.Ulong)
//go:linkname X__GmpRandclear C.__gmp_randclear
func X__GmpRandclear(GmpRandstatePtr)
//go:linkname X__GmpUrandombUi C.__gmp_urandomb_ui
func X__GmpUrandombUi(GmpRandstatePtr, c.Ulong) c.Ulong
//go:linkname X__GmpUrandommUi C.__gmp_urandomm_ui
func X__GmpUrandommUi(GmpRandstatePtr, c.Ulong) c.Ulong
//go:linkname X__GmpAsprintf C.__gmp_asprintf
func X__GmpAsprintf(__llgo_arg_0 **int8, __llgo_arg_1 *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname X__GmpPrintf C.__gmp_printf
func X__GmpPrintf(__llgo_arg_0 *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname X__GmpSnprintf C.__gmp_snprintf
func X__GmpSnprintf(__llgo_arg_0 *int8, __llgo_arg_1 uintptr, __llgo_arg_2 *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname X__GmpSprintf C.__gmp_sprintf
func X__GmpSprintf(__llgo_arg_0 *int8, __llgo_arg_1 *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname X__GmpScanf C.__gmp_scanf
func X__GmpScanf(__llgo_arg_0 *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname X__GmpSscanf C.__gmp_sscanf
func X__GmpSscanf(__llgo_arg_0 *int8, __llgo_arg_1 *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname X__GmpzRealloc C.__gmpz_realloc
func X__GmpzRealloc(MpzPtr, MpSizeT) unsafe.Pointer
//go:linkname X__GmpzAbs C.__gmpz_abs
func X__GmpzAbs(MpzPtr, MpzSrcptr)
//go:linkname X__GmpzAdd C.__gmpz_add
func X__GmpzAdd(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzAddUi C.__gmpz_add_ui
func X__GmpzAddUi(MpzPtr, MpzSrcptr, c.Ulong)
//go:linkname X__GmpzAddmul C.__gmpz_addmul
func X__GmpzAddmul(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzAddmulUi C.__gmpz_addmul_ui
func X__GmpzAddmulUi(MpzPtr, MpzSrcptr, c.Ulong)
//go:linkname X__GmpzAnd C.__gmpz_and
func X__GmpzAnd(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzArrayInit C.__gmpz_array_init
func X__GmpzArrayInit(MpzPtr, MpSizeT, MpSizeT)
//go:linkname X__GmpzBinUi C.__gmpz_bin_ui
func X__GmpzBinUi(MpzPtr, MpzSrcptr, c.Ulong)
//go:linkname X__GmpzBinUiui C.__gmpz_bin_uiui
func X__GmpzBinUiui(MpzPtr, c.Ulong, c.Ulong)
//go:linkname X__GmpzCdivQ C.__gmpz_cdiv_q
func X__GmpzCdivQ(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzCdivQ2exp C.__gmpz_cdiv_q_2exp
func X__GmpzCdivQ2exp(MpzPtr, MpzSrcptr, MpBitcntT)
//go:linkname X__GmpzCdivQUi C.__gmpz_cdiv_q_ui
func X__GmpzCdivQUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzCdivQr C.__gmpz_cdiv_qr
func X__GmpzCdivQr(MpzPtr, MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzCdivQrUi C.__gmpz_cdiv_qr_ui
func X__GmpzCdivQrUi(MpzPtr, MpzPtr, MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzCdivR C.__gmpz_cdiv_r
func X__GmpzCdivR(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzCdivR2exp C.__gmpz_cdiv_r_2exp
func X__GmpzCdivR2exp(MpzPtr, MpzSrcptr, MpBitcntT)
//go:linkname X__GmpzCdivRUi C.__gmpz_cdiv_r_ui
func X__GmpzCdivRUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzCdivUi C.__gmpz_cdiv_ui
func X__GmpzCdivUi(MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzClear C.__gmpz_clear
func X__GmpzClear(MpzPtr)
//go:linkname X__GmpzClears C.__gmpz_clears
func X__GmpzClears(__llgo_arg_0 MpzPtr, __llgo_va_list ...interface{})
//go:linkname X__GmpzClrbit C.__gmpz_clrbit
func X__GmpzClrbit(MpzPtr, MpBitcntT)
//go:linkname X__GmpzCmp C.__gmpz_cmp
func X__GmpzCmp(MpzSrcptr, MpzSrcptr) c.Int
//go:linkname X__GmpzCmpD C.__gmpz_cmp_d
func X__GmpzCmpD(MpzSrcptr, float64) c.Int
//go:linkname X__GmpzCmpSi C.__gmpz_cmp_si
func X__GmpzCmpSi(MpzSrcptr, c.Long) c.Int
//go:linkname X__GmpzCmpUi C.__gmpz_cmp_ui
func X__GmpzCmpUi(MpzSrcptr, c.Ulong) c.Int
//go:linkname X__GmpzCmpabs C.__gmpz_cmpabs
func X__GmpzCmpabs(MpzSrcptr, MpzSrcptr) c.Int
//go:linkname X__GmpzCmpabsD C.__gmpz_cmpabs_d
func X__GmpzCmpabsD(MpzSrcptr, float64) c.Int
//go:linkname X__GmpzCmpabsUi C.__gmpz_cmpabs_ui
func X__GmpzCmpabsUi(MpzSrcptr, c.Ulong) c.Int
//go:linkname X__GmpzCom C.__gmpz_com
func X__GmpzCom(MpzPtr, MpzSrcptr)
//go:linkname X__GmpzCombit C.__gmpz_combit
func X__GmpzCombit(MpzPtr, MpBitcntT)
//go:linkname X__GmpzCongruentP C.__gmpz_congruent_p
func X__GmpzCongruentP(MpzSrcptr, MpzSrcptr, MpzSrcptr) c.Int
//go:linkname X__GmpzCongruent2expP C.__gmpz_congruent_2exp_p
func X__GmpzCongruent2expP(MpzSrcptr, MpzSrcptr, MpBitcntT) c.Int
//go:linkname X__GmpzCongruentUiP C.__gmpz_congruent_ui_p
func X__GmpzCongruentUiP(MpzSrcptr, c.Ulong, c.Ulong) c.Int
//go:linkname X__GmpzDivexact C.__gmpz_divexact
func X__GmpzDivexact(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzDivexactUi C.__gmpz_divexact_ui
func X__GmpzDivexactUi(MpzPtr, MpzSrcptr, c.Ulong)
//go:linkname X__GmpzDivisibleP C.__gmpz_divisible_p
func X__GmpzDivisibleP(MpzSrcptr, MpzSrcptr) c.Int
//go:linkname X__GmpzDivisibleUiP C.__gmpz_divisible_ui_p
func X__GmpzDivisibleUiP(MpzSrcptr, c.Ulong) c.Int
//go:linkname X__GmpzDivisible2expP C.__gmpz_divisible_2exp_p
func X__GmpzDivisible2expP(MpzSrcptr, MpBitcntT) c.Int
//go:linkname X__GmpzDump C.__gmpz_dump
func X__GmpzDump(MpzSrcptr)
//go:linkname X__GmpzExport C.__gmpz_export
func X__GmpzExport(unsafe.Pointer, *uintptr, c.Int, uintptr, c.Int, uintptr, MpzSrcptr) unsafe.Pointer
//go:linkname X__GmpzFacUi C.__gmpz_fac_ui
func X__GmpzFacUi(MpzPtr, c.Ulong)
//go:linkname X__Gmpz2facUi C.__gmpz_2fac_ui
func X__Gmpz2facUi(MpzPtr, c.Ulong)
//go:linkname X__GmpzMfacUiui C.__gmpz_mfac_uiui
func X__GmpzMfacUiui(MpzPtr, c.Ulong, c.Ulong)
//go:linkname X__GmpzPrimorialUi C.__gmpz_primorial_ui
func X__GmpzPrimorialUi(MpzPtr, c.Ulong)
//go:linkname X__GmpzFdivQ C.__gmpz_fdiv_q
func X__GmpzFdivQ(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzFdivQ2exp C.__gmpz_fdiv_q_2exp
func X__GmpzFdivQ2exp(MpzPtr, MpzSrcptr, MpBitcntT)
//go:linkname X__GmpzFdivQUi C.__gmpz_fdiv_q_ui
func X__GmpzFdivQUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzFdivQr C.__gmpz_fdiv_qr
func X__GmpzFdivQr(MpzPtr, MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzFdivQrUi C.__gmpz_fdiv_qr_ui
func X__GmpzFdivQrUi(MpzPtr, MpzPtr, MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzFdivR C.__gmpz_fdiv_r
func X__GmpzFdivR(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzFdivR2exp C.__gmpz_fdiv_r_2exp
func X__GmpzFdivR2exp(MpzPtr, MpzSrcptr, MpBitcntT)
//go:linkname X__GmpzFdivRUi C.__gmpz_fdiv_r_ui
func X__GmpzFdivRUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzFdivUi C.__gmpz_fdiv_ui
func X__GmpzFdivUi(MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzFibUi C.__gmpz_fib_ui
func X__GmpzFibUi(MpzPtr, c.Ulong)
//go:linkname X__GmpzFib2Ui C.__gmpz_fib2_ui
func X__GmpzFib2Ui(MpzPtr, MpzPtr, c.Ulong)
//go:linkname X__GmpzFitsSintP C.__gmpz_fits_sint_p
func X__GmpzFitsSintP(MpzSrcptr) c.Int
//go:linkname X__GmpzFitsSlongP C.__gmpz_fits_slong_p
func X__GmpzFitsSlongP(MpzSrcptr) c.Int
//go:linkname X__GmpzFitsSshortP C.__gmpz_fits_sshort_p
func X__GmpzFitsSshortP(MpzSrcptr) c.Int
//go:linkname X__GmpzFitsUintP C.__gmpz_fits_uint_p
func X__GmpzFitsUintP(MpzSrcptr) c.Int
//go:linkname X__GmpzFitsUlongP C.__gmpz_fits_ulong_p
func X__GmpzFitsUlongP(MpzSrcptr) c.Int
//go:linkname X__GmpzFitsUshortP C.__gmpz_fits_ushort_p
func X__GmpzFitsUshortP(MpzSrcptr) c.Int
//go:linkname X__GmpzGcd C.__gmpz_gcd
func X__GmpzGcd(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzGcdUi C.__gmpz_gcd_ui
func X__GmpzGcdUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzGcdext C.__gmpz_gcdext
func X__GmpzGcdext(MpzPtr, MpzPtr, MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzGetD C.__gmpz_get_d
func X__GmpzGetD(MpzSrcptr) float64
//go:linkname X__GmpzGetD2exp C.__gmpz_get_d_2exp
func X__GmpzGetD2exp(*c.Long, MpzSrcptr) float64
//go:linkname X__GmpzGetSi C.__gmpz_get_si
func X__GmpzGetSi(MpzSrcptr) c.Long
//go:linkname X__GmpzGetStr C.__gmpz_get_str
func X__GmpzGetStr(*int8, c.Int, MpzSrcptr) *int8
//go:linkname X__GmpzGetUi C.__gmpz_get_ui
func X__GmpzGetUi(MpzSrcptr) c.Ulong
//go:linkname X__GmpzGetlimbn C.__gmpz_getlimbn
func X__GmpzGetlimbn(MpzSrcptr, MpSizeT) MpLimbT
//go:linkname X__GmpzHamdist C.__gmpz_hamdist
func X__GmpzHamdist(MpzSrcptr, MpzSrcptr) MpBitcntT
//go:linkname X__GmpzImport C.__gmpz_import
func X__GmpzImport(MpzPtr, uintptr, c.Int, uintptr, c.Int, uintptr, unsafe.Pointer)
//go:linkname X__GmpzInit C.__gmpz_init
func X__GmpzInit(MpzPtr)
//go:linkname X__GmpzInit2 C.__gmpz_init2
func X__GmpzInit2(MpzPtr, MpBitcntT)
//go:linkname X__GmpzInits C.__gmpz_inits
func X__GmpzInits(__llgo_arg_0 MpzPtr, __llgo_va_list ...interface{})
//go:linkname X__GmpzInitSet C.__gmpz_init_set
func X__GmpzInitSet(MpzPtr, MpzSrcptr)
//go:linkname X__GmpzInitSetD C.__gmpz_init_set_d
func X__GmpzInitSetD(MpzPtr, float64)
//go:linkname X__GmpzInitSetSi C.__gmpz_init_set_si
func X__GmpzInitSetSi(MpzPtr, c.Long)
//go:linkname X__GmpzInitSetStr C.__gmpz_init_set_str
func X__GmpzInitSetStr(MpzPtr, *int8, c.Int) c.Int
//go:linkname X__GmpzInitSetUi C.__gmpz_init_set_ui
func X__GmpzInitSetUi(MpzPtr, c.Ulong)
//go:linkname X__GmpzInvert C.__gmpz_invert
func X__GmpzInvert(MpzPtr, MpzSrcptr, MpzSrcptr) c.Int
//go:linkname X__GmpzIor C.__gmpz_ior
func X__GmpzIor(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzJacobi C.__gmpz_jacobi
func X__GmpzJacobi(MpzSrcptr, MpzSrcptr) c.Int
//go:linkname X__GmpzKroneckerSi C.__gmpz_kronecker_si
func X__GmpzKroneckerSi(MpzSrcptr, c.Long) c.Int
//go:linkname X__GmpzKroneckerUi C.__gmpz_kronecker_ui
func X__GmpzKroneckerUi(MpzSrcptr, c.Ulong) c.Int
//go:linkname X__GmpzSiKronecker C.__gmpz_si_kronecker
func X__GmpzSiKronecker(c.Long, MpzSrcptr) c.Int
//go:linkname X__GmpzUiKronecker C.__gmpz_ui_kronecker
func X__GmpzUiKronecker(c.Ulong, MpzSrcptr) c.Int
//go:linkname X__GmpzLcm C.__gmpz_lcm
func X__GmpzLcm(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzLcmUi C.__gmpz_lcm_ui
func X__GmpzLcmUi(MpzPtr, MpzSrcptr, c.Ulong)
//go:linkname X__GmpzLucnumUi C.__gmpz_lucnum_ui
func X__GmpzLucnumUi(MpzPtr, c.Ulong)
//go:linkname X__GmpzLucnum2Ui C.__gmpz_lucnum2_ui
func X__GmpzLucnum2Ui(MpzPtr, MpzPtr, c.Ulong)
//go:linkname X__GmpzMillerrabin C.__gmpz_millerrabin
func X__GmpzMillerrabin(MpzSrcptr, c.Int) c.Int
//go:linkname X__GmpzMod C.__gmpz_mod
func X__GmpzMod(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzMul C.__gmpz_mul
func X__GmpzMul(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzMul2exp C.__gmpz_mul_2exp
func X__GmpzMul2exp(MpzPtr, MpzSrcptr, MpBitcntT)
//go:linkname X__GmpzMulSi C.__gmpz_mul_si
func X__GmpzMulSi(MpzPtr, MpzSrcptr, c.Long)
//go:linkname X__GmpzMulUi C.__gmpz_mul_ui
func X__GmpzMulUi(MpzPtr, MpzSrcptr, c.Ulong)
//go:linkname X__GmpzNeg C.__gmpz_neg
func X__GmpzNeg(MpzPtr, MpzSrcptr)
//go:linkname X__GmpzNextprime C.__gmpz_nextprime
func X__GmpzNextprime(MpzPtr, MpzSrcptr)
//go:linkname X__GmpzPrevprime C.__gmpz_prevprime
func X__GmpzPrevprime(MpzPtr, MpzSrcptr) c.Int
//go:linkname X__GmpzPerfectPowerP C.__gmpz_perfect_power_p
func X__GmpzPerfectPowerP(MpzSrcptr) c.Int
//go:linkname X__GmpzPerfectSquareP C.__gmpz_perfect_square_p
func X__GmpzPerfectSquareP(MpzSrcptr) c.Int
//go:linkname X__GmpzPopcount C.__gmpz_popcount
func X__GmpzPopcount(MpzSrcptr) MpBitcntT
//go:linkname X__GmpzPowUi C.__gmpz_pow_ui
func X__GmpzPowUi(MpzPtr, MpzSrcptr, c.Ulong)
//go:linkname X__GmpzPowm C.__gmpz_powm
func X__GmpzPowm(MpzPtr, MpzSrcptr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzPowmSec C.__gmpz_powm_sec
func X__GmpzPowmSec(MpzPtr, MpzSrcptr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzPowmUi C.__gmpz_powm_ui
func X__GmpzPowmUi(MpzPtr, MpzSrcptr, c.Ulong, MpzSrcptr)
//go:linkname X__GmpzProbabPrimeP C.__gmpz_probab_prime_p
func X__GmpzProbabPrimeP(MpzSrcptr, c.Int) c.Int
//go:linkname X__GmpzRandom C.__gmpz_random
func X__GmpzRandom(MpzPtr, MpSizeT)
//go:linkname X__GmpzRandom2 C.__gmpz_random2
func X__GmpzRandom2(MpzPtr, MpSizeT)
//go:linkname X__GmpzRealloc2 C.__gmpz_realloc2
func X__GmpzRealloc2(MpzPtr, MpBitcntT)
//go:linkname X__GmpzRemove C.__gmpz_remove
func X__GmpzRemove(MpzPtr, MpzSrcptr, MpzSrcptr) MpBitcntT
//go:linkname X__GmpzRoot C.__gmpz_root
func X__GmpzRoot(MpzPtr, MpzSrcptr, c.Ulong) c.Int
//go:linkname X__GmpzRootrem C.__gmpz_rootrem
func X__GmpzRootrem(MpzPtr, MpzPtr, MpzSrcptr, c.Ulong)
//go:linkname X__GmpzRrandomb C.__gmpz_rrandomb
func X__GmpzRrandomb(MpzPtr, GmpRandstatePtr, MpBitcntT)
//go:linkname X__GmpzScan0 C.__gmpz_scan0
func X__GmpzScan0(MpzSrcptr, MpBitcntT) MpBitcntT
//go:linkname X__GmpzScan1 C.__gmpz_scan1
func X__GmpzScan1(MpzSrcptr, MpBitcntT) MpBitcntT
//go:linkname X__GmpzSet C.__gmpz_set
func X__GmpzSet(MpzPtr, MpzSrcptr)
//go:linkname X__GmpzSetD C.__gmpz_set_d
func X__GmpzSetD(MpzPtr, float64)
//go:linkname X__GmpzSetF C.__gmpz_set_f
func X__GmpzSetF(MpzPtr, MpfSrcptr)
//go:linkname X__GmpzSetQ C.__gmpz_set_q
func X__GmpzSetQ(MpzPtr, MpqSrcptr)
//go:linkname X__GmpzSetSi C.__gmpz_set_si
func X__GmpzSetSi(MpzPtr, c.Long)
//go:linkname X__GmpzSetStr C.__gmpz_set_str
func X__GmpzSetStr(MpzPtr, *int8, c.Int) c.Int
//go:linkname X__GmpzSetUi C.__gmpz_set_ui
func X__GmpzSetUi(MpzPtr, c.Ulong)
//go:linkname X__GmpzSetbit C.__gmpz_setbit
func X__GmpzSetbit(MpzPtr, MpBitcntT)
//go:linkname X__GmpzSize C.__gmpz_size
func X__GmpzSize(MpzSrcptr) uintptr
//go:linkname X__GmpzSizeinbase C.__gmpz_sizeinbase
func X__GmpzSizeinbase(MpzSrcptr, c.Int) uintptr
//go:linkname X__GmpzSqrt C.__gmpz_sqrt
func X__GmpzSqrt(MpzPtr, MpzSrcptr)
//go:linkname X__GmpzSqrtrem C.__gmpz_sqrtrem
func X__GmpzSqrtrem(MpzPtr, MpzPtr, MpzSrcptr)
//go:linkname X__GmpzSub C.__gmpz_sub
func X__GmpzSub(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzSubUi C.__gmpz_sub_ui
func X__GmpzSubUi(MpzPtr, MpzSrcptr, c.Ulong)
//go:linkname X__GmpzUiSub C.__gmpz_ui_sub
func X__GmpzUiSub(MpzPtr, c.Ulong, MpzSrcptr)
//go:linkname X__GmpzSubmul C.__gmpz_submul
func X__GmpzSubmul(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzSubmulUi C.__gmpz_submul_ui
func X__GmpzSubmulUi(MpzPtr, MpzSrcptr, c.Ulong)
//go:linkname X__GmpzSwap C.__gmpz_swap
func X__GmpzSwap(MpzPtr, MpzPtr)
//go:linkname X__GmpzTdivUi C.__gmpz_tdiv_ui
func X__GmpzTdivUi(MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzTdivQ C.__gmpz_tdiv_q
func X__GmpzTdivQ(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzTdivQ2exp C.__gmpz_tdiv_q_2exp
func X__GmpzTdivQ2exp(MpzPtr, MpzSrcptr, MpBitcntT)
//go:linkname X__GmpzTdivQUi C.__gmpz_tdiv_q_ui
func X__GmpzTdivQUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzTdivQr C.__gmpz_tdiv_qr
func X__GmpzTdivQr(MpzPtr, MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzTdivQrUi C.__gmpz_tdiv_qr_ui
func X__GmpzTdivQrUi(MpzPtr, MpzPtr, MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzTdivR C.__gmpz_tdiv_r
func X__GmpzTdivR(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzTdivR2exp C.__gmpz_tdiv_r_2exp
func X__GmpzTdivR2exp(MpzPtr, MpzSrcptr, MpBitcntT)
//go:linkname X__GmpzTdivRUi C.__gmpz_tdiv_r_ui
func X__GmpzTdivRUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong
//go:linkname X__GmpzTstbit C.__gmpz_tstbit
func X__GmpzTstbit(MpzSrcptr, MpBitcntT) c.Int
//go:linkname X__GmpzUiPowUi C.__gmpz_ui_pow_ui
func X__GmpzUiPowUi(MpzPtr, c.Ulong, c.Ulong)
//go:linkname X__GmpzUrandomb C.__gmpz_urandomb
func X__GmpzUrandomb(MpzPtr, GmpRandstatePtr, MpBitcntT)
//go:linkname X__GmpzUrandomm C.__gmpz_urandomm
func X__GmpzUrandomm(MpzPtr, GmpRandstatePtr, MpzSrcptr)
//go:linkname X__GmpzXor C.__gmpz_xor
func X__GmpzXor(MpzPtr, MpzSrcptr, MpzSrcptr)
//go:linkname X__GmpzLimbsRead C.__gmpz_limbs_read
func X__GmpzLimbsRead(MpzSrcptr) MpSrcptr
//go:linkname X__GmpzLimbsWrite C.__gmpz_limbs_write
func X__GmpzLimbsWrite(MpzPtr, MpSizeT) MpPtr
//go:linkname X__GmpzLimbsModify C.__gmpz_limbs_modify
func X__GmpzLimbsModify(MpzPtr, MpSizeT) MpPtr
//go:linkname X__GmpzLimbsFinish C.__gmpz_limbs_finish
func X__GmpzLimbsFinish(MpzPtr, MpSizeT)
//go:linkname X__GmpzRoinitN C.__gmpz_roinit_n
func X__GmpzRoinitN(MpzPtr, MpSrcptr, MpSizeT) MpzSrcptr
//go:linkname X__GmpqAbs C.__gmpq_abs
func X__GmpqAbs(MpqPtr, MpqSrcptr)
//go:linkname X__GmpqAdd C.__gmpq_add
func X__GmpqAdd(MpqPtr, MpqSrcptr, MpqSrcptr)
//go:linkname X__GmpqCanonicalize C.__gmpq_canonicalize
func X__GmpqCanonicalize(MpqPtr)
//go:linkname X__GmpqClear C.__gmpq_clear
func X__GmpqClear(MpqPtr)
//go:linkname X__GmpqClears C.__gmpq_clears
func X__GmpqClears(__llgo_arg_0 MpqPtr, __llgo_va_list ...interface{})
//go:linkname X__GmpqCmp C.__gmpq_cmp
func X__GmpqCmp(MpqSrcptr, MpqSrcptr) c.Int
//go:linkname X__GmpqCmpSi C.__gmpq_cmp_si
func X__GmpqCmpSi(MpqSrcptr, c.Long, c.Ulong) c.Int
//go:linkname X__GmpqCmpUi C.__gmpq_cmp_ui
func X__GmpqCmpUi(MpqSrcptr, c.Ulong, c.Ulong) c.Int
//go:linkname X__GmpqCmpZ C.__gmpq_cmp_z
func X__GmpqCmpZ(MpqSrcptr, MpzSrcptr) c.Int
//go:linkname X__GmpqDiv C.__gmpq_div
func X__GmpqDiv(MpqPtr, MpqSrcptr, MpqSrcptr)
//go:linkname X__GmpqDiv2exp C.__gmpq_div_2exp
func X__GmpqDiv2exp(MpqPtr, MpqSrcptr, MpBitcntT)
//go:linkname X__GmpqEqual C.__gmpq_equal
func X__GmpqEqual(MpqSrcptr, MpqSrcptr) c.Int
//go:linkname X__GmpqGetNum C.__gmpq_get_num
func X__GmpqGetNum(MpzPtr, MpqSrcptr)
//go:linkname X__GmpqGetDen C.__gmpq_get_den
func X__GmpqGetDen(MpzPtr, MpqSrcptr)
//go:linkname X__GmpqGetD C.__gmpq_get_d
func X__GmpqGetD(MpqSrcptr) float64
//go:linkname X__GmpqGetStr C.__gmpq_get_str
func X__GmpqGetStr(*int8, c.Int, MpqSrcptr) *int8
//go:linkname X__GmpqInit C.__gmpq_init
func X__GmpqInit(MpqPtr)
//go:linkname X__GmpqInits C.__gmpq_inits
func X__GmpqInits(__llgo_arg_0 MpqPtr, __llgo_va_list ...interface{})
//go:linkname X__GmpqInv C.__gmpq_inv
func X__GmpqInv(MpqPtr, MpqSrcptr)
//go:linkname X__GmpqMul C.__gmpq_mul
func X__GmpqMul(MpqPtr, MpqSrcptr, MpqSrcptr)
//go:linkname X__GmpqMul2exp C.__gmpq_mul_2exp
func X__GmpqMul2exp(MpqPtr, MpqSrcptr, MpBitcntT)
//go:linkname X__GmpqNeg C.__gmpq_neg
func X__GmpqNeg(MpqPtr, MpqSrcptr)
//go:linkname X__GmpqSet C.__gmpq_set
func X__GmpqSet(MpqPtr, MpqSrcptr)
//go:linkname X__GmpqSetD C.__gmpq_set_d
func X__GmpqSetD(MpqPtr, float64)
//go:linkname X__GmpqSetDen C.__gmpq_set_den
func X__GmpqSetDen(MpqPtr, MpzSrcptr)
//go:linkname X__GmpqSetF C.__gmpq_set_f
func X__GmpqSetF(MpqPtr, MpfSrcptr)
//go:linkname X__GmpqSetNum C.__gmpq_set_num
func X__GmpqSetNum(MpqPtr, MpzSrcptr)
//go:linkname X__GmpqSetSi C.__gmpq_set_si
func X__GmpqSetSi(MpqPtr, c.Long, c.Ulong)
//go:linkname X__GmpqSetStr C.__gmpq_set_str
func X__GmpqSetStr(MpqPtr, *int8, c.Int) c.Int
//go:linkname X__GmpqSetUi C.__gmpq_set_ui
func X__GmpqSetUi(MpqPtr, c.Ulong, c.Ulong)
//go:linkname X__GmpqSetZ C.__gmpq_set_z
func X__GmpqSetZ(MpqPtr, MpzSrcptr)
//go:linkname X__GmpqSub C.__gmpq_sub
func X__GmpqSub(MpqPtr, MpqSrcptr, MpqSrcptr)
//go:linkname X__GmpqSwap C.__gmpq_swap
func X__GmpqSwap(MpqPtr, MpqPtr)
//go:linkname X__GmpfAbs C.__gmpf_abs
func X__GmpfAbs(MpfPtr, MpfSrcptr)
//go:linkname X__GmpfAdd C.__gmpf_add
func X__GmpfAdd(MpfPtr, MpfSrcptr, MpfSrcptr)
//go:linkname X__GmpfAddUi C.__gmpf_add_ui
func X__GmpfAddUi(MpfPtr, MpfSrcptr, c.Ulong)
//go:linkname X__GmpfCeil C.__gmpf_ceil
func X__GmpfCeil(MpfPtr, MpfSrcptr)
//go:linkname X__GmpfClear C.__gmpf_clear
func X__GmpfClear(MpfPtr)
//go:linkname X__GmpfClears C.__gmpf_clears
func X__GmpfClears(__llgo_arg_0 MpfPtr, __llgo_va_list ...interface{})
//go:linkname X__GmpfCmp C.__gmpf_cmp
func X__GmpfCmp(MpfSrcptr, MpfSrcptr) c.Int
//go:linkname X__GmpfCmpZ C.__gmpf_cmp_z
func X__GmpfCmpZ(MpfSrcptr, MpzSrcptr) c.Int
//go:linkname X__GmpfCmpD C.__gmpf_cmp_d
func X__GmpfCmpD(MpfSrcptr, float64) c.Int
//go:linkname X__GmpfCmpSi C.__gmpf_cmp_si
func X__GmpfCmpSi(MpfSrcptr, c.Long) c.Int
//go:linkname X__GmpfCmpUi C.__gmpf_cmp_ui
func X__GmpfCmpUi(MpfSrcptr, c.Ulong) c.Int
//go:linkname X__GmpfDiv C.__gmpf_div
func X__GmpfDiv(MpfPtr, MpfSrcptr, MpfSrcptr)
//go:linkname X__GmpfDiv2exp C.__gmpf_div_2exp
func X__GmpfDiv2exp(MpfPtr, MpfSrcptr, MpBitcntT)
//go:linkname X__GmpfDivUi C.__gmpf_div_ui
func X__GmpfDivUi(MpfPtr, MpfSrcptr, c.Ulong)
//go:linkname X__GmpfDump C.__gmpf_dump
func X__GmpfDump(MpfSrcptr)
//go:linkname X__GmpfEq C.__gmpf_eq
func X__GmpfEq(MpfSrcptr, MpfSrcptr, MpBitcntT) c.Int
//go:linkname X__GmpfFitsSintP C.__gmpf_fits_sint_p
func X__GmpfFitsSintP(MpfSrcptr) c.Int
//go:linkname X__GmpfFitsSlongP C.__gmpf_fits_slong_p
func X__GmpfFitsSlongP(MpfSrcptr) c.Int
//go:linkname X__GmpfFitsSshortP C.__gmpf_fits_sshort_p
func X__GmpfFitsSshortP(MpfSrcptr) c.Int
//go:linkname X__GmpfFitsUintP C.__gmpf_fits_uint_p
func X__GmpfFitsUintP(MpfSrcptr) c.Int
//go:linkname X__GmpfFitsUlongP C.__gmpf_fits_ulong_p
func X__GmpfFitsUlongP(MpfSrcptr) c.Int
//go:linkname X__GmpfFitsUshortP C.__gmpf_fits_ushort_p
func X__GmpfFitsUshortP(MpfSrcptr) c.Int
//go:linkname X__GmpfFloor C.__gmpf_floor
func X__GmpfFloor(MpfPtr, MpfSrcptr)
//go:linkname X__GmpfGetD C.__gmpf_get_d
func X__GmpfGetD(MpfSrcptr) float64
//go:linkname X__GmpfGetD2exp C.__gmpf_get_d_2exp
func X__GmpfGetD2exp(*c.Long, MpfSrcptr) float64
//go:linkname X__GmpfGetDefaultPrec C.__gmpf_get_default_prec
func X__GmpfGetDefaultPrec() MpBitcntT
//go:linkname X__GmpfGetPrec C.__gmpf_get_prec
func X__GmpfGetPrec(MpfSrcptr) MpBitcntT
//go:linkname X__GmpfGetSi C.__gmpf_get_si
func X__GmpfGetSi(MpfSrcptr) c.Long
//go:linkname X__GmpfGetStr C.__gmpf_get_str
func X__GmpfGetStr(*int8, *MpExpT, c.Int, uintptr, MpfSrcptr) *int8
//go:linkname X__GmpfGetUi C.__gmpf_get_ui
func X__GmpfGetUi(MpfSrcptr) c.Ulong
//go:linkname X__GmpfInit C.__gmpf_init
func X__GmpfInit(MpfPtr)
//go:linkname X__GmpfInit2 C.__gmpf_init2
func X__GmpfInit2(MpfPtr, MpBitcntT)
//go:linkname X__GmpfInits C.__gmpf_inits
func X__GmpfInits(__llgo_arg_0 MpfPtr, __llgo_va_list ...interface{})
//go:linkname X__GmpfInitSet C.__gmpf_init_set
func X__GmpfInitSet(MpfPtr, MpfSrcptr)
//go:linkname X__GmpfInitSetD C.__gmpf_init_set_d
func X__GmpfInitSetD(MpfPtr, float64)
//go:linkname X__GmpfInitSetSi C.__gmpf_init_set_si
func X__GmpfInitSetSi(MpfPtr, c.Long)
//go:linkname X__GmpfInitSetStr C.__gmpf_init_set_str
func X__GmpfInitSetStr(MpfPtr, *int8, c.Int) c.Int
//go:linkname X__GmpfInitSetUi C.__gmpf_init_set_ui
func X__GmpfInitSetUi(MpfPtr, c.Ulong)
//go:linkname X__GmpfIntegerP C.__gmpf_integer_p
func X__GmpfIntegerP(MpfSrcptr) c.Int
//go:linkname X__GmpfMul C.__gmpf_mul
func X__GmpfMul(MpfPtr, MpfSrcptr, MpfSrcptr)
//go:linkname X__GmpfMul2exp C.__gmpf_mul_2exp
func X__GmpfMul2exp(MpfPtr, MpfSrcptr, MpBitcntT)
//go:linkname X__GmpfMulUi C.__gmpf_mul_ui
func X__GmpfMulUi(MpfPtr, MpfSrcptr, c.Ulong)
//go:linkname X__GmpfNeg C.__gmpf_neg
func X__GmpfNeg(MpfPtr, MpfSrcptr)
//go:linkname X__GmpfPowUi C.__gmpf_pow_ui
func X__GmpfPowUi(MpfPtr, MpfSrcptr, c.Ulong)
//go:linkname X__GmpfRandom2 C.__gmpf_random2
func X__GmpfRandom2(MpfPtr, MpSizeT, MpExpT)
//go:linkname X__GmpfReldiff C.__gmpf_reldiff
func X__GmpfReldiff(MpfPtr, MpfSrcptr, MpfSrcptr)
//go:linkname X__GmpfSet C.__gmpf_set
func X__GmpfSet(MpfPtr, MpfSrcptr)
//go:linkname X__GmpfSetD C.__gmpf_set_d
func X__GmpfSetD(MpfPtr, float64)
// llgo:link MpBitcntT.X__GmpfSetDefaultPrec C.__gmpf_set_default_prec
func (recv_ MpBitcntT) X__GmpfSetDefaultPrec() {
}
//go:linkname X__GmpfSetPrec C.__gmpf_set_prec
func X__GmpfSetPrec(MpfPtr, MpBitcntT)
//go:linkname X__GmpfSetPrecRaw C.__gmpf_set_prec_raw
func X__GmpfSetPrecRaw(MpfPtr, MpBitcntT)
//go:linkname X__GmpfSetQ C.__gmpf_set_q
func X__GmpfSetQ(MpfPtr, MpqSrcptr)
//go:linkname X__GmpfSetSi C.__gmpf_set_si
func X__GmpfSetSi(MpfPtr, c.Long)
//go:linkname X__GmpfSetStr C.__gmpf_set_str
func X__GmpfSetStr(MpfPtr, *int8, c.Int) c.Int
//go:linkname X__GmpfSetUi C.__gmpf_set_ui
func X__GmpfSetUi(MpfPtr, c.Ulong)
//go:linkname X__GmpfSetZ C.__gmpf_set_z
func X__GmpfSetZ(MpfPtr, MpzSrcptr)
//go:linkname X__GmpfSize C.__gmpf_size
func X__GmpfSize(MpfSrcptr) uintptr
//go:linkname X__GmpfSqrt C.__gmpf_sqrt
func X__GmpfSqrt(MpfPtr, MpfSrcptr)
//go:linkname X__GmpfSqrtUi C.__gmpf_sqrt_ui
func X__GmpfSqrtUi(MpfPtr, c.Ulong)
//go:linkname X__GmpfSub C.__gmpf_sub
func X__GmpfSub(MpfPtr, MpfSrcptr, MpfSrcptr)
//go:linkname X__GmpfSubUi C.__gmpf_sub_ui
func X__GmpfSubUi(MpfPtr, MpfSrcptr, c.Ulong)
//go:linkname X__GmpfSwap C.__gmpf_swap
func X__GmpfSwap(MpfPtr, MpfPtr)
//go:linkname X__GmpfTrunc C.__gmpf_trunc
func X__GmpfTrunc(MpfPtr, MpfSrcptr)
//go:linkname X__GmpfUiDiv C.__gmpf_ui_div
func X__GmpfUiDiv(MpfPtr, c.Ulong, MpfSrcptr)
//go:linkname X__GmpfUiSub C.__gmpf_ui_sub
func X__GmpfUiSub(MpfPtr, c.Ulong, MpfSrcptr)
//go:linkname X__GmpfUrandomb C.__gmpf_urandomb
func X__GmpfUrandomb(MpfPtr, GmpRandstatePtr, MpBitcntT)
//go:linkname X__GmpnAdd C.__gmpn_add
func X__GmpnAdd(MpPtr, MpSrcptr, MpSizeT, MpSrcptr, MpSizeT) MpLimbT
//go:linkname X__GmpnAdd1 C.__gmpn_add_1
func X__GmpnAdd1(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT
//go:linkname X__GmpnAddN C.__gmpn_add_n
func X__GmpnAddN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT) MpLimbT
//go:linkname X__GmpnAddmul1 C.__gmpn_addmul_1
func X__GmpnAddmul1(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT
//go:linkname X__GmpnCmp C.__gmpn_cmp
func X__GmpnCmp(MpSrcptr, MpSrcptr, MpSizeT) c.Int
//go:linkname X__GmpnZeroP C.__gmpn_zero_p
func X__GmpnZeroP(MpSrcptr, MpSizeT) c.Int
//go:linkname X__GmpnDivexact1 C.__gmpn_divexact_1
func X__GmpnDivexact1(MpPtr, MpSrcptr, MpSizeT, MpLimbT)
//go:linkname X__GmpnDivexactBy3c C.__gmpn_divexact_by3c
func X__GmpnDivexactBy3c(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT
//go:linkname X__GmpnDivrem C.__gmpn_divrem
func X__GmpnDivrem(MpPtr, MpSizeT, MpPtr, MpSizeT, MpSrcptr, MpSizeT) MpLimbT
//go:linkname X__GmpnDivrem1 C.__gmpn_divrem_1
func X__GmpnDivrem1(MpPtr, MpSizeT, MpSrcptr, MpSizeT, MpLimbT) MpLimbT
//go:linkname X__GmpnDivrem2 C.__gmpn_divrem_2
func X__GmpnDivrem2(MpPtr, MpSizeT, MpPtr, MpSizeT, MpSrcptr) MpLimbT
//go:linkname X__GmpnDivQr1 C.__gmpn_div_qr_1
func X__GmpnDivQr1(MpPtr, *MpLimbT, MpSrcptr, MpSizeT, MpLimbT) MpLimbT
//go:linkname X__GmpnDivQr2 C.__gmpn_div_qr_2
func X__GmpnDivQr2(MpPtr, MpPtr, MpSrcptr, MpSizeT, MpSrcptr) MpLimbT
//go:linkname X__GmpnGcd C.__gmpn_gcd
func X__GmpnGcd(MpPtr, MpPtr, MpSizeT, MpPtr, MpSizeT) MpSizeT
// llgo:link MpLimbT.X__GmpnGcd11 C.__gmpn_gcd_11
func (recv_ MpLimbT) X__GmpnGcd11(MpLimbT) MpLimbT {
	return 0
}
//go:linkname X__GmpnGcd1 C.__gmpn_gcd_1
func X__GmpnGcd1(MpSrcptr, MpSizeT, MpLimbT) MpLimbT
// llgo:link (*MpLimbSignedT).X__GmpnGcdext1 C.__gmpn_gcdext_1
func (recv_ *MpLimbSignedT) X__GmpnGcdext1(*MpLimbSignedT, MpLimbT, MpLimbT) MpLimbT {
	return 0
}
//go:linkname X__GmpnGcdext C.__gmpn_gcdext
func X__GmpnGcdext(MpPtr, MpPtr, *MpSizeT, MpPtr, MpSizeT, MpPtr, MpSizeT) MpSizeT
//go:linkname X__GmpnGetStr C.__gmpn_get_str
func X__GmpnGetStr(*int8, c.Int, MpPtr, MpSizeT) uintptr
//go:linkname X__GmpnHamdist C.__gmpn_hamdist
func X__GmpnHamdist(MpSrcptr, MpSrcptr, MpSizeT) MpBitcntT
//go:linkname X__GmpnLshift C.__gmpn_lshift
func X__GmpnLshift(MpPtr, MpSrcptr, MpSizeT, c.Uint) MpLimbT
//go:linkname X__GmpnMod1 C.__gmpn_mod_1
func X__GmpnMod1(MpSrcptr, MpSizeT, MpLimbT) MpLimbT
//go:linkname X__GmpnMul C.__gmpn_mul
func X__GmpnMul(MpPtr, MpSrcptr, MpSizeT, MpSrcptr, MpSizeT) MpLimbT
//go:linkname X__GmpnMul1 C.__gmpn_mul_1
func X__GmpnMul1(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT
//go:linkname X__GmpnMulN C.__gmpn_mul_n
func X__GmpnMulN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnSqr C.__gmpn_sqr
func X__GmpnSqr(MpPtr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnNeg C.__gmpn_neg
func X__GmpnNeg(MpPtr, MpSrcptr, MpSizeT) MpLimbT
//go:linkname X__GmpnCom C.__gmpn_com
func X__GmpnCom(MpPtr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnPerfectSquareP C.__gmpn_perfect_square_p
func X__GmpnPerfectSquareP(MpSrcptr, MpSizeT) c.Int
//go:linkname X__GmpnPerfectPowerP C.__gmpn_perfect_power_p
func X__GmpnPerfectPowerP(MpSrcptr, MpSizeT) c.Int
//go:linkname X__GmpnPopcount C.__gmpn_popcount
func X__GmpnPopcount(MpSrcptr, MpSizeT) MpBitcntT
//go:linkname X__GmpnPow1 C.__gmpn_pow_1
func X__GmpnPow1(MpPtr, MpSrcptr, MpSizeT, MpLimbT, MpPtr) MpSizeT
//go:linkname X__GmpnPreinvMod1 C.__gmpn_preinv_mod_1
func X__GmpnPreinvMod1(MpSrcptr, MpSizeT, MpLimbT, MpLimbT) MpLimbT
//go:linkname X__GmpnRandom C.__gmpn_random
func X__GmpnRandom(MpPtr, MpSizeT)
//go:linkname X__GmpnRandom2 C.__gmpn_random2
func X__GmpnRandom2(MpPtr, MpSizeT)
//go:linkname X__GmpnRshift C.__gmpn_rshift
func X__GmpnRshift(MpPtr, MpSrcptr, MpSizeT, c.Uint) MpLimbT
//go:linkname X__GmpnScan0 C.__gmpn_scan0
func X__GmpnScan0(MpSrcptr, MpBitcntT) MpBitcntT
//go:linkname X__GmpnScan1 C.__gmpn_scan1
func X__GmpnScan1(MpSrcptr, MpBitcntT) MpBitcntT
//go:linkname X__GmpnSetStr C.__gmpn_set_str
func X__GmpnSetStr(MpPtr, *int8, uintptr, c.Int) MpSizeT
//go:linkname X__GmpnSizeinbase C.__gmpn_sizeinbase
func X__GmpnSizeinbase(MpSrcptr, MpSizeT, c.Int) uintptr
//go:linkname X__GmpnSqrtrem C.__gmpn_sqrtrem
func X__GmpnSqrtrem(MpPtr, MpPtr, MpSrcptr, MpSizeT) MpSizeT
//go:linkname X__GmpnSub C.__gmpn_sub
func X__GmpnSub(MpPtr, MpSrcptr, MpSizeT, MpSrcptr, MpSizeT) MpLimbT
//go:linkname X__GmpnSub1 C.__gmpn_sub_1
func X__GmpnSub1(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT
//go:linkname X__GmpnSubN C.__gmpn_sub_n
func X__GmpnSubN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT) MpLimbT
//go:linkname X__GmpnSubmul1 C.__gmpn_submul_1
func X__GmpnSubmul1(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT
//go:linkname X__GmpnTdivQr C.__gmpn_tdiv_qr
func X__GmpnTdivQr(MpPtr, MpPtr, MpSizeT, MpSrcptr, MpSizeT, MpSrcptr, MpSizeT)
//go:linkname X__GmpnAndN C.__gmpn_and_n
func X__GmpnAndN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnAndnN C.__gmpn_andn_n
func X__GmpnAndnN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnNandN C.__gmpn_nand_n
func X__GmpnNandN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnIorN C.__gmpn_ior_n
func X__GmpnIorN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnIornN C.__gmpn_iorn_n
func X__GmpnIornN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnNiorN C.__gmpn_nior_n
func X__GmpnNiorN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnXorN C.__gmpn_xor_n
func X__GmpnXorN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnXnorN C.__gmpn_xnor_n
func X__GmpnXnorN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnCopyi C.__gmpn_copyi
func X__GmpnCopyi(MpPtr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnCopyd C.__gmpn_copyd
func X__GmpnCopyd(MpPtr, MpSrcptr, MpSizeT)
//go:linkname X__GmpnZero C.__gmpn_zero
func X__GmpnZero(MpPtr, MpSizeT)
// llgo:link MpLimbT.X__GmpnCndAddN C.__gmpn_cnd_add_n
func (recv_ MpLimbT) X__GmpnCndAddN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT) MpLimbT {
	return 0
}
// llgo:link MpLimbT.X__GmpnCndSubN C.__gmpn_cnd_sub_n
func (recv_ MpLimbT) X__GmpnCndSubN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT) MpLimbT {
	return 0
}
//go:linkname X__GmpnSecAdd1 C.__gmpn_sec_add_1
func X__GmpnSecAdd1(MpPtr, MpSrcptr, MpSizeT, MpLimbT, MpPtr) MpLimbT
// llgo:link MpSizeT.X__GmpnSecAdd1Itch C.__gmpn_sec_add_1_itch
func (recv_ MpSizeT) X__GmpnSecAdd1Itch() MpSizeT {
	return 0
}
//go:linkname X__GmpnSecSub1 C.__gmpn_sec_sub_1
func X__GmpnSecSub1(MpPtr, MpSrcptr, MpSizeT, MpLimbT, MpPtr) MpLimbT
// llgo:link MpSizeT.X__GmpnSecSub1Itch C.__gmpn_sec_sub_1_itch
func (recv_ MpSizeT) X__GmpnSecSub1Itch() MpSizeT {
	return 0
}
// llgo:link MpLimbT.X__GmpnCndSwap C.__gmpn_cnd_swap
func (recv_ MpLimbT) X__GmpnCndSwap(*MpLimbT, *MpLimbT, MpSizeT) {
}
//go:linkname X__GmpnSecMul C.__gmpn_sec_mul
func X__GmpnSecMul(MpPtr, MpSrcptr, MpSizeT, MpSrcptr, MpSizeT, MpPtr)
// llgo:link MpSizeT.X__GmpnSecMulItch C.__gmpn_sec_mul_itch
func (recv_ MpSizeT) X__GmpnSecMulItch(MpSizeT) MpSizeT {
	return 0
}
//go:linkname X__GmpnSecSqr C.__gmpn_sec_sqr
func X__GmpnSecSqr(MpPtr, MpSrcptr, MpSizeT, MpPtr)
// llgo:link MpSizeT.X__GmpnSecSqrItch C.__gmpn_sec_sqr_itch
func (recv_ MpSizeT) X__GmpnSecSqrItch() MpSizeT {
	return 0
}
//go:linkname X__GmpnSecPowm C.__gmpn_sec_powm
func X__GmpnSecPowm(MpPtr, MpSrcptr, MpSizeT, MpSrcptr, MpBitcntT, MpSrcptr, MpSizeT, MpPtr)
// llgo:link MpSizeT.X__GmpnSecPowmItch C.__gmpn_sec_powm_itch
func (recv_ MpSizeT) X__GmpnSecPowmItch(MpBitcntT, MpSizeT) MpSizeT {
	return 0
}
// llgo:link (*MpLimbT).X__GmpnSecTabselect C.__gmpn_sec_tabselect
func (recv_ *MpLimbT) X__GmpnSecTabselect(*MpLimbT, MpSizeT, MpSizeT, MpSizeT) {
}
//go:linkname X__GmpnSecDivQr C.__gmpn_sec_div_qr
func X__GmpnSecDivQr(MpPtr, MpPtr, MpSizeT, MpSrcptr, MpSizeT, MpPtr) MpLimbT
// llgo:link MpSizeT.X__GmpnSecDivQrItch C.__gmpn_sec_div_qr_itch
func (recv_ MpSizeT) X__GmpnSecDivQrItch(MpSizeT) MpSizeT {
	return 0
}
//go:linkname X__GmpnSecDivR C.__gmpn_sec_div_r
func X__GmpnSecDivR(MpPtr, MpSizeT, MpSrcptr, MpSizeT, MpPtr)
// llgo:link MpSizeT.X__GmpnSecDivRItch C.__gmpn_sec_div_r_itch
func (recv_ MpSizeT) X__GmpnSecDivRItch(MpSizeT) MpSizeT {
	return 0
}
//go:linkname X__GmpnSecInvert C.__gmpn_sec_invert
func X__GmpnSecInvert(MpPtr, MpPtr, MpSrcptr, MpSizeT, MpBitcntT, MpPtr) c.Int
// llgo:link MpSizeT.X__GmpnSecInvertItch C.__gmpn_sec_invert_itch
func (recv_ MpSizeT) X__GmpnSecInvertItch() MpSizeT {
	return 0
}

const (
	GMPERRORNONE                c.Int = 0
	GMPERRORUNSUPPORTEDARGUMENT c.Int = 1
	GMPERRORDIVISIONBYZERO      c.Int = 2
	GMPERRORSQRTOFNEGATIVE      c.Int = 4
	GMPERRORINVALIDARGUMENT     c.Int = 8
	GMPERRORMPZOVERFLOW         c.Int = 16
)

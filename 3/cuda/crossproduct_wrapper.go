package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
	"sync"
	"unsafe"
)

// CUDA handle for crossproduct kernel
var crossproduct_code cu.Function

// Stores the arguments for crossproduct kernel invocation
type crossproduct_args_t struct {
	arg_dstx unsafe.Pointer
	arg_dsty unsafe.Pointer
	arg_dstz unsafe.Pointer
	arg_ax   unsafe.Pointer
	arg_ay   unsafe.Pointer
	arg_az   unsafe.Pointer
	arg_bx   unsafe.Pointer
	arg_by   unsafe.Pointer
	arg_bz   unsafe.Pointer
	arg_N    int
	argptr   [10]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for crossproduct kernel invocation
var crossproduct_args crossproduct_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	crossproduct_args.argptr[0] = unsafe.Pointer(&crossproduct_args.arg_dstx)
	crossproduct_args.argptr[1] = unsafe.Pointer(&crossproduct_args.arg_dsty)
	crossproduct_args.argptr[2] = unsafe.Pointer(&crossproduct_args.arg_dstz)
	crossproduct_args.argptr[3] = unsafe.Pointer(&crossproduct_args.arg_ax)
	crossproduct_args.argptr[4] = unsafe.Pointer(&crossproduct_args.arg_ay)
	crossproduct_args.argptr[5] = unsafe.Pointer(&crossproduct_args.arg_az)
	crossproduct_args.argptr[6] = unsafe.Pointer(&crossproduct_args.arg_bx)
	crossproduct_args.argptr[7] = unsafe.Pointer(&crossproduct_args.arg_by)
	crossproduct_args.argptr[8] = unsafe.Pointer(&crossproduct_args.arg_bz)
	crossproduct_args.argptr[9] = unsafe.Pointer(&crossproduct_args.arg_N)
}

// Wrapper for crossproduct CUDA kernel, asynchronous.
func k_crossproduct_async(dstx unsafe.Pointer, dsty unsafe.Pointer, dstz unsafe.Pointer, ax unsafe.Pointer, ay unsafe.Pointer, az unsafe.Pointer, bx unsafe.Pointer, by unsafe.Pointer, bz unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("crossproduct")
	}

	crossproduct_args.Lock()
	defer crossproduct_args.Unlock()

	if crossproduct_code == 0 {
		crossproduct_code = fatbinLoad(crossproduct_map, "crossproduct")
	}

	crossproduct_args.arg_dstx = dstx
	crossproduct_args.arg_dsty = dsty
	crossproduct_args.arg_dstz = dstz
	crossproduct_args.arg_ax = ax
	crossproduct_args.arg_ay = ay
	crossproduct_args.arg_az = az
	crossproduct_args.arg_bx = bx
	crossproduct_args.arg_by = by
	crossproduct_args.arg_bz = bz
	crossproduct_args.arg_N = N

	args := crossproduct_args.argptr[:]
	cu.LaunchKernel(crossproduct_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("crossproduct")
	}
}

// maps compute capability on PTX code for crossproduct kernel.
var crossproduct_map = map[int]string{0: "",
	30: crossproduct_ptx_30,
	32: crossproduct_ptx_32,
	35: crossproduct_ptx_35,
	37: crossproduct_ptx_37,
	50: crossproduct_ptx_50,
	52: crossproduct_ptx_52,
	53: crossproduct_ptx_53,
	60: crossproduct_ptx_60,
	61: crossproduct_ptx_61,
	62: crossproduct_ptx_62,
	70: crossproduct_ptx_70,
	72: crossproduct_ptx_72,
	75: crossproduct_ptx_75}

// crossproduct PTX code for various compute capabilities.
const (
	crossproduct_ptx_30 = `
.version 6.5
.target sm_30
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.f32 	%f1, [%rd22];
	ld.global.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.f32 	%f4, [%rd20];
	ld.global.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_32 = `
.version 6.5
.target sm_32
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_35 = `
.version 6.5
.target sm_35
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_37 = `
.version 6.5
.target sm_37
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_50 = `
.version 6.5
.target sm_50
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_52 = `
.version 6.5
.target sm_52
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_53 = `
.version 6.5
.target sm_53
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_60 = `
.version 6.5
.target sm_60
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_61 = `
.version 6.5
.target sm_61
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_62 = `
.version 6.5
.target sm_62
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_70 = `
.version 6.5
.target sm_70
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_72 = `
.version 6.5
.target sm_72
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
	crossproduct_ptx_75 = `
.version 6.5
.target sm_75
.address_size 64

	// .globl	crossproduct

.visible .entry crossproduct(
	.param .u64 crossproduct_param_0,
	.param .u64 crossproduct_param_1,
	.param .u64 crossproduct_param_2,
	.param .u64 crossproduct_param_3,
	.param .u64 crossproduct_param_4,
	.param .u64 crossproduct_param_5,
	.param .u64 crossproduct_param_6,
	.param .u64 crossproduct_param_7,
	.param .u64 crossproduct_param_8,
	.param .u32 crossproduct_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<29>;


	ld.param.u64 	%rd1, [crossproduct_param_0];
	ld.param.u64 	%rd2, [crossproduct_param_1];
	ld.param.u64 	%rd3, [crossproduct_param_2];
	ld.param.u64 	%rd4, [crossproduct_param_3];
	ld.param.u64 	%rd5, [crossproduct_param_4];
	ld.param.u64 	%rd6, [crossproduct_param_5];
	ld.param.u64 	%rd7, [crossproduct_param_6];
	ld.param.u64 	%rd8, [crossproduct_param_7];
	ld.param.u64 	%rd9, [crossproduct_param_8];
	ld.param.u32 	%r2, [crossproduct_param_9];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd4;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	cvta.to.global.u64 	%rd13, %rd5;
	add.s64 	%rd14, %rd13, %rd11;
	cvta.to.global.u64 	%rd15, %rd6;
	add.s64 	%rd16, %rd15, %rd11;
	cvta.to.global.u64 	%rd17, %rd7;
	add.s64 	%rd18, %rd17, %rd11;
	cvta.to.global.u64 	%rd19, %rd8;
	add.s64 	%rd20, %rd19, %rd11;
	cvta.to.global.u64 	%rd21, %rd9;
	add.s64 	%rd22, %rd21, %rd11;
	ld.global.nc.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd18];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd12];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	cvta.to.global.u64 	%rd23, %rd1;
	add.s64 	%rd24, %rd23, %rd11;
	st.global.f32 	[%rd24], %f7;
	cvta.to.global.u64 	%rd25, %rd2;
	add.s64 	%rd26, %rd25, %rd11;
	st.global.f32 	[%rd26], %f12;
	cvta.to.global.u64 	%rd27, %rd3;
	add.s64 	%rd28, %rd27, %rd11;
	st.global.f32 	[%rd28], %f15;

BB0_2:
	ret;
}


`
)

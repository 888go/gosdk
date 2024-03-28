
<原文开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// The functions in pcg implement a 32 bit PRNG with a 64 bit period: pcg xsh rr
// 64 32. See https://www.pcg-random.org/ for more information. This
// implementation is geared specifically towards the needs of fuzzing: Simple
// creation and use, no reproducibility, no concurrency safety, just the
// necessary methods, optimized for speed.
<原文结束>

# <翻译开始>
// The functions in pcg implement a 32 bit PRNG with a 64 bit period: pcg xsh rr
// 64 32. See https://www.pcg-random.org/ for more information. This
// implementation is geared specifically towards the needs of fuzzing: Simple
// creation and use, no reproducibility, no concurrency safety, just the
// necessary methods, optimized for speed.
# <翻译结束>


<原文开始>
// pcgRand is a PRNG. It should not be copied or shared. No Rand methods are
// concurrency safe.
<原文结束>

# <翻译开始>
// pcgRand is a PRNG. It should not be copied or shared. No Rand methods are
// concurrency safe.
# <翻译结束>


<原文开始>
// newPcgRand generates a new, seeded Rand, ready for use.
<原文结束>

# <翻译开始>
// newPcgRand generates a new, seeded Rand, ready for use.
# <翻译结束>


<原文开始>
// uint32 returns a pseudo-random uint32.
<原文结束>

# <翻译开始>
// uint32 returns a pseudo-random uint32.
# <翻译结束>


<原文开始>
// intn returns a pseudo-random number in [0, n).
// n must fit in a uint32.
<原文结束>

# <翻译开始>
// intn returns a pseudo-random number in [0, n).
// n must fit in a uint32.
# <翻译结束>


<原文开始>
// uint32n returns a pseudo-random number in [0, n).
//
// For implementation details, see:
// https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction
// https://lemire.me/blog/2016/06/30/fast-random-shuffling
<原文结束>

# <翻译开始>
// uint32n returns a pseudo-random number in [0, n).
//
// For implementation details, see:
// https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction
// https://lemire.me/blog/2016/06/30/fast-random-shuffling
# <翻译结束>


<原文开始>
// exp2 generates n with probability 1/2^(n+1).
<原文结束>

# <翻译开始>
// exp2 generates n with probability 1/2^(n+1).
# <翻译结束>


<原文开始>
// bool generates a random bool.
<原文结束>

# <翻译开始>
// bool generates a random bool.
# <翻译结束>


<原文开始>
// noCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
<原文结束>

# <翻译开始>
// noCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
# <翻译结束>


<原文开始>
// lock is a no-op used by -copylocks checker from `go vet`.
<原文结束>

# <翻译开始>
// lock is a no-op used by -copylocks checker from `go vet`.
# <翻译结束>



<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// byteSliceRemoveBytes removes a random chunk of bytes from b.
<原文结束>

# <翻译开始>
// byteSliceRemoveBytes removes a random chunk of bytes from b.
# <翻译结束>


<原文开始>
// byteSliceInsertRandomBytes inserts a chunk of random bytes into b at a random
// position.
<原文结束>

# <翻译开始>
// byteSliceInsertRandomBytes inserts a chunk of random bytes into b at a random
// position.
# <翻译结束>


<原文开始>
// byteSliceDuplicateBytes duplicates a chunk of bytes in b and inserts it into
// a random position.
<原文结束>

# <翻译开始>
// byteSliceDuplicateBytes duplicates a chunk of bytes in b and inserts it into
// a random position.
# <翻译结束>


<原文开始>
	// Use the end of the slice as scratch space to avoid doing an
	// allocation. If the slice is too small abort and try something
	// else.
<原文结束>

# <翻译开始>
	// Use the end of the slice as scratch space to avoid doing an
	// allocation. If the slice is too small abort and try something
	// else.
# <翻译结束>


<原文开始>
	// Increase the size of b to fit the duplicated block as well as
	// some extra working space
<原文结束>

# <翻译开始>
	// Increase the size of b to fit the duplicated block as well as
	// some extra working space
# <翻译结束>


<原文开始>
	// Copy the block of bytes we want to duplicate to the end of the
	// slice
<原文结束>

# <翻译开始>
	// Copy the block of bytes we want to duplicate to the end of the
	// slice
# <翻译结束>


<原文开始>
	// Shift the bytes after the splice point n positions to the right
	// to make room for the new block
<原文结束>

# <翻译开始>
	// Shift the bytes after the splice point n positions to the right
	// to make room for the new block
# <翻译结束>


<原文开始>
// Insert the duplicate block into the splice point
<原文结束>

# <翻译开始>
// Insert the duplicate block into the splice point
# <翻译结束>


<原文开始>
// byteSliceOverwriteBytes overwrites a chunk of b with another chunk of b.
<原文结束>

# <翻译开始>
// byteSliceOverwriteBytes overwrites a chunk of b with another chunk of b.
# <翻译结束>


<原文开始>
// byteSliceBitFlip flips a random bit in a random byte in b.
<原文结束>

# <翻译开始>
// byteSliceBitFlip flips a random bit in a random byte in b.
# <翻译结束>


<原文开始>
// byteSliceXORByte XORs a random byte in b with a random value.
<原文结束>

# <翻译开始>
// byteSliceXORByte XORs a random byte in b with a random value.
# <翻译结束>


<原文开始>
	// In order to avoid a no-op (where the random value matches
	// the existing value), use XOR instead of just setting to
	// the random value.
<原文结束>

# <翻译开始>
	// In order to avoid a no-op (where the random value matches
	// the existing value), use XOR instead of just setting to
	// the random value.
# <翻译结束>


<原文开始>
// byteSliceSwapByte swaps two random bytes in b.
<原文结束>

# <翻译开始>
// byteSliceSwapByte swaps two random bytes in b.
# <翻译结束>


<原文开始>
// byteSliceArithmeticUint8 adds/subtracts from a random byte in b.
<原文结束>

# <翻译开始>
// byteSliceArithmeticUint8 adds/subtracts from a random byte in b.
# <翻译结束>


<原文开始>
// byteSliceArithmeticUint16 adds/subtracts from a random uint16 in b.
<原文结束>

# <翻译开始>
// byteSliceArithmeticUint16 adds/subtracts from a random uint16 in b.
# <翻译结束>


<原文开始>
// byteSliceArithmeticUint32 adds/subtracts from a random uint32 in b.
<原文结束>

# <翻译开始>
// byteSliceArithmeticUint32 adds/subtracts from a random uint32 in b.
# <翻译结束>


<原文开始>
// byteSliceArithmeticUint64 adds/subtracts from a random uint64 in b.
<原文结束>

# <翻译开始>
// byteSliceArithmeticUint64 adds/subtracts from a random uint64 in b.
# <翻译结束>


<原文开始>
// byteSliceOverwriteInterestingUint8 overwrites a random byte in b with an interesting
// value.
<原文结束>

# <翻译开始>
// byteSliceOverwriteInterestingUint8 overwrites a random byte in b with an interesting
// value.
# <翻译结束>


<原文开始>
// byteSliceOverwriteInterestingUint16 overwrites a random uint16 in b with an interesting
// value.
<原文结束>

# <翻译开始>
// byteSliceOverwriteInterestingUint16 overwrites a random uint16 in b with an interesting
// value.
# <翻译结束>


<原文开始>
// byteSliceOverwriteInterestingUint32 overwrites a random uint16 in b with an interesting
// value.
<原文结束>

# <翻译开始>
// byteSliceOverwriteInterestingUint32 overwrites a random uint16 in b with an interesting
// value.
# <翻译结束>


<原文开始>
// byteSliceInsertConstantBytes inserts a chunk of constant bytes into a random position in b.
<原文结束>

# <翻译开始>
// byteSliceInsertConstantBytes inserts a chunk of constant bytes into a random position in b.
# <翻译结束>


<原文开始>
	// TODO(rolandshoemaker,katiehockman): 4096 was mainly picked
	// randomly. We may want to either pick a much larger value
	// (AFL uses 32768, paired with a similar impl to chooseLen
	// which biases towards smaller lengths that grow over time),
	// or set the max based on characteristics of the corpus
	// (libFuzzer sets a min/max based on the min/max size of
	// entries in the corpus and then picks uniformly from
	// that range).
<原文结束>

# <翻译开始>
	// TODO(rolandshoemaker,katiehockman): 4096 was mainly picked
	// randomly. We may want to either pick a much larger value
	// (AFL uses 32768, paired with a similar impl to chooseLen
	// which biases towards smaller lengths that grow over time),
	// or set the max based on characteristics of the corpus
	// (libFuzzer sets a min/max based on the min/max size of
	// entries in the corpus and then picks uniformly from
	// that range).
# <翻译结束>


<原文开始>
// byteSliceOverwriteConstantBytes overwrites a chunk of b with constant bytes.
<原文结束>

# <翻译开始>
// byteSliceOverwriteConstantBytes overwrites a chunk of b with constant bytes.
# <翻译结束>


<原文开始>
// byteSliceShuffleBytes shuffles a chunk of bytes in b.
<原文结束>

# <翻译开始>
// byteSliceShuffleBytes shuffles a chunk of bytes in b.
# <翻译结束>


<原文开始>
	// Start at the end of the range, and iterate backwards
	// to dst, swapping each element with another element in
	// dst:dst+n (Fisher-Yates shuffle).
<原文结束>

# <翻译开始>
	// Start at the end of the range, and iterate backwards
	// to dst, swapping each element with another element in
	// dst:dst+n (Fisher-Yates shuffle).
# <翻译结束>


<原文开始>
// byteSliceSwapBytes swaps two chunks of bytes in b.
<原文结束>

# <翻译开始>
// byteSliceSwapBytes swaps two chunks of bytes in b.
# <翻译结束>


<原文开始>
	// Choose the random length as len(b) - max(src, dst)
	// so that we don't attempt to swap a chunk that extends
	// beyond the end of the slice
<原文结束>

# <翻译开始>
	// Choose the random length as len(b) - max(src, dst)
	// so that we don't attempt to swap a chunk that extends
	// beyond the end of the slice
# <翻译结束>


<原文开始>
	// Check that neither chunk intersect, so that we don't end up
	// duplicating parts of the input, rather than swapping them
<原文结束>

# <翻译开始>
	// Check that neither chunk intersect, so that we don't end up
	// duplicating parts of the input, rather than swapping them
# <翻译结束>


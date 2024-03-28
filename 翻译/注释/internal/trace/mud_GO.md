
<原文开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// mud is an updatable mutator utilization distribution.
//
// This is a continuous distribution of duration over mutator
// utilization. For example, the integral from mutator utilization a
// to b is the total duration during which the mutator utilization was
// in the range [a, b].
//
// This distribution is *not* normalized (it is not a probability
// distribution). This makes it easier to work with as it's being
// updated.
//
// It is represented as the sum of scaled uniform distribution
// functions and Dirac delta functions (which are treated as
// degenerate uniform distributions).
<原文结束>

# <翻译开始>
// mud is an updatable mutator utilization distribution.
//
// This is a continuous distribution of duration over mutator
// utilization. For example, the integral from mutator utilization a
// to b is the total duration during which the mutator utilization was
// in the range [a, b].
//
// This distribution is *not* normalized (it is not a probability
// distribution). This makes it easier to work with as it's being
// updated.
//
// It is represented as the sum of scaled uniform distribution
// functions and Dirac delta functions (which are treated as
// degenerate uniform distributions).
# <翻译结束>


<原文开始>
	// trackMass is the inverse cumulative sum to track as the
	// distribution is updated.
<原文结束>

# <翻译开始>
	// trackMass is the inverse cumulative sum to track as the
	// distribution is updated.
# <翻译结束>


<原文开始>
	// trackBucket is the bucket in which trackMass falls. If the
	// total mass of the distribution is < trackMass, this is
	// len(hist).
<原文结束>

# <翻译开始>
	// trackBucket is the bucket in which trackMass falls. If the
	// total mass of the distribution is < trackMass, this is
	// len(hist).
# <翻译结束>


<原文开始>
	// trackSum is the cumulative sum of hist[:trackBucket]. Once
	// trackSum >= trackMass, trackBucket must be recomputed.
<原文结束>

# <翻译开始>
	// trackSum is the cumulative sum of hist[:trackBucket]. Once
	// trackSum >= trackMass, trackBucket must be recomputed.
# <翻译结束>


<原文开始>
// hist is a hierarchical histogram of distribution mass.
<原文结束>

# <翻译开始>
// hist is a hierarchical histogram of distribution mass.
# <翻译结束>


<原文开始>
	// mudDegree is the number of buckets in the MUD summary
	// histogram.
<原文结束>

# <翻译开始>
	// mudDegree is the number of buckets in the MUD summary
	// histogram.
# <翻译结束>


<原文开始>
// At x, the function increases by y.
<原文结束>

# <翻译开始>
// At x, the function increases by y.
# <翻译结束>


<原文开始>
// Additionally at x is a Dirac delta function with area dirac.
<原文结束>

# <翻译开始>
// Additionally at x is a Dirac delta function with area dirac.
# <翻译结束>


<原文开始>
// add adds a uniform function over [l, r] scaled so the total weight
// of the uniform is area. If l==r, this adds a Dirac delta function.
<原文结束>

# <翻译开始>
// add adds a uniform function over [l, r] scaled so the total weight
// of the uniform is area. If l==r, this adds a Dirac delta function.
# <翻译结束>


<原文开始>
			// The tracked mass now falls in a different
			// bucket. Recompute the inverse cumulative sum.
<原文结束>

# <翻译开始>
			// The tracked mass now falls in a different
			// bucket. Recompute the inverse cumulative sum.
# <翻译结束>


<原文开始>
// setTrackMass sets the mass to track the inverse cumulative sum for.
//
// Specifically, mass is a cumulative duration, and the mutator
// utilization bounds for this duration can be queried using
// approxInvCumulativeSum.
<原文结束>

# <翻译开始>
// setTrackMass sets the mass to track the inverse cumulative sum for.
//
// Specifically, mass is a cumulative duration, and the mutator
// utilization bounds for this duration can be queried using
// approxInvCumulativeSum.
# <翻译结束>


<原文开始>
	// Find the bucket currently containing trackMass by computing
	// the cumulative sum.
<原文结束>

# <翻译开始>
	// Find the bucket currently containing trackMass by computing
	// the cumulative sum.
# <翻译结束>


<原文开始>
// mass falls in bucket i.
<原文结束>

# <翻译开始>
// mass falls in bucket i.
# <翻译结束>


<原文开始>
// approxInvCumulativeSum is like invCumulativeSum, but specifically
// operates on the tracked mass and returns an upper and lower bound
// approximation of the inverse cumulative sum.
//
// The true inverse cumulative sum will be in the range [lower, upper).
<原文结束>

# <翻译开始>
// approxInvCumulativeSum is like invCumulativeSum, but specifically
// operates on the tracked mass and returns an upper and lower bound
// approximation of the inverse cumulative sum.
//
// The true inverse cumulative sum will be in the range [lower, upper).
# <翻译结束>


<原文开始>
// Merge with sorted edges.
<原文结束>

# <翻译开始>
// Merge with sorted edges.
# <翻译结束>


<原文开始>
// Traverse edges in order computing a cumulative sum.
<原文结束>

# <翻译开始>
// Traverse edges in order computing a cumulative sum.
# <翻译结束>


<原文开始>
			// y was exceeded between the previous edge
			// and this one.
<原文结束>

# <翻译开始>
			// y was exceeded between the previous edge
			// and this one.
# <翻译结束>


<原文开始>
				// Anywhere between prevX and
				// e.x will do. We return e.x
				// because that takes care of
				// the y==0 case naturally.
<原文结束>

# <翻译开始>
				// Anywhere between prevX and
				// e.x will do. We return e.x
				// because that takes care of
				// the y==0 case naturally.
# <翻译结束>


<原文开始>
// y was exceeded by the Dirac delta at e.x.
<原文结束>

# <翻译开始>
// y was exceeded by the Dirac delta at e.x.
# <翻译结束>



<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2022 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Pod encapsulates a set of files emitted during the executions of a
// coverage-instrumented binary. Each pod contains a single meta-data
// file, and then 0 or more counter data files that refer to that
// meta-data file. Pods are intended to simplify processing of
// coverage output files in the case where we have several coverage
// output directories containing output files derived from more
// than one instrumented executable. In the case where the files that
// make up a pod are spread out across multiple directories, each
// element of the "Origins" field below will be populated with the
// index of the originating directory for the corresponding counter
// data file (within the slice of input dirs handed to CollectPods).
// The ProcessIDs field will be populated with the process ID of each
// data file in the CounterDataFiles slice.
<原文结束>

# <翻译开始>
// Pod 封装了一组在执行了覆盖率检测的二进制文件后产生的文件。每个 Pod 包含一个元数据文件，以及0个或多个引用该元数据文件的计数器数据文件。Pod 的设计目的是简化处理多个包含来自多个已检测可执行文件输出结果的覆盖率输出目录的情况。当组成 Pod 的文件分散在多个目录中时，"Origins"字段下的每个元素将被填充为对应计数器数据文件来源目录（在传递给 CollectPods 函数的输入目录切片中的索引）。"ProcessIDs"字段将被填充为 CounterDataFiles 切片中每个数据文件的进程 ID。
# <翻译结束>


<原文开始>
// CollectPods visits the files contained within the directories in
// the list 'dirs', collects any coverage-related files, partitions
// them into pods, and returns a list of the pods to the caller, along
// with an error if something went wrong during directory/file
// reading.
//
// CollectPods skips over any file that is not related to coverage
// (e.g. avoids looking at things that are not meta-data files or
// counter-data files). CollectPods also skips over 'orphaned' counter
// data files (e.g. counter data files for which we can't find the
// corresponding meta-data file). If "warn" is true, CollectPods will
// issue warnings to stderr when it encounters non-fatal problems (for
// orphans or a directory with no meta-data files).
<原文结束>

# <翻译开始>
// CollectPods 函数会遍历列表 'dirs' 中的目录及其包含的文件，收集所有与覆盖率相关的文件，将它们划分为多个组（pods），并将这些组作为一个列表返回给调用者。如果在读取目录或文件过程中发生错误，还会同时返回错误信息。
//
// CollectPods 会跳过任何与覆盖率无关的文件（例如，忽略非元数据文件或计数器数据文件）。此外，它也会跳过“孤立”的计数器数据文件，即那些找不到对应元数据文件的计数器数据文件。若参数 "warn" 为 true，当遇到非致命问题时（如孤儿文件或不含元数据文件的目录），CollectPods 将向标准错误输出警告信息。
# <翻译结束>


<原文开始>
// CollectPodsFromFiles functions the same as "CollectPods" but
// operates on an explicit list of files instead of a directory.
<原文结束>

# <翻译开始>
// CollectPodsFromFiles 函数与 "CollectPods" 功能相同，但
// 它是针对显式文件列表操作，而不是针对目录。
# <翻译结束>


<原文开始>
// collectPodsImpl examines the specified list of files and picks out
// subsets that correspond to coverage pods. The first stage in this
// process is collecting a set { M1, M2, ... MN } where each M_k is a
// distinct coverage meta-data file. We then create a single pod for
// each meta-data file M_k, then find all of the counter data files
// that refer to that meta-data file (recall that the counter data
// file name incorporates the meta-data hash), and add the counter
// data file to the appropriate pod.
//
// This process is complicated by the fact that we need to keep track
// of directory indices for counter data files. Here is an example to
// motivate:
//
//	directory 1:
//
// M1   covmeta.9bbf1777f47b3fcacb05c38b035512d6
// C1   covcounters.9bbf1777f47b3fcacb05c38b035512d6.1677673.1662138360208416486
// C2   covcounters.9bbf1777f47b3fcacb05c38b035512d6.1677637.1662138359974441782
//
//	directory 2:
//
// M2   covmeta.9bbf1777f47b3fcacb05c38b035512d6
// C3   covcounters.9bbf1777f47b3fcacb05c38b035512d6.1677445.1662138360208416480
// C4   covcounters.9bbf1777f47b3fcacb05c38b035512d6.1677677.1662138359974441781
// M3   covmeta.a723844208cea2ae80c63482c78b2245
// C5   covcounters.a723844208cea2ae80c63482c78b2245.3677445.1662138360208416480
// C6   covcounters.a723844208cea2ae80c63482c78b2245.1877677.1662138359974441781
//
// In these two directories we have three meta-data files, but only
// two are distinct, meaning that we'll wind up with two pods. The
// first pod (with meta-file M1) will have four counter data files
// (C1, C2, C3, C4) and the second pod will have two counter data files
// (C5, C6).
<原文结束>

# <翻译开始>
// collectPodsImpl examines the specified list of files and picks out
// subsets that correspond to coverage pods. The first stage in this
// process is collecting a set { M1, M2, ... MN } where each M_k is a
// distinct coverage meta-data file. We then create a single pod for
// each meta-data file M_k, then find all of the counter data files
// that refer to that meta-data file (recall that the counter data
// file name incorporates the meta-data hash), and add the counter
// data file to the appropriate pod.
//
// This process is complicated by the fact that we need to keep track
// of directory indices for counter data files. Here is an example to
// motivate:
//
//	directory 1:
//
// M1   covmeta.9bbf1777f47b3fcacb05c38b035512d6
// C1   covcounters.9bbf1777f47b3fcacb05c38b035512d6.1677673.1662138360208416486
// C2   covcounters.9bbf1777f47b3fcacb05c38b035512d6.1677637.1662138359974441782
//
//	directory 2:
//
// M2   covmeta.9bbf1777f47b3fcacb05c38b035512d6
// C3   covcounters.9bbf1777f47b3fcacb05c38b035512d6.1677445.1662138360208416480
// C4   covcounters.9bbf1777f47b3fcacb05c38b035512d6.1677677.1662138359974441781
// M3   covmeta.a723844208cea2ae80c63482c78b2245
// C5   covcounters.a723844208cea2ae80c63482c78b2245.3677445.1662138360208416480
// C6   covcounters.a723844208cea2ae80c63482c78b2245.1877677.1662138359974441781
//
// In these two directories we have three meta-data files, but only
// two are distinct, meaning that we'll wind up with two pods. The
// first pod (with meta-file M1) will have four counter data files
// (C1, C2, C3, C4) and the second pod will have two counter data files
// (C5, C6).
# <翻译结束>


<原文开始>
			// We need to allow for the possibility of duplicate
			// meta-data files. If we hit this case, use the
			// first encountered as the canonical version.
<原文结束>

# <翻译开始>
			// We need to allow for the possibility of duplicate
			// meta-data files. If we hit this case, use the
			// first encountered as the canonical version.
# <翻译结束>


<原文开始>
			// FIXME: should probably check file length and hash here for
			// the duplicate.
<原文结束>

# <翻译开始>
			// FIXME: should probably check file length and hash here for
			// the duplicate.
# <翻译结束>


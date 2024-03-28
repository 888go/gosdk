
<原文开始>
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// preEncode populates the unexported fields to be used by encode
// (with suffix X) from the corresponding exported fields. The
// exported fields are cleared up to facilitate testing.
<原文结束>

# <翻译开始>
// preEncode populates the unexported fields to be used by encode
// (with suffix X) from the corresponding exported fields. The
// exported fields are cleared up to facilitate testing.
# <翻译结束>


<原文开始>
// repeated ValueType sample_type = 1
<原文结束>

# <翻译开始>
// repeated ValueType sample_type = 1
# <翻译结束>


<原文开始>
// repeated Sample sample = 2
<原文结束>

# <翻译开始>
// repeated Sample sample = 2
# <翻译结束>


<原文开始>
// repeated Mapping mapping = 3
<原文结束>

# <翻译开始>
// repeated Mapping mapping = 3
# <翻译结束>


<原文开始>
// repeated Location location = 4
<原文结束>

# <翻译开始>
// repeated Location location = 4
# <翻译结束>


<原文开始>
// repeated Function function = 5
<原文结束>

# <翻译开始>
// repeated Function function = 5
# <翻译结束>


<原文开始>
// repeated string string_table = 6
<原文结束>

# <翻译开始>
// repeated string string_table = 6
# <翻译结束>


<原文开始>
// repeated int64 drop_frames = 7
<原文结束>

# <翻译开始>
// repeated int64 drop_frames = 7
# <翻译结束>


<原文开始>
// repeated int64 keep_frames = 8
<原文结束>

# <翻译开始>
// repeated int64 keep_frames = 8
# <翻译结束>


<原文开始>
// repeated int64 time_nanos = 9
<原文结束>

# <翻译开始>
// repeated int64 time_nanos = 9
# <翻译结束>


<原文开始>
// repeated int64 duration_nanos = 10
<原文结束>

# <翻译开始>
// repeated int64 duration_nanos = 10
# <翻译结束>


<原文开始>
// optional string period_type = 11
<原文结束>

# <翻译开始>
// optional string period_type = 11
# <翻译结束>


<原文开始>
// repeated int64 period = 12
<原文结束>

# <翻译开始>
// repeated int64 period = 12
# <翻译结束>


<原文开始>
// repeated int64 comment = 13
<原文结束>

# <翻译开始>
// repeated int64 comment = 13
# <翻译结束>


<原文开始>
// int64 defaultSampleType = 14
<原文结束>

# <翻译开始>
// int64 defaultSampleType = 14
# <翻译结束>


<原文开始>
// postDecode takes the unexported fields populated by decode (with
// suffix X) and populates the corresponding exported fields.
// The unexported fields are cleared up to facilitate testing.
<原文结束>

# <翻译开始>
// postDecode takes the unexported fields populated by decode (with
// suffix X) and populates the corresponding exported fields.
// The unexported fields are cleared up to facilitate testing.
# <翻译结束>


<原文开始>
// optional int64 type = 1
<原文结束>

# <翻译开始>
// optional int64 type = 1
# <翻译结束>


<原文开始>
// optional int64 unit = 2
<原文结束>

# <翻译开始>
// optional int64 unit = 2
# <翻译结束>


<原文开始>
// repeated uint64 location = 1
<原文结束>

# <翻译开始>
// repeated uint64 location = 1
# <翻译结束>


<原文开始>
// repeated int64 value = 2
<原文结束>

# <翻译开始>
// repeated int64 value = 2
# <翻译结束>


<原文开始>
// repeated Label label = 3
<原文结束>

# <翻译开始>
// repeated Label label = 3
# <翻译结束>


<原文开始>
// optional uint64 memory_offset = 2
<原文结束>

# <翻译开始>
// optional uint64 memory_offset = 2
# <翻译结束>


<原文开始>
// optional uint64 memory_limit = 3
<原文结束>

# <翻译开始>
// optional uint64 memory_limit = 3
# <翻译结束>


<原文开始>
// optional uint64 file_offset = 4
<原文结束>

# <翻译开始>
// optional uint64 file_offset = 4
# <翻译结束>


<原文开始>
// optional int64 filename = 5
<原文结束>

# <翻译开始>
// optional int64 filename = 5
# <翻译结束>


<原文开始>
// optional int64 build_id = 6
<原文结束>

# <翻译开始>
// optional int64 build_id = 6
# <翻译结束>


<原文开始>
// optional bool has_functions = 7
<原文结束>

# <翻译开始>
// optional bool has_functions = 7
# <翻译结束>


<原文开始>
// optional bool has_filenames = 8
<原文结束>

# <翻译开始>
// optional bool has_filenames = 8
# <翻译结束>


<原文开始>
// optional bool has_line_numbers = 9
<原文结束>

# <翻译开始>
// optional bool has_line_numbers = 9
# <翻译结束>


<原文开始>
// optional bool has_inline_frames = 10
<原文结束>

# <翻译开始>
// optional bool has_inline_frames = 10
# <翻译结束>


<原文开始>
// optional uint64 id = 1;
<原文结束>

# <翻译开始>
// optional uint64 id = 1;
# <翻译结束>


<原文开始>
// optional uint64 mapping_id = 2;
<原文结束>

# <翻译开始>
// optional uint64 mapping_id = 2;
# <翻译结束>


<原文开始>
// optional uint64 address = 3;
<原文结束>

# <翻译开始>
// optional uint64 address = 3;
# <翻译结束>


<原文开始>
// optional uint64 function_id = 1
<原文结束>

# <翻译开始>
// optional uint64 function_id = 1
# <翻译结束>


<原文开始>
// optional int64 line = 2
<原文结束>

# <翻译开始>
// optional int64 line = 2
# <翻译结束>


<原文开始>
// optional int64 function_name = 2
<原文结束>

# <翻译开始>
// optional int64 function_name = 2
# <翻译结束>


<原文开始>
// optional int64 function_system_name = 3
<原文结束>

# <翻译开始>
// optional int64 function_system_name = 3
# <翻译结束>


<原文开始>
// repeated int64 filename = 4
<原文结束>

# <翻译开始>
// repeated int64 filename = 4
# <翻译结束>


<原文开始>
// optional int64 start_line = 5
<原文结束>

# <翻译开始>
// optional int64 start_line = 5
# <翻译结束>


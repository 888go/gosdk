
<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// This package implements string table writer and reader utilities,
// for use in emitting and reading/decoding coverage meta-data and
// counter-data files.
<原文结束>

# <翻译开始>
// This package implements string table writer and reader utilities,
// for use in emitting and reading/decoding coverage meta-data and
// counter-data files.
# <翻译结束>


<原文开始>
// Writer implements a string table writing utility.
<原文结束>

# <翻译开始>
// Writer implements a string table writing utility.
# <翻译结束>


<原文开始>
// InitWriter initializes a stringtab.Writer.
<原文结束>

# <翻译开始>
// InitWriter initializes a stringtab.Writer.
# <翻译结束>


<原文开始>
// Nentries returns the number of strings interned so far.
<原文结束>

# <翻译开始>
// Nentries returns the number of strings interned so far.
# <翻译结束>


<原文开始>
// Lookup looks up string 's' in the writer's table, adding
// a new entry if need be, and returning an index into the table.
<原文结束>

# <翻译开始>
// Lookup looks up string 's' in the writer's table, adding
// a new entry if need be, and returning an index into the table.
# <翻译结束>


<原文开始>
// Size computes the memory in bytes needed for the serialized
// version of a stringtab.Writer.
<原文结束>

# <翻译开始>
// Size computes the memory in bytes needed for the serialized
// version of a stringtab.Writer.
# <翻译结束>


<原文开始>
// Write writes the string table in serialized form to the specified
// io.Writer.
<原文结束>

# <翻译开始>
// Write writes the string table in serialized form to the specified
// io.Writer.
# <翻译结束>


<原文开始>
// Freeze sends a signal to the writer that no more additions are
// allowed, only lookups of existing strings (if a lookup triggers
// addition, a panic will result). Useful as a mechanism for
// "finalizing" a string table prior to writing it out.
<原文结束>

# <翻译开始>
// Freeze sends a signal to the writer that no more additions are
// allowed, only lookups of existing strings (if a lookup triggers
// addition, a panic will result). Useful as a mechanism for
// "finalizing" a string table prior to writing it out.
# <翻译结束>


<原文开始>
// Reader is a helper for reading a string table previously
// serialized by a Writer.Write call.
<原文结束>

# <翻译开始>
// Reader is a helper for reading a string table previously
// serialized by a Writer.Write call.
# <翻译结束>


<原文开始>
// NewReader creates a stringtab.Reader to read the contents
// of a string table from 'r'.
<原文结束>

# <翻译开始>
// NewReader creates a stringtab.Reader to read the contents
// of a string table from 'r'.
# <翻译结束>


<原文开始>
// Read reads/decodes a string table using the reader provided.
<原文结束>

# <翻译开始>
// Read reads/decodes a string table using the reader provided.
# <翻译结束>


<原文开始>
// Entries returns the number of decoded entries in a string table.
<原文结束>

# <翻译开始>
// Entries returns the number of decoded entries in a string table.
# <翻译结束>


<原文开始>
// Get returns string 'idx' within the string table.
<原文结束>

# <翻译开始>
// Get returns string 'idx' within the string table.
# <翻译结束>


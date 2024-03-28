
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Package xcoff implements access to XCOFF (Extended Common Object File Format) files.
<原文结束>

# <翻译开始>
// Package xcoff implements access to XCOFF (Extended Common Object File Format) files.
# <翻译结束>


<原文开始>
// SectionHeader holds information about an XCOFF section header.
<原文结束>

# <翻译开始>
// SectionHeader holds information about an XCOFF section header.
# <翻译结束>


<原文开始>
// AuxiliaryCSect holds information about an XCOFF symbol in an AUX_CSECT entry.
<原文结束>

# <翻译开始>
// AuxiliaryCSect holds information about an XCOFF symbol in an AUX_CSECT entry.
# <翻译结束>


<原文开始>
// AuxiliaryFcn holds information about an XCOFF symbol in an AUX_FCN entry.
<原文结束>

# <翻译开始>
// AuxiliaryFcn holds information about an XCOFF symbol in an AUX_FCN entry.
# <翻译结束>


<原文开始>
// ImportedSymbol holds information about an imported XCOFF symbol.
<原文结束>

# <翻译开始>
// ImportedSymbol holds information about an imported XCOFF symbol.
# <翻译结束>


<原文开始>
// FileHeader holds information about an XCOFF file header.
<原文结束>

# <翻译开始>
// FileHeader holds information about an XCOFF file header.
# <翻译结束>


<原文开始>
// A File represents an open XCOFF file.
<原文结束>

# <翻译开始>
// A File represents an open XCOFF file.
# <翻译结束>


<原文开始>
// Open opens the named file using os.Open and prepares it for use as an XCOFF binary.
<原文结束>

# <翻译开始>
// Open opens the named file using os.Open and prepares it for use as an XCOFF binary.
# <翻译结束>


<原文开始>
// Close closes the File.
// If the File was created using NewFile directly instead of Open,
// Close has no effect.
<原文结束>

# <翻译开始>
// Close closes the File.
// If the File was created using NewFile directly instead of Open,
// Close has no effect.
# <翻译结束>


<原文开始>
// Section returns the first section with the given name, or nil if no such
// section exists.
// Xcoff have section's name limited to 8 bytes. Some sections like .gosymtab
// can be trunked but this method will still find them.
<原文结束>

# <翻译开始>
// Section returns the first section with the given name, or nil if no such
// section exists.
// Xcoff have section's name limited to 8 bytes. Some sections like .gosymtab
// can be trunked but this method will still find them.
# <翻译结束>


<原文开始>
// SectionByType returns the first section in f with the
// given type, or nil if there is no such section.
<原文结束>

# <翻译开始>
// SectionByType returns the first section in f with the
// given type, or nil if there is no such section.
# <翻译结束>


<原文开始>
// cstring converts ASCII byte sequence b to string.
// It stops once it finds 0 or reaches end of b.
<原文结束>

# <翻译开始>
// cstring converts ASCII byte sequence b to string.
// It stops once it finds 0 or reaches end of b.
# <翻译结束>


<原文开始>
// getString extracts a string from an XCOFF string table.
<原文结束>

# <翻译开始>
// getString extracts a string from an XCOFF string table.
# <翻译结束>


<原文开始>
// NewFile creates a new File for accessing an XCOFF binary in an underlying reader.
<原文结束>

# <翻译开始>
// NewFile creates a new File for accessing an XCOFF binary in an underlying reader.
# <翻译结束>


<原文开始>
// Read XCOFF target machine
<原文结束>

# <翻译开始>
// Read XCOFF target machine
# <翻译结束>


<原文开始>
// Read string table (located right after symbol table).
<原文结束>

# <翻译开始>
// Read string table (located right after symbol table).
# <翻译结束>


<原文开始>
// The first 4 bytes contain the length (in bytes).
<原文结束>

# <翻译开始>
// The first 4 bytes contain the length (in bytes).
# <翻译结束>


<原文开始>
// Symbol map needed by relocation
<原文结束>

# <翻译开始>
// Symbol map needed by relocation
# <翻译结束>


<原文开始>
// Must have at least one csect auxiliary entry.
<原文结束>

# <翻译开始>
// Must have at least one csect auxiliary entry.
# <翻译结束>


<原文开始>
		// If this symbol is a function, it must retrieve its size from
		// its AUX_FCN entry.
		// It can happen that a function symbol doesn't have any AUX_FCN.
		// In this case, needAuxFcn is false and their size will be set to 0.
<原文结束>

# <翻译开始>
		// If this symbol is a function, it must retrieve its size from
		// its AUX_FCN entry.
		// It can happen that a function symbol doesn't have any AUX_FCN.
		// In this case, needAuxFcn is false and their size will be set to 0.
# <翻译结束>


<原文开始>
// Read csect auxiliary entry (by convention, it is the last).
<原文结束>

# <翻译开始>
// Read csect auxiliary entry (by convention, it is the last).
# <翻译结束>


<原文开始>
	// Read relocations
	// Only for .data or .text section
<原文结束>

# <翻译开始>
	// Read relocations
	// Only for .data or .text section
# <翻译结束>


<原文开始>
// zeroReaderAt is ReaderAt that reads 0s.
<原文结束>

# <翻译开始>
// zeroReaderAt is ReaderAt that reads 0s.
# <翻译结束>


<原文开始>
// ReadAt writes len(p) 0s into p.
<原文结束>

# <翻译开始>
// ReadAt writes len(p) 0s into p.
# <翻译结束>


<原文开始>
// Data reads and returns the contents of the XCOFF section s.
<原文结束>

# <翻译开始>
// Data reads and returns the contents of the XCOFF section s.
# <翻译结束>


<原文开始>
// CSect reads and returns the contents of a csect.
<原文结束>

# <翻译开始>
// CSect reads and returns the contents of a csect.
# <翻译结束>


<原文开始>
	// There are many other DWARF sections, but these
	// are the ones the debug/dwarf package uses.
	// Don't bother loading others.
<原文结束>

# <翻译开始>
	// There are many other DWARF sections, but these
	// are the ones the debug/dwarf package uses.
	// Don't bother loading others.
# <翻译结束>


<原文开始>
// readImportID returns the import file IDs stored inside the .loader section.
// Library name pattern is either path/base/member or base/member
<原文结束>

# <翻译开始>
// readImportID returns the import file IDs stored inside the .loader section.
// Library name pattern is either path/base/member or base/member
# <翻译结束>


<原文开始>
// Read loader import file ID table
<原文结束>

# <翻译开始>
// Read loader import file ID table
# <翻译结束>


<原文开始>
// First import file ID is the default LIBPATH value
<原文结束>

# <翻译开始>
// First import file ID is the default LIBPATH value
# <翻译结束>


<原文开始>
// ImportedSymbols returns the names of all symbols
// referred to by the binary f that are expected to be
// satisfied by other libraries at dynamic load time.
// It does not return weak symbols.
<原文结束>

# <翻译开始>
// ImportedSymbols returns the names of all symbols
// referred to by the binary f that are expected to be
// satisfied by other libraries at dynamic load time.
// It does not return weak symbols.
# <翻译结束>


<原文开始>
// Read loader section string table
<原文结束>

# <翻译开始>
// Read loader section string table
# <翻译结束>


<原文开始>
// Read imported libraries
<原文结束>

# <翻译开始>
// Read imported libraries
# <翻译结束>


<原文开始>
// Read loader symbol table
<原文结束>

# <翻译开始>
// Read loader symbol table
# <翻译结束>


<原文开始>
// ImportedLibraries returns the names of all libraries
// referred to by the binary f that are expected to be
// linked with the binary at dynamic link time.
<原文结束>

# <翻译开始>
// ImportedLibraries returns the names of all libraries
// referred to by the binary f that are expected to be
// linked with the binary at dynamic link time.
# <翻译结束>


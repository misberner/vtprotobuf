package descriptors

import (
	"sort"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// topSortRec performs a topological sort via depth-first search. This implementation is intentionally kept simplistic,
// in particular, it does not check for cycles.
func topSortRec(out *[]protoreflect.FileDescriptor, curr protoreflect.FileDescriptor, files map[string]protoreflect.FileDescriptor) {
	if _, found := files[curr.Path()]; !found {
		return // already covered
	}

	delete(files, curr.Path())
	for i := 0; i < curr.Imports().Len(); i++ {
		f := files[curr.Imports().Get(i).Path()]
		if f != nil {
			topSortRec(out, f, files)
		}
	}
	*out = append(*out, curr)
}

// topSortFiles returns a topological ordering of the given file descriptors, in which each file appears before it is
// imported.
func topSortFiles(files []protoreflect.FileDescriptor) []protoreflect.FileDescriptor {
	filesByPath := make(map[string]protoreflect.FileDescriptor, len(files))
	filesSorted := make([]protoreflect.FileDescriptor, 0, len(files))
	for _, f := range files {
		filesByPath[f.Path()] = f
		filesSorted = append(filesSorted, f)
	}

	// Sort the files by path. This ensures a deterministic traversal order and thus a stable sort.
	sort.Slice(filesSorted, func(i, j int) bool {
		return filesSorted[i].Path() < filesSorted[j].Path()
	})

	result := make([]protoreflect.FileDescriptor, 0, len(files))
	for _, f := range filesSorted {
		topSortRec(&result, f, filesByPath)
	}
	return result
}

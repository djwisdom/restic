package fs

import (
	"os"
	"path/filepath"

	"github.com/restic/restic/internal/restic"
)

// Local is the local file system. Most methods are just passed on to the stdlib.
type Local struct{}

// statically ensure that Local implements FS.
var _ FS = &Local{}

// VolumeName returns leading volume name. Given "C:\foo\bar" it returns "C:"
// on Windows. Given "\\host\share\foo" it returns "\\host\share". On other
// platforms it returns "".
func (fs Local) VolumeName(path string) string {
	return filepath.VolumeName(path)
}

// OpenFile is the generalized open call; most users will use Open
// or Create instead.  It opens the named file with specified flag
// (O_RDONLY etc.) and perm, (0666 etc.) if applicable.  If successful,
// methods on the returned File can be used for I/O.
// If there is an error, it will be of type *PathError.
func (fs Local) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	f, err := os.OpenFile(fixpath(name), flag, perm)
	if err != nil {
		return nil, err
	}
	_ = setFlags(f)
	return f, nil
}

// Stat returns a FileInfo describing the named file. If there is an error, it
// will be of type *PathError.
func (fs Local) Stat(name string) (os.FileInfo, error) {
	return os.Stat(fixpath(name))
}

// Lstat returns the FileInfo structure describing the named file.
// If the file is a symbolic link, the returned FileInfo
// describes the symbolic link.  Lstat makes no attempt to follow the link.
// If there is an error, it will be of type *PathError.
func (fs Local) Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(fixpath(name))
}

// DeviceID extracts the DeviceID from the given FileInfo. If the fs does
// not support a DeviceID, it returns an error instead
func (fs Local) DeviceID(fi os.FileInfo) (id uint64, err error) {
	return deviceID(fi)
}

// ExtendedStat converts the give FileInfo into ExtendedFileInfo.
func (fs Local) ExtendedStat(fi os.FileInfo) ExtendedFileInfo {
	return ExtendedStat(fi)
}

func (fs Local) NodeFromFileInfo(path string, fi os.FileInfo, ignoreXattrListError bool) (*restic.Node, error) {
	return nodeFromFileInfo(path, fi, ignoreXattrListError)
}

// Join joins any number of path elements into a single path, adding a
// Separator if necessary. Join calls Clean on the result; in particular, all
// empty strings are ignored. On Windows, the result is a UNC path if and only
// if the first path element is a UNC path.
func (fs Local) Join(elem ...string) string {
	return filepath.Join(elem...)
}

// Separator returns the OS and FS dependent separator for dirs/subdirs/files.
func (fs Local) Separator() string {
	return string(filepath.Separator)
}

// IsAbs reports whether the path is absolute.
func (fs Local) IsAbs(path string) bool {
	return filepath.IsAbs(path)
}

// Abs returns an absolute representation of path. If the path is not absolute
// it will be joined with the current working directory to turn it into an
// absolute path. The absolute path name for a given file is not guaranteed to
// be unique. Abs calls Clean on the result.
func (fs Local) Abs(path string) (string, error) {
	return filepath.Abs(path)
}

// Clean returns the cleaned path. For details, see filepath.Clean.
func (fs Local) Clean(p string) string {
	return filepath.Clean(p)
}

// Base returns the last element of path.
func (fs Local) Base(path string) string {
	return filepath.Base(path)
}

// Dir returns path without the last element.
func (fs Local) Dir(path string) string {
	return filepath.Dir(path)
}

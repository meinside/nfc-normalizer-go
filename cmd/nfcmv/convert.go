package main

import (
	"fmt"
	"os"
	"path/filepath"

	nfcnorm "github.com/meinside/nfc-normalizer-go"
)

// check if given filepath is a directory
func isDir(filepath string) bool {
	if info, err := os.Stat(filepath); err == nil {
		return info.IsDir()
	}
	return false
}

// convert all items
//
// FIXME: not to recurse
func convertAll(paths []string, recursive, dryrun, interactive, forceReplace bool) (converted uint) {
	//  NOTE: depth-first traverse
	for _, path := range paths {
		if isDir(path) && recursive {
			converted += recurse(path, dryrun, interactive, forceReplace)
		}

		converted += convert(path, dryrun, interactive, forceReplace)
	}

	return converted
}

// recurse into given directory
func recurse(dir string, dryrun, interactive, forceReplace bool) uint {
	if entries, err := os.ReadDir(dir); err == nil {
		paths := []string{}
		for _, entry := range entries {
			paths = append(paths, filepath.Join(dir, entry.Name()))
		}

		return convertAll(paths, true, dryrun, interactive, forceReplace)
	} else {
		_v("failed to read recurse directory '%s': %s", dir, err)
	}

	return 0
}

// confirm user (y/N)
func confirm(message string) bool {
	fmt.Print(message)

	var input string
	if n, err := fmt.Scanln(&input); err == nil {
		if n > 0 && input == "y" || input == "Y" {
			return true
		}
	}

	return false
}

// do the actual conversion
func convert(original string, dryrun, interactive, forceReplace bool) uint {
	// NOTE: convert only the last element of filepath
	dir, filename := filepath.Split(original)

	// skip non-normalizable(that is, already normalized) one
	if nfcnorm.Normalizable(filename) {
		normalized := nfcnorm.Normalize(filename)
		destination := filepath.Join(dir, normalized)

		if dryrun {
			_p("[dryrun] converting: '%s{%s => %s}'", dir, filename, normalized)

			return 1
		} else {
			confirmed := true

			if fileExists(destination) && (interactive || !forceReplace) {
				confirmed = confirm(fmt.Sprintf("┍ overwrite already existing file: '%s'? [y/N]: ", destination))
			} else if interactive {
				confirmed = confirm(fmt.Sprintf("┍ convert this file?: '%s{%s => %s}' [y/N]: ", dir, filename, normalized))
			}

			if confirmed {
				return move(dir, filename, normalized, original, destination)
			} else {
				_p("┕ skipped '%s'", original)
			}
		}
	} else {
		_v("not a normalizable filename: %s", filename)
	}

	return 0
}

// move file at path `original` to `detination`
func move(dir, filename, normalized, original, destination string) uint {
	_v("moving file '%s' => '%s'", original, destination)

	err := os.Rename(original, destination)
	if err == nil {
		_p("┕ converted: '%s{%s => %s}'", dir, filename, normalized)

		return 1
	} else {
		_p("┕ failed to move file: %s", original)
	}

	return 0
}

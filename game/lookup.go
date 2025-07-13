package game

import (
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
)

var cardJSONMap map[string]string

// InitCardJSONLookup scans dir/*.json once, stripping trailing _<digits>
func InitCardJSONLookup(dir string) error {
	cardJSONMap = make(map[string]string)
	re := regexp.MustCompile(`^(?P<base>.+?)_\d+$`)
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(d.Name(), ".json") {
			return err
		}
		name := strings.TrimSuffix(d.Name(), ".json")
		base := name
		if m := re.FindStringSubmatch(name); m != nil {
			base = m[1]
		}
		key := strings.ToLower(base)
		cardJSONMap[key] = d.Name()
		return nil
	})
}

func normalizeName(human string) string {
	s := strings.ReplaceAll(human, "’", "")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, " ", "_")
	return strings.ToLower(s)
}

func LookupCardJSON(humanName string) (string, bool) {
	fname, ok := cardJSONMap[normalizeName(humanName)]
	return fname, ok
}

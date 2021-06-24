package path

import "strings"

const (
	defaultPathSeparator = "/"
)

var (
	wildcardChars = []uint8{'*', '?', '{'}
)

func NewAntPathMatcherDefault() *AntPathMatcher {
	return NewAntPathMatcher(defaultPathSeparator)
}

func NewAntPathMatcher(pathSeparator string) *AntPathMatcher {
	return &AntPathMatcher{
		pathSeparator: pathSeparator}
}

type AntPathMatcher struct {
	pathSeparator string
	trimTokens    bool
}

func (m *AntPathMatcher) SetPathSeparator(pathSeparator string) {
	m.pathSeparator = pathSeparator
}

func (m *AntPathMatcher) SetTrimTokens(trimTokens bool) {
	m.trimTokens = trimTokens
}

func (m *AntPathMatcher) IsPattern(path string) bool {
	return strings.Index(path, "*") != -1 || strings.Index(path, "?") != -1
}

func (m *AntPathMatcher) Match(pattern string, path string) bool {
	return m.doMatch(pattern, path, true)
}

func (m *AntPathMatcher) doMatch(pattern string, path string, fullMatch bool) bool {
	if strings.HasPrefix(path, m.pathSeparator) != strings.HasPrefix(pattern, m.pathSeparator) {
		return false
	}

	pattDirs := tokenizeToStringArray(pattern, m.pathSeparator, m.trimTokens, true)
	if fullMatch && !m.isPotentialMatch(path, pattDirs) {
		return false
	}

	pathDirs := tokenizeToStringArray(path, m.pathSeparator, m.trimTokens, true)

	var pattIdxStart = 0
	var pattIdxEnd = len(pattDirs) - 1
	var pathIdxStart = 0
	var pathIdxEnd = len(pathDirs) - 1

	// Match all elements up to the first **
	for pattIdxStart <= pattIdxEnd && pathIdxStart <= pathIdxEnd {
		patDir := pattDirs[pattIdxStart]
		if "**" == patDir {
			break
		}
		if !matchStrings(patDir, pathDirs[pathIdxStart]) {
			return false
		}
		pattIdxStart++
		pathIdxStart++
	}

	if pathIdxStart > pathIdxEnd {
		// Path is exhausted, only Match if rest of pattern is * or **'s
		if pattIdxStart > pattIdxEnd {
			if strings.HasSuffix(pattern, m.pathSeparator) {
				return strings.HasSuffix(path, m.pathSeparator)
			} else {
				return !strings.HasSuffix(path, m.pathSeparator)
			}
		}
		if !fullMatch {
			return true
		}
		if (pattIdxStart == pattIdxEnd) && ("*" == pattDirs[pattIdxStart]) && strings.HasSuffix(path, m.pathSeparator) {
			return true
		}
		for i := pattIdxStart; i <= pattIdxEnd; i++ {
			if "**" != pattDirs[i] {
				return false
			}
		}
		return true
	} else if pattIdxStart > pattIdxEnd {
		// string not exhausted, but pattern is. Failure.
		return false
	} else if !fullMatch && ("**" == pattDirs[pattIdxStart]) {
		// Path start definitely matches due to "**" in pattern.
		return true
	}

	// up to last '**'
	for pattIdxStart <= pattIdxEnd && pathIdxStart <= pathIdxEnd {
		patDir := pattDirs[pattIdxEnd]
		if "**" == patDir {
			break
		}
		if !matchStrings(patDir, pathDirs[pathIdxEnd]) {
			return false
		}
		pattIdxEnd--
		pathIdxEnd--
	}
	if pathIdxStart > pathIdxEnd {
		// string is exhausted
		for i := pattIdxStart; i <= pattIdxEnd; i++ {
			if !(pattDirs[i] == "**") {
				return false
			}
		}
		return true
	}

	for pattIdxStart != pattIdxEnd && pathIdxStart <= pathIdxEnd {
		patIdxTmp := -1
		for i := pattIdxStart + 1; i <= pattIdxEnd; i++ {
			if "*" == pattDirs[i] {
				patIdxTmp = i
				break
			}
		}
		if patIdxTmp == pattIdxStart+1 {
			// '**/**' situation, so skip one
			pattIdxStart++
			continue
		}
		// Find the pattern between padIdxStart & padIdxTmp in str between
		// strIdxStart & strIdxEnd
		patLength := patIdxTmp - pattIdxStart - 1
		strLength := pathIdxEnd - pathIdxStart - 1
		foundIdx := -1

	strLoop:
		for i := 0; i <= strLength-patLength; i++ {
			for j := 0; j < patLength; j++ {
				subPat := pattDirs[pattIdxStart+j+i]
				subStr := pathDirs[pattIdxStart+i+j]
				if !matchStrings(subPat, subStr) {
					continue strLoop
				}
			}
			foundIdx = pathIdxStart + i
			break
		}
		if foundIdx == -1 {
			return false
		}
		pattIdxStart = patIdxTmp
		pathIdxStart = foundIdx + patLength
	}
	for i := pattIdxStart; i <= pattIdxEnd; i++ {
		if !("**" == pattDirs[i]) {
			return false
		}
	}
	return true
}

func (m *AntPathMatcher) isPotentialMatch(path string, pattDirs []string) bool {
	if !m.trimTokens {
		pos := 0
		for _, pattDir := range pattDirs {
			skipped := skipSeparator(path, pos, m.pathSeparator)
			pos += skipped
			skipped = skipSegment(path, pos, pattDir)
			if skipped < len(pattDir) {
				return skipped > 0 || (len(pattDir) > 0 && isWildcardChar(pattDir[0]))
			}
			pos += skipped
		}
	}
	return true
}

func skipSegment(path string, pos int, prefix string) int {
	skipped := 0
	for index := range prefix {
		if isWildcardChar(prefix[index]) {
			return skipped
		}
		currPos := pos + skipped
		if currPos >= len(path) {
			return 0
		}
		if prefix[index] == path[currPos] {
			skipped += 1
		}
	}
	return skipped
}

func skipSeparator(path string, pos int, separator string) int {
	skipped := 0
	sepLn := len(separator)
	for strings.Index(path, separator) == pos {
		skipped += sepLn
	}
	return skipped
}

func isWildcardChar(c uint8) bool {
	for index := range wildcardChars {
		if c == wildcardChars[index] {
			return true
		}
	}
	return false
}

func tokenizeToStringArray(str string, delimiters string, trimTokens bool, ignoreEmptyTokens bool) []string {
	tokens := strings.Split(str, delimiters)
	for i := len(tokens) - 1; i >= 0; i -= 1 {
		if trimTokens {
			tokens[i] = strings.TrimSpace(tokens[i])
		}
		if ignoreEmptyTokens && len(tokens[i]) == 0 {
			tokens = append(tokens[:i], tokens[i+1:]...)
		}

	}
	return tokens
}

func matchStrings(pattern string, str string) bool {
	patArr := []byte(pattern)
	strArr := []byte(str)
	patIdxStart := 0
	patIdxEnd := len(patArr) - 1
	strIdxStart := 0
	strIdxEnd := len(strArr) - 1

	var b byte

	containsStar := false
	for _, aPatArr := range patArr {
		if aPatArr == '*' {
			containsStar = true
			break
		}
	}

	if !containsStar {
		// No '*'s, so we make a shortcut
		if patIdxEnd != strIdxEnd {
			// Pattern and string do not have the same size
			return false
		}
		for i := 0; i <= patIdxEnd; i++ {
			b := patArr[i]
			if b != '?' {
				if b != strArr[i] {
					// Character mismatch
					return false
				}
			}
		}
		return true
	}

	if patIdxEnd == 0 {
		// Pattern contains only '*', which matches anything
		return true
	}

	// Process characters before first star

	b = patArr[patIdxStart]
	for (b != '*') && strIdxStart <= strIdxEnd {
		if b != '?' {
			if b != strArr[strIdxStart] {
				return false
			}
		}
		patIdxStart++
		strIdxStart++
		b = patArr[patIdxStart]
	}
	if strIdxStart > strIdxEnd {
		// All characters in the string are used, check if only '*'s are
		// left in the pattern. If so, we succeeded.Otherwise failure
		for i := patIdxStart; i < patIdxEnd; i++ {
			if patArr[i] != '*' {
				return false
			}
		}
		return true
	}

	// Process characters after last star
	b = patArr[patIdxEnd]
	for (b != '*') && strIdxStart <= strIdxEnd {
		if b != '?' {
			if b != strArr[strIdxEnd] {
				return false
			}
		}
		patIdxEnd--
		strIdxEnd--
		b = patArr[patIdxEnd]
	}
	if strIdxStart > strIdxEnd {
		// All characters in the string are used, check if only '*'s are
		// left in the pattern. If so, we succeeded.Otherwise failure
		for i := patIdxStart; i < patIdxEnd; i++ {
			if patArr[i] != '*' {
				return false
			}
		}
		return true
	}

	for patIdxStart != patIdxEnd && strIdxStart <= strIdxEnd {
		patIdxTmp := -1
		for i := patIdxStart; i <= patIdxEnd; i++ {
			if patArr[i] == '*' {
				patIdxTmp = i
				break
			}
		}

		if patIdxTmp == patIdxStart+1 {
			// Two stars next to each other, skip the first one
			patIdxStart++
			continue
		}
		// Find the pattern between padIdxStart & padIdxTmp in str between
		// strIdxStart & strIdxEnd
		patLength := patIdxTmp - patIdxStart - 1
		strLength := strIdxEnd - strIdxStart - 1
		foundIdx := -1
	strLoop:
		for i := 0; i <= strLength-patLength; i++ {
			for j := 0; j < patLength; j++ {
				b = patArr[patIdxStart+j+1]
				if b != '?' {
					if b != strArr[strIdxStart+i+j] {
						continue strLoop
					}
				}
			}
			foundIdx = strIdxStart + i
			break
		}
		if foundIdx == -1 {
			return false
		}
		patIdxStart = patIdxTmp
		strIdxStart = foundIdx + patLength
	}

	// All characters in the string are used. Check if only '*'s are left
	// in the pattern. If so, we succeeded. Otherwise failure
	for i := patIdxStart; i <= patIdxEnd; i++ {
		if patArr[i] != '*' {
			return false
		}
	}
	return true
}

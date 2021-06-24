package path

//
//import (
//	"github.com/xuzhuoxi/infra-go/stringx"
//	"regexp"
//	"strings"
//)
//
//// AntPathMatcher专用
//var (
//	defaultPathSeparator2  = "/"
//	cacheTurnOffThreshold2 = 65536
//	variablePattern2       = regexp.MustCompile(`\{[^/]+?}`)
//	wildcardChars2         = []uint8{'*', '?', '{'}
//)
//
//// antPathStringMatcher专用
//var (
//	globPattern2            = regexp.MustCompile(`\?|\*|\{((?:\{[^/]+?}|[^/{}]|\\[{}])+?)}`)
//	defaultVariablePattern2 = "(.*)"
//)
//
//func NewAntPathMatcher2() *AntPathMatcher2 {
//	return NewAntPathMatcher2BySeparator(defaultPathSeparator2)
//}
//
//func NewAntPathMatcher2BySeparator(pathSeparator string) *AntPathMatcher2 {
//	return &AntPathMatcher2{
//		pathSeparator:             pathSeparator,
//		pathSeparatorPatternCache: newPathSeparatorPatternCache(pathSeparator)}
//}
//
//type AntPathMatcher2 struct {
//	pathSeparator             string
//	pathSeparatorPatternCache *pathSeparatorPatternCache
//	caseSensitive             bool
//	trimTokens                bool
//	cachePatterns             bool
//	tokenizedPatternCache     map[string][]string
//	stringMatcherCache        map[string]*antPathStringMatcher2
//}
//
//func (m *AntPathMatcher2) SetPathSeparator(pathSeparator string) {
//	m.pathSeparator = pathSeparator
//	m.pathSeparatorPatternCache = newPathSeparatorPatternCache(pathSeparator)
//}
//
//func (m *AntPathMatcher2) SetCaseSensitive(caseSensitive bool) {
//	m.caseSensitive = caseSensitive
//}
//
//func (m *AntPathMatcher2) SetTrimTokens(trimTokens bool) {
//	m.trimTokens = trimTokens
//}
//
//func (m *AntPathMatcher2) SetCachePatterns(cachePatterns bool) {
//	m.cachePatterns = cachePatterns
//}
//
//func (m *AntPathMatcher2) IsPattern(path string) bool {
//	if len(path) == 0 {
//		return false
//	}
//	uriVar := false
//	for index := range path {
//		if path[index] == '*' || path[index] == '?' {
//			return true
//		}
//		if path[index] == '{' {
//			uriVar = true
//			continue
//		}
//		if path[index] == '}' && uriVar {
//			return true
//		}
//	}
//	return false
//}
//
//func (m *AntPathMatcher2) Match(pattern string, path string) bool {
//	return m.doMatch(pattern, path, true, nil)
//}
//
//func (m *AntPathMatcher2) MatchStart(pattern string, path string) bool {
//	return m.doMatch(pattern, path, false, nil)
//}
//
//func (m *AntPathMatcher2) doMatch(pattern string, path string, fullMatch bool, uriTemplateVariables map[string]string) bool {
//	if len(path) == 0 || strings.Index(path, m.pathSeparator) != strings.Index(pattern, m.pathSeparator) {
//		return false
//	}
//
//	pattDirs := m.tokenizePattern(pattern)
//	if fullMatch && m.caseSensitive && !m.isPotentialMatch(path, pattDirs) {
//		return false
//	}
//
//	pathDirs := m.tokenizePath(path)
//	pattIdxStart := 0
//	pattIdxEnd := len(pattDirs) - 1
//	pathIdxStart := 0
//	pathIdxEnd := len(pattDirs) - 1
//
//	// Match all elements up to the first **
//	for pattIdxStart <= pattIdxEnd && pathIdxStart <= pathIdxEnd {
//		pattDir := pattDirs[pattIdxStart]
//		if "**" == pattDir {
//			break
//		}
//		if !m.matchStrings(pattDir, pathDirs[pathIdxStart], uriTemplateVariables) {
//			return false
//		}
//		pattIdxStart++
//		pathIdxStart++
//	}
//
//	if pathIdxStart > pathIdxEnd {
//		// Path is exhausted, only match if rest of pattern is * or **'s
//		if pattIdxStart > pattIdxEnd {
//			return stringx.EndWith(pattern, m.pathSeparator) == stringx.EndWith(path, m.pathSeparator)
//		}
//		if !fullMatch {
//			return true
//		}
//		if pattIdxStart == pattIdxEnd && pattDirs[pattIdxStart] == "*" && stringx.EndWith(path, m.pathSeparator) {
//			return true
//		}
//		for i := pattIdxStart; i <= pattIdxEnd; i += 1 {
//			if pattDirs[i] != "**" {
//				return false
//			}
//		}
//		return true
//	} else if pattIdxStart > pattIdxEnd {
//		// String not exhausted, but pattern is. Failure.
//		return false
//	} else if !fullMatch && "**" == pattDirs[pattIdxStart] {
//		// Path start definitely matches due to "**" part in pattern.
//		return true
//	}
//
//	// up to last '**'
//	for pattIdxStart <= pattIdxEnd && pathIdxStart <= pathIdxEnd {
//		pattDir := pattDirs[pattIdxEnd]
//		if pattDir == "**" {
//			break
//		}
//		if !m.matchStrings(pattDir, pathDirs[pathIdxEnd], uriTemplateVariables) {
//			return false
//		}
//		pattIdxEnd--
//		pathIdxEnd--
//	}
//
//	if pathIdxStart > pathIdxEnd {
//		// String is exhausted
//		for i := pattIdxStart; i <= pattIdxEnd; i += 1 {
//			if pattDirs[i] != "**" {
//				return false
//			}
//		}
//		return true
//	}
//
//	for pattIdxStart != pattIdxEnd && pathIdxStart <= pathIdxEnd {
//		patIdxTmp := -1
//		for i := pattIdxStart + 1; i <= pattIdxEnd; i += 1 {
//			if pattDirs[i] == "**" {
//				patIdxTmp = i
//				break
//			}
//		}
//		if patIdxTmp == pattIdxStart+1 {
//			// '**/**' situation, so skip one
//			pattIdxStart++
//			continue
//		}
//		// Find the pattern between padIdxStart & padIdxTmp in str between
//		// strIdxStart & strIdxEnd
//		patLength := (patIdxTmp - pattIdxStart - 1)
//		strLength := (pathIdxEnd - pathIdxStart + 1)
//		foundIdx := -1
//
//	strLoop:
//		for i := 0; i <= strLength-patLength; i += 1 {
//			for j := 0; j < patLength; j += 1 {
//				subPat := pattDirs[pattIdxStart+j+1]
//				subStr := pathDirs[pathIdxStart+i+j]
//				if !m.matchStrings(subPat, subStr, uriTemplateVariables) {
//					continue strLoop
//				}
//			}
//			foundIdx = pathIdxStart + i
//			break
//		}
//
//		if foundIdx == -1 {
//			return false
//		}
//
//		pattIdxStart = patIdxTmp
//		pathIdxStart = foundIdx + patLength
//	}
//
//	for i := pattIdxStart; i <= pattIdxEnd; i += 1 {
//		if pattDirs[i] != "**" {
//			return false
//		}
//	}
//
//	return true
//}
//
//func (m *AntPathMatcher2) matchStrings(pattern string, str string, uriTemplateVariables map[string]string) bool {
//	return m.getStringMatcher(pattern).MatchStrings(str, uriTemplateVariables)
//}
//
//func (m *AntPathMatcher2) getStringMatcher(pattern string) *antPathStringMatcher2 {
//	if m.cachePatterns {
//		if cacheMatcher, ok := m.stringMatcherCache[pattern]; ok {
//			return cacheMatcher
//		}
//	}
//	matcher := newAntPathStringMatcherByCaseSensitive(pattern, m.caseSensitive)
//	if m.cachePatterns && len(m.stringMatcherCache) >= cacheTurnOffThreshold2 {
//		// Try to adapt to the runtime situation that we're encountering:
//		// There are obviously too many different patterns coming in here...
//		// So let's turn off the cache since the patterns are unlikely to be reoccurring.
//		m.deactivatePatternCache()
//		return matcher
//	}
//	if m.cachePatterns {
//		m.stringMatcherCache[pattern] = matcher
//	}
//	return matcher
//}
//
//func (m *AntPathMatcher2) isPotentialMatch(path string, pattDirs []string) bool {
//	if !m.trimTokens {
//		pos := 0
//		for _, pattDir := range pattDirs {
//			skipped := m.skipSeparator(path, pos, m.pathSeparator)
//			pos += skipped
//			skipped = m.skipSegment(path, pos, pattDir)
//			if skipped < len(pattDir) {
//				return skipped > 0 || (len(pattDir) > 0 && m.isWildcardChar(pattDir[0]))
//			}
//			pos += skipped
//		}
//	}
//	return true
//}
//
//func (m *AntPathMatcher2) skipSegment(path string, pos int, prefix string) int {
//	skipped := 0
//	for index := range prefix {
//		if m.isWildcardChar(prefix[index]) {
//			return skipped
//		}
//		currPos := pos + skipped
//		if currPos >= len(path) {
//			return 0
//		}
//		if prefix[index] == path[currPos] {
//			skipped += 1
//		}
//	}
//	return skipped
//}
//
//func (m *AntPathMatcher2) isWildcardChar(c uint8) bool {
//	for index := range wildcardChars2 {
//		if c == wildcardChars2[index] {
//			return true
//		}
//	}
//	return false
//}
//
//func (m *AntPathMatcher2) skipSeparator(path string, pos int, separator string) int {
//	skipped := 0
//	sepLn := len(separator)
//	for strings.Index(path, separator) == pos {
//		skipped += sepLn
//	}
//	return skipped
//}
//
//func (m *AntPathMatcher2) tokenizePattern(pattern string) []string {
//	var tokens []string = nil
//	if m.cachePatterns {
//		if v, ok := m.tokenizedPatternCache[pattern]; ok {
//			tokens = v
//		}
//	}
//	if nil == tokens {
//		tokens = m.tokenizePath(pattern)
//		if m.cachePatterns && len(m.tokenizedPatternCache) >= cacheTurnOffThreshold2 {
//			m.deactivatePatternCache()
//			return tokens
//		}
//		if m.cachePatterns {
//			m.tokenizedPatternCache[pattern] = tokens
//		}
//	}
//	return tokens
//}
//
//func (m *AntPathMatcher2) deactivatePatternCache() {
//	m.cachePatterns = false
//	m.tokenizedPatternCache = make(map[string][]string)
//	m.stringMatcherCache = make(map[string]*antPathStringMatcher2)
//}
//
//func (m *AntPathMatcher2) tokenizePath(path string) []string {
//	return stringx.SplitToken(path, m.pathSeparator, m.trimTokens, true)
//}
//
////----------------
//
//func newAntPathStringMatcher(pattern string) *antPathStringMatcher2 {
//	return newAntPathStringMatcherByCaseSensitive(pattern, true)
//}
//
//func newAntPathStringMatcherByCaseSensitive(pattern string, caseSensitive bool) *antPathStringMatcher2 {
//	rs := &antPathStringMatcher2{}
//	rs.init(pattern, caseSensitive)
//	return rs
//}
//
//type antPathStringMatcher2 struct {
//	pattern       *regexp.Regexp
//	variableNames []string
//}
//
//func (m *antPathStringMatcher2) init(pattern string, caseSensitive bool) {
//}
//
//func (m *antPathStringMatcher2) quote(s string, start int, end int) string {
//	if start == end {
//		return ""
//	}
//	return P
//}
//
//func (m *antPathStringMatcher2) MatchStrings(str string, uriTemplateVariables map[string]string) bool {
//}
//
//func newPathSeparatorPatternCache(pathSeparator string) *pathSeparatorPatternCache {
//	return &pathSeparatorPatternCache{
//		endsOnWildCard:       pathSeparator + "*",
//		endsOnDoubleWildCard: pathSeparator + "**"}
//}
//
//type pathSeparatorPatternCache struct {
//	endsOnWildCard       string
//	endsOnDoubleWildCard string
//}
//
//func (c *pathSeparatorPatternCache) GetEndsOnWildCard() string {
//	return c.endsOnWildCard
//}
//
//func (c *pathSeparatorPatternCache) GetEndsOnDoubleWildCard() string {
//	return c.endsOnDoubleWildCard
//}

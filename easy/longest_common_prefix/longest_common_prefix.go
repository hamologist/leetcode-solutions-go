package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := ""
	baseStr := strs[0]
	for strIndex := 0; strIndex < len(baseStr); strIndex += 1 {
		for _, str := range strs[1:] {
			if strIndex >= len(str) || baseStr[strIndex] != str[strIndex] {
				return prefix
			}
		}

		prefix += string(baseStr[strIndex])
	}

	return prefix
}

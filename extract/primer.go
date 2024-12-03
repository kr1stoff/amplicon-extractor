package extract

/*
@description: expandDegenerateBases 将一个简并的 DNA 序列扩展为所有可能的序列
@param:

	sequence: 需要扩展的简并的 DNA 序列

@return:

	[]string: 包含所有可能的序列的切片
*/
func ExpandDegenerateBases(sequence string) []string {
	degenerateMap := map[rune][]rune{
		'R': {'A', 'G'},
		'Y': {'C', 'T'},
		'M': {'A', 'C'},
		'K': {'G', 'T'},
		'S': {'G', 'C'},
		'W': {'A', 'T'},
		'H': {'A', 'C', 'T'},
		'B': {'C', 'G', 'T'},
		'V': {'A', 'C', 'G'},
		'D': {'A', 'G', 'T'},
		'N': {'A', 'C', 'G', 'T'},
	}

	var expand func(seq string, idx int, current []string) []string

	/*
		@description: expand 函数用于递归地扩展简并碱基, 生成所有可能的序列
		@param:
			seq: 需要扩展的简并的 DNA 序列
			idx: 当前处理的索引位置
			current: 当前已生成的序列切片
		@return:
			[]string: 包含所有可能的序列的切片
	*/
	expand = func(seq string, idx int, current []string) []string {
		// 递归终止条件, 当索引到达序列末尾时返回当前结果
		if idx == len(seq) {
			return current
		}

		var expanded []string
		bases, ok := degenerateMap[rune(seq[idx])]

		if ok {
			// 如果当前字符是简并碱基, 则进行扩展
			if len(current) == 0 {
				// 第一个字符是简并碱基, 则将其扩展到当前结果, 数组大小 +1
				for _, base := range bases {
					expanded = append(expanded, string(base))
				}
			} else {
				// 后续字符是简并碱基, 则将其扩展到当前结果
				for _, s := range current {
					for _, base := range bases {
						expanded = append(expanded, s+string(base))
					}
				}
			}
		} else {
			// 如果当前字符不是简并碱基, 则直接添加到当前结果
			if len(current) == 0 {
				// 第一个字符不是简并碱基, 则直接添加到当前结果
				expanded = []string{string(seq[idx])}
			} else {
				// 后续字符不是简并碱基, 则直接添加到当前结果
				for i := range current {
					current[i] += string(seq[idx])
				}
				expanded = current
			}
		}

		// fmt.Println(expanded)
		// 递归调用, 索引 +1
		return expand(seq, idx+1, expanded)
	}

	return expand(sequence, 0, nil)
}

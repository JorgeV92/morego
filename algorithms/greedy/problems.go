package greedy

type Problem string

const (
	OneSwap      Problem = "one-swap"
	LargestSubsK Problem = "largest-subs-k"
)

func ProblemNames() []string {
	return []string{
		string(OneSwap),
		string(LargestSubsK),
	}
}

func LargestSwap(s string) string {
	b := []byte(s)
	maxDigit := byte('0')
	maxIdx := -1
	l, r := -1, -1

	for i := len(b) - 1; i >= 0; i-- {
		if b[i] > maxDigit {
			maxDigit = b[i]
			maxIdx = i
		} else if b[i] < maxDigit {
			l = i
			r = maxIdx
		}
	}
	if l == -1 {
		return s
	}
	b[l], b[r] = b[r], b[l]
	return string(b)
}

func LargestSubsAtLeastK(s string, k int) string {
	b := []byte(s)
	start := 0
	ans := make([]byte, len(b))
	for start < len(b) {
		var cnt [26]int
		var last [26]int
		for i := 0; i < 26; i++ {
			last[i] = -1
		}
		for i := start; i < len(b); i++ {
			idx := b[i] - 'a'
			cnt[idx]++
			last[idx] = i
		}
		ch := -1
		for c := 25; c >= 0; c-- {
			if cnt[c] >= k {
				ch = c
				break
			}
		}
		if ch == -1 {
			break
		}
		target := byte(ch) + 'a'
		for i := start; i <= last[ch]; i++ {
			if b[i] == target {
				ans = append(ans, b[i])
			}
		}
		start = last[ch] + 1
	}
	return string(ans)
}

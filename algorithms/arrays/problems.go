package arrays

import "sort"

type Problem string

const (
	ThreeSumProblem    Problem = "three-sum"
	ArrayChangeProblem Problem = "array-change"
)

func ProblemNames() []string {
	return []string{
		string(ThreeSumProblem),
		string(ArrayChangeProblem),
	}
}

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	values := append([]int(nil), nums...)
	sort.Ints(values)

	result := make([][]int, 0)

	for i := 0; i < len(values)-2 && values[i] <= 0; i++ {
		if i > 0 && values[i] == values[i-1] {
			continue
		}

		left, right := i+1, len(values)-1
		for left < right {
			sum := values[i] + values[left] + values[right]

			switch {
			case sum < 0:
				left++
			case sum > 0:
				right--
			default:
				result = append(result, []int{values[i], values[left], values[right]})
				left++
				right--

				for left < right && values[left] == values[left-1] {
					left++
				}
				for left < right && values[right] == values[right+1] {
					right--
				}
			}
		}
	}

	return result
}

func ThreeSumTargetExists(nums []int, target int) bool {
	if len(nums) < 3 {
		return false
	}

	values := append([]int(nil), nums...)
	sort.Ints(values)

	for i := 0; i < len(values)-2; i++ {
		left, right := i+1, len(values)-1

		for left < right {
			sum := values[i] + values[left] + values[right]

			switch {
			case sum < target:
				left++
			case sum > target:
				right--
			default:
				return true
			}
		}
	}

	return false
}

func ArrayChange(nums []int, operations [][2]int) []int {
	result := append([]int(nil), nums...)
	indexByValue := make(map[int]int, len(result))

	for index, value := range result {
		indexByValue[value] = index
	}

	for _, operation := range operations {
		from, to := operation[0], operation[1]
		index := indexByValue[from]
		result[index] = to
		indexByValue[to] = index
		delete(indexByValue, from)
	}

	return result
}

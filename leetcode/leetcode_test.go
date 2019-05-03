package leetcode

import (
	"math/rand"
	"testing"
)

func TestLeetCode(t *testing.T) {
	// t.Run("test873", test873)
	// t.Run("test874", test874)
	// t.Run("test877", test877)
	// t.Run("test878", test878)
	// t.Run("test879", test879)
	// t.Run("test880", test880)
	// t.Run("test883", test883)
	// t.Run("test885", test885)
	t.Run("test887", test887)
}

func test873(t *testing.T) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	result := 6
	return_result := openLock(deadends, target)
	if result != return_result {
		t.Fatalf("deadends: %v target: %v result: %d return_result: %d", deadends, target, result, return_result)
	}

	// deadends := []string{"8888"}
	// target := "0009"
	// result := 1
	// return_result := openLock(deadends, target)
	// if result != return_result {
	// 	t.Fatalf("deadends: %v target: %v result: %d return_result: %d", deadends, target, result, return_result)
	// }

	// deadends := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	// target := "8888"
	// result := -1
	// return_result := openLock(deadends, target)
	// if result != return_result {
	// 	t.Fatalf("deadends: %v target: %v result: %d return_result: %d", deadends, target, result, return_result)
	// }

	// deadends := []string{"0000"}
	// target := "8888"
	// result := -1
	// return_result := openLock(deadends, target)
	// if result != return_result {
	// 	t.Fatalf("deadends: %v target: %v result: %d return_result: %d", deadends, target, result, return_result)
	// }

	// deadends := []string{"6586", "6557", "0399", "3436", "1106", "4255", "1161", "7546", "2375", "5535", "7623", "0805", "7045", "8244", "1804", "1777", "5152", "7241", "4488", "3653", "7485", "9103", "2726", "4624", "8654", "1404", "9321", "5145", "4237", "5423", "9350", "3383", "8658", "2601", "2446", "1605", "6804", "1521", "0832", "5555", "6710", "3851", "6370", "0069", "7369", "6352", "4165", "4327", "9727", "1363", "9906", "9463", "8628", "5239", "0009", "2743", "0419", "4722", "7251", "5645", "5159", "4040", "1406", "5836", "0623", "9851", "2970", "0479", "1707", "5248", "0135", "8840", "9395", "1068", "9653", "4461", "6830", "7851", "7798", "3745", "1608", "2061", "5404", "3536", "3875", "3552", "8430", "0846", "5575", "2835", "1777", "5848", "5181", "8129", "2408", "3257", "9168", "3279", "4705", "9799", "1592", "7849", "4934", "1210", "0384", "3946", "5200", "3702", "4792", "1363", "0340", "4623", "9837", "0798", "2400", "0859", "3002", "1819", "2925", "8966", "7065", "3310", "1415", "9986", "7612", "1233", "9681", "6869", "5324", "4271", "1632", "2947", "8829", "9102", "9502", "4896", "2556", "4998", "7642", "8477", "4439", "8391", "7171", "2081", "5401", "0369", "4498", "1269", "2535", "7805", "6611", "1605", "1432", "6237", "5565", "9618", "2123", "5178", "3649", "8657", "6236", "6737", "1561", "1802", "1349", "9738", "6245", "7202", "8442", "7183", "5105", "7963", "0259", "5622", "3098", "0664", "7366", "1556", "5711", "9981", "4607", "2063", "7540", "1818", "7320", "8505", "1028", "6127", "1816", "8961", "7126", "4739", "4050", "7729", "5887", "4836", "1244", "2697", "3937", "9817", "2759", "9536", "0154", "7214", "5688", "1284", "5434", "7103", "2704", "6790", "3244", "8797", "3860", "1988", "1458", "4268", "1901", "4787", "7599", "6672", "3579", "3726", "6670", "1603", "3332", "7249", "0984", "6783", "4456", "0023", "2678", "0167", "8626", "6080", "5716", "5083", "6135", "8700", "7890", "8683", "2089", "0264", "2123", "0787", "3056", "2647", "4645", "8748", "6936", "6899", "0031", "4934", "0221", "9481", "9959", "1386", "7695", "2034", "0466", "0809", "9166", "6381", "6937", "0744", "8059", "8498", "5772", "8379", "4448", "5794", "7423", "2568", "4671", "6408", "4335", "1655", "3662", "1250", "5262", "7197", "6831", "8004", "0575", "8784", "2920", "0869", "7157", "0153", "7255", "1541", "1247", "5498", "0566", "6632", "7640", "1733", "2546", "5110", "2852", "8042", "8175", "0284", "8589", "8918", "5755", "2289", "0812", "4850", "4650", "9018", "6649", "5099", "6532", "9891", "8675", "1718", "5442", "6786", "8915", "3710", "3833", "2659", "7040", "3959", "2505", "7574", "1199", "3465", "4557", "7230", "9161", "5177", "7815", "4564", "1470", "8051", "6287", "2504", "4025", "8911", "6158", "6857", "8948", "7991", "3670", "3413", "0423", "7184", "7921", "1351", "8908", "1921", "1685", "5579", "4641", "0286", "6410", "2800", "7018", "1402", "7410", "3471", "1312", "9530", "4581", "5364", "4820", "8192", "3088", "4714", "2255", "2342", "5042", "8673", "9788", "2203", "0879", "2345", "9712", "2008", "0652", "0939", "0720", "2954", "4482", "2390", "0807", "4634", "6266", "5222", "6898", "7491", "0294", "1811", "0667", "8282", "5754", "1841", "9518", "9093", "7904", "4902", "0068", "5157", "7823", "8073", "8801", "8179", "1402", "9977", "2332", "9448", "2251", "8455", "6157", "1878", "4183", "3331", "8047", "1254", "9639", "2156", "5780", "7359", "0260", "9683", "6842", "1098", "6495", "2057", "6583", "0932", "2577", "1818", "6042", "8358", "1833", "5512", "4529", "0583", "9955", "9205", "6055", "3322", "2232", "5372", "5835", "2202", "9696", "1596", "3424", "3696", "5695", "1365", "6432", "0327", "1565", "8509", "6936", "3363", "3007", "3107", "0410", "6258", "2492", "0300", "1255", "1664", "8666", "6826", "9961", "5782", "0140", "5567", "9596", "1680", "1892", "5016", "8804", "4962", "9318", "4540", "5044", "0979", "2004", "4265", "7689", "0289", "3434", "6090", "1375", "3135", "3935", "5140", "9255", "3997", "3482", "8150", "8164", "0787"}
	// target := "8828"
	// result := 8
	// return_result := openLock(deadends, target)
	// if result != return_result {
	// 	t.Fatalf("deadends: %v target: %v result: %d return_result: %d", deadends, target, result, return_result)
	// }

	// deadends := []string{"0101", "1001", "0001", "1100"}
	// target := "2101"
	// result := 4
	// return_result := openLock(deadends, target)
	// if result != return_result {
	// 	t.Fatalf("deadends: %v target: %v result: %d return_result: %d", deadends, target, result, return_result)
	// }

	// deadends := []string{"7867", "6676", "8687", "7886", "6768", "8877", "6767", "6676", "6666", "7876", "6688", "6677", "6877", "7786", "6778", "6868", "6868", "7867", "7668", "7666", "8868", "7887", "6788", "7687", "7788", "7877", "6867", "6867", "7876", "8787", "8878", "6668", "6878", "6766", "8667", "8688", "6788", "7687", "8887", "8766", "6867", "8867", "7866", "7866", "6686", "7776", "8687", "7888", "6777", "6678", "6678", "6686", "6677", "7886", "6876", "8666", "6667", "7768", "7688", "7668", "6786", "7766", "7867", "8866", "7887", "6676", "8776", "6867", "8888", "6678", "8687", "6868", "7888", "8666", "6678", "6668", "7678", "7667", "8786", "8768", "6766", "8776", "8677", "7788", "7868", "7878", "6786", "6678", "6876", "7667", "8866", "8666", "8768", "8886", "8787", "8688", "8766", "8867", "7886", "6876", "7776", "7867", "8668", "7777", "8888", "7767", "8778", "8888", "6876", "8777", "7877", "8866", "8668", "8878", "7678", "8787", "7788", "8887", "8667", "7887", "6686", "8778", "7768", "8787", "7677", "6768", "7877", "7788", "7768", "6768", "6786", "7887", "7768", "6676", "6777", "8686", "7867", "8788", "8887", "8776", "7677", "8786", "8678", "7666", "8776", "7676", "6767", "8776", "8888", "8766", "8876", "7777", "7677", "6767", "7878", "7868", "8677", "7677", "8788", "6667", "8866", "8887", "6686", "6777", "6676", "8787", "6788", "8866", "6767", "8676", "8868", "8768", "8888", "7866", "7877", "7768", "7686", "7888", "6666", "6887", "6787", "7667", "6676", "8666", "8886", "8878", "8678", "8868", "8888", "8867", "7878", "7787", "8776", "7877", "6788", "8778", "6768", "8677", "8678", "6778", "7888", "6866", "6768", "6666", "6887", "8866", "7676", "7866", "7876", "7678", "7686", "8887", "7676", "6788", "8787", "6666", "8866", "6876", "8676", "8688", "8887", "7887", "7777", "8887", "8688", "6668", "6686", "6887", "7677", "6867", "6786", "6877", "7788", "6667", "8778", "8786", "8767", "7778", "8867", "8877", "6668", "8886", "7888", "7767", "7666", "8678", "8668", "8767", "7666", "6787", "6886", "8787", "6886", "8768", "8767", "8676", "6767", "8776", "8768", "8687", "8778", "7888", "6768", "7878", "6668", "7688", "6687", "7866", "8878", "6877", "7667", "8886", "7876", "6667", "8877", "7666", "7668", "7676", "6888", "6686", "7666", "7688", "7666", "6678", "6676", "7678", "8788", "7667", "7767", "8766", "6867", "8767", "8676", "8786", "8667", "6678", "6778", "8877", "8788", "6866", "7687", "6876", "8878", "8866", "6788", "6877", "8768", "8778", "8778", "8866", "7866", "7887", "7878", "8766", "8778", "7868", "8787", "6676", "8668", "7866", "8787", "8767", "6876", "8867", "6688", "6886", "6668", "6878", "7866", "8678", "8867", "7667", "7878", "8778", "8777", "7866", "8878", "7868", "6876", "7688", "7677", "7678", "7777", "8888", "8776", "8688", "6878", "8877", "7678", "7777", "7878", "6678", "6688", "6868", "8876", "6668", "8877", "8786", "6688", "8766", "8887", "6678", "8886", "8876", "8888", "8878", "6786", "7686", "7867", "7767", "7888", "8866", "6876", "7767", "6687", "6687", "6688", "6868", "8668", "6886", "8686", "7766", "8777", "8667", "8886", "7676", "7768", "6788", "8688", "7676", "7686", "8777", "7886", "7788", "6666", "7687", "6676", "6777", "6866", "6767", "7787", "7877", "6777", "6886", "7877", "7787", "7787", "8768", "7787", "8778", "6766", "7677", "6788", "6786", "6767", "8687", "6687", "8668", "6876", "6666", "7676", "8667", "6688", "6766", "6677", "7667", "8668", "8866", "7686", "8866", "8687", "8866", "8768", "7886", "6877", "8877", "6676", "6887", "6788", "8877", "8887", "8886", "8887", "6676", "8867", "6867", "7768", "8868", "6668", "7878", "7887", "8768", "6876", "7787", "7876", "8886", "6778", "7778", "7687", "6686", "7787", "8767", "8668", "7686", "7678", "8788", "6687", "8666", "7877", "6668", "7686", "6866", "6888", "8786", "7778", "7786", "8787", "6777", "6867", "7787", "7777", "6766", "8666", "6778", "6867", "8668", "8667", "7678", "8668", "7677", "8787", "6876", "6668", "7788", "7688", "7687", "8778", "8787", "8688", "8867"}
	// target := "6776"
	// result := 16
	// return_result := openLock(deadends, target)
	// if result != return_result {
	// 	t.Fatalf("deadends: %v target: %v result: %d return_result: %d", deadends, target, result, return_result)
	// }
}

func test874(t *testing.T) {
	start := 23
	max := 23
	for i := start; i <= max; i++ {
		n := numSquares(i)
		t.Logf("Input: %d output: %d\n", i, n)
	}
}

func test877(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	min := minStack.GetMin()
	t.Logf("min: %d\n", min)
	minStack.Pop()
	minStack.Top()
	min2 := minStack.GetMin()
	t.Logf("min2: %d\n", min2)
}

func test878(t *testing.T) {
	t.Logf("Input: %s Result: %v", "()[]{}", isValid("()[]{}"))
	t.Logf("Input: %s Result: %v", "(]", isValid("(]"))
	t.Logf("Input: %s Result: %v", "([)]", isValid("([)]"))
	t.Logf("Input: %s Result: %v", "{[]}", isValid("{[]}"))
	t.Logf("Input: %s Result: %v", "]", isValid("]"))
}

func test879(t *testing.T) {
	t.Logf("Input: %v Result: %v", []int{73, 74, 75, 71, 69, 72, 76, 73}, dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	t.Logf("Input: %v Result: %v", []int{73, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 89}, dailyTemperatures([]int{73, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 89}))
}

func test880(t *testing.T) {
	t.Logf("Input: %v Result: %v", []string{"2", "1", "+", "3", "*"}, evalRPN([]string{"2", "1", "+", "3", "*"}))
	t.Logf("Input: %v Result: %v", []string{"4", "13", "5", "/", "+"}, evalRPN([]string{"4", "13", "5", "/", "+"}))
	t.Logf("Input: %v Result: %v", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func test883(t *testing.T) {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	t.Logf("Input: %s output %d\n", grid, numIslands(grid))

	grid2 := [][]byte{}
	t.Logf("Input: %s output %d\n", grid2, numIslands(grid2))

	grid3 := [][]byte{{'0'}}
	t.Logf("Input: %s output %d\n", grid3, numIslands(grid3))

	grid4 := [][]byte{{'1'}}
	t.Logf("Input: %s output %d\n", grid4, numIslands(grid4))

	grid5 := [][]byte{{'1', '0', '0'}, {'0', '1', '0'}, {'0', '0', '1'}}
	t.Logf("Input: %s output %d\n", grid4, numIslands(grid5))
}

func test885(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	s := 3
	t.Logf("Input: nums: %v s:%d Output: %d\n", nums, s, findTargetSumWays(nums, s))

	nums2 := []int{1, 1, 1, 1, 1}
	s2 := 0
	t.Logf("Input: nums: %v s:%d Output: %d\n", nums2, s2, findTargetSumWays(nums2, s2))

	nums3 := []int{}
	s3 := 0
	t.Logf("Input: nums: %v s:%d Output: %d\n", nums3, s3, findTargetSumWays(nums3, s3))

	nums4 := []int{}
	for i := 0; i < 20; i++ {
		nums4 = append(nums4, rand.Intn(1000))
	}
	s4 := 10
	t.Logf("Input: nums: %v s:%d Output: %d\n", nums4, s4, findTargetSumWays(nums4, s4))
}

func test887(t *testing.T) {
	root := &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}}
	t.Logf("input: %v Output: %v", root, inorderTraversal(root))

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 10}}}},
	}
	t.Logf("input: %v Output: %v", root2, inorderTraversal(root2))

	root3 := &TreeNode{Val: 1}
	t.Logf("input: %v Output: %v", root3, inorderTraversal(root3))

	t.Logf("input: %v Output: %v", []int{}, inorderTraversal(nil))

}

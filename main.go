package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	for _, soln := range solns {
		fmt.Printf("%d. %s\n", soln.id, soln.desc)
		fmt.Printf("Answer: %s\n", soln.fn())
	}
}

type soln struct {
	id   int
	desc string
	fn   func() string
}

var solns = []soln{
	{
		id:   1,
		desc: "Multiples of 3 and 5",
		fn: func() string {
			var sum int
			for i := 3; i < 1000; i++ {
				if i%3 == 0 || i%5 == 0 {
					sum += i
				}
			}
			return fmt.Sprintf("%d", sum)
		},
	},
	{
		id:   2,
		desc: "Even Fibonacci numbers",
		fn: func() string {
			var sum int64 = 2
			for i, j, curr := 1, 2, 3; curr <= 4000000; i, j = j, curr {
				curr = i + j
				if curr%2 == 0 {
					sum += int64(curr)
				}
			}
			return fmt.Sprintf("%d", sum)
		},
	},
	{
		id:   3,
		desc: "Largest prime factor",
		fn: func() string {
			// TODO: implement solution
			return "unsolved"
		},
	},
	{
		id:   4,
		desc: "Largest palindrome product",
		fn: func() string {
			// TODO: implement solution
			return "unsolved"
		},
	},
	{
		id:   5,
		desc: "Smallest multiple",
		fn: func() string {
			var gcd = func(a int64, b int64) int64 {
				var t int64
				for b != 0 {
					t = b
					b, a = a%b, t
				}
				return a
			}

			var lcm = func(a int64, b int64) int64 {
				return (a * b / gcd(a, b))
			}

			var lcmm func(...int64) int64
			lcmm = func(ns ...int64) int64 {
				if len(ns) == 2 {
					return lcm(ns[0], ns[1])
				}
				newNs := ns[1:]
				return lcm(ns[0], lcmm(newNs...))
			}

			min := 1
			max := 20
			var ns []int64
			for i := min; i <= max; i++ {
				ns = append(ns, int64(i))
			}
			return fmt.Sprintf("%d", lcmm(ns...))
		},
	},
	{
		id:   6,
		desc: "Sum square difference",
		fn: func() string {
			var squareSum func(int64) int64
			squareSum = func(n int64) int64 {
				if n == 0 {
					return 0
				}
				return (n * n) + squareSum(n-1)
			}

			var naturalSum func(int64) int64
			naturalSum = func(n int64) int64 {
				if n == 0 {
					return 0
				}
				return n + naturalSum(n-1)
			}

			var n int64 = 100
			sqSum := squareSum(n)
			natSum := naturalSum(n)
			natSumSq := natSum * natSum
			diff := natSumSq - sqSum
			return fmt.Sprintf("%d", diff)
		},
	},
	{
		id:   7,
		desc: "10001st prime",
		fn: func() string {
			var isPrime = func(n int64) bool {
				if n <= 1 {
					return false
				} else if n <= 3 {
					return true
				} else if n%2 == 0 || n%3 == 0 {
					return false
				}

				var i int64 = 5
				for i*i <= n {
					if n%i == 0 || n%(i+2) == 0 {
						return false
					}
					i += 6
				}

				return true
			}

			var n int64 = 10001
			var k int64 = 0
			var cnt int64 = 1
			for k < n {
				if isPrime(cnt) {
					k++
				}
				cnt++
			}

			return fmt.Sprintf("%d", cnt-1)
		},
	},
	{
		id:   8,
		desc: "Largest product in a series",
		fn: func() string {

			var getProd = func(segment string) int {
				parts := strings.Split(segment, "")
				sum := 1
				for _, s := range parts {
					n, _ := strconv.Atoi(s)
					sum *= n
				}
				return sum
			}

			var spaceMap = func(str string) string {
				return strings.Map(func(r rune) rune {
					if unicode.IsSpace(r) {
						return -1
					}
					return r
				}, str)
			}

			var getGreatestProd = func(digits string, cnt int) int {
				//trim white spaces
				digits = spaceMap(digits)
				length := len(digits)
				great := 0
				for i := 0; i < length && i+cnt < length; i++ {
					segment := string(digits[i : i+cnt])
					if strings.IndexByte(segment, '0') < 0 {
						prod := getProd(segment)
						if prod > great {
							great = prod
						}
					}
				}
				return great
			}

			const digits = `
				73167176531330624919225119674426574742355349194934
				96983520312774506326239578318016984801869478851843
				85861560789112949495459501737958331952853208805511
				12540698747158523863050715693290963295227443043557
				66896648950445244523161731856403098711121722383113
				62229893423380308135336276614282806444486645238749
				30358907296290491560440772390713810515859307960866
				70172427121883998797908792274921901699720888093776
				65727333001053367881220235421809751254540594752243
				52584907711670556013604839586446706324415722155397
				53697817977846174064955149290862569321978468622482
				83972241375657056057490261407972968652414535100474
				82166370484403199890008895243450658541227588666881
				16427171479924442928230863465674813919123162824586
				17866458359124566529476545682848912883142607690042
				24219022671055626321111109370544217506941658960408
				07198403850962455444362981230987879927244284909188
				84580156166097919133875499200524063689912560717606
				05886116467109405077541002256983155200055935729725
				71636269561882670428252483600823257530420752963450
			`

			return fmt.Sprintf("%d", getGreatestProd(digits, 4))
		},
	},
}

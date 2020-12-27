###  本周总结

本周主要对DP问题进行了一个汇总，但是还没整理完

##### 说在前面

定义状态还是不算难的，但是如何推导状态转移方程，我觉得很难找到一个很好的切入点。网上也看了很多高赞回答，虽然他们的题解头头是道，讲的比较让初学者接受，但终究不能解决我心中的疑惑——“为什么他们会这么想，而我却毫无思路？”。事实上，很难有一篇题解能告诉你从无到有的这个思考过程。这个问题贯穿了我整个刷题过程，尤其是处理DP问题的时候，直到有一天我在查资料时看到了[这篇文章](https://oi-wiki.org/dp/basic/)，不知为何，我突然就想明白了：事实上我们根本不需要怀疑自己的智商，会的人之所以会，因为他们压根不是第一次遇到这类问题（这里的第一次是真正意义上的第一次，没有任何前置知识）。看到ACM选手整理的那篇文章，总结了如此之多的DP套路、模板，我们现在做的leetcode题，充其量只是ACM训练的入门题（当然也是经典题）。只有经过了大量的训练，才能对这一类问题形成自己的思路。现在我没有过这方面的训练，又怎么可能有思路呢？所以对于DP入门级选手，重要的不是自己绞劲脑汁想思路，而是记住这些经典问题的思路（当然了，能完全自己想出来那就更好了）。所以，刷题的思路要转变过来。


已整理题目：
- [一维DP](#一维dp)
    - [509. 斐波那契数](#509-斐波那契数)
    - [70. 爬楼梯](#70-爬楼梯)
    - [746. 使用最小花费爬楼梯](#746-使用最小花费爬楼梯)
    - [322. 零钱兑换](#322-零钱兑换)
    - [198. 打家劫舍](#198-打家劫舍)
    - [213. 打家劫舍 II](#213-打家劫舍-ii)
    - [53. 连续子数组的最大和](#53-连续子数组的最大和)
    - [152. 连续子数组的最大乘积](#152-连续子数组的最大乘积)
- [二维DP](#二维dp)
  - [字符串处理DP](#字符串处理dp)
    - [5. 最长回文子串](#5-最长回文子串)
    - [516. 最长回文子序列](#516-最长回文子序列)
    - [300. 最长递增子序列](#300-最长递增子序列)
    - [1143. 最长公共子序列](#1143-最长公共子序列)
  - [常见背包问题](#常见背包问题)
  - [股票系列问题](#股票系列问题)


待整理题目列表

3. [300. Longest Increasing Subsequence](https://leetcode-cn.com/problems/longest-increasing-subsequence/) [五星]
4. [1143. Longest Common Subsequence](https://leetcode-cn.com/problems/longest-common-subsequence/) [五星]
5. [10. Regular Expression Matching](https://leetcode-cn.com/problems/regular-expression-matching/) [五星+++]
6. [44. Wildcard Matching](https://leetcode-cn.com/problems/wildcard-matching/) [五星+++]
9. [62. Unique Paths](https://leetcode-cn.com/problems/unique-paths/)
10. [63. Unique Paths II](https://leetcode-cn.com/problems/unique-paths-ii/)
11. [64. Minimum Path Sum](https://leetcode-cn.com/problems/minimum-path-sum/) [和62题是一样的]
12. [174. Dungeon Game](https://leetcode-cn.com/problems/dungeon-game/) 
13. [741. Cherry Pickup](https://leetcode-cn.com/problems/cherry-pickup/) [round trip, 暂时放弃...]
14. [120. Triangle](https://leetcode-cn.com/problems/triangle/) [第64题的变形, 自底向上/自顶向下]
18. [84. Largest Rectangle in Histogram](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/) [这一题不属于dp问题, 但是和85题是强相关的]
19. [85. Maximal Rectangle](https://leetcode-cn.com/problems/maximal-rectangle/) [五星++, 84题的follow-up]
20. [221. Maximal Square](https://leetcode-cn.com/problems/maximal-square/)
21. [32. Longest Valid Parentheses](https://leetcode-cn.com/problems/longest-valid-parentheses/) [这一题比较独立, 五星++]


---

### 动态规划关键点

* 状态：F(i)
* 最优子结构：F(i) = best_of(F(i-1), F(i-2),...)
* 状态转移方程：

---



# 一维DP



### [509. 斐波那契数](https://leetcode-cn.com/problems/fibonacci-number/)

斐波那契数，通常用 `F(n)` 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

```
F(0) = 0
F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
给定 N（0 ≤ N ≤ 30 ），计算 F(N)
```

分析：这是学习「递归」最经典的习题。我们从”暴力递归 --> 递归+备忘录 --> 动态规划 --> 动态规划+空间压缩“这4层来加以说明。

**方法1**：暴力递归，直接根据斐波那契数列的定义写出代码。这种方法会出现大量的冗余计算，需要指数级时间，时间复杂度为`O(2^n)`。

```java
class Solution {
    public int fib(int n) {
        if(n <= 1) return n;
        return fib(n-1) + fib(n-2);
    }
}
```

**方法2**：由于暴力递归会出现大量重复计算，因此，我们可以考虑通过”备忘录“把已经计算过的值记录下来，当再次需要它的时候，就直接返回结果。

```java
class Solution {
    public int fib(int n) {
        int[] cache = new int[100];
        return helper(n, cache);
    }

    private int helper(int n, int[] cache) {
        if(n <= 1) return n;
        if(cache[n] != 0) return cache[n]; // 直接返回结果
        cache[n] = helper(n-1, cache) + helper(n-2, cache); // 先把本次计算的结果存放在缓存中
        return cache[n];
    }
}
```

**方法3**：动态规划。递归的做法本质是「自顶向下」，而动态规划的做法其实是「自底向上」。

所谓自顶向下，比如说，我要求`fib(5)`，发现`fib(5)`需要用到`fib(4)`和`fib(3)`，那么接着去求`fib(4)`和`fib(3)`，一层一层下去直到遇到递归出口。就是说，「为了完成当前的任务，我需要材料a，但是发现材料a还没有，于是就先去准备材料a；为了准备材料a，发现它又需要材料b，而材料b又tm没有，于是又得先去准备材料b...」。而自底向上表示的是，「我事先把材料a, b都准备好了，从而可以制作出材料c，而利用材料b和c又可以制作出材料d，...，于是完成了最终的任务」。

```java
// time: O(n)
// space: O(n)
class Solution {
    public int fib(int n) {
        if(n <= 1) return n;
        int[] dp = new int[n+1];
        dp[0] = 0;
        dp[1] = 1;
        for(int i = 2; i <= n; i++) {
            dp[i] = dp[i-1] + dp[i-2];
        }
        return dp[n];
    }
}
```

**方法4**：动态规划+空间压缩。我们发现在方法3中，更新`dp[i]`只需要`dp[i-1]` 和` dp[i-2]`，也就是说，实际上不需要一整个dp数组，因此可以把空间复杂度从`O(n)`降为`O(1)`。

```java
// time: O(n)
// space: O(1)
class Solution {
    public int fib(int n) {
        if(n <= 1) return n;
        int pre = 0, cur = 1;
        for(int i = 2; i <= n; i++) {
            int tmp = cur;
            cur = pre + cur;
            pre = tmp;
        }
        return cur;
    }
}
```



### [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

假设你正在爬楼梯。需要 *n* 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？**注意：**给定 *n* 是一个正整数。

```java
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
```

分析：本题本质上就是斐波那契数列问题。

首先，定义状态。令`dp[i]`表示到达第 `i` 个台阶的方法总数。

其次，初始化。`dp[0]`表示到达第 0 个台阶的方法总数，其实就是站在起点处嘛，这里初始化为`dp[0]=1`（从实际情况来说，站着也是一种方法嘛，难不成还算0种方法？）；而`dp[1]`表示到达第 1 个台阶的方法数，显然只能是1，即通过爬1个台阶到达。**初始化的时候，对于一些比较模糊的边界情况，可以从n=2,3的情况返回去倒推**，加以确定。

最后，确定状态转移方程。由于规定每次可以爬1或2个台阶，因此，只可能从第`i-1`个台阶向上跨1步，或从第`i-2`个台阶向上跨2步到达第`i`个台阶。因此，也就推导出了状态转移方程（这不就是斐波那契数列的定义嘛~）：

```
dp[i] = dp[i-1] + dp[i-2]
```

实现如下：

```java
// time: O(n)
// space: O(n)
class Solution {
    public int climbStairs(int n) {
        if(n <= 1) return 1;
        int[] dp = new int[n+1];
        dp[0] = 1;
        dp[1] = 1;
        for(int i = 2; i <= n; i++) {
            dp[i] = dp[i-1] + dp[i-2];
        }
        return dp[n];
    }
}

// time: O(n)
// space: O(1)
class Solution {
    public int climbStairs(int n) {
        if(n <= 1) return 1;
        int pre= 1, cur = 1;
        for(int i = 2; i <= n; i++) {
            int tmp = cur;
            cur = pre + cur;
            pre = tmp;
        }
        return cur;
    }
}
```



### [746. 使用最小花费爬楼梯](https://leetcode-cn.com/problems/min-cost-climbing-stairs/)

数组的每个索引作为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 `cost[i]`(索引从0开始)。每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。

```
输入: cost = [10, 15, 20]
输出: 15
解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。

输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出: 6
解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费6。
```

**注意：**

1. `cost` 的长度将会在 `[2, 1000]`。
2. 每一个 `cost[i]` 将会是一个Integer类型，范围为 `[0, 999]`。

分析：如下图所示：

![lc746](../img/lc746.png)

本题是「70. 爬楼梯」的进阶版。对于”如何爬“的逻辑是不变的，即要到达第 i 级台阶，仅可能从第 i-1 级台阶向上爬，或从第 i-2 级台阶向上爬。但是本题还要考虑另外一个变量，即爬楼梯的花费值，如果从第 i 级台阶向上跳（不管跳1级还是2级），其花费值为`cost[i]`。

首先，定义状态。我们令`dp[i]`表示到达第 i 级台阶所需的花费，显然，最终的结果存放在`dp[n]`中。

其次，初始化。本题已经说明，”可以选择从索引为 0 或 1 的元素作为初始阶梯“，因此可以得出`dp[0] = 0`和`dp[1] = 0`。

最后，确定状态转移方程。我们还是从第70题的逻辑出发，想要到达第 i 级台阶，仅可能从第 i-1 级台阶向上爬，或从第 i-2 级台阶向上爬。如果从第 i-1 级台阶向上爬，那么所需要的花费是 `dp[i-1] + cost[i-1]`；如果从第 i-2 级台阶向上爬，则需要的花费是`dp[i-2] + cost[i-2]`。由于本题考虑的是最小花费，而不是方法总数，因此两者取较小者，于是得到状态转移方程：

```
dp[i] = min{dp[i-1]+cost[i-1], dp[i-2]+cost[i-2]}
```

代码如下：

```java
// time: O(n)
// space: O(n)
class Solution {
    public int minCostClimbingStairs(int[] cost) {
        int[] dp = new int[cost.length+1];
        dp[0] = 0;
        dp[1] = 0;
        for(int i = 2; i <= cost.length; i++) {
            dp[i] = Math.min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2]);
        }
        return dp[cost.length];
    }
}

// time: O(n)
// space: O(1)
class Solution {
    public int minCostClimbingStairs(int[] cost) {
        int pre2 = 0, pre1 = 0, cur = 0;
        for(int i = 2; i <= cost.length; i++) {
            cur = Math.min(pre1+cost[i-1], pre2+cost[i-2]);
            pre2 = pre1;
            pre1 = cur;
        }
        return cur;
    }
}
```



### [322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)

给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的**最少的硬币个数**。如果没有任何一种硬币组合能组成总金额，返回 -1。

你可以认为**每种硬币的数量是无限的**。

```
输入：coins = [1, 2, 5], amount = 11
输出：3 
解释：11 = 5 + 5 + 1
```

分析：本题是「背包问题」的经典例题。但是背包的解法会在背包问题系列中统一介绍，这里从其他几个角度进行说明。

**方法1**：从爬楼梯问题延伸而来。

在「70. 爬楼梯」问题中，为了到达第 i 级台阶，可以从第 i-1 级台阶向上爬1级，或第 i-2 级台阶向上爬2级。类似的，在零钱兑换问题中，也可以采用同样的思路。以`coins = [1, 2, 5]`为例，**为了凑成金额 i，那么可以由”金额为 i-1 加上1，或金额 i-2 加上2，或金额 i-5 加上5“拼凑而来**。区别在于，爬楼梯问题中，当求 i 时，只有 i-1 或 i-2 两种情况需要考虑，而在本问题中，需要考虑n种情况，即考虑一个数组coins。

首先，定义状态。令`dp[i]`表示凑成金额 i 需要的硬币个数。显然，最终的答案存放在`dp[amount]`中。

其次，初始化。`dp[0] = 0`表示凑成金额0需要0个硬币，这应该不难理解。

最后，确定状态转移方程。由上面的分析，可以类比爬楼梯问题写出状态转移方程。还是以`coins = [1, 2, 5]`为例，即：

```
dp[i] = min{dp[i-1], dp[i-2], dp[i-5]} + 1
```

推广之，得到：

```
dp[i] = min{dp[i-coins[k]]} + 1, where 0 ≤ k ≤ n and i-coins[k] ≥ 0
```

代码如下：

```java
// time: O(amount × coins.length)
// space: O(amount)
class Solution {
    private final int MAX = 999999;
    public int coinChange(int[] coins, int amount) {
        int[] dp = new int[amount+1];
        dp[0] = 0;
        for(int i = 1; i <= amount; i++) {
            int minCount = MAX;
            for(int k = 0; k < coins.length; k++) {
                if(i - coins[k] < 0) continue;
                minCount = Math.min(minCount, dp[i-coins[k]]);
            }
            if(minCount == MAX){
                dp[i] = MAX; // 如果minCount没有变化，说明数组coins中无法凑成金额i，于是令dp[i]=MAX，表示”无法凑成“
            }else {
                dp[i] = minCount + 1;
            }
        }
        return dp[amount] == MAX ? -1 : dp[amount];
    }
}
```



### [198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/)

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，**如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警**。给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

```
输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。偷窃到的最高金额 = 2 + 9 + 1 = 12 。
```

分析：对于每一个元素，都存在偷或不偷两种选择，但需要特别注意的是，在本题的情况下，如果第 i 个元素已经被偷了，那么第 i+1 个元素就肯定不能偷了。

首先，定义状态。**令`F(i)`表示到第 i 个元素为止所能偷到的最大金额**，对于第 i 个元素，既可以是偷的，也可以是不偷的，`F(i)`表示的是考虑了这两种情况下的最大值。因此，最终结果存放在`F(n-1)`处。

其次，初始化。当只有 1 个元素时，肯定是偷的，即 `F(0) = nums[0]`；当有 2 个元素时，由于不能偷取相邻的两个元素（即只能2选1），因此此时选择较大的元素，即`F(1) = max{nums[0], nums[1]}`。

最后，定义状态转移方程。对于第 i 个元素：

* 如果偷它，那么前一个元素肯定是不能偷的，即：`F(i) = F(i-2) + nums[i]`
* 如果不偷它，那么：`F(i) = F(i-1)`

状态转移方程如下：

```
F(i) = max{F(i-2) + nums[i], F(i-1)}
```

Java版

```Java
// time: O(n)
// space: O(n)
class Solution {
    public int rob(int[] nums) {
        if(nums.length == 0) return 0;
        if(nums.length == 1) return nums[0];
        int[] dp = new int[nums.length];
        dp[0] = nums[0];
        dp[1] = Math.max(nums[0], nums[1]);
        for(int i = 2; i < nums.length; i++) {
            dp[i] = Math.max(dp[i-2]+nums[i], dp[i-1]);
        }
        return dp[dp.length-1];
    }
}
```

Go版

```go
func rob(nums []int) int {
	// 特判
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	// 初始化
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	// 状态转移
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(dp)-1]
}

func max(x, y int) int {
	if x >= y {
		return x
	}else {
		return y
	}
}
```

**改进**：注意观察上面的状态转移方程，当更新`dp[i]`时，我们只用到了`dp[i-1]`和`dp[i-2]`。因此，可以进一步压缩空间，使空间复杂度降为常数级别。如下，改进后的代码更加简洁，但可能对于初学者不太好理解。

```Java
// time: O(n)
// space: O(1)
class Solution {
    public int rob(int[] nums) {
        if(nums.length == 0) return 0;
        if(nums.length == 1) return nums[0];
        int pre = 0, cur = nums[0];
        for(int i = 1; i < nums.length; i++){
            int tmp = cur;
            cur = Math.max(pre + nums[i], cur);
            pre = tmp;
        }
        return cur;
    }
}
```



### [213. 打家劫舍 II](https://leetcode-cn.com/problems/house-robber-ii/)

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。**这个地方所有的房屋都围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的**。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下 ，能够偷窃到的最高金额。

```
输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
```

分析：本题属于「198. 打家劫舍」的进阶版，在前一题中，元素是单向排列的；而在本题中，元素是环装排列的。由于首尾相连，选择了第一个元素就势必无法选择最后一个元素，选择了最后一个元素就一定无法选择第一个元素。因此，我们可以把问题拆分成两个子问题，即分别对数组`nums[0:n-2]`和数组`nums[1:n-1]`进行考虑，然后选择较大值就是本题的解了。

```java
class Solution {
    public int rob(int[] nums) {
        if(nums.length == 0) return 0;
        if(nums.length == 1) return nums[0];
        return Math.max(helper(nums, 0, nums.length-2),
                        helper(nums, 1, nums.length-1));
    }
		// 即「198. 打家劫舍」问题
    private int helper(int[] nums, int begin, int end) {
        int pre = 0, cur = nums[begin];
        for(int i = begin+1; i <= end; i++) {
            int tmp = cur;
            cur = Math.max(pre + nums[i], cur);
            pre = tmp;
        }
        return cur;
    }
}
```



### [53. 连续子数组的最大和](https://leetcode-cn.com/problems/maximum-subarray/)

给定一个整数数组 `nums` ，找到一个具有**最大和的连续子数组**（子数组最少包含一个元素），返回其最大和。

```
输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
```

分析：对于子数组之和的问题，我们首先需要认识到：

* 如果数组中的元素**全是非负数**，问题就变得没什么意义了，因为具有最大和的连续子数组就是整个数组。
* 如果数组中的元素**全是负数**，问题也很简单，具有最大和的连续子数组**一定只包含一个元素**，也就是最大的那个元素。（如果题目允许子数组为空，那么具有最大和的子数组就是空数组）
* 相同的子数组之和可能对应有多个子数组。

首先，定义状态。令`dp[i]`表示到下标 i 为止（包括`nums[i]`）的子数组的最大和。通过遍历整个数组，找到最大值。

其次，初始化。对于第 0 个元素，显然有`dp[0] = nums[0]`。

最后，确定状态转移方程。当计算以第 i 个元素结尾的子数组之和时：

* 如果之前的子数组之和（即`dp[i-1]`）是负的，那么显然**不会**加上之前的部分，否则越加越小了啊，而只取当前这一个元素作为子数组，即`dp[i] = nums[i]`；
* 如果之前的子数组之和（即`dp[i-1]`）是正的，那么会加上之前的部分，即`dp[i] = dp[i-1] + nums[i]`。

因此，可推导出状态转移方程如下：

```
dp[i] = max{dp[i-1]+nums[i], nums[i]}
```

代码实现

```java
// time: O(n)
// space: O(n)
class Solution {
    public int maxSubArray(int[] nums) {
        if(nums.length == 0) return 0;
        int[] dp = new int[nums.length];
        dp[0] = nums[0];
        int max = dp[0];
        for(int i = 1; i < nums.length; i++) {
            dp[i] = Math.max(dp[i-1]+nums[i], nums[i]);
            max = Math.max(max, dp[i]);
        }
        return max;
    }
}
```

**进阶**：观察发现，当更新`dp[i]`时，只需要`dp[i-1]`，因此我们可以压缩空间，将空间复杂度由`O(n)`降为`O(1)`。如下所示：

```java
// time: O(n)
// space: O(1)
class Solution {
    public int maxSubArray(int[] nums) {
        if(nums.length == 0) return 0;
        int pre = nums[0], max = nums[0];
        for(int i = 1; i < nums.length; i++) {
            int cur = Math.max(pre+nums[i], nums[i]);
            max = Math.max(max, cur);
            pre = cur;
        }
        return max;
    }
}
```

另外，这里再介绍一下 [Kadane's algorithm](https://en.wikipedia.org/wiki/Maximum_subarray_problem#Kadane's_algorithm) ，该算法用于求解子数组的最大和问题。

该算法由左至右扫描数组，当扫描至第 i 个元素时，计算以 `nums[i]` 为结尾的子数组最大和，由变量`currSum`表示。

在循环结构中，当计算第 i 步时，旧的`currSum`值维护了以`nums[i-1]`为结尾的子数组最大和，其中`currSum = nums[j]+nums[j+1]+...+nums[i-1], j ∈ [0,i-1]`。因此，`currSum+nums[i]`表示的是以`nums[i]`为结尾的子数组最大和，其中`currSum = nums[j]+nums[j+1]+...+nums[i], j ∈ [0,i]`。注意，当 j 取 i 的时候，即表示`currSum = nums[i]`，因此`currSum`应该表示为：

```
currSum = max{nums[i], currSum+nums[i]}
```

实现如下：（其实本质上就是DP方法的空间压缩版）

```java
// time: O(n)
// space: O(1)
class Solution {
    public int maxSubArray(int[] nums){
        int max = Integer.MIN_VALUE, currSum = 0;
        for(int num : nums){
            currSum = Math.max(num, currSum + num);
            max = Math.max(max, currSum);
        }
        return max;
    }
}
```

**再次进阶：除了求解子数组的最大和，如果还要知道最大和子数组的区间范围，该如何求解呢？**

```java
// time: O(n)
// space: O(1)
class Solution {
    public int maxSubArray(int[] nums){
        int maxSum = Integer.MIN_VALUE, curSum = 0;
        int maxBegin = 0, maxEnd = 0, curBegin = 0;
        for(int i = 0; i < nums.length; i++) {
            if(curSum <= 0) { // 以当前元素为首重新开始一个子数组
                curBegin = i;
                curSum = nums[i];
            }else { // 扩展原区间
                curSum += nums[i];
            }
            // 更新最大值
            if(curSum > maxSum) {
                maxSum = curSum;
                maxBegin = curBegin;
                maxEnd = i;
            }
        }
        System.out.printf("max sub array: [%d, %d], maxSubSum = %d\n", maxBegin, maxEnd, maxSum);
        return maxSum;
    }
}
```



### [152. 连续子数组的最大乘积](https://leetcode-cn.com/problems/maximum-product-subarray/)

给你一个整数数组 nums ，请你找出数组中**乘积最大的连续子数组**（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

```
示例 1:
输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。

示例 2:
输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
```

分析：对于子数组的乘积问题，首先需要明白：

* 如果上一轮的最大乘积`curProduct`是一个负数，当前值`nums[i]`也是一个负数，那么相乘之后的值可能是一个很大的数；
* 如果上一轮的最大乘积`curProduct`是一个正(负)数，而当前值`nums[i]`是一个负(正)数，那么相乘之后的值可能是一个很小的数；

**由于相乘会出现”负负得正“的情况**，我们不能再像「53. 连续子数组的最大和」问题一样只维护一个当前状态了。

由左至右扫描数组，当扫描至第 i 个元素时，分别计算以 `nums[i]` 为结尾的子数组的最大乘积（`curMax`）和最小乘积（`curMin`）。其中，

```
curMax = max(nums[j] × nums[j+1] × ... × nums[i]), where j ∈ [0,i]
curMin = min(nums[k] × nums[k+1] × ... × nums[i]), where k ∈ [0,i]
```

在循环结构中，当计算第 i 步时：

* 若 nums[i] < 0, curMax = curMin × nums[i], curMin = curMax_origin × nums[i]
* 若 nums[i] > 0, curMax = curMax × nums[i], curMin = curMin × nums[i]

实现如下：

Java版

```java
// time: O(n)
// space: O(1)
class Solution {
    public int maxProduct(int[] nums) {
        int curMin = 1, curMax = 1, maxProduct = Integer.MIN_VALUE;
        for(int num : nums) {
            if(num <= 0) {
                int curMaxOrigin = curMax;
                curMax = Math.max(num, curMin * num);
                curMin = Math.min(num, curMaxOrigin * num);
            }else {
                curMax = Math.max(num, curMax * num);
                curMin = Math.min(num, curMin * num);
            }
            maxProduct = Math.max(maxProduct, curMax);
        }
        return maxProduct;
    }
}
```

Go版

```go
func maxProduct(nums []int) int {
	curMin, curMax, maxProd := 1, 1, math.MinInt64 // 注意math.MinInt64的表达
	for _, num := range nums {
		if num <= 0 {
			curMaxOrigin := curMax
			curMax = max(num, curMin*num)
			curMin = min(num, curMaxOrigin*num)
		}else {
			curMax = max(num, curMax*num)
			curMin = min(num, curMin*num)
		}
		maxProd = max(maxProd, curMax)
	}
	return maxProd
}

func max(x, y int) int {
    if x > y {
        return x
    }else {
        return y
    }
}

func min(x, y int) int {
    if x < y {
        return x
    }else {
        return y
    }
}
```



# 二维DP

## 字符串处理DP

### [5. 最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

```
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
```

分析：

**方法1：动态规划**

首先，定义状态。令`dp[i][j]`表示子串`s[i...j]`是否为回文子串。如果是，则标记为true，反之标记为false。

其次，初始化。当只有一个字符时，显然是回文子串。即`dp[i][i] = true`；当有两个字符时，`dp[i][j] = s[i]==s[j]`。

最后，确定状态转移方程。对于任意一个位置 i 和 j，有：

```
         / true,  when s[i]==s[j] and dp[i+1][j-1] == true
dp[i][j]=
         \ false, when s[i]!=s[j]
```

Java版

```java
// time: O(n^2)
// space: O(n^2)
class Solution {
    public String longestPalindrome(String s) {
        int n = s.length(), begin = 0, end = 0;
        boolean[][] dp = new boolean[n][n];
        for(int i = n-1; i >= 0; i--) { // 这里必须从后向前
            for(int j = i; j < n; j++) { // 这里必须从前向后
                // 初始化
                if(j - i <= 1) {
                    dp[i][j] = s.charAt(i) == s.charAt(j);
                    if(dp[i][j] && j - i > end - begin) {
                        begin = i;
                        end = j;
                    }
                    continue;
                }
                // 状态转移
                if(s.charAt(i) == s.charAt(j) && dp[i+1][j-1]) {
                    dp[i][j] = true;
                }else {
                    dp[i][j] = false;
                }
                // 更新结果
                if(dp[i][j] && j - i > end - begin) {
                    begin = i;
                    end = j;
                }
            }
        }
        //System.out.printf("[%d, %d]\n", begin, end);
        return s.substring(begin, end+1);
    }
}
```

Go版

```go
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n) // go语言中二维数组的声明比较烦，也可以使用var dp [x][y]bool，但x,y必须是constant
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	begin, end := 0, 0
	for i := n-1; i >= 0; i-- {
		for j := i; j < n; j++ {
			// 初始化
			if j - i <= 1 {
				dp[i][j] = s[i] == s[j]
				if dp[i][j] && j - i > end - begin {
					begin = i
					end = j
				}
				continue
			}
			// 状态转移
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}else {
				dp[i][j] = false
			}
			// 更新结果
			if dp[i][j] && j - i > end - begin {
				begin = i
				end = j
			}
		}
	}
	return s[begin:end+1]
}
```



**方法2：中心扩展法**。该方法更加符合直觉，外层循环作为中心点，内层循环分别向两侧扩散，如果符合回文串的特性，则一直扩展下去。需要考虑的点在于——如何处理偶数个、奇数个元素的问题？代码实现起来也非常的简洁。

```java
// time: O(n^2)
// space: O(1)
class Solution {
    private int begin = 0, end = 0;
    public String longestPalindrome(String s) {
        for(int i = 0; i < s.length(); i++) {
            extend(s, i, i); // 考虑奇数个元素
            extend(s, i, i+1); // 考虑偶数个元素
        }
        return s.substring(begin, end+1);
    }

    private void extend(String s, int l, int r) {
        while(l >= 0 && r < s.length() && s.charAt(l) == s.charAt(r)) {
            l--;
            r++;
        }
        l++;
        r--;
        if(r - l > end - begin) { // 更新结果
            begin = l;
            end = r;
        }
    }
}
```



**方法3：[Manacher 算法](https://en.wikipedia.org/wiki/Longest_palindromic_substring)**

该方法可以在 O(n) 时间复杂度下完成，解法暂时略。



### [516. 最长回文子序列](https://leetcode-cn.com/problems/longest-palindromic-subsequence/)

给定一个字符串 `s` ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 `s` 的最大长度为 `1000` 。

```
示例 1:
输入:"bbbab"
输出:4
最长回文子序列为 "bbbb"。 // 注意，子序列不要求连续，子串要求连续
```

分析：

首先，定义状态。令`dp[i][j]`表示子序列`s[i...j]`中最大回文子序列的长度。那么，最终结果存放在`dp[0][n-1]`中。

其次，初始化。令`dp[i][i] = 1`

最后，确定状态转移方程。

* 如果`s[i]==s[j]`，显然有`dp[i][j] = dp[i+1][j-1]+2`。
* 如果`s[i]!=s[j]`，则`dp[i][j] = max{dp[i+1][j], dp[i][j-1]}`，考虑到这一点是本题的关键！

如下图所示：

![lc516](../img/lc516.png)

从而可推导出状态转移方程为：

```
					/ dp[i+1][j-1] + 2,            when s[i] == s[j] 
dp[i][j] = 
          \ max{dp[i+1][j], dp[i][j-1]}, when s[i] != s[j] 
```

首先，可以明确的是，由于`dp[i][j]`表示子序列`s[i...j]`中回文子序列的长度，因此一定有`j ≥ i`。另外，从状态转移方程可以知道，要求解`dp[i][j]`，需要提前知道`dp[i+1][j-1]`、`dp[i+1][j]`和`dp[i][j-1]`，这为我们如何编写代码指明了方向。如下图所示：

<img src="../img/lc516_2.png" alt="lc516_2" style="zoom:50%;" />

Java版

```java
// time: O(n^2)
// space: O(n^2)
class Solution {
    public int longestPalindromeSubseq(String s) {
        int n = s.length();
        int[][] dp = new int[n][n];
        char[] ss = s.toCharArray(); 
        for(int i = n-1; i >= 0; i--) { // 从下至上
            for(int j = i; j < n; j++) { // 从左向右
                // 初始化
                if(j == i) { 
                    dp[i][j] = 1;
                    continue;
                }
                // 状态转移
                if(ss[i] == ss[j]) { // faster than s.charAt(i) == s.charAt(j)
                    dp[i][j] = dp[i+1][j-1] + 2;
                }else {
                    dp[i][j] = Math.max(dp[i+1][j], dp[i][j-1]);
                }
            }
        }
        return dp[0][n-1];
    }
}
```

Go版

```go
func longestPalindromeSubseq(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
    }

    for i := n-1; i >=0; i-- {
        for j := i; j < n; j++ {
            // 初始化
            if j == i {
                dp[i][j] = 1
                continue
            }
            // 状态转移
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1] + 2
            }else {
                dp[i][j] = max(dp[i+1][j], dp[i][j-1])
            }
        }
    }
    return dp[0][n-1]
}

func max(x, y int) int {
    if x > y {
        return x
    }else {
        return y
    }
}
```



进阶：采用**滚动数组**压缩空间，将空间复杂度由$O(n^2)$下降为$O(n)$。

```java
// time: O(n^2)
// space: O(n)
class Solution {
    public int longestPalindromeSubseq(String s) {
        int n = s.length();
        int[] dp = new int[n];
        char[] ss = s.toCharArray(); 
        for(int i = n-1; i >= 0; i--) { // 从下至上
            int pre = dp[i]; // 关键
            for(int j = i; j < n; j++) { // 从左向右
                // 初始化
                if(j == i) { 
                    dp[j] = 1;
                    continue;
                }
                // 状态转移
                int tmp = dp[j]; // 关键
                if(ss[i] == ss[j]) { 
                    dp[j] = pre + 2;
                }else {
                    dp[j] = Math.max(dp[j], dp[j-1]);
                }
                pre = tmp;
            }
        }
        return dp[n-1];
    }
}
```



### [300. 最长递增子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)

给你一个整数数组 nums ，找到其中最长**严格递增**子序列的长度。

```
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
```

提示：

* 1 <= nums.length <= 2500
* $-10^4$ <= nums[i] <= $10^4$


进阶：

* 你可以设计时间复杂度为 $O(n^2)$ 的解决方案吗？你能将算法的时间复杂度降低到 $O(n log(n))$ 吗?

分析：





### [1143. 最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)

给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。

```
示例 1:
输入：text1 = "abcde", text2 = "ace" 
输出：3  
解释：最长公共子序列是 "ace"，它的长度为 3。

示例 2:
输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0。
```

分析：





## 常见背包问题





## 股票系列问题





---

参考：<https://www.zhihu.com/collection/168876650>


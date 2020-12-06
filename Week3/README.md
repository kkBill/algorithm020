## 本周学习笔记

（1）递归代码模板
```java
public void recursion(int level, int param) {
    // 1. 递归边界
    if(level > MAX_LEVEL) {
        // 处理结果
        return
    }

    // 2. 处理当前层的逻辑
    process(level, param)

    // 3. 进入到下一层
    recursion(level+1, newParam)

    // 4. 如果有必要，在这里进行状态重置或者资源清理
    ...
}
```

（2）分治代码模板
```java
public void divide_conquer(Problem problem, int param) {
    // 1. 递归边界
    if 达到边界条件 {
        // 处理结果
        return
    }

    // 2. 处理当前问题，将问题分割成多个子问题（相当于进入”下一层“）
    subResult1 = divide_conquer(subProblem1, newParam1)
    subResult2 = divide_conquer(subProblem2, newParam2)
    ...

    // 3. 合并结果
    result = process_result(subResult1, subResult2, ...)

    // 4. 如果有必要，在这里进行状态重置或者资源清理
    ...
}
```

关键点：
* 尽量避免人肉递归
* 寻找可重复性（最小重复子问题）
* 数学归纳法思维（从i=1,2,3的情况开始思考，然后思考i=n的情况）

---

本周主要整理了**回溯**相关的题目。笔记主要包括两个方面，首先，找了很多资料去补充理解回溯思想的本质，最后整理在自己的github仓库中。其次，整理了LeetCode中有关回溯的15道题。如下：

1. [回溯思想](https://github.com/kkBill/algorithm/blob/master/note/%E5%9B%9E%E6%BA%AF%E6%80%9D%E6%83%B3.md)
2. [常见回溯例题整理](https://github.com/kkBill/algorithm/blob/master/note/Backtracking.md)

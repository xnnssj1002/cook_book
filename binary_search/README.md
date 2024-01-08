## 二分搜索
### 二分搜索的经典写法。需要注意的三点：
1. 循环退出条件，注意是 low <= high，而不是 low < high。
2. mid 的取值，mid := low + (high-low)>>1
3. low 和 high 的更新，low = mid + 1，high = mid - 1。

### 二分搜索的变种写法。有 4 个基本变种:
1. 查找第一个与 target 相等的元素，时间复杂度 O(logn)
2. 查找最后一个与 target 相等的元素，时间复杂度 O(logn)
3. 查找第一个大于等于 target 的元素，时间复杂度 O(logn)
4. 查找最后一个小于等于 target 的元素，时间复杂度 O(logn)
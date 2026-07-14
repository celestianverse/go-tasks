# 88. Merge Sorted Array

## Description

Given two integer arrays `nums1` and `nums2`, both sorted in non-decreasing order, and two integers `m` and `n`:

- `m` is the number of initialized elements in `nums1`
- `n` is the number of elements in `nums2`

`nums1` has a total length of `m + n`, where the first `m` elements should be merged and the last `n` elements are set to `0` and should be ignored.

Merge `nums2` into `nums1` so that `nums1` becomes a single sorted array in non-decreasing order.

## Examples

### Example 1

Input:
```text
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6], n = 3
```

Output:
```text
[1,2,2,3,5,6]
```

Explanation:
- Merge `[1,2,3]` and `[2,5,6]`
- Result is `[1,2,2,3,5,6]`

### Example 2

Input:
```text
nums1 = [1], m = 1
nums2 = [], n = 0
```

Output:
```text
[1]
```

### Example 3

Input:
```text
nums1 = [0], m = 0
nums2 = [1], n = 1
```

Output:
```text
[1]
```

Explanation:
- `m = 0` means `nums1` has no valid initialized elements
- The extra `0` is reserved space for the merged result

## Constraints

- `nums1.length == m + n`
- `nums2.length == n`
- `0 <= m, n <= 200`
- `1 <= m + n <= 200`
- `-10^9 <= nums1[i], nums2[j] <= 10^9`

## Follow-up

Can you implement an algorithm that runs in `O(m + n)` time and uses `O(1)` extra space?

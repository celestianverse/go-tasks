# 94. Binary Tree Inorder Traversal

## Description

Given the root of a binary tree, return the inorder traversal of its nodes' values.

## Examples

### Example 1

Input:
```text
root = [1,null,2,3]
```

Output:
```text
[1,3,2]
```

### Example 2

Input:
```text
root = [1,2,3,4,5,null,8,null,null,6,7,9]
```

Output:
```text
[4,2,6,5,7,1,3,9,8]
```

### Example 3

Input:
```text
root = []
```

Output:
```text
[]
```

### Example 4

Input:
```text
root = [1]
```

Output:
```text
[1]
```

## Constraints

- The number of nodes in the tree is in the range $[0, 100]$
- $-100 <= Node.val <= 100$

## Follow-up

A recursive solution is trivial. Can you implement it iteratively?

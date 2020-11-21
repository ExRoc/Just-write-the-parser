# 练习

忽略 `/*` 与 `*/` 之间的注释内容解析。

## version1

实现 `/*` 与 `*/` 的可嵌套注释，如：

```
// valid expression, the result is 4:
2 /* first comment /* second comment */ * 2 */ * 2

// invalid expression:
2 /* first comment /* second comment */ * 2 * 2
```

## version2

实现 `/*` 与 `*/` 的不可嵌套注释，如：

```
// valid expression, the result is 8:
2 /* first comment /* second comment */ * 2 * 2

// invalid expression:
2 /* first comment /* second comment */ * 2 */ * 2
```


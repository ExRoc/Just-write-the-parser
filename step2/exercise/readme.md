# 练习

## 题目

构造足够多的测试用例，使得使用 `split` 和 `single pass` 两种实现的运行时长达到毫秒级，并比较这两种实现 的效率。

## 比较结果

```
Time of ParsingBySplit: 495.392285ms
Time of ParsingByASinglePass: 984.182014ms
```

生成 $10^6$ 组满足条件的表达式，发现使用 `single pass` 方式实现的解析时长大约是 `split` 的两倍。


- [题目](https://www.acwing.com/problem/content/description/792/)
- [提交记录](https://www.acwing.com/problem/content/submission/792/)
- [讨论](https://www.acwing.com/problem/content/discussion/index/792/1/)
- [题解](https://www.acwing.com/problem/content/solution/792/1/)
- [视频讲解](https://www.acwing.com/problem/content/video/792/)

给定一个浮点数n，求它的三次方根。

### **输入格式**

共一行，包含一个浮点数n。

### **输出格式**

共一行，包含一个浮点数，表示问题的解。

注意，结果保留6位小数。

### **数据范围**

−10000≤n≤10000−10000≤n≤10000

### **输入样例：**

```
1000.00

```

### **输出样例：**

```
10.000000
```

## 思路

1. 考察浮点数二分. 
2. 不断缩小 l, r 的范围
3. 当求的值的精度大于 1e-8的时候， 可以认为是准确的

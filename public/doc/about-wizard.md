# 求解 Fibonacci 的正确姿势

![在这里插入图片描述](https://www.hostinger.com/tutorials/wp-content/uploads/sites/2/2017/06/what-is-403-forbidden-error-and-how-to-fix-it.jpg)

## 前言

吼。很久没有产出博客了，这几天都闲的很，可是产出博客的好时机啊，可是写什么呢？脑子里过的第一个是 fibonacci，好啦，就写它吧！

## Fibonacci (递归)

这玩意儿算是老朋友了，在大学，或者是刚学计算机基础的时候，递归的经典案例就是它，别被它的名字吓到，它就是普通的数列而已，它的通项公式是

```javascript
Fibonacci(n) = Fibonacci(n - 1) + Fibonacci(n - 2)
```

是很简单吧，随手就能写一个求 Fibonacci 第 n 项的方法

```javascript
const fibonacci = n => n <= 2 ? 1 : fibonacci(n - 1) + fibonacci(n - 2)
```

这里用了 lambda 表达式，没什么为什么，因为这样简洁一点哈哈哈

或者是骚一点能够用 Y不动点组合子

```javascript
(f => n => n < 2 ? 1 : f(f)(n - 1) + f(f)(n - 2))(f => n => n < 2 ? 1 : f(f)(n - 1) + f(f)(n - 2))(3)
```

一切看起来十分美好（一般这样说的话就是有事情要发生了），但是你会发现时间复杂度是`指数级`！！what ？ 你不关心时间复杂度？那好，跑一下` fibonacci(100)`吧

```js
fibonacci(100)
```

你会发现过了 `1min`，`2min`, `3min`，都没有跑出来，反正我测试的时候是耐心都没了，直接 `ctrl + c` 了，不过算一些靠前的项还比较好

## Fibonacci (动态规划)

递归的好处之一就是写起来比较方便，emmmm，但是递归的过程中可能会遇到函数栈桢的`push`，`pop`，这样的也是一笔小开销呢，如果有尾递归优化的话，在这里也没什么很大的作用。

看看这里的动态规划，那上面用了递归，动态规划我就用递推吧。

Emmmmm, 先看看为什么适合用动态规划，我得先找个工具画图

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190220175449360.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0hhb0RhV2FuZw==,size_16,color_FFFFFF,t_70)

因为懒，我只画了三层，但是足够表达意思了，我用相同颜色标注了计算重复的项，这里只有三层，后面还有更多，这就意味着，有很多个项我们已经计算过了！我们没有必要去计算重复的项，现在，我们知道了有最优子结构，重叠问题还有状态转移方程，于是，用动态规划吧。

```js
function fibonacciDP(n) {
  if (n < 3) {
    return 1;
  }

  let result,
    item1 = 1,
    item2 = 1;

  while (n-- >= 3) {
    result = item1 + item2;
    item2 = item1;
    item1 = result;
  }

  return result;
}
```

这里用了递推，用递归的话可能看上去更简洁一点。

下面来试试跑起来需要用多少时间吧，为了让 v8 预热一下（这里的环境是 node，为了防止 JIT 还没跑起来测试就结束了，需要进行一次预热），于是写了一个循环，根据随机数返回不同类型的值（为了让 V8 执行去优化）

```js
let n;
for (let i = 0; i < 1000000; i++) {
  n = (Math.random() * 100 + 1) & 1 ? "a" : 1;
}
```

测试

```js
console.time("DP");
fibonacciDP(10000000); // DP: 44.883ms
console.timeEnd("DP");
```

让上面的递归也跑一遍？不不不，那样得等到明天了。

## Fibonacci （矩阵的快速幂）

这个东西关于基础知识这里不再赘述了，有兴趣的少年可以去看看线代。

然后快速幂这个玩意儿的话，能够把 pow 变为 `log` 级别的时间复杂度，通常我们见到的`log`级别的时间复杂度大概在二分法见的比较多嘛，这里也有一部分相似。

看一个普通的快速幂（O(logN)）

```js
function pow(base, n) {
  let result = 1;

  while (n != 0) {
    if (n % 2) {
      result *= base;
    }

    base *= base;
    n = parseInt(n / 2);
  }
  return result;
}
```

普通的幂运算（O(N)）

```js
function normalPow(base, n) {
  let result = 1;

  while (n--) {
    result *= base;
  }

  return result;
}
```

矩阵的快速幂跟普通的快速幂的原理是一样的，但是操作对象变为了矩阵。

区别是，普通的幂运算的初始值为1，但是矩阵中，可以用单元矩阵来替代，在矩阵中，任何矩阵 * 单元矩阵都为 `这个矩阵本身`

```js
matrix * cellMatrix = matrix
```

然后再说 Fibonacci 跟矩阵的关系，Fibonacci 的递推关系文章最开头已经给出了，于是可以得到

```js
Fibonacci(n) = 1 * Fibonacci(n - 1) + 1 * Fibonacci(n - 2)
Fibonacci(n - 1) = 0 * Fibonacci(n - 2) + 1 * Fibonacci(n - 1)
```

于是得到（写到代码块里将就看吧）

```js
0    1       Fibonacci(n - 2)     Fibonacci(n - 1)
         *                     =
1    1       Fibonacci(n - 1)     Fibonacci(n)
```

这是递归式，于是可以得到 （ps：括号里的 `n - 1` 是前面那个矩阵的幂）

```js
0    1  (n - 1)      Fibonacci(0)     Fibonacci(n - 1)
                 *                 =
1    1               Fibonacci(1)     Fibonacci(n)
```

根据已知 `Fibonacci(0) = Fibonacci(1) = 1` 得到

```js
0    1  (n - 1)      1         Fibonacci(n - 1)
                 *        =
1    1               1         Fibonacci(n)
```

呼，这不就是前面那个矩阵的 `n - 1`次幂么

```js
0    1  (n - 1)

1    1
```

下面可以写一个快速幂版的 fibonacci 了（O(logN)）

```js
function fibonacci(n) {
  let unitMatrix = [[1, 0], [0, 1]];
  let originMatrix = [[0, 1], [1, 1]];
  let result = unitMatrix;

  while (n) {
    if (n & 1) {
      result = matrixMultiplication(result, originMatrix);
    }

    originMatrix = matrixMultiplication(originMatrix, originMatrix);

    n /= 2;
  }

  return result[1][0];
}
```

`ps 这里直接算了 n 次，取了 f(n)，matrixMultiplication 是一个矩阵相乘的函数，这个可以根据矩阵的乘法原则来自己写，这里就不再赘述了`

## 运行时间的比较

由于第一个的递归太慢了，就懒得跑了，主要是跑动态规划和快速幂的对比，当然我自己试了几遍，由于两个频度函数的交叉点太远了…… 导致需要很大的数才能让动态规划花费的时间比快速幂要多

下面为 `10000k` 项的时候

```js
DP: 55.604ms
fast power: 4.686ms
```


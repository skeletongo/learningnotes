//
// Created by admin on 2022/2/25.
//

// 预处理器

// 使用 C99 标准编译
#pragma c9x on

// 加载代码
// 形式一
#include <stdio.h> // 加载系统提供的文件
#include <math.h>
// 形式二
// #include "foo.h" // 加载用户提供的文件

// 宏定义(总的来说就是代码替换)

// 单行定义
#define HELLO "hello world"

// 多行定义
#define Hi "I am \
Tom"

// 使用宏定义宏
#define TWO 2
#define FOUR TWO*TWO

// 参数宏
#define SQUARE(X) X*X

// 获取输入的第一个字节
#define getchar() getc(stdin)

// 取两个数中最大值
#define MAX(x, y) ((x)>(y)?(x):(y))

// 偶数判断
#define IS_EVEN(n) ((n)%2==0)

// 多行定义，加上{} 防止变量污染
#define PRINT_NUMS_TO_PRODUCT(a, b) { \
  printf("hhhhh\n");   \
  int product = (a) * (b); \
  for (int i = 0; i < product; i++) { \
    printf("%d\n", i); \
  } \
}

// 宏调用宏函数
#define QUADP(a, b, c) ((-(b) + sqrt((b) * (b) - 4 * (a) * (c))) / (2 * (a)))
#define QUADM(a, b, c) ((-(b) - sqrt((b) * (b) - 4 * (a) * (c))) / (2 * (a)))
#define QUAD(a, b, c) QUADP(a, b, c), QUADM(a, b, c)

// # 操作符
// 使用#符号，指定返回值为字符串
#define STR(x) "x"#x
// 等同于 printf("%s\n", "3.14159");

// ## 操作符
// 根据参数确定一个标识符
#define MK_ID(n) i##n

// 不定参数
#define X(a, b, ...) (10*(a) + 20*(b)), __VA_ARGS__
// 字符串
#define XS(...) #__VA_ARGS__

#define LIMIT 400
// 删除宏定义LIMIT; 相同名称的宏不能定义不同的值
#undef LIMIT

void pretest() {
    printf("%s\n", HELLO);
    printf("%s\n", Hi);
    printf("%d\n", FOUR);
    printf("SQUARE = %d\n", SQUARE(3 + 4));  // 输出19 3+4*3+4
    //printf("%d\n",getchar()); // 会等待输入
    printf("MAX = %d\n", MAX(1, 2));
    printf("IS_EVEN = %d\n", IS_EVEN(2));
    PRINT_NUMS_TO_PRODUCT(1, 2);
    printf("QUAD = %d,%d\n", QUAD(10, 20, 10));
    printf("STR = %s\n", STR(3.14));
    int i3 = 3;
    printf("MK_ID = %d\n", MK_ID(3));
    printf("X = %d,%d,%d\n", X(1, 2, 3, 4));
    printf("XS = %s\n", XS(1, 2, 3));
    //printf("LIMIT = %d\n",LIMIT); // LIMIT宏已经不存在了

    // if...else... 判断
    // 未定义的宏值为0，编译时可以定义宏
#if HAPPY_FACTOR == 0
    printf("I'm not happy!\n");
#elif HAPPY_FACTOR == 1
    printf("I'm just regular\n");
#else
  printf("I'm extra happy!\n");
#endif

    // 判断宏是否定义过
#ifdef EXTRA_HAPPY
    printf("I'm extra happy!\n");
#else
    printf("I'm just regular\n");
#endif

    // #if defined 和 #ifdef 等价
#if defined FOO
    printf("I'm FOO\n");
#elif defined BAR
    printf("I'm BAR\n");
#else
    printf("I'm Empty\n");
#endif

    // #ifndef 和 #if !defined 等价
#ifndef FOO
#endif
// 等同于
#if !defined FOO
#endif

    // 预定义宏，可直接使用
    printf("time = %s %s\n", __DATE__, __TIME__); // 时间
    printf("version = %ld\n", __STDC_VERSION__); // c语言版本

    // 预处理错误
#if __STDC_VERSION__ != 199901L //201112L
#error Not C11
#endif

}

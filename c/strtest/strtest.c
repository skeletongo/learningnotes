//
// Created by admin on 2022/1/25.
//

#include <string.h>
#include <stdio.h>

void strtest() {
    // 字符串有两种表示形式，指针和数组，字符串其实就是字节数组
    char *a = "hello";
    printf("a=%s\n", a);

    char b[] = "world";
    printf("b=%s\n", b);

    // 字符串拷贝
    char c[6];
    // 字符串复制：返回第一个参数，连同\0也会复制过去;但不会自动添加\0
    strcpy(c, b);
    printf("c=%s\n", c);
    printf("len(c) = %d\n", strlen(c)); // 获取长度不带\0,单位字节

    char d[] = "hello world";
    char e[] = "hello";
    char *f = strcpy(d + 6, e); // 返回值为拷贝的数据
    printf("f=%s,d=%s,e=%s\n", f, d, e);

    char g[] = "abcd";
    char h[] = "efgh";
    char *i = strncpy(g, h, 2); // 可以指定拷贝几个字节
    printf("g=%s,h=%s,i=%s\n", g, h, i);

    // 字符串拼接：返回第一个参数,会自动添加\0
    // 确认第一个字符串可以容纳拼接的字符串
    char m[11] = "hello";
    char n[] = "world";
    char *o = strncat(m, n, 5);
    printf("m=%s, n=%s, o=%s m==o:%d\n", m, n, o, m == o);

    // 字符串比较，按顺序比字节，相等为0
    int p = strncmp("hello", "hello world", 5);
    printf("strncmp(%d)\n", p);

    // 给字符串赋值
    // 参数：字符串地址，写入多少字节（不算\0）,格式化字符串，要写入的字符串
    char q[40];
    int r = snprintf(q, 12, "%s %s", "hello", "world");
    printf("q=%s, r=%d\n", q, r);

    // 字符串数组
    char *s[] = {"a", "b"};
    printf("s[0]=%s\n", s[0]);

    char ss[][10] = {"hello","world"}; // 字符串为第二层的char数组
    printf("%s %s\n",ss[0],ss[1]);
}
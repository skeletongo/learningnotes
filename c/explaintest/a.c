//
// Created by admin on 2022/5/2.
//

#include <stdio.h>
//#include "a.h" // 引入自己的头文件，编译器会检查声明和定义是否匹配

static int one = 1;

static int adda(int a, int b) {
    return a + b;
}

int hello() {
    char *a = "hello";
   // static char *b = a; 编译错误
}

void SayHello() {
    puts("hello");
}
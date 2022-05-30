//
// Created by admin on 2022/5/2.
//

#include <stdio.h> // 头文件中可以引入头文件

#ifndef DEMO_A_H
#define DEMO_A_H

int h = 5; // 头文件定义变量

int adda(int,int); // 这里声明也不行，因为函数被static修饰

void SayHello();

void SayWorld() { // 头文件定义函数
    printf("head hello");
}

#endif //DEMO_A_H

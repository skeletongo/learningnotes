//
// Created by admin on 2022/5/2.
//

#include <stdio.h>
#include "a.h"

// 变量说明符

void explaintest() {
    // const 禁止修改
    const int a = 1;
    const char b[10];
    const int *c; // c指针指向的数据不能修改
    int *const d; // d指针本身不能修改

    // static
    // 1.修饰局部变量时变量只初始化一次，并且值一直保留
    // 2.修饰全局变量时表示全局变量的作用域只在当前源码文件
    //printf("%d",one); // 编译错误；one在a.c中被static修饰
    // 3.修饰方法同修饰全局变量，方法的作用域只在当前源码文件
    //printf("%d", adda(1, 1)); // 编译错误；adda在a.c中被static修饰
    // 4.变量初始化必须是常量

    // extern 外部定义
    // 1.修饰的变量从当前作用域以外获取
    // 2.被修饰的变量不能初始化值，否则相当于没有被修饰
    // 3.函数声明默认有此修饰符
    //void SayHello(); // 相等 extern void SayHello();
    SayHello();
    printf("%d\n",h);
    SayWorld();

    // register 使用cpu中的寄存器存储变量,也可能不一定生效，被放进寄存器的变量不能取地址
    //register int a;
    //int *p = &a; // 编译器报错

    // volatile 禁止优化当前变量，每次使用都从内存获取

    // restrict
    // 1.只能修饰指针；被修饰的指针是访问当前数据的唯一方式
    // 2.在参数列表中使用的意思是多个参数指针内存地址没有重叠
}
//
// Created by admin on 2022/1/28.
//

#include <stdio.h>

void uniontest() {
    // 和struct基本一样，区别就是union用同一块内存空间表示多个数据
    union foo {
        int a;
        float b;
    } x;

    int *foo_int_p = (int *) &x;
    float *foo_float_p = (float *) &x;

    x.a = 12;
    printf("%d\n", x.a);           // 12
    printf("%d\n", *foo_int_p);    // 12

    x.b = 3.141592;
    printf("%f\n", x.b);           // 3.141592
    printf("%f\n", *foo_float_p);  // 3.141592

    typedef struct {
        int a;
        union {
            int i;
            float f;
        };
    } One;
    One p = {.a=9, .i=8};
    printf("a=%d, union=%d %d\n", p.a, p.i, p.f);
}
//
// Created by admin on 2022/1/17.
//

#include <stdio.h>
#include <stdint.h>

void ptrtest() {
    // 指针和数字加减
    int *i = (int *) 0x123;
    printf("int *i:%p\n", i);
    i = i + 1; // 增加一个单位，单位等于sizeof(指针类型)
    printf("i + 1:%p\n", i);
    i = i - 1;
    printf("i - 1:%p\n", i);

    long long *j = (long long *)0x1;
    printf("long long *j:%p\n",j);
    j = j + 1;
    printf("j + 1:%p\n", j);
    j = j - 1;
    printf("j + 1:%p\n", j);

    // 指针和指针不能加

    // 指针和指针减
    int32_t *a = 0x123;
    int32_t *b = 0x127;
    printf("b-a %d\n", b - a); // 距离几个sizeof(指针类型)
    printf("a-b %d\n", a - b); // 距离几个sizeof(指针类型)

    int c = 0x1;
    int *d = &c;
    *d = 2;
    printf("c=%d\n", *d);
}
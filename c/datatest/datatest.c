//
// Created by admin on 2022/1/17.
//

#include <stdio.h>
#include <stdint.h>

void datatest() {
    char c;
    printf("char %d\n", sizeof(c));
    short s;
    printf("short %d\n", sizeof(s));
    int i;
    printf("int %d\n", sizeof(i));
    long l;
    printf("long %d\n", sizeof(l));
    long long ll;
    printf("long long %d\n", sizeof(ll));
    float f;
    printf("float %d\n", sizeof(f));
    double d;
    printf("double %d\n", sizeof(d));
    int32_t i32;
    printf("int32_t %d\n", sizeof(i32));
}

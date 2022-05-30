#include <stdarg.h>
#include <stdio.h>

int *add(int a) { // 方法是指针
    int b;
    b = a + 1;
    return &b; // 不能返回局部变量的指针
}

void add2(int *a) {
    *a = *a + 1;
}

int add3(int a) {
    return a + 1;
}

int count() {
    static int a; // 定义一次
    a++;
    return a;
}

void local(const int *a) {
    int b = 0;
    // *a = b; // const修饰的值不能修改
}

int avg(int i, ...) {
    va_list ls;
    va_start(ls, i);

    int total = 0;
    for (int j = 0; j <= i; j++) {
        total += va_arg(ls, int);
    }
    va_end(ls);
    return total / i;
}

void functest() {
    printf("%d\n", add(1)); // 0
    int a = 1;
    add2(&a);
    printf("%d\n", a); // 2

    int (*add_ptr)(int) = &add3; // 函数名称 == 函数指针
    printf("%d\n", (*add_ptr)(2)); // 3
    // 等于
    // int (*add_ptr)(int) = add3;
    // printf("%d", add_ptr(2));

    printf("static %d\n", count());
    printf("static %d\n", count());

    // printf("test %d", test());

    int i = 0;
    for (; i < 2; i++) {
        printf("for i++ %d\n", i);
    }
    printf("for end %d\n", i);

    printf("avg %d", avg(4, 1, 2, 3, 4, 5));
}


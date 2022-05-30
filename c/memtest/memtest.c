//
// Created by admin on 2022/1/25.
//

#include <malloc.h>
#include <stdio.h>
#include <string.h>

void memtest() {
    // void* 指针，无类型指针
    int x = 1;
    void* xx = &x;
    int *x1 = xx;
    printf("x1=%d\n",*x1); // x1=1

    // 内存分配，不会修改原来内存中的值
    int *a = malloc(sizeof(int));
    if (a == NULL) {
        // 分配失败
    }
    *a = 2;
    printf("a=%d\n", *a);

    // 会初始化内存中的数据都为0
    int *b = calloc(10, sizeof(int));
    for (int i = 0; i < 10; i++) {
        b[i] = i;
        printf("b[*]=%d\n", b[i]);
    }

    // 调整已经分配的内存空间，增加或缩小，不会初始化内存值
    // 增加内存可能会导内存复制
    b = realloc(b, sizeof(int) * 5);
    for (int i = 0; i < 5; i++) {
        printf("b[*]=%d\n", b[i]);
    }

    // 内存拷贝
    // restrict 唯一访问方式说明符，只能由原指针访问
    int *restrict c = malloc(sizeof(int));
    *c = 1;

    // 非共享内存复制
    memcpy(c, a, 1);
    printf("c=%d\n", *c);

    // 可以是同一块内存
    memmove(b, b + 3, sizeof(int) * 2);
    for (int i = 0; i < 5; i++) {
        printf("b[*]=%d\n", b[i]);
    }

    // memcmp 内存数据比较, 相等为0
    // 参数：内存1,内存2,比前多少个字节
    printf("memcmp %d\n", memcmp("hello", "helld", 5));

    // memset 修改内存值
    // 参数：内存地址，值，修改多少个字节
    memset(b, 0, sizeof(int) * 5);
    for (int i = 0; i < 5; i++) {
        printf("b[*]=%d\n", b[i]);
    }

    // 内存释放
    free(a);
    free(b);
}

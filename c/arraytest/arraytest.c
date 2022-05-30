//
// Created by admin on 2022/1/22.
//

#include <stdio.h>
#include <string.h>

void modify(int arr[]) {
//void modify(int *arr) { // 等价,数组名和指针等价
    arr[0] = 2;
}

void modify2(int *a) {
    // 解指针修改数据
    *a = 4;
}

void modify3(int *a) {
    int c = 4;
    a = &c; // 没有解指针，不能修改指针对应内存的数据
}

int sum(int *start, int *end) {
    int total = 0;

    while (start < end) {
        total += *start;
        start++;
    }

    return total;
}

void arraytest() {
    int arr[5];
    for (int i = 0; i < 5; i++) {
        arr[i] = i;
    }
    for (int i = 0; i < 5; ++i) {
        printf("%d\n", arr[i]);
    }
    // 数组长度
    printf("%d\n", sizeof(arr) / sizeof(int));

    // 用指针求数组元素和
    printf("sum=%d\n", sum(arr, arr + 5));

    // 指针获取元素值;
    //int *a = &arr[0];
    int *a = arr;
    printf("arr[0]=%d\n", *a);
    printf("arr[1]=%d\n", *(a + 1));

    //
    int *a1 = &arr[5];
    int *a2 = &arr[3];
    printf("arr[5]-arr[3] = %d\n", a2 - a1); // -2
    printf("arr[5]-arr[2] = %d\n", arr[5] - arr[2]); // 不能这么用
    for (int i = 0; i < 5; i++) {
        printf("arr[*] = %d\n", arr[i]);
    }

    // 数组参数
    modify(arr);
    printf("modify arr[0]=%d\n", arr[0]);

    int d = 1;
    modify2(&d);
    printf("d=%d\n", d);

    int e = 1;
    modify3(&e);
    printf("e=%d\n", e);

    // 二维数组
    int f[2][2] = {{1, 0},
                   {0, 1}};
    printf("**f = %d\n", **f);

    // 数组复制
    int arr2[5];
    memcpy(arr2, arr, sizeof(arr)); // 内存拷贝
    for (int i = 0; i < 5; i++) {
        printf("arr2[*] = %d\n", arr2[i]);
    }
    int arr3[5];
    for (int i = 0; i < 5; i++) {
        arr3[i] = arr[i];
        printf("arr3[*] = %d\n", arr3[i]);
    }

    // 字面量数组
    modify((int[]) {1, 2});

    // 字符串数组:char name[20]，不能指向另一个地址，已经和内存地址绑定，内容可修改
    char aa[2] = {'h','i'};
    aa[1]= '0';
    printf("%s\n",aa);

    // 字符串指针:char *name，可以修指向的内存，但字符串内容不可修改
    char *name = "hello";
    name = "world";
    printf("%s\n",name);
}

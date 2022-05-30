//
// Created by admin on 2022/1/26.
//

#include <stdio.h>
#include <string.h>

void structtest() {
    // 类型定义
    struct Slice { // 结构体名称
        int cap;
        int n;
        void *arr;
    };
    // 使用类型时要带 struct 关键字，比较麻烦
    struct Slice s = {1, 2, (int[]) {1, 2}};
    printf("cap:%d n:%d arr:%d\n", s.cap, s.n, s.arr);
    for (int i = 0; i < s.n; i++) {
        int *arr = s.arr;
        printf("arr[*]=%d\n", arr[i]);
    }

    // 类型和变量同时声明
    struct People {
        char name[20];
        int age;
    } p1 = {"tom", 18};
    printf("name:%s\n", p1.name);
    printf("age:%d\n", p1.age);

    // 使用类型别名
    typedef struct /*这个地方的结构体名称可省略*/{
        char name[20];
        int age;
    } Person;
    Person p = {"jack", 20};
    printf("name:%s\n", p.name);
    printf("age:%d\n", p.age);

    // 复制
    // 字符串数组:char name[20]
    // 字符串指针:char *name
    // 复制结构体时，字符串数组可以被复制，字符串指针只是复制指针值
    struct Arrtest {
        char name[20];
    };
    struct Arrptr {
        char *name;
    };
    struct Arrtest a1 = {"tom"};
    struct Arrtest a2 = a1;
    a2.name[0] = 'k'; // 字符串数组不可重新赋值，但可以修改
    printf("a1.name=%s\n", a1.name);
    printf("a2.name=%s\n", a2.name);

    struct Arrptr a3 = {"tom"};
    struct Arrptr a4 = a3;
    a4.name = "jack"; // 字符串指针可以被重新赋值
    printf("a3.name=%s\n", a3.name);
    printf("a4.name=%s\n", a4.name);

    // 从指针取属性
    typedef struct {
        int a;
    } PP;
    PP pp = {10};
    PP *pp1 = &pp;
    pp1->a = 20;
    printf("pp1->a:%d\n", pp1->a);

    // 结构体数组
    typedef struct {
        int a;
    } ele;
    ele* arr[10];
    for (int i = 0; i < 10; i++) {
        arr[i] = NULL;
    }
    ele e1 = {1};
    arr[0] = &e1;
    for (int i = 0; i < 10; i++) {
        if (arr[i] == NULL) {
            printf("isNull\n");
        } else {
            printf("e: %d\n", arr[i]->a);
        };
    }

    // 结构体嵌套，其实就是属性也是结构体
    typedef struct {
        int a;
    } A;
    typedef struct {
        A b;
        A *br;
    } B;
    A a = {1};
    B b = {a, &a};
    printf("a=%d\n", a.a);
    printf("b.b.a=%d\n", b.b.a);
    printf("b.br->a=%d\n", b.br->a);

    // 位字段结构
    typedef struct {
        // 由上到下，低位到高位
        unsigned int a: 1;
        unsigned int b: 1;
        unsigned int c: 1;
        unsigned int d: 1;
        unsigned int : 0;
    } L;
    printf("sizeof(L)=%d\n", sizeof(L)); // 占用 int 数据类型整数倍的空间
    L bi = {1, 1, 1, 1};
    printf("bi=%d\n", bi); // 15

    typedef struct {
        unsigned int a: 1;
        unsigned int b: 2;
        unsigned int c: 1;
        unsigned int : 0; // 将剩余的28位都占用
        unsigned int d: 1; // 设置第5个字节的低第一位
        unsigned int : 0; // 将剩余所有位设置为0
    } LL;
    printf("sizeof(LL)=%d\n", sizeof(LL));
    LL bii = {1, 1, 0};
    bii.d = 1;
    printf("LL=%d,%d,%d,%d\n",bii.a,bii.b,bii.c,bii.d);
    long long LLI= 0;
    memcpy(&LLI,&bii,8); // 这里不知道为什么好像只复制了4个字节的数据
    printf("LLI=%d,sizeof(LLI)=%d\n",LLI, sizeof(LLI));

    typedef struct {
        int e;
    } es;

    typedef struct {
        int f;
        struct {
            int n;
        } d;
        es p;
    } f;

    f ff = {.f=1, .d.n=2, .p.e=4};
    printf("ff=%d %d %d\n", ff.f, ff.d.n, ff.p.e);
}
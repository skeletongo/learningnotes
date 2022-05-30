//
// Created by admin on 2022/1/28.
//

#include <stdio.h>
#include <stdlib.h>
#include "datatest/datatest.h"
#include "arraytest/arraytest.h"
#include "functest/functest.h"
#include "memtest/memtest.h"
#include "ptrtest/ptrtest.h"
#include "strtest/strtest.h"
#include "structtest/structtest.h"
#include "uniontest/uniontest.h"
#include "pretest/pretest.h"
#include "iotest/iotest.h"
#include "explaintest/explain.h"

int main(int argc, char* argv[]) {
    printf("argc=%d\n",argc); // 程序参数 + 1
    for (int i=0; i < argc;i++) {
        printf("argv[%d] = %s\n",i,argv[i]);
    }

    // 获取环境变量
    printf("env:PATH=%s\n",getenv("PATH"));

//    datatest();
//    arraytest();
//    functest();
//    memtest();
//    ptrtest();
//    strtest();
//    structtest();
//    uniontest();
//    pretest();
//    iotest();
//    explaintest();

//    return 0; // 只有main函数可以省略
}

// 阶段编译：(为了实现只编译修改的文件)
// 1.每个文件可以单独生成对象文件
// gcc -c main.c ...
// 2.用所有的对象文件生成可执行文件
// gcc -o main.o ...
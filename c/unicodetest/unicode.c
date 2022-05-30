//
// Created by admin on 2022/5/2.
//

#include <stdio.h>
#include <locale.h>
#include <wchar.h>
#include <stdlib.h>

int main() {
    // unicode编码
    // \123：以八进制值表示一个字符，斜杠后面需要三个数字。
    // \x4D：以十六进制表示一个字符，\x后面是十六进制整数。
    // \u2620：以 Unicode 码点表示一个字符（不适用于 ASCII 字符），码点以十六进制表示，\u后面需要4个字符。
    // \U0001243F：以 Unicode 码点表示一个字符（不适用于 ASCII 字符），码点以十六进制表示，\U后面需要8个字符。

    // 都输出 ABC
    printf("ABC\n");
    printf("\101\102\103\n");
    printf("\x41\x42\x43\n");

    // 都输出 好 Bullet 1
    printf("\u597D Bullet 1\n");
    printf("\U0000597D Bullet 1\n");

    char* s = "你好";
    printf("%s\n",s); // 你好

    // 注意，\u + 码点和\U + 码点的写法，不能用来表示 ASCII 码字符（码点小于0xA0的字符）
    // 只有三个字符除外：0x24（$），0x40（@）和0x60（`）。
    char* ss = "\u0024\u0040\u0060";
    printf("%s\n", ss);  // @$`

//    // 宽字符，一个字符占用多个字节; 没测试成功？？？
//    wchar_t c = L'牛';
//    printf("%lc\n", c);
//
//    wchar_t* cs = L"春天";
//    printf("%ls\n", cs);
//
//    char* mbs1 = "春天";
//    printf("%d\n", mblen(mbs1, MB_CUR_MAX)); // 3
//
//    char* mbs2 = "abc";
//    printf("%d\n", mblen(mbs2, MB_CUR_MAX)); // 1

}
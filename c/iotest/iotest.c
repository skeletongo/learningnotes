//
// Created by admin on 2022/2/25.
//

#include <stdio.h>

void iotest() {
//    // 读取键盘输入；以空白字符结束或不属于当前读取的类型结束
//    int i;
//    // 输入：13 5 或 13\n5 是一样的
//    printf("res=%d\n",scanf("%d",&i));
//    printf("out:%d\n", i); // 13
//    printf("res=%d\n",scanf("%d",&i));
//    printf("out:%d\n", i); // 5
//    // 输入：13.5 读取到13
//    printf("res=%d\n",scanf("%d",&i));
//    printf("out:%d\n", i); // 13
//
//    // 读取字符时不会忽略空白字符
//    char a;
//    scanf(" %c", &a); // 前面加一个空格表示忽略开始的空白字符
//    printf("%c",a);
//
//    // 字符串读取，从当前第一个非空白字符开始读起，直到遇到空白字符为止
//    char s[11];
//    scanf("%10s",s); // 数字代表读取最多几个字符
//
//    // 赋值或略符 %*c
//    int year,month,day;
//    int num = scanf("%d%*c%d%*c%d", &year, &month, &day); // 输入2001-10-10，不会检查匹配“-”，“-”可以替换成其他任意单个字符
//    printf("res=%d\n",num); // 3
//
//    // 从字符串读取数据
//    char *str = "12 13";
//    int i,j;
//    int num = sscanf(str, "%d%d", &i, &j);
//    printf("%d %d\n",i,j); // 12 13
//    printf("res=%d\n",num); // 2
//
//    // 读取失败返回EOF, EOF等于-1
//
//    // 读取一个任意字符,包括空白字符
//    char ch;
//    ch = getchar();
//    printf("getchar=%c\n",ch);
//    // 等同于
//    //scanf("%c", &ch);
//
//    // 打印一个任意字符
//    putchar('a');
//    // 等同于
//    //printf("%c", ch);
//
//    // 输出一行字符串带换行
//    puts("hello world!");
//    puts("hello world!");
//
//    // 读取一行
//    char str[12];
//    // 从标准输入读取一行（以换行符区分一行），fgets会去掉换行符，并加上\0
//    fgets(str, sizeof(str), stdin); // 输入 hello world
//    printf("%s",str);
}

#include <stdio.h>
#define NDEBUG // 关闭assert检测
#include <assert.h>

int add(int a, int b) {
    static_assert(2<1,"fail"); // 编译期检测条件是否成立，不成立就报错
    return a+b;
}

int main() {
    assert(1<2); // 运行期检测条件是否成立，不成立就报错
    puts("success");
    assert(2<1);
    puts("success");
}
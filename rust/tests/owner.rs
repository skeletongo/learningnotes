// 所有权概念
// 1.所有变量都有所有者
// 2.任意时刻变量只有一个所有者
// 3.变量离开作用域，变量会被释放，如果当做返回值所有权会重新修改

// 作用域：花括号
// 有效：可以访问
// 无效：不可用
// 移动：变量赋值到其它变量，变量没有实现Copy,原变量失效
// 复制：变量赋值到其它变量，变量实现了Copy，原变量仍然可用
// 所有权：所有变量各自都记录自己唯一的拥有者；移动就是改变了所有权
// 引用：变量取地址
// 可变引用：默认引用对应的变量的值不可修改，可变引用对应变量的值可以修改
// 悬垂引用：只有引用，引用的值已经被释放了

// Clone: 深拷贝
// Copy: 浅拷贝

#[test]
fn owner1() {
    let s = String::from("hello");
    let s2 = s;
    //println!("{}",s); // String没有实现Copy,所以会发生移动，移动之后原变量就不可用了
    println!("{}", s2)
}

#[test]
fn owner2() {
    let s = String::from("hello");
    hello(s); // 发生移动
              //println!("{}",s); // 编译错误

    let x = 1;
    number(x); // 发生复制
    println!("x={}", x);
}

fn hello(s: String) {}

fn number(n: i32) {}

#[test]
fn test1() {
    fn aa(s: &String) {
        println!("{}", s)
    }

    let s = String::from("hehe");
    let ss = &s; // 借用（引用）不会发生移动
    aa(ss);
    println!("{}", s);
    println!("{}", ss); // 引用不会发生移动是复制
}

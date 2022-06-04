// 基本语法，基本数据类型

#[test]
fn test1() {
    let a = 1; // 不可变变量
    println!("{}", a);
    let mut b = 1; // 可变变量
    println!("{}", b);
    b = 2;
    println!("{}", b);
}

#[test]
fn test2() {
    let arr = [1, 2, 3]; // 数组
    for e in arr {
        println!("arr {}", e);
    }

    let arrint: [isize; 2] = [1, 2]; // 数组声明,类型 isize,长度为2
    for e in arrint {
        println!("arrint {}", e);
    }

    let arr5 = [5; 2]; // 创建一个长度为2，元素值都为5的数组
    for e in arr5 {
        println!("arr5 {}", e);
    }
}

#[test]
fn testfunc() {
    fn five() -> i32 {
        // 箭头就是多余，小括号已经可以区分后面是返回值了
        5 // 用return返回不行吗，非要来个小语法
    }
    println!("five {}", five());

    fn add(a: i32, b: i32) -> (i32, bool) {
        // 原来元组是为了实现多值返回的
        (a + b, true)
    }
    let (r, b) = add(1, 1); // 接收竟然还需要括号
    println!("add {}, {}", r, b)
}

#[test]
fn testfor() {
    // 控制语句
    if true {
    } else if true {
    } else {
    }

    loop {
        println!("无限循环");
        break;
    }

    let mut i = 0;
    while i < 2 {
        i += 1;
    }

    let arr = [1; 3];
    for e in arr {
        println!("{}", e);
    }
}

#[test]
fn testlist() {
    // 斐波那锲数列
    fn get(n: i32) -> i32 {
        let mut a = 0;
        let mut b = 1;
        let mut c = 0;
        let mut i = 0;
        loop {
            a = b;
            b = c;
            c = a + b;
            i += 1;
            if i >= n {
                break;
            }
        }
        return c;
    }

    let mut i = 1;
    while i < 10 {
        println!("index {} = {}", i, get(i));
        i += 1;
    }
}

#[test]
fn teststring() {
    // String类型
    let s = String::from("abcd efgh");
    for (i, e) in s.as_bytes().iter().enumerate() {
        println!("{}, {}", i, e);
    }

    // 字符串切片（字符串的部分引用）
    let slice = "hello";
    println!("slice = {}", slice);

    let ss = &s[..]; // 和其它语言的切片差不多，只是符号不一样（想不通为什么要换符号）
    println!("ss = {}", ss)
}

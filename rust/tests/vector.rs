// 动态数组

#[test]
fn test1() {
    // 宏
    let r = vec![1, 2, 3];
    for e in r {
        // e 是元素值
        println!("{}", e);
    }

    let rr = vec![4, 5, 6];
    for e in &rr {
        // e 是元素值的引用
        println!("{}", e);
    }

    let mut ss = vec![6, 7, 8];
    for e in &mut ss {
        // e 是元素值的引用
        *e += 1;
        println!("{}", e);
    }

    // Vector
    let mut dd: Vec<&str> = Vec::new(); // 在返回值中定义泛型的方式还有点不太适应
    dd.push("a");
    dd.push("b");
    dd.push("c");
    for e in dd {
        println!("{}", e);
    }
}

use std::collections::HashMap;

#[test]
fn test1() {
    let s = String::from("hehe"); // String
    let ss = &s; // &String
    let sss = s.as_str(); // &str 字符串切片
    println!("{},{},{}", s, ss, sss);

    let c = &s[0..1]; // 字符串截取
    println!("{}", c);

    let sl = "hehe"; // 字符串切片
    let s = sl.to_string(); // 切片创建字符串对象
    println!("{},{}", sl, s);

    // 字符遍历
    for c in s.chars() {
        println!("{}", c);
    }
    // 字节遍历
    for b in s.bytes() {
        println!("{}", b);
    }
}

#[test]
fn testhashmap() {
    // 哈希表
    use std::collections::HashMap;

    let mut m = HashMap::new();
    // 添加键值对，直接覆盖
    m.insert("a", 1);
    m.insert("a", 2);
    // 没有才添加
    m.entry("b").or_insert(3);
    m.entry("b").or_insert(4);
    for (k, v) in m {
        println!("{},{}", k, v);
    }

    // 更新
    let mut mm = HashMap::new();
    let v = mm.entry("a").or_insert(1);
    *v += 1;
    println!("{}", v);
}

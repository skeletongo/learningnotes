// 生命周期

// 生命周期省略规则
// 1.引用的参数都有它自己的生命周期参数
// 2.如果只有一个输入生命周期参数，那么它被赋予所有输出生命周期参数
// 3.如果方法有多个输入生命周期参数并且其中一个参数是 &self 或 &mut self，说明是个对象的方法(method)
// 那么所有输出生命周期参数被赋予 self 的生命周期

// longest 函数返回的引用的生命周期应该与传入参数的生命周期中较短那个保持一致
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

#[test]
fn test1() {
    let s = String::from("aaa");
    let ss = "bbbb";
    let r = longest(s.as_str(), ss);
    println!("{}", r)
}

#[test]
fn test2() {
    let string1 = String::from("long string is long");
    let result;
    {
        let string2 = String::from("xyz");
        result = longest(string1.as_str(), string2.as_str()); // string2的生命周期短于result最后使用的地方
    }
    println!("The longest string is {}", result);
}

// 枚举类型

#[test]
fn test1() {
    // 普通枚举
    enum IpAddr {
        V4,
        V6,
    }
    let v4 = IpAddr::V4;
    let v6 = IpAddr::V6;

    // 绑定数据的枚举
    enum IpAddr2 {
        V4(u8, u8, u8, u8),
        V6(String),
    }
    let ipv4 = IpAddr2::V4(127, 0, 0, 1);
    let ipv6 = IpAddr2::V6(String::from(":::1"));

    enum Message {
        Quit,                       // 相当于类单元结构体
        Move { x: i32, y: i32 },    // 相当于普通结构体
        Write(String),              // 相当于元组结构体
        ChangeColor(i32, i32, i32), // 相当于元组结构体
    }

    //todo 相当于一个类型拥有多种结构体,多个结构体属于同一个类型
    // 是不是用其它特性也能实现
}

#[test]
fn test2() {
    // match

    #[derive(Debug)]
    enum Hello {
        A,
        B,
        C(State), // 绑定数据
    }
    #[derive(Debug)]
    enum State {
        D,
        E,
    }
    fn switch1(t: Hello) {
        // 必须包含所有枚举
        match t {
            Hello::A => {
                println!("switch1 {:?}", t);
            }
            Hello::B => {
                println!("switch1 {:?}", t);
            }
            Hello::C(s) => {
                println!("switch1 state={:?}", s);
            }
        }
    }
    switch1(Hello::A);
    switch1(Hello::C(State::D));

    // 其它语言中的switch
    fn switch2(n: i32) {
        match n {
            0 => {
                println!("switch2 0")
            }
            // other => {
            //     println!("switch2 {}", other)
            // }
            other => number(other),
        }
    }
    fn number(n: i32) {
        println!("number {}", n)
    }
    switch2(0);
    switch2(1);
}

#[test]
fn test3() {
    // Option<T>

    fn number(n: i32) -> Option<i32> {
        return Some(n);
    }

    let five = number(5);

    fn add(o: Option<i32>) -> Option<i32> {
        match o {
            None => None,
            Some(n) => Some(n + 1),
        }
    }

    match add(five) {
        None => {
            println!("None");
        }
        Some(t) => {
            println!("{}", t);
        }
    }
}

#[test]
fn test4() {
    // if let, match的另一种写法
    let a = Some(1);
    if let Some(0) = a {
        println!("0")
    } else if let Some(1) = a {
        println!("1")
    } else {
        println!("2")
    }
}

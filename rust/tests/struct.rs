// 结构体数据类型

// 结构体中数据的所有权有两种情况
// 1.属于结构体
// 2.不属于结构体(引用)，需要加生命周期

#[test]
fn test1() {
    // 三种结构体

    // 结构体
    struct User {
        name: String,
        age: u8,
    }
    let u = User {
        name: String::from("tom"),
        age: 18,
    };
    println!("name = {}, age = {}", u.name, u.age);

    // 元组结构体
    struct Color(u8, u8, u8);
    let c = Color(10, 10, 10);
    println!("color {},{},{}", c.0, c.1, c.2);

    // 类单元结构体（没有字段的结构体）
    struct Nothing;
    let _ = Nothing;
}

#[test]
fn testprint() {
    // 三种打印方式

    #[derive(Debug)]
    struct R {
        width: u32,
        height: u32,
    }

    fn area(r: R) -> u32 {
        r.width * r.height
    }

    // 第一种
    println!(
        "area={}",
        area(R {
            width: 5,
            height: 5,
        })
    );

    let r = R {
        width: dbg!(2 * 5), // dbg!() 打印文件名及行号还有表达式的值到stderr,并返回表达式的值
        height: 2,
    };

    // debug模式打印,需要添加 #[derive(Debug)] 注释
    println!("area={:?}", r); // 第二种
    println!("area={:#?}", r); // 第三种
}

#[test]
fn testprint2() {
    // 结构体的关联函数
    // 1.方法：第一个参数是结构体对象的函数
    // 2.函数：不是方法的函数

    struct R {
        width: u32,
        height: u32,
    }

    impl R {
        // 方法
        fn area(&self) -> u32 {
            self.width * self.height
        }

        // 多参数方法
        fn can_hold(&self, other: &R) -> bool {
            self.width >= other.width && self.height >= other.height
        }

        // 函数
        fn new(width: u32, height: u32) -> R {
            R { width, height }
        }
    }

    let r = R::new(5, 5);
    println!("area = {}", r.area());
    
    let r2 = R::new(2, 2);
    println!("hold {}",r.can_hold(&r2));
}

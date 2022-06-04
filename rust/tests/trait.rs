use std::fmt::{Display, Formatter};

#[test]
fn test1() {
    trait Person {
        fn name(&self) -> &String;
        fn age(&self) -> i32 {
            18
        }
        fn run(&self) {
            println!("run")
        }
    }

    struct Tom {
        name: String,
    }

    impl Tom {
        fn name(&self) -> &String {
            println!("Tom fn");
            &self.name
        }

        fn run(&self) {
            println!("tom run")
        }
    }

    impl Person for Tom {
        fn name(&self) -> &String {
            println!("trait fn");
            &self.name
        }
    }

    let p1 = Tom {
        name: String::from("tom"),
    };
    // 作为参数类型要加impl,作为泛型类型时不用加
    fn myname(p: &impl Person) {
        p.name();
    }
    myname(&p1); // trait fn

    fn myname2(p: &Tom) {
        p.name();
    }
    myname2(&p1); // Tom fn

    println!("{}", p1.age()); // trait 默认方法
    p1.run(); // 覆盖默认实现
}

#[test]
fn test2() {
    // 声明式类型接口
    trait Person {
        fn name(&self) -> &String;
    }

    struct Tom {
        name: String,
    }
    impl Tom {
        fn name(&self) -> &String {
            &self.name
        }
    }

    fn myname(p: &impl Person) {
        p.name();
    }
    let p = Tom {
        name: String::from("tom"),
    };
    //myname(&p); // 编译出错，类型实现trait必须显示声明，我不太喜欢这种方式

    // impl Person for Tom {
    //     fn name(&self) -> &String {
    //         &self.name
    //     }
    // }
}

#[test]
fn test3() {
    // 泛型使用trait
    fn f<T: PartialOrd + Copy>() -> i32 {
        1
    }

    // where语法
    fn ff<T>() -> i32
    where
        T: PartialOrd + Copy,
    {
        1
    }
}

#[test]
fn test4() {
    // trait实现规则
    // 1.不能给外部类型实现外部trait
    // 2.本地类型可以实现外部trait
    // 3.外部类型可以实现本地trait
}

#[test]
fn test5() {
    //todo 感觉不太符合直觉
    trait Comm {}
    struct A {}
    struct B {}
    impl Comm for A {}
    impl Comm for B {}

    // fn f(b: bool) -> impl Comm {
    //     if b {
    //         A {}
    //     } else {
    //         B {} // 编译出错
    //     }
    // }
}

#[test]
fn test6() {
    // 为泛型类型中子类型实现了特定trait的类型实现特定的方法

    trait MyString {
        fn show(&self);
    }
    struct Point<T> {
        x: T,
        y: T,
    }

    // 两类：
    // 1.实现trait
    // 2.实现方法

    // 1.为所有类型实现方法
    impl<T> Point<T> {
        fn new(x: T, y: T) -> Self {
            Self { x, y }
        }
        // 等价
        // fn new(x: T, y: T) -> Point<T> {
        //     Point { x, y }
        // }

        fn one(&self) {}
    }
    let p1 = Point::new(1, 2);

    // 为泛型的某些类型实现trait
    impl MyString for Point<i32> {
        fn show(&self) {}
    }
    // 为泛型的某些trait类型实现方法
    impl<T: MyString> Point<T> {
        fn hehe(&self) {}
    }

    fn f(p: Point<i32>) {
        p.show();
    }

    fn f1<T>(p: Point<T>) {
        p.one();
        //p.show(); // 编译错误，没有实现这个方法
        //p.hehe(); // 编译错误，没有实现这个方法
    }

    fn f2(p: Point<impl MyString>) {
        p.hehe();
    }

    // 2.为所有类型实现trait
    // impl<T> MyString for Point<T> {
    //     fn show(&self) {}
    // }
    // fn f(p: &impl MyString) {
    //     p.show();
    // }
    // f(&p1);
}

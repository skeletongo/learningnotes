// 泛型
// 原理：单态化，编译时生成需要的具体类型的代码

use std::cmp::Ordering;

#[test]
fn test1() {
    // 泛型函数
    fn find_max<T: PartialOrd + Copy>(list: &[T]) -> &T {
        let mut max_value = &list[0];

        for e in list {
            match max_value.partial_cmp(e) {
                Some(order) => match order {
                    Ordering::Less => {
                        max_value = e;
                    }
                    _ => (),
                },
                _ => (),
            }
        }

        max_value
    }

    let a = vec![1, 2, 3];
    println!("{}", find_max(&a));

    let b = vec![1.1, 2.2, 3.3];
    println!("{}", find_max(&b));
}

#[test]
fn test2() {
    // 泛型结构体
    struct Point<T, Y> {
        x: T,
        y: Y,
    }

    let a = Point { x: 1, y: 2.2 };
    println!("{},{}", a.x, a.y);

    // 泛型枚举
    enum Number<T, Y> {
        A(T),
        B(Y),
    }
    fn bb<A, B>(a: A, b: B) -> Number<A, B> {
        Number::A(a)
    }
    let r = match bb(1, 2) {
        Number::A(a) => a,
        Number::B(b) => b,
    };
    println!("{}", r);
}

#[test]
fn test3() {
    // 泛型方法
    struct Point<T> {
        x: T,
        y: T,
    }

    // 第一个T是泛型，第二个是类型名称的一部分
    impl<T> Point<T> {
        pub fn x(&self) -> &T {
            // 因为T的类型不确定，不知道是否实现了Copy
            &self.x
        }
    }

    // 为某些类型实现方法
    impl Point<f32> {
        pub fn distance(&self) -> f32 {
            (self.x.powi(2) + self.y.powi(2)).sqrt()
        }
    }
}

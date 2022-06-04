// 作为公共库

pub mod hehe; // 相当于将hehe模块代码添加到当前模块

pub fn setup() {
    hehe::hi();
}

pub mod One {}

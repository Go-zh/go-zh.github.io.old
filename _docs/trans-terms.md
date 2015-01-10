---
layout: default
title: 术语列表
summary: 列出的术语的翻译及详情。
---
# {{ page.title }}

## 术语翻译

单词              | 译法              | 词性   | 类别 | 备注
----------------- | ----------------- | ------ | ---- | ----
approximation     | 近似[值]/逼近式   | n      |      |
argument          | 实参              | n      |      |
assignment        | 赋值              | n      |      |
block             | 块/阻塞           | n      |      | “阻塞”仅用于信道
body              | 执行体            | n      |      |
case              | 情况/写法         | n      |      | 在用作大小写时译作“写法”，因为有些字母还有其它多种写法
channel           | 信道              | n      |      |
char/character    | 字符              | n      |      |
code point        | 码点              | n      |      |
coefficient       | 系数              | n      |      |
commit            | 提交              | v      |      | 即直接向repo提交代码
compatibility     | 兼容性            | n/adj  |      |
complex           | 复数              | n      |      |
constant          | 常量              | n      |      |
convention        | 约定              | n      |      |
defer             | 推迟              | v      |      |
degree            | 阶                | n      |      | 仅用于多项式
distribute        | 分发              | v      |      |
distribution      | 分发              | n      |      |
error             | 错误/误差         | n      |      | “误差”用于数学
evaluation        | 求值              | n      |      |
even              | 偶(数)            | n/adj  |      |
expression        | 表达式            | n      |      |
flag              | 标志              | n      |      |
floating-point    | 浮点数            | n      |      |
form              | 形式              | n      |      |
function          | 函数              | n      |      |
implementation    | 实现              | n      |      |
integer           | 整数              | n      |      |
introduction      | 引言              | n      |      |
label             | 标签              | n/v    |      |
lock              | 锁                | n/v    |      |
method            | 方法              | n      |      |
mutex             | 互斥锁            | n      |      |
name space        | 命名空间          | n      |      |
normalize         | 规范化            | v      |      |
odd               | 奇(数)            | n/adj  |      |
panic             |（保留）           | n/v    |      |
parameter         | 形参              | n      |      |
pending           | 待定/挂起         | adj    |      |
polynomial        | 多项式            | n      |      |
profile/profiling | 评估              | n/v    |      |
race              | 竞争/竞态         | n/v    |      |
recover           | 恢复              | v      |      |
reduction         | 换算              | n      |      |
reference         | 引用              | n/v    |      |
repository        | [源码]仓库/源码库 | n      |      |
round             | 舍入              | v/adj  |      | 在数学中为舍入
rune              | 符文              | n      |      |
scope             | 作用域            | n      |      |
script            | 脚本/书写[系统]   | n      |      | 在Unicode中译作“书写[系统]”
source code       | 源码/源代码       | n      |      |
statement         | 语句              | n      |      |
stride            | 间距              | n      |      | 用作两个码点的间距。如从A(0x41)到a(0x61)的间距为32(0x20)
struct            | 结构体            | n      |      |
submit            | 递交              | v      |      | 指递交至主代码树
tag               | 标记              | n      |      | 多指struct tag，上下文只与xml/html相关时仍作“标签”
token             | [词法]标记        | n      |      |
ulp               | 末尾单元          | n      |      | Unit in the Last Place 的缩写
universe block    | 全域块            | n      |      |
variable          | 变量              | n      |      |
buffer            | 缓冲区            | n      |      |
cache             | 缓存              | n      |      |
constructor       | 构造函数          | n      |      |
embedding         | 内嵌              | n/v    |      | 用作类型，与“嵌入式”分开

## 有争议的术语

容易有争议的翻译   | 最终结果                  | 供讨论的选项                                          | 附注/用法实例
------------------ | ------------------------- | ----------------------------------------------------- | -------------
goroutine          |                           | Go程                                                  |
map                | 字典                      | 一般翻译为映射，但是映射是个数据结构么？字典/关联数组 |
tag                |                           |                                                       | 注意，新版 go1compat 上改用 (un)keyed struct literal，所以不再有 tagged struct literal这种用法了 (但是 keyed struct literal 咋翻译？)
label              |                           | 标记                                                  | 和 tag 区分，tag 用得多 struct tag, build tag, tagged struct，label 似乎只在一个地方有用
channel            | 信道                      | 通道/信道？                                           | receive/send-only channel？
panic              |                           |                                                       | panic
recover            |                           |                                                       | recover
embedding          |                           | 是嵌入好内嵌好？                                      | struct embedding/embedded struct
nil                | 无值                      | nil/空值/空/无值                                      | nil 无值，empty 空(值)，zero-value 零值
buitin function    | 内建函数                  | 内建函数？                                            |
argument/parameter | 实(际)参(数)/形(式)参(数) | 参数？实参/形参？                                     |
receiver           |                           | 接收器？太生硬了                                      |
method             | 方法                      | 方法                                                  |
interface          | 接口                      | 接口                                                  |
variadic           | 变参(函数)                | 变参函数 （变的是参数还是参数的个数？）               | variadic function, Google Translate 翻译做 可变参数
unaddressable      | 无法取址                  | 无法取址的？                                          | unaddressable value
switch             |                           | 多项选择？                                            |
case               | 分支                      | 情况？分支？                                          |
fallthrough        | fallthrough               | 穿透/下穿                                             |
untyped            | 未定型                    | 无类型/未定型（类型尚未确定）？？                     |
field              |                           |（结构体的）字段                                       |
reference          |                           | 引用？                                                | argument must be a field reference

说明: 如果对翻译有更好的建议请[在此（需梯子）](https://docs.google.com/spreadsheets/d/1oBEYj0TUCuTDy5vLGsjtMQ_Sd-_FASf393bhC-rfbGQ/edit?usp=sharing)参与讨论.

# 概述

## 操作

### 基础操作
**List**包含如下操作

*   构建一个空白的List
*   判断是否为空
*   添加元素到List开始位置
*   添加元素到List结束位置
*   获取第一个元素
*   获取List所有数据(去掉开头元素)


#### 位置操作

*   第一个元素
*   最后一个元素
*   下一个元素
*   上一个元素

### 扩展操作

*   元素数量
*   文字化表示
*   清空列表

### 可选操作

*   替换
*   变换
*   过滤
*   排序

## 接口定义

```golang
type List interface {
	New() List
	IsEmpty() bool
	Append(value ...Element) bool
	Prepend(value ...Element) bool
	Iterator() Iterator

	Tail() List
	Size() int
	String() string
	Clear()
}

type Element interface {
	Less(another Element) int
}

type Iterator interface {
	Head() (Element, bool)
	End() (Element, bool)
	Next() (Element, bool)
	Prev() (Element, bool)
}

```

### 数据结构关系
存在以下数据概念

*   基础List 通过New()生成
*   子List   通过List.Tail()生成
*   迭代器    通过List.Iterator() 生成
*   子迭代器    子List的迭代器

#### 子List和基础List

|   模式  |   优点  |  缺点   |
|   --- |   ---     |   --- |
|   共享数据    |  创建子List速度快   |   修改性操作互相影响，并且存在较明显的资源竞争  |
|   独立数据    |  修改性操作互不影响，创建完成后不存在资源竞争   |   创建子List速度慢  |

**考虑到大量的修改性操作，应当选择独立数据模式**

#### List和迭代器关系


### 实现方式
可选的实现方式有

*   数组
*   单链表
*   双链表
*   自平衡二叉查找树


#### 数组实现

##### 和迭代器的关联操作

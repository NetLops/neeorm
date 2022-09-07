## 支持的特性
- 表的创建、删除、迁移。
- 记录的增删查改，查询条件的链式操作。
- 单一主键的设置(primary key)。
- 钩子(在创建/更新/删除/查找之前或之后)
- 事务(transaction)。
## 实现
### Dialect
> 使用`Dialect`隔离不同数据库之间的差异，便于扩展。

| go               | sql     |
|------------------|---------|
| `int8、int16、int32` | `integer` |

### Schema
> 对象（Object） <<====>> 表（Table）

>给定一个任意的对象，转换为关系型数据库中的表结构

- 表名(table name) -- 结构体名(struct name)
- 字段名和字段类型 -- 成员变量和类型。
- 额外的约束条件（例如非空、主键等）-- 成员变量的Tag（Python、Java 使用注解实现）
```go
type User struct {
	Name string `neeorm:"PRIMARY KEY"`
	Age int
}
```
==>
```sqlite
CREATE TABLE `User` (`Name` text PRIMARY KEY ,`Age` integer)
```

### hooks
- 通过反射(`reflect`)获取结构体绑定的钩子(`hooks`)、并调用
- 支持增删查改(`CRUD`)前后调用钩子
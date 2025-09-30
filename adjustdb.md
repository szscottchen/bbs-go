# 数据库字段变更规范 (adjustdb.md)

## 1. 基本原则
1. **无迁移脚本原则**：禁止创建独立迁移脚本，所有变更必须通过模型修改实现
2. **向后兼容原则**：
   - 新增字段必须支持NULL值
   - 必须设置合理的SQL默认值（如适用）
3. **最小影响窗口**：变更应在系统维护时段执行

## 2. 字段定义规范
### 必须属性
```go
// 标准字段定义模板
FieldName DataType `gorm:"size:<长度>;<约束>" json:"<字段名>" form:"<字段名>"`
```

### 允许的数据类型
| 存储类型 | Go类型 | 空值处理 | 示例 |
|---------|--------|----------|------|
| VARCHAR | `sql.NullString` | `NOT NULL`需特殊审批 | `sql.NullString` |
| INT     | `int`  | 必须设置默认值 | `int` |
| BOOLEAN | `bool` | 默认`false` | `bool` |
| TEXT    | `string` | 允许NULL | `string` |

## 3. 操作流程
### 模型修改
1. 修改`server/internal/models/models.go`：
```go
type User struct {
    // ...
    NewField1 sql.NullString `gorm:"size:64" json:"newField1"`
    NewField2 int           `gorm:"default:0" json:"newField2"`
}
```

### 初始化逻辑
2. 在`server/internal/install/install.go`中添加：
```go
// 在InitDB函数中添加
if err := db.AutoMigrate(&models.User{}); err != nil {
    slog.Error("User表迁移失败", slog.Any("error", err))
}
```

## 4. 验证要求
1. 必须包含的测试项：
   - 空值写入测试
   - 批量数据更新测试
   - 索引性能测试（如添加索引）

## 5. 回滚方案
```go
// 回滚步骤：
// 1. 移除模型中的字段定义
// 2. 执行手动SQL删除列（如已部署）
ALTER TABLE t_user DROP COLUMN new_field;
```

## 6. 变更记录
所有变更必须记录在：
`server/docs/db-changelog.md`
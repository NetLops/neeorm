package schema

import (
	"go/ast"
	"neeorm/dialect"
	"reflect"
)

// Field represents a column of database
type Field struct {
	Name string
	Type string
	Tag  string // constraint condition // 约束条件
}

// Schema represents a table of database
type Schema struct {
	Model      interface{}       // Mapping Object // 映射对象 Model
	Name       string            // TableName // 表名
	Fields     []*Field          // 字段
	FieldNames []string          // 素有字段名（列名）
	fieldMap   map[string]*Field // 记录字段名和 Filed 的映射关系，
}

// GetField returns field by name
func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// Values return the values of dest`s member variables
func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}

type ITableName interface {
	TableName() string
}

// Parse
// TypeOf返回入参的类型 ValueOf 返回入参的值
// reflect.Indirect() 获取指针只指向的实例
// modelType.Name() 获取到结构体的名称作为表名
// NumField() 获取实例的字段的个数，然后通过下标获取到特定字段 p := modelType.Field(i);
// p.Name 字段名,p.Type字段类型,p.Tag 额外的约束条件
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	var tableName string
	t, ok := dest.(ITableName)
	if !ok {
		tableName = modelType.Name()
	} else {
		tableName = t.TableName()
	}
	schema := &Schema{
		Model:    dest,
		Name:     tableName,
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		pty := reflect.Indirect(reflect.New(p.Type))
		//fmt.Println(p, pty, p.Type)
		//fmt.Println(reflect.TypeOf(pty), pty)
		//fmt.Println(d, reflect.TypeOf(d))
		if !p.Anonymous && ast.IsExported(p.Name) { //  ast.IsExported 检查是否是公开字段
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(pty),
			}
			if v, ok := p.Tag.Lookup("neeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}

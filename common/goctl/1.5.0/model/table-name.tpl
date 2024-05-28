func (m *default{{.upperStartCamelObject}}Model) tableName() string {
	return m.table
}
// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *default{{.upperStartCamelObject}}Model) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}

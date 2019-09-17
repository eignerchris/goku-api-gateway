package plugin_entity

var GlobalPlugin map[string]*Plugin = make(map[string]*Plugin)

// 插件
type Plugin struct {
	PluginID     int    `json:"pluginID"`
	PluginName   string `json:"pluginName"`
	PluginStatus int    `json:"pluginStatus"`
	PluginIndex  int    `json:"pluginPriority"`
	PluginConfig string `json:"pluginConfig"`
	PluginInfo   string `json:"pluginInfo"`
	IsStop       int    `json:"isStop"`
}


type PluginParams struct {
	PluginName   string `json:"pluginName"`
	PluginConfig string `json:"pluginConfig"`
	PluginIndex  int    `json:"pluginPriority"`
	//PluginInfo   string `json:"pluginInfo"`
	IsStop       int    `json:"isStop"`
}

type PluginSlice []*Plugin

func (p PluginSlice) Len() int { // 重写 Len() 方法
	return len(p)
}
func (p PluginSlice) Swap(i, j int) { // 重写 Swap() 方法
	p[i], p[j] = p[j], p[i]
}
func (p PluginSlice) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	return p[i].PluginIndex < p[j].PluginIndex
}

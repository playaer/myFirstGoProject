package utils

type TemplateVars struct {
	Vars map[string]interface{}
	Data interface{}
}

func (self *TemplateVars) SetData(data interface{}) {
	self.Data = data
}

func (self *TemplateVars) AddVar(name string, value interface{}) {
	if self.Vars == nil {
		self.Vars = map[string]interface{}{
			name: value,
		}
	} else {
		self.Vars[name] = value
	}
}
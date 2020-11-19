package dbzapi

type ConfigUser struct {
	name       string
	powerlevel int
}

func (cu *ConfigUser) GetTotalPL() (int, string) {
	if cu.powerlevel > 6000 {
		return cu.powerlevel, "Es un guerrero de clase alta"
	} else {
		return cu.powerlevel, "Es un guerrero de clase baja, Insecto!"
	}
}

// FUNCTIONS
func GetLevel(imp GetConf) (int, string) {
	i, s := imp.GetTotalPL()
	return i, s
}

// Interface
type GetConf interface {
	GetTotalPL() (int, string)
}

func SetConfig(name string, plevel int) *ConfigUser {
	//fmt.Println(name)
	return &ConfigUser{name: name, powerlevel: plevel}
}

package models

type CommonDao interface {
	Save() error
	//UpdateField(field string) error
	//Remove(field string) error
	//FindAll() error
}

func Save(common CommonDao) (err error) {
	err = common.Save()
	return
}

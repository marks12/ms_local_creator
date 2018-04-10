package common

type singleton struct {

}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}   // <--- NOT THREAD SAFE
	}
	return instance
}

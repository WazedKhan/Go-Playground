package repository

var activeStorage StorageSelector

func SetStorage(s StorageSelector) {
	activeStorage = s
}

func GetStorage() StorageSelector {
	return activeStorage
}

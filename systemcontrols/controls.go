package systemcontrols

type Controller interface {
	SetBrightness(level int) error
	GetBrightness() (int, error)
	SetVolume(level int) error
	GetVolume() (int, error)
}

func New() (Controller, error) {
	return newPlatformController()
}

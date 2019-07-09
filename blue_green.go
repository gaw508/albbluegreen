package albbluegreen

type BlueGreenStatus string

const Blue BlueGreenStatus = "blue"
const Green BlueGreenStatus = "green"

type BlueGreenService interface {
	Status() (status BlueGreenStatus, err error)
	SetStatus(status BlueGreenStatus) error
	Toggle() (newStatus BlueGreenStatus, err error)
}

func InvertStatus(status BlueGreenStatus) BlueGreenStatus {
	if status == Blue {
		return Green
	}
	return Blue
}

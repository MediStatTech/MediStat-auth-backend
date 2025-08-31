package enums


type PersonalStatusEnum string

const (
	PersonalStatusEnumActive   PersonalStatusEnum = "active"
	PersonalStatusEnumInactive PersonalStatusEnum = "inactive"
)

func (r *PersonalStatusEnum) IsValid() bool {
	switch *r {
	case PersonalStatusEnumActive, PersonalStatusEnumInactive:
		return true
	}
	return false
}

func (r PersonalStatusEnum) String() string {
	return string(r)
}

type PersonalDepartureEnum string

const (
	PersonalDepartureEnumHospital PersonalDepartureEnum = "hospital"
	PersonalDepartureEnumDispatcher PersonalDepartureEnum = "dispatcher"
)

func (r *PersonalDepartureEnum) IsValid() bool {
	switch *r {
	case PersonalDepartureEnumHospital, PersonalDepartureEnumDispatcher:
		return true
	}
	return false
}

func (r PersonalDepartureEnum) String() string {
	return string(r)
}
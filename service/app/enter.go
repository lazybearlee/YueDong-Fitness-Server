package appservice

type ServiceGroup struct {
	ExerciseRecordService
	HealthStatusService
	RankService
	ExercisePlanService
	HeartRateService
	BloodPressureService
}

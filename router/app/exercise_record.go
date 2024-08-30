package approuter

import "github.com/gin-gonic/gin"

type ExerciseRecordRouter struct{}

func (e ExerciseRecordRouter) InitExerciseRecordRouter(Router *gin.RouterGroup) {
	exerciseRecordRouter := Router.Group("record")
	{
		exerciseRecordRouter.POST("insert_exercise_record", exerciseRecordApi.InsertExerciseRecord)
		exerciseRecordRouter.PUT("update_exercise_record", exerciseRecordApi.UpdateExerciseRecord)
		exerciseRecordRouter.DELETE("delete_exercise_record", exerciseRecordApi.DeleteExerciseRecord)
		exerciseRecordRouter.DELETE("delete_exercise_records", exerciseRecordApi.DeleteExerciseRecords)
		exerciseRecordRouter.GET("get_exercise_record", exerciseRecordApi.GetExerciseRecord)
		exerciseRecordRouter.POST("get_exercise_record_list", exerciseRecordApi.GetExerciseRecordList)
		exerciseRecordRouter.GET("get_all_exercise_record_of_user", exerciseRecordApi.GetAllExerciseRecordListOfUser)
	}
}

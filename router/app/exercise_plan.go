package approuter

import "github.com/gin-gonic/gin"

type ExercisePlanRouter struct{}

func (e *ExercisePlanRouter) InitExercisePlanRouter(Router *gin.RouterGroup) {
	exercisePlanRouter := Router.Group("plan")
	{
		exercisePlanRouter.POST("create_exercise_plan", exercisePlanApi.CreateExercisePlan)
		exercisePlanRouter.POST("get_exercise_plans", exercisePlanApi.GetExercisePlans)
		exercisePlanRouter.PUT("update_exercise_plan", exercisePlanApi.UpdateExercisePlan)
		exercisePlanRouter.DELETE("delete_exercise_plans", exercisePlanApi.DeleteExercisePlans)
		exercisePlanRouter.GET("get_all_exercise_plans", exercisePlanApi.GetAllExercisePlans)
		exercisePlanRouter.GET("get_uncompleted_exercise_plans", exercisePlanApi.GetUnCompletedExercisePlans)
		exercisePlanRouter.GET("get_started_exercise_plans", exercisePlanApi.GetStartedExercisePlans)
	}
}

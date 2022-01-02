package handlers
import(
	"database/sql"
	"net/http"
	"strconv"
	"go-echo-vue/models"
	"github.com/labstack/echo/v4"
)
 
type H map[string]interface{}

//GetTasks endpoint

func GetTasks(db *sql.DB) echo.HandlerFunc{
	return func(c echo.Context)error{
		// return c.JSON(http.StatusOK,"tasks")
		//Fetch tasks using our new model
		return c.JSON(http.StatusOK,models.GetTasks(db))
	}
}

//PutTask endpoint

// func PutTask(db *sql.DB) echo.HandlerFunc{
// 	return func(c echo.Context) error{
// 		return c.JSON(http.StatusOK,H{
// 			"created":123,
// 		})
// 	}
// }
func PutTask(db *sql.DB)echo.HandlerFunc{
	return func(c echo.Context) error{
		//Instantiate a new task
		var task models.Task
		//map incoming JSON body to the new Task
		c.Bind(&task)
		//Add a task using our new model
		id,err:=models.PutTask(db,task.Name)
		//Return a JSON response if successful
		if err==nil{
			return c.JSON(http.StatusCreated,H{
				"created":id,
			})
			//Handle any errors
		}else{
			return err
		}
	}
}
//DeleteTask endpoints

// func DeleteTask (db *sql.DB) echo.HandlerFunc{
// 	return func(c echo.Context) error{
// 		id,_:=strconv.Atoi(c.Param("id"))
// 		return c.JSON(http.StatusOK,H{
// 			"deleted":id,
// 		})
// 	}
// }

func DeleteTask(db *sql.DB) echo.HandlerFunc{
	return func(c echo.Context) error{
		id,_:=strconv.Atoi(c.Param("id"))
		//Use our new model to delete a task
		_,err:=models.DeleteTask(db,id)
		if err==nil{
			return c.JSON(http.StatusOK,H{
				"deleted":id,
			})
			//Handle errors
		}else{
			return err
		}
	}
}
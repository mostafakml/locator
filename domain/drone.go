package domain

type Drone struct {
	XValue string `json:"x_value" binding:"required" validate:"is_convertible"`
	YValue string `json:"y_value" binding:"required" validate:"is_convertible"`
	ZValue string `json:"z_value" binding:"required" validate:"is_convertible"`
	Vel    string 	  `json:"vel" binding:"required" validate:"is_convertible"`
}




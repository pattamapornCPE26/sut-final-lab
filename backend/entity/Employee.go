package entity

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name 	string		
	URL		string			`valid:"url~URL is invalid"`
	Employee	string		`valid:"matches(EM), stringlength(12|12)~EmployeeID not true"`
}

// `valid:"required:Name is required"`
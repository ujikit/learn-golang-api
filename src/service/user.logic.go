package service

import (
	context "context"

	// Interface
	ServiceInterface "learn-golang-api/src/service/interface"
	
	// Logic
	LanguageLogic "learn-golang-api/src/helper/language"
)

func UserDetail (ctx context.Context) (ServiceInterface.UserData){
	// static status code, for testing purpose
	statusCode := 200
	var UserData ServiceInterface.UserData
	
	if statusCode == 200 {
		
		user1 := ServiceInterface.User{"Ashley Cole", 21}
		statusCode := ServiceInterface.StatusCode{statusCode,LanguageLogic.Translate(ctx, "Response_DataSuccess")}
		users := []ServiceInterface.User{user1}
		response := ServiceInterface.UserData{StatusCode: statusCode, Data: users}
		
		return response
		} else if statusCode == 500 {
			
		statusCode := ServiceInterface.StatusCode{statusCode,LanguageLogic.Translate(ctx, "Response_ServerError")}
		response := ServiceInterface.UserData{StatusCode: statusCode}

		return response
	}
	return UserData
}
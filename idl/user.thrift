namespace go user

struct User{
    1:i64 id
    2:string name
    3:string password
}

struct ServiceResBasic {
    1:i32 http_code
    2:i32 status_code
    3:string status_message
}

struct ServiceReqRegister{
    1:string username
    2:string password
}

struct ServiceResRegister{
    1:ServiceResBasic base_res
    2:i64 user_id
    3:string token
}

struct ServiceReqSignin{
    1:string username
    2:string password
}

struct ServiceResSignin{
    1:ServiceResBasic base_res
    2:i64 user_id
    3:string token
}

struct ServiceReqGetUserById{
    1:i64 id
}

struct ServiceResGetUserById{
    1:ServiceResBasic base_res
    2:User user
}

service UserService {
    ServiceResRegister Register(1:ServiceReqRegister ServiceReqRegister)
    ServiceResSignin Signin(1:ServiceReqSignin ServiceReqSignin)
    ServiceResGetUserById GetUserById(1:ServiceReqGetUserById ServiceReqGetUserById)
}
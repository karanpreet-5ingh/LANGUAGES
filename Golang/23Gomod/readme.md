#go mod init github.com/karanpreet-5ingh/mymodules 

this should be the server where you want to place you code 

go.dev/blog/using-go-modules

github.com/gorilla/mux

require github.com/gorilla/mux v1.8.0 // indirect

indirect means that this module is not yet used in the code anywhere


go.sum file check sum of the file whether it is from the authorized user or not




go env 



go build .




go mod tidy


go mod verify  //to verify the go.sum file that every thing is working fine or not 




go list


go list all // dunmp all the packages and dependency 

go list  -m -versions github.com/gorilla/mux   // throw  all the version of gorilla mux



 go mod tidy // tide up all the package which i am using also remove all the package that i am not using also try to bring all the required packages  





go mod why  github.com/gorilla/mux
// gives the reasi=on why i am using this module and which package is dependent on what 



go mod graph  




go mod vendor 

so vendor jbb aajata h tbbhame internet se module download krne ki jzurat nhi padti
so ek vendor  naam ka folder ban jata h 
 


 now do this, 
 go run -mod=vendor main.go



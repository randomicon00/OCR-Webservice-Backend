package main 
//Conversion - Model of a common conversion orc -> text
type Conversion struct {
	IPAddress	string	
	Filename	string
	Language 	string
	Output		string
	Status   	string
	Success		bool
	Error		string	
}

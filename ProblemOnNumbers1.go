/************************************************/
//
// Auther : Manali Kulkarni
// Date : 19/07/2022
// Subject : Problems on Numbers
//
/***********************************************/

/************* Imports *******************/
package main

import "fmt"

/*************** Helper Functions ************/

//Program to add 2 numbers
func Sum(iNo1 int, iNo2 int)int{
  var iRet int;
  iRet = iNo1 + iNo2;
  return iRet;
}

//Progeam to display class from percentage
func displayPer(fPer float64){
  if fPer >= 70.00{
    fmt.Println("Distinction");
  }else if fPer >= 60.00 && fPer < 70.00{
    fmt.Println("First Class");
  }else if fPer >= 50.00 && fPer < 60.00{
    fmt.Println("Second Class");
  }else if fPer >= 35.00 && fPer < 50.00{
    fmt.Println("Pass Class");
  }else{
    fmt.Println("Fail");
  }
}

//Program to dispaly City Name using city code
func cityCode(iCode int){
  switch iCode {
  case 01:
    fmt.Println("Mumbai");
  case 02:
    fmt.Println("Pune");
  case 03:
    fmt.Println("Solapur");
  default:
    fmt.Println("Invalid Code");

  }
}

//WAP which accepts a number from user and dispaly it for 5 times
func displayNum(iNo int){
  for iCnt := 0; iCnt < 5;iCnt++{
    fmt.Println(iNo);
  }
}

//Accept no. from user & check wethwer it is dividble by 5 and 3
func displayDiv(iNo int){
  if iNo%3==0 && iNo%5==0{
    fmt.Println("True");
  }else{
    fmt.Println("False")
  }
}

/*************** Driver Function ************/
func main(){
  /*fmt.Println("Enter Numbers for addtion:");
  var iNo1, iNo2 int;
  iNo1 = 0; iNo2 = 0;
  fmt.Scanf("%d%d",&iNo1,&iNo2);
  fmt.Println("Sum is : ",Sum(iNo1,iNo2));*/

  /*fmt.Println("Enter marks in Percentage : ");
  var fPer float64;
  fmt.Scanf("%f",&fPer);
  displayPer(fPer);*/

  /*fmt.Println("Enter City Code : ");
  var iCode int;
  fmt.Scanf("%d",&iCode);
  cityCode(iCode);*/

  fmt.Println("Enter Number : ");
  var iNo int;
  fmt.Scanf("%d",&iNo);
  //displayNum(iNo);
  displayDiv(iNo);
}

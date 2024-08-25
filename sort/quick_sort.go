package main


import "fmt"


func main(){

	 var arraySize int
	 fmt.Printf("enter the array size")
	 fmt.Scanf("%d",&arraySize)

	 arr := make([] int ,arraySize)

	 acceptInput(arr)
	 quick_sort(arr)
	 displayArray(arr)	 
}

func acceptInput(arr [] int ){

	for index,_:=range arr{
		fmt.Printf("enter data:")
		fmt.Scanf("%d",&arr[index])
	}
}

func quick_sort(arr [] int){

	int i,j,mid

	i=0
	j=len(arr)-1

	mid = (j+i)/2
}

func displayArray(arr [] int){

	for _,value:=range arr{
		fmt.Printf("value is %d\n",value)
	}
}
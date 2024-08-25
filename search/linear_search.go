package main


import "fmt"

func main(){


	var n ,key int 

	fmt.Printf("enter the array size:")
	fmt.Scanf("%d",&n)

	arr:= make([] int ,n)

	acceptInput(arr)
	displayArray(arr)

	fmt.Printf("enter data to search")
	fmt.Scanf("%d",&key)

	pos:=linearSearch(arr,key)

	if pos==-1{
		fmt.Printf("data is not found");
	}else{
		fmt.Printf("data is %d found at index %d ",key,pos)
	}


}


func acceptInput(arr [] int){

	for i:=0;i<len(arr);i++{

		fmt.Printf("enter data")
		fmt.Scanf("%d",&arr[i])

	}
}

func displayArray(arr [] int){

	for index,value :=range arr{
		fmt.Printf(" %v ::DATA %d\n",index+1,value)
	}
}

func linearSearch(arr [] int, key int)int{

	for index,_:=range arr{

		if arr[index]==key{
			return index
		}
	}
	return -1
}
package main


import "fmt"


func main(){

	var arry_size int
	
	fmt.Printf("how many data you want to store")
	fmt.Scanf("%v",&arry_size)

	arr := make([] int ,arry_size)

	acceptInput(arr,arry_size)
	insertion_Sort(arr,arry_size)
	displayArray(arr,arry_size)
}


func acceptInput(arr [] int ,arry_size int ){
	for i:=0;i<arry_size;i++{
		fmt.Printf("enter the element %v:",i+1)
		fmt.Scanf("%v",&arr[i])
	}
}

func insertion_Sort(arr [] int ,arry_size int){
	
	var temp_memory int

	for i:=1;i<arry_size;i++{
		var j int
		temp_memory=arr[i];

		for j = i-1; j>= 0; j-- {
			if arr[j]>temp_memory{
				arr[j+1]=arr[j];
			}else{
				break;
			}
		}
		arr[j+1]=temp_memory
	}

}

func displayArray(arr [] int,arry_size int){
	for i := 0; i < arry_size; i++{
		fmt.Println("your array ",arr[i])
	}
}



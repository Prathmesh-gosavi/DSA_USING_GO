package main

import "fmt"


func main(){

	var arry_size int
	
	fmt.Printf("how many data you want to store")
	fmt.Scanf("%v",&arry_size)


	arr := make([] int ,arry_size)

	acceptInput(arr);
	selectionSort(arr);
	displayArray(arr);
}


func acceptInput(arr [] int ){

	for index,_:=range arr{
		fmt.Print("enter the data")
		fmt.Scanf("%d",&arr[index])
	}
}





func selectionSort(arr [] int ){

	n:=len(arr)
	for i:=0;i<n;i++{

		for j:=i+1;j<n;j++{

			if arr[i]>arr[j]{

				arr[j],arr[i]=arr[i],arr[j]
			}
		}
	}
}

func displayArray(arr [] int ){

	for index,value:=range arr{

		fmt.Printf("index is %v,value is %v\n",index,value)
	}
}



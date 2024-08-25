package main


import "fmt"

func main(){


	var n ,key int 

	fmt.Printf("enter the array size:")
	fmt.Scanf("%d",&n)

	arr:= make([] int ,n)

	acceptInput(arr)
	bobbleSort(arr)
	displayArray(arr)

	fmt.Printf("enter data to search")
	fmt.Scanf("%d",&key)

	pos:=binarySearch(arr,key)

	if pos==-1{
		fmt.Printf("data is not found\n");
	}else{
		fmt.Printf("data is %d found at index %d \n",key,pos)
	}


}


func acceptInput(arr [] int){

	for i:=0;i<len(arr);i++{

		fmt.Printf("enter data")
		fmt.Scanf("%d",&arr[i])

	}
}

func bobbleSort(arr [] int){

	for i := len(arr)-1; i >= 0 ; i--{
		for j:=0;j<i;j++{

			if arr[j]>arr[j+1]{

				arr[j+1],arr[j]=arr[j],arr[j+1]
			}
		}

		}
}

func displayArray(arr [] int){

	for index,value :=range arr{
		fmt.Printf(" %v ::DATA %d\n",index+1,value)
	}
}

func binarySearch(arr [] int, key int)int{

	i:=0;
	j:=len(arr)-1;
		
	for i <= j{

		 mid := (j+i)/2

		if arr[mid]==key{

			return mid

		}else if arr[mid]>key{
			
			j=mid-1
		}else{
			
			i=mid+1
		}
	}


	return -1
}
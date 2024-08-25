package main

import "fmt"


func main(){

	var arry_size int
	
	fmt.Printf("how many data you want to store")
	fmt.Scanf("%d",&arry_size)


	arr := make([] int ,arry_size)

	acceptInput(arr);
	mergeSort(arr,0,arry_size-1);
	displayArray(arr);
}


func acceptInput(arr [] int ){

	for index,_:=range arr{
		fmt.Print("enter the data")
		fmt.Scanf("%d",&arr[index])
	}
}


func merge(arr [] int,lowerbound int ,mid int ,upperbound int){

	temparray := make([]int, upperbound-lowerbound+1)
	i:=lowerbound;
	j:=mid+1;
	k:=0;


	for i <= mid && j<=upperbound {

		if arr[i]<arr[j]{

			temparray[k]=arr[i]
			i++

		}else{

			temparray[k]=arr[j]	
			j++			
			
			}
			k++
		}

	

	for i <= mid {

		temparray[k]=arr[i]
		i++ 
		k++
	}

	for j <= upperbound{

		temparray[k]=arr[j]
		k++
		j++
	}

	for i := 0; i < len(temparray); i++ {
         arr[lowerbound+i] = temparray[i]
    }
	

}




func mergeSort(arr [] int ,lowerbound int , upperbound int){


	if lowerbound<upperbound{

		mid :=(lowerbound+upperbound)/2;

		mergeSort(arr,lowerbound,mid);
		mergeSort(arr,mid+1,upperbound);
		merge(arr,lowerbound,mid,upperbound);

	}

	
}

func displayArray(arr [] int ){

	for index,value:=range arr{

		fmt.Printf("index is %d,value is %d\n",index,value)
	}
}



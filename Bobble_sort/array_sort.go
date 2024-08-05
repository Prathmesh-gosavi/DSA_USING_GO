package main


import "fmt"


func main(){


	var n int

	fmt.Printf("enter the value of n:")
	fmt.Scanf("%d",&n)


	a:=make([] int ,n)

	accept(a,n)
	sorting(a,n)
	display(a,n)

}


func accept(a[] int,n int ){

	var i int

	for i=0;i<n;i++{
		fmt.Printf("enter the data")
		fmt.Scanf("%d",&a[i])
	}
}

func sorting(a[]int,n int){

	var i,j int

	//var temp int

	for i=0;i<n-1;i++{
		for j=0;j<n-i-1;j++{
			if a[j]>a[j+1]{

				/*temp=a[j]; error generate temp declere but not use
				a[j]=a[j+1];
				a[j+1]=temp;*/

				a[j],a[j+1]=a[j+1],a[j]//using go lang asine two variable at a time
			}
		}
	}

}

func display(a[]int ,n int ){

	var i int

	fmt.Println("sorted array")
	for i=0;i<n;i++{
		fmt.Println(a[i])
	}

}
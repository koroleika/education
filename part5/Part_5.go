package main

import "fmt"

func main(){
	var (
		arr [8][8]int
		dist int
		v1 int
		v2 int
		d[8] int
		v[8] bool
	)

	for i:=0;i<8;i++{
		arr[i][i]=0
		for j:=i+1;j<8;j++{
			fmt.Printf("Введите расстояние %d - %d: ", i+1, j+1)
			fmt.Scanf("%d", &dist)
			arr[i][j] = dist
			arr[j][i] = dist
		}
	}
	for i:=0;i<8;i++{
		for j:=0;j<8;j++{
			fmt.Printf("%5d %s", arr[i][j], "    ")
		}
		fmt.Println()
	}

	fmt.Print("Введите номера вершин для поиска минимального расстояния: ")
	fmt.Scanf("%d", &v1,&v2)
	//v1 = 0

	for i:=0;i<8;i++{
		d[i] = 10000
		v[i] = false
	}
	d[v1] = 0
	for{
		minindex := 10000
		min:= 10000

		for i:=0; i<8;i++{
			if (v[i] == false) && (d[i]<min){
				min = d[i]
				minindex = i
			}
		}

		if minindex != 10000{
			for i:=0; i<8;i++{
				if (arr[minindex][i] > 0){
					dist = min+arr[minindex][i]
					if dist < d[i]{
						d[i] = dist
					}
				}
			}
			v[minindex] = true
		}

		if !(minindex < 10000){
			break
		}
	}

	fmt.Print("Кратчайшие расстояние до вершин  ")
	for i:=0;i<8;i++{
		fmt.Print(d[i], "    ")
	}

}

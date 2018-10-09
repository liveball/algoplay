package core

//BubbleSort for bubble sort
func BubbleSort(data interface{}) {
	d := data.([]interface{})

	for i := 0; i < len(d)-1; i++ { //control  outer layer cycle
		for j := 0; j < len(d)-i-1; j++ { //swap two adjacent elements

			switch d[j].(type) {
			case int:
				m := d[j].(int)
				n := d[j+1].(int)
				if m < n {
					t := d[j]
					d[j] = d[j+1]
					d[j+1] = t
				}

			case string:
				m := d[j].(string)
				n := d[j+1].(string)
				if m < n {
					t := d[j]
					d[j] = d[j+1]
					d[j+1] = t
				}
			}

		}
	}
}

func SelectionSort(data interface{}) {
	d := data.([]interface{})
	for i := 0; i < len(d)-1; i++ {
		min := i
		for j := i + 1; j < len(d); j++ { //search the min elem
			switch d[j].(type) {
			case int:
				m := d[j].(int)
				n := d[min].(int)
				if m < n {
					min = j //save the min elem index
				}
			case string:
				m := d[j].(string)
				n := d[min].(string)
				if m < n {
					min = j
				}
			}
		}
		//keep order zone by min
		t := d[i]
		d[i] = d[min]
		d[min] = t
	}
}

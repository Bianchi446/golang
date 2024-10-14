/* Within the Go runtime, a map is implemented as a pointer to a struct; 
	Passing a map to a function means you are copying a pointer

	1. Rather than passing a map around, use a struct 
	2. Passing a slice to a function -- Any modifications in the contents of the slice 
		is reflected in the original value. 

	3. Also, changes to the length and capacity are not reflected back in the original. 
*/






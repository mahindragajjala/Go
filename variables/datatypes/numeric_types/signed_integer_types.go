package numerictypes

import "fmt"

/*

Type	Size	                         Range
int8	8-bit	                        -128 to 127
int16	16-bit	                        -32,768 to 32,767
int32	32-bit	                        -2,147,483,648 to 2,147,483,647
int64	64-bit	                        -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
int	    Depends on system (32 or 64-bit)	Platform dependent

*/
func Signed_integer_type() {
	var a int8 = -128                 // Minimum for int8
	var b int16 = 32767               // Maximum for int16
	var c int32 = -2147483648         // Minimum for int32
	var d int64 = 9223372036854775807 // Maximum for int64
	var e int = -123456               // Platform-dependent (usually 64-bit)

	fmt.Println("int8:", a)
	fmt.Println("int16:", b)
	fmt.Println("int32:", c)
	fmt.Println("int64:", d)
	fmt.Println("int:", e)
}

package stringtypes

/*

What is a string in Go?
In Go, string is a basic data type used to store textual data.

- A string represents a sequence of bytes — and each byte usually
holds a single character (if it's plain ASCII).

But Go strings can hold any data, even UTF-8 encoded text (which might use multiple
bytes per character for non-ASCII characters).

----------------------------------Strings are immutable-----------------------------------------
Once a string is created, you cannot change it.
If you want to "modify" a string, you need to create a new one.
									s := "hello"
									s[0] = 'H'  // ❌ Invalid — won't compile

-----------------------------------String length (len)------------------------------------------------------
The len() function returns the number of bytes in the string, not the number of characters.
If the string is ASCII, bytes == characters.
If the string uses UTF-8 characters (like emojis or Hindi text),
the byte count will be larger than the character count.
									s := "hello"
									fmt.Println(len(s))  // Output: 5

									s := "हेलो"  // Hindi word "Hello"
									fmt.Println(len(s))  // Output: 9 (each character here uses multiple bytes)

----------------------------------- Accessing string elements (s[i])------------------------------------------------------
You can access a string's bytes by using indexing.
Indices start from 0 to len(s)-1.
Each s[i] returns a byte (alias: uint8).
									s := "hello"
									fmt.Println(s[0])         // Output: 104 (ASCII for 'h')
									fmt.Printf("%c\n", s[0])  // Output: h
Extra Info: Constant Length
If your string is a constant, its length is determined at compile time.

const greeting = "hello"
fmt.Println(len(greeting))  // Compiler knows this is 5 at compile time.
If it's a variable, the length is known at runtime.

*/

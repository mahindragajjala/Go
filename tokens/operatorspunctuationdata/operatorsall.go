package operatorspunctuationdata

/*

&= - ADDRESS
() - parentheses
|= - OR and =
[] - square brackets
* - pointer,multiplication
^ - RAND
* -  a = a*b
^ -  a = a^b
{} - curly brackets
++ - increment
, - Not operator its "separator"
; - semicolon
>>= - right shift and equal to
-- - decrement
... - ellipsis
. - access operator
: - colon
&^ - bitclear
&^= - AND_NOT_assignment_operator



Arithmetic - Performs basic math such as addition, subtraction, multiplication
+ - addition
- - subtraction
* - multiplication
/ - division
% - remainder

Assignment - Assigns or updates variable values
=  assignment operator
+= - compound assignment operator
-= - ""
*= - ""
/= - ""
%= - ""

Relational - Compares values and returns a truth value
< - left operand is less than the right operand
> - > operator is a comparison operator used to check if the left operand is greater than the right operand.
<= - less than or equal to
>= - greater than or equal to
== - comparsion checking both elements or same or not

!= - Not equal to ( != ): Returns true if the operands are not equal.


Logical	- Combines multiple conditions or inverts them

&& - OR && is used in boolean expressions to combine conditions.
|| - logical_operator
! - NOT

Bitwise	Works - directly on the binary representation of integers
& - & operator is used to get the memory address of a variable,
	i.e., to obtain a pointer to that variable.
| - bitwise_operator

^ - kar-et - the ^ operator has different meanings depending on the context —
			 it’s not like in some other languages where it's used for exponentiation
~ - not there

Why ^ instead of ~?
Go is designed to be simple and consistent. It avoids multiple symbols for the same operation.

So:

^ between two values = bitwise XOR

^ before a value = bitwise NOT

<< -  bitwise left shift operator
>> -  bitwise right shift operator

Special -	Covers unique tasks like size checking, pointer access, and more
?: - ternary operator - not there in golang
. - access operator
<- - used in channel .

*/

func AllOperator() {
	//Addition_operator()
	//Passing_by_Reference_to_Functions()
	//Ampersand_with_New()
	//Ampersand_Linking_Data_Structures()
	//Compound_assignment_operator()
	//BitwiseOperator()
	//	Pointer_operator()
	//ModulesOperator()
	//Assignment_operator
	//Logical_NOT_operator()
	//BitwiseOperator()
	Variadic_Function_Parameters()
}

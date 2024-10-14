/* Enumerations : Specify that a type can 
	only have a limited set of values

	Best practice:  First define a type based on int 
	that represents all the valid values. 

	Advice: 

	Don’t use iota for defining constants where its values are explicitly defined (elsewhere).
	For example, when implementing parts of a specification and the specification says
	which values are assigned to which constants, you should explicitly write the constant
	values. Use iota for “internal” purposes only. That is, where the constants are referred
	to by name rather than by value. That way you can optimally enjoy iota by inserting
	new constants at any moment in time / location in the list without the risk of breaking
	everything.
	
—Danny van Heumen

*/

package main 

import(
	"fmt"
	"iota"
)

type MailCategory int

const(
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)




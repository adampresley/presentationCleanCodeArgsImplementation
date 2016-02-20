Notes
=====

* This implements an argument parser
	* A string of arguments
	* Each argument ID is a single letter
	* The type is denoted by
		* * == string
		* # = integer
		* ## = double
* Start by analyzing source code organization
	* Observe there is a usage in main
	* A primary class to interface with
	* An interface that defines a marshaller
	* Various marshal implementations
	* Exception classes
* Starting with the interface. Seems like the simplest item to begin with
	* Go's idiom of interfaces is similar to Java
	* Same basic principles apply. It defines a contract that structures can implement
	* The interface uses a lower case function name for something being used publicly. Will have to upper case for Go
	* The interface calls for an interator of type string as the argument to Set. Not sure if it is necessary.
		* Considered a byte buffer
		* Considered a string reader
		* Not sure I need that
* Realized the method on the interface expects to throw an exception. In Go we will return an error
	* Pivoting to make the ArgException structure
	* Think that the ArgsException structure should implement the error interface
	* See http://golang.org/pkg/builtin/#error for information about the error interface
	* Wrote first test. ```An ArgsException implements the error interface```
	* The ArgsException class expects an error code, which is an enum. Choosing to make a package to hold constants. Simple
	* While writing tests, decided getters and setters is way too Java... don't need it
	* Wrote tests to cover the various types of error messages to expect
* Implemented error codes. It is a **custom type** with constants
* Ready to start implementing the struct
	* Implemented constructors and struct
	* Implemented Error() function with returning blank string.
	* All tests fail. Expected
	* Filled in the values. Had some tests with bad copy/paste jobs. Fixed
* NOW time for the interface
* Testing BooleanArgumentMarshaler
* Implemented BooleanArgumentMarshaler, StringArgumentMarshaler, and IntegerArgumentMarshaler
* During implementation I strugged with the conversion. It didn't feel very **Go**... to Java
*

Observations
------------

* Book talks about what makes clean code
	* Breaks each concept down and gives examples
	* Will provide examples of non-clean code, or trivial implementations
* Specific details include
	* Naming is very important
	* How much time is spent reading code vs writing code?
		* [Peter Hallam states he spends 78% of his time understanding code](http://blogs.msdn.com/b/peterhal/archive/2006/01/04/509302.aspx)
			* Peter Hallam worked on Visual Studio, C# compiler
		* Robert Martin claims the ratio of time spent reading code to writing code is 10 to 1
			* Clean Code book
		* [http://www.itworld.com/article/2828299/it-management/software-engineers-spend-lots-of-time-not-building-software.html](Electric Cloud surveyed 443 developers in 2013 and found that only half their time was spent actually coding during a week)
		* [http://www.program-comprehension.org/](There's a conference every year dedicated to the science and art of program comprehension)
		* IBM performed a study in 1989 and determined that more than half the effort in accomplishing a programming task is toward understanding the system
		* [Bell Labs in 1992](http://www.manclswx.com/papers/evolving.pdf) - A research paper is published that analyzed how much time a developer spends in "discovery" over various phases of a project, coupled with experience in the project. 60-80% in discovery early in a project
		*
	* Function composition is important
		* Single responsibility
		* The less parameters the better. The more parameters there are
			* The harder it is to reason about
			* Difficult to remember what the parameters are, or what they do
		* Flag arguments (booleans) are a bad idea. You will forget what the true/false is for, or what it does. It may also signal that the function does more than one thing
	*
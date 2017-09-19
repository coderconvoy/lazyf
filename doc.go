//Lazyf format is intended to provide a simple way of getting data from a non technical or lazy user, into a format that golang can work with.
//Files are broken up in LZ objects, and their properties tabbed in and ":" separated for name:value
//eg:
//"LZob1
//"    propertyname:propertyvalue
//"	   anotherproperty:anothervalue
//
//In this case quotes are not needed, the owner object decides what type it should be at calltime using the PString,PInt,PFloat methods to extract the desired data
package lazyf

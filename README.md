# go-deser-problem

The purpose of this repo is to demonstrate a problem that I ran into while trying to support
the OpenAPI `oneOf` feature in Go.

In the problem I'm trying to solve, I have an example struct defined with two fields like this:

```
type Resource struct{
        ID *string`json:"id" validate:"required"`
        Info InfoIntf `json:"info,omitempty"`
}
```

plus three other structs:

```
type Info struct{
        Foo *string`json:"foo,omitempty"`
        Bar *string`json:"bar,omitempty"`
}
type Foo struct{
        Foo *string`json:"foo" validate:"required"`
}
type Bar struct{
        Bar *string`json:"bar" validate:"required"`
}
```

The "Info" field in the Resource struct models a schema property that uses "oneOf" 
(Info is expressed as being "one of" either Foo or Bar) and so in this example I want the user to be able to 
supply an instance of the Foo struct or an instance of the Bar struct wherever the Info struct is called for.
So to accomplish this, I define an interface (InfoIntf) along with the Info, Foo and Bar structs which 
all "implement" the interface.   This allows for the correct serialization of the Resource struct regardless
of whether the Info field is a pointer to a Foo or a Bar.

Unfortunately (and unsurprisingly), the deserialization doesn't work because the default json deserializer 
doesn't know how to deserialize the "Info" field into the "InfoIntf" interface.
To support the deserialization of a schema property that uses oneOf such as the "Resource.Info" field, we 
generate the "parent" Info struct (i.e. "parent" of both Foo and Bar) so that it contains the union of the properties
found in the Foo and Bar structs.  Assuming that we have a JSON string that represents a valid instance of 
the Resource schema (i.e. the "info" field contains either a "Foo" or "Bar" instance), we should be able to 
deserialize into the Info struct.   The problem is that the default json deserializer gets tripped up
because of the use of the interface (InfoIntf) when declaring the Info field (which is done to support the 
serialization side of the equation).

So... the crix of the problem is... how can I implement a custom deserializer/unmarshaller that will accomplish this?
 
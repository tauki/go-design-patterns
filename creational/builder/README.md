# Builder design pattern

A builder pattern is designed to provide a flexible way to solve object creation problems. It deals with problems like how a same process can construct different representation of a complex object or, simplifying the process of creating an object. It usually does so by separating the responsibility of creating and assembling the parts of an object by separate `builder` object. 

> **The intent of the Builder design pattern is to separate the construction of a complex object from its representation. By doing so the same construction process can create different representations.**
>  &mdash; Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides (1994). Design Patterns: Elements of Reusable Object-Oriented Software. Addison Wesley. pp. 97ff. ISBN 0-201-63361-2.

A `builder pattern` has it's own pros and cons, like every other design pattern, it is designed to solve a particular set of problems while having its own trade offs. For example, the design pattern encapsulates construction process and representation, allows you to vary an objects internal representation, but it also requires you to create separate concrete builders for each type of objects.

The pattern tries to bring in an abstraction so there is an separation between the object user and the creation. At the same time it tries to reuse object creation algorithm for different objects. It has been described as the relationship between a director, Builders, and the product they build.

In the particular exercise a `vehicle manufacturing` process is implemented using a builder pattern. The requirements are as follows:

> It must have a manufacturing type that constructs everything that a vehicle needs.

> When using a car builder, the `VehicleProduct` with four wheels, five seats, and a structure defined as `Car` must be returned.

> When using a motorbike builder, the `VehicleProduct` with two wheels, two seats, and a structure defined as `Motorbike` must be returned.

> A `VehicleProduct` built by any `BuildProcess` builder must be open to modifications.

There is a test included which tests if the requirements are fulfilled.


---
* _Builder:_
Abstract interface for creating objects (product).
* _ConcreteBuilder:_ 
Provides implementation for Builder. It is an object able to construct other objects. Constructs and assembles parts to build the objects.

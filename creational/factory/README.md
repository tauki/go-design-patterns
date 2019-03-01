# Factory Design Pattern

Factory design pattern is a creational design pattern and one of the most popular design pattern. It deals with creating objects without specifying the type of the object. This is done by delegating the object creation to a factory and the factory will expose an interface that meets the user's need. By this it also enables the implementation to be flexible to upgrading and downgrading if needed and provides an extra layer of encapsulation so the growth of the object is controlled. The factory method defers object installation to the user.

For example, suppose there's an app for a food chain. When you order a pizza online there's an abstraction to keep the user uncluttered from how and where the pizza is being made and the factory, from user's input decides which branch the pizza will be made in and returns the appropriate object for the user to use.

Or, imagine an app that selects shape according to what an object looks like. There is a shape interface and using that interface providing the information of the object the user gets the appropriate object using the shapeFactory. For a round object, using the shapeFactory will return an object that represents the round or circle shape.

Factory design pattern grants us an extra level of encapsulation so that our program can grow in a controlled environment.

> **Define an interface for creating an object, but let subclasses decide which class to instantiate. The Factory method lets a class defer instantiation it uses to subclasses.**
>  &mdash; Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides (1994). Design Patterns: Elements of Reusable Object-Oriented Software. Addison Wesley.

We are going to implement a payment method factory here as an example
The acceptance criteria are as follows:
- To have a common method for every payment method called `Pay`
- To be able to delegate the creation of payments methods to the Factory
- To be able to add more payment methods to the library by just adding it to the factory method

There is a test included which tests if the requirements are fulfilled.
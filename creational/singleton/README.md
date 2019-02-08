# Singleton Design Pattern

As the name implies, Singleton is constitutes a single instance of an object and restricts duplication. This is useful when exactly one object is needed to coordinate actions across the system. 

In the particular exercise, a singleton is implemented for a `unique counter` which will increment the count upon used.

the _requirements_ are as follows:

> When no counter has been created before, a new one is created with the value 0

> If a counter has already been created, return this instance that holds the actual count

> If we call the method AddOne, the count must be incremented by 1

There's a test included which tests if the requirements are fulfilled.

_This implementation is not thread safe_
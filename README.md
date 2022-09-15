#CLEAN_ARCHITECTURE_CRUD

This repo is aimed at providing a simple CRUD application using Clean Architecture, which basically means organizing the application codebase in a way that is independent of frameworks, UI, databases, and is testable.

- More About Clean Architecture:

Clean Architecture is a software design philosophy that separates the elements of a design into ring levels. By doing this, the dependencies between the rings are forced to go only inward. This means that the dependencies are forced to be one way, from the outermost ring to the innermost ring. This makes the code more flexible, easier to test, and easier to maintain.

![](https://github.com/binadam1983/CLEAN_ARCHITECTURE_CRUD/blob/958903e5e6bf49bbae9b57ca8278a56b24a41524/CleanArchitecture.jpg)

The rings are as follows:

1. Entities/Models/Domain
2. Use Cases/ Services
3. Controller/ Presentation Layer/ UI Layer/ Delivery
4. Repository/ Data Access Layer

![](https://github.com/binadam1983/CLEAN_ARCHITECTURE_CRUD/blob/958903e5e6bf49bbae9b57ca8278a56b24a41524/clean-arch.png)

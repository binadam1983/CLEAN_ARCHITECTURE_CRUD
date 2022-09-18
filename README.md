# THE_CLEAN_ARCHITECTURE_CRUD

This work-in-progress repo is aimed at providing a simple CRUD application using Clean Architecture, which basically means organizing the application codebase in a way that is independent of frameworks, UI, databases, and is testable.

- More About Clean Architecture:

Clean Architecture is a software design philosophy that separates the elements of a design into ring levels. By doing this, the dependencies between the rings are forced to go only inward. This means that the dependencies are forced to be one way, from the outermost ring to the innermost ring. This makes the code more flexible, easier to test, and easier to maintain.

The rings are as follows:

1. Entities/Models/Domain
2. Use Cases/ Services
3. Controller/ Presentation Layer/ UI Layer/ Delivery
4. Repository/ Data Access Layer

![](https://github.com/binadam1983/THE_CLEAN_ARCHITECTURE_CRUD/blob/8eca1a5d4c39454e56a1662e67de4ed11cd1512f/images/clean-arch.png)

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

Moreover, I have used Golang's GIN web framework and MySQL as persistent DB, but the app is not dependent on them, given the arch..

![](https://github.com/binadam1983/THE_CLEAN_ARCHITECTURE_CRUD/blob/e939ec1b5372af164252a95286f2beaf395e6233/images/TCAC.mp4)

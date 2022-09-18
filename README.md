# CLEAN_ARCHITECTURE_CRUD

This work-in-progress repo is aimed at providing a simple CRUD application using Clean Architecture, which basically means organizing the application codebase in a way that is independent of frameworks, UI, databases, and is testable.

- More About Clean Architecture:

Clean Architecture is a software design philosophy that separates the elements of a design into ring levels. By doing this, the dependencies between the rings are forced to go only inward. This means that the dependencies are forced to be one way, from the outermost ring to the innermost ring. This makes the code more flexible, easier to test, and easier to maintain.

The rings are as follows:

1. Entities/Models/Domain
2. Use Cases/ Services
3. Controller/ Presentation Layer/ UI Layer/ Delivery
4. Repository/ Data Access Layer

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

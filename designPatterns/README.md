# Structural Design Patterns

Structural patterns explain how to assemble objects and classes into larger structures while keeping these structures flexible and efficient.

It describes the relationship with entities, it deals with the way the architecture or organization of data are created. They are used to form large structures using classes and objects.

## Gang of Four Structural Design Patterns

### Adapter

It is a pattern that allows incompatible interfaces or objects, interact together.

The Adapter acts as a wrapper between two objects. It catches calls for one object and transforms them to format and interface recognizable by the second object.

The adapter pattern comprises the target, adaptee, adapter, and client:

- Target is the interface that the client calls and invokes methods on the adapter and adaptee.

- The client wants the incompatible interface implemented by the adapter.

- The adapter translates the incompatible interface of the adaptee into an interface that the client wants.

- Bridge
- Composite
- Decorator
- Facade
- Flyweight
- Private class data
- proxy

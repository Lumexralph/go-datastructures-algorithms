# Flyweight

It is a structural design pattern that allows programs to support vast quantities of objects by keeping their memory consumption low.

This constant data of an object is usually called the intrinsic state. It lives within the object; other objects can only read it, not change it. The rest of the object’s state, often altered “from the outside” by other objects, is called the extrinsic state.

The Flyweight pattern suggests that you stop storing the extrinsic state inside the object. Instead, you should pass this state to specific methods which rely on it. Only the intrinsic state stays within the object, letting you reuse it in different contexts. As a result, you’d need fewer of these objects since they only differ in the intrinsic state, which has much fewer variations than the extrinsic. [Source](https://refactoring.guru/design-patterns/flyweight)

The object that only stores the intrinsic state is called flyweight. It stores the constant state of the application. The extrinsic state can be added to another class which will be recreated throughout the application but flyweight object (holding the intrinsic data) is created once and should be immutable. It should only initialize its state once through a constructor.

## Benefits of this pattern

The benefit of applying the pattern depends heavily on how and where it’s used. It’s most useful when:

- an application needs to spawn a huge number of similar objects
- this drains all available RAM on a target device
- the objects contain duplicate states which can be extracted and shared between multiple objects

NB: If your program doesn’t struggle with a shortage of RAM, then you might just ignore this pattern for a while.

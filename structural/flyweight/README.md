# Flyweight

It is a structural design pattern. This pattern is used when a large number of similar objects need to be created. These objects are called flyweight objects and are immutable.

When to Use?
When the objects have some intrinsic properties which can be shared.
As in the above example, dress is the intrinsic property that was taken out and shared.
Use flyweight when a large number of objects needs to be created which can cause memory issue. In case figure out all the common or intrinsic state and create flyweight objects for that.

UML Diagram:
![](./../../images/Flyweight-Design-Pattern.png)
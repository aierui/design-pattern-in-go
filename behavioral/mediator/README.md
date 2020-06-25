# Mediator

Mediator design pattern is a behavioral design pattern. This pattern suggests creating a mediator object to prevent direct communication among objects so that direct dependencies between them is avoided.

One very good example of a mediator patter is the railway system platform.  Two trains never communicate between themselves for the availability of the platform. The stationManager acts as a mediator and makes the platform available to only one of the trains. The train connects with stationManager and acts accordingly. It maintains a queue of waiting trains. In case of any train leaving a platform, it notifies one of the train to arrive on the platform next.


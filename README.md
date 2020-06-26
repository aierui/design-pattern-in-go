# Go 语言设计模式

在软件工程中，设计模式（design pattern）是对软件设计中普遍存在（反复出现）的各种问题，所提出的解决方案。这个术语是由埃里希·伽玛（Erich Gamma）等人在1990年代从建筑设计领域引入到计算机科学的。

设计模式并不直接用来完成代码的编写，而是描述在各种不同情况下，要怎么解决问题的一种方案。面向对象设计模式通常以类别或对象来描述其中的关系和相互作用，但不涉及用来完成应用程序的特定类别或对象。设计模式能使不稳定依赖于相对稳定、具体依赖于相对抽象，避免会引起麻烦的紧耦合，以增强软件设计面对并适应变化的能力。

并非所有的软件模式都是设计模式，设计模式特指软件“设计”层次上的问题。还有其他非设计模式的模式，如架构模式。同时，算法不能算是一种设计模式，因为算法主要是用来解决计算上的问题，而非设计上的问题。

随着软件开发社群对设计模式的兴趣日益增长，已经出版了一些相关的专著，定期召开相应的研讨会，而且沃德·坎宁安（Ward Cunningham）为此发明了WikiWiki用来交流设计模式的经验。

-- 摘录 wikipedia

## 设计模式类型

《设计模式》一书原先把设计模式分为创建型模式、结构型模式、行为型模式，把它们通过授权、聚合、诊断的概念来描述。

### 创建型模式

- [抽象工厂模式(Abstract Factory)](./creational/abstract-factory/)
- [工厂方法模式(Factory Method)](./creational/factory-method/)
- [生成器模式(Builder)](./creational/builder/)
- [惰性初始模式]()
- [对象池模式]()
- [原型模式(Prototype)](./creational/prototype/)
- [单例模式(Singleton)](./creational/singleton/)
- [多例模式]()

### 结构型模式

- [适配器模式(Adapter)](./structural/adapter/)
- [桥接模式(Bridge)](./structural/bridge/)
- [组合模式(Composite)](./structural/composite/)
- [修饰模式(Decorator)](./structural/decorator/)
- [外观模式(Facade)](./structural/facade/)
- [享元模式(Flywight)](./structural/flyweight/)
- [代理模式(Proxy)](./structural/proxy/)

### 行为型模式

- [责任链模式(Chain of Responsibility)](./behavioral/chain-of-responsibility/)
- [命令模式(Command)](./behavioral/command/)
- [解释器模式]()
- [迭代器模式(Iterator)](./behavioral/iterator/)
- [中介者模式(Mediator)](./behavioral/mediator/)
- [备忘录模式(Memento)](./behavioral/memento/)
- [空对象模式]()
- [观察者模式(Observer)](./behavioral/observer/)
- [策略模式(Strategy)](./behavioral/strategy/)
- [状态模式(State)](./behavioral/state/)
- [访问者模式(Vistor)](./behavioral/vistor/)
- [模版方法模式(Template Method)](./behavioral/template-method/)

## 设计模式的关系

## 设计模式六大原则

### 开闭原则（Open Close Principle）

开闭原则的意思是：对扩展开放，对修改关闭。在程序需要进行拓展的时候，不能去修改原有的代码，实现一个热插拔的效果。简言之，是为了使程序的扩展性好，易于维护和升级。想要达到这样的效果，我们需要使用接口和抽象类，后面的具体设计中我们会提到这点。

### 里氏代换原则（Liskov Substitution Principle）

里氏代换原则是面向对象设计的基本原则之一。 里氏代换原则中说，任何基类可以出现的地方，子类一定可以出现。LSP 是继承复用的基石，只有当派生类可以替换掉基类，且软件单位的功能不受到影响时，基类才能真正被复用，而派生类也能够在基类的基础上增加新的行为。里氏代换原则是对开闭原则的补充。实现开闭原则的关键步骤就是抽象化，而基类与子类的继承关系就是抽象化的具体实现，所以里氏代换原则是对实现抽象化的具体步骤的规范。

### 依赖倒转原则（Dependence Inversion Principle）

这个原则是开闭原则的基础，具体内容：针对接口编程，依赖于抽象而不依赖于具体。

### 接口隔离原则（Interface Segregation Principle）

这个原则的意思是：使用多个隔离的接口，比使用单个接口要好。它还有另外一个意思是：降低类之间的耦合度。由此可见，其实设计模式就是从大型软件架构出发、便于升级和维护的软件设计思想，它强调降低依赖，降低耦合。

### 迪米特法则，又称最少知道原则（Demeter Principle）

最少知道原则是指：一个实体应当尽量少地与其他实体之间发生相互作用，使得系统功能模块相对独立。

### 合成复用原则（Composite Reuse Principle）

合成复用原则是指：尽量使用合成/聚合的方式，而不是使用继承。

## 参考链接

- [Design Patterns](https://sourcemaking.com/design_patterns)
- [Design Patterns](https://refactoring.guru/design-patterns)
- [All Design Patterns in Go](https://golangbyexample.com/all-design-patterns-golang/)
- [设计模式](https://zh.wikipedia.org/wiki/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F_(%E8%AE%A1%E7%AE%97%E6%9C%BA))
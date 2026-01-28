> When we talk about software architecture, software is recursive and fractal in nature,

> Not only does a good architecture meet the needs of its users, developers, and owners at a given point in time, but it also meets them over time.

> The only way to go fast, is to go well.

> When software is done right, it requires a fraction of the human resources to create and maintain.

> The goal of software architecture is to minimize the human resources required to build and maintain the required system.

> These developers buy into a familiar lie: “We can clean it up later; we just have to get to market first!”

> Notice also that work on the TDD days proceeded approximately 10% faster than work on the non-TDD days, and that even the slowest TDD day was faster than the fastest non-TDD day.

> The developers may think that the answer is to start over from scratch and redesign the whole system—but that’s just the Hare talking again. The same overconfidence that led to the mess is now telling them that they can build it better if only they can start the race over.

> Every software system provides two different values to the stakeholders: behavior and structure.

> When the stakeholders change their minds about a feature, that change should be simple and easy to make. The difficulty in making such a change should be proportional only to the scope of the change, and not to the shape of the change.

> If you give me a program that works perfectly but is impossible to change, then it won’t work when the requirements change, and I won’t be able to make it work. Therefore the program will become useless. If you give me a program that does not work but is easy to change, then I can make it work, and keep it working as requirements change. Therefore the program will remain continually useful.

> The dilemma for software developers is that business managers are not equipped to evaluate the importance of architecture. That’s what software developers were hired to do. Therefore it is the responsibility of the software development team to assert the importance of architecture over the urgency of features.

> Remember, as a software developer, you are a stakeholder.

> Structured programming imposes discipline on direct transfer of control.  
> Object-oriented programming imposes discipline on indirect transfer of control.  
> Functional programming imposes discipline upon assignment.  
> Each of the paradigms removes capabilities from the programmer. None of them adds new capabilities. Each imposes some kind of extra discipline that is negative in its intent. The paradigms tell us what not to do, more than they tell us what to do.

> We use polymorphism as the mechanism to cross architectural boundaries; we use functional programming to impose discipline on the location of and access to data; and we use structured programming as the algorithmic foundation of our modules.

> A program of any complexity contains too many details for a human brain to manage without help. Overlooking just one small detail results in programs that may seem to work, but fail in surprising ways.

> That is the nature of scientific theories and laws: They are falsifiable but not provable.

> Dijkstra once said, “Testing shows the presence, not the absence, of bugs.”

> [S]oftware is like a science. We show correctness by failing to prove incorrectness, despite our best efforts.

> It’s fair to say that while OO languages did not give us something completely brand new, it did make the masquerading of data structures significantly more convenient.

> The bottom line is that polymorphism is an application of pointers to functions. Programmers have been using pointers to functions to achieve polymorphic behavior since Von Neumann architectures were first implemented in the late 1940s.

> Using an OO language makes polymorphism trivial. That fact provides an enormous power that old C programmers could only dream of.

> The fact that OO languages provide safe and convenient polymorphism means that any source code dependency, no matter where it is, can be inverted.

> OO is the ability, through the use of polymorphism, to gain absolute control over every source code dependency in the system. It allows the architect to create a plugin architecture, in which modules that contain high-level policies are independent of modules that contain low-level details. The low-level details are relegated to plugin modules that can be deployed and developed independently from the modules that contain high-level policies.

> All race conditions, deadlock conditions, and concurrent update problems are due to mutable variables. You cannot have a race condition or a concurrent update problem if no variable is ever updated. You cannot have deadlocks without mutable locks.

> Structured programming is discipline imposed upon direct transfer of control. Object-oriented programming is discipline imposed upon indirect transfer of control. Functional programming is discipline imposed upon variable assignment.

> The rules of software are the same today as they were in 1946, when Alan Turing wrote the very first code that would execute in an electronic computer. The tools have changed, and the hardware has changed, but the essence of software remains the same.

> The number of functions required to calculate pay, generate a report, or save the data is likely to be large in each case. Each of those classes would have many private methods in them. Each of the classes that contain such a family of methods is a scope.

> The Single Responsibility Principle is about functions and classes—but it reappears in a different form at two more levels. At the level of components, it becomes the Common Closure Principle. At the architectural level, it becomes the Axis of Change responsible for the creation of Architectural Boundaries.

> This is how the OCP works at the architectural level. Architects separate functionality based on how, why, and when it changes, and then organize that separated functionality into a hierarchy of components. Higher-level components in that hierarchy are protected from the changes made to lower-level components.

> The OCP is one of the driving forces behind the architecture of systems. The goal is to make the system easy to extend without incurring a high impact of change. This goal is accomplished by partitioning the system into components, and arranging those components into a dependency hierarchy that protects higher-level components from changes in lower-level components.

> It is the volatile concrete elements of our system that we want to avoid depending on. Those are the modules that we are actively developing, and that are undergoing frequent change.

> Every change to an abstract interface corresponds to a change to its concrete implementations. Conversely, changes to concrete implementations do not always, or even usually, require changes to the interfaces that they implement. Therefore interfaces are less volatile than implementations.

> DIP violations cannot be entirely removed, but they can be gathered into a small number of concrete components and kept separate from the rest of the system.

> Regardless of how they are eventually deployed, well-designed components always retain the ability to be independently deployable and, therefore, independently developable.

> We are now living in the age of software reuse—a fulfillment of one of the oldest promises of the object-oriented model.

> The Reuse/Release Equivalence Principle (REP) is a principle that seems obvious, at least in hindsight. People who want to reuse software components cannot, and will not, do so unless those components are tracked through a release process and are given release numbers.

> Classes and modules that are grouped together into a component should be releasable together. The fact that they share the same version number and the same release tracking, and are included under the same release documentation, should make sense both to the author and to the users.

> The Common Closure Principle — Gather into components those classes that change for the same reasons and at the same times. Separate into different components those classes that change at different times and for different reasons.

> Just as the SRP says that a class should not contain multiples reasons to change, so the Common Closure Principle (CCP) says that a component should not have multiple reasons to change.

> For most applications, maintainability is more important than reusability.

> Gather together those things that change at the same times and for the same reasons. Separate those things that change at different times or for different reasons.

> The Common Reuse Principle — Don’t force users of a component to depend on things they don’t need.

> The REP and CCP are inclusive principles: Both tend to make components larger. The CRP is an exclusive principle, driving components to be smaller. It is the tension between these principles that good architects seek to resolve.

> The Acyclic Dependencies Principle — Allow no cycles in the component dependency graph.

> It is always possible to break a cycle of components and reinstate the dependency graph as a DAG.

> Indeed, as the application grows, the component dependency structure jitters and grows. Thus the dependency structure must always be monitored for cycles. When cycles occur, they must be broken somehow.

> The Stable Dependencies — Principle Depend in the direction of stability.

> A component with lots of incoming dependencies is very stable because it requires a great deal of work to reconcile any changes with all the dependent components.

> Three components depend on X, so it has three good reasons not to change. We say that X is responsible to those three components. Conversely, X depends on nothing, so it has no external influence to make it change. We say it is independent.

> No other components depend on Y, so we say that it is irresponsible. Y also has three components that it depends on, so changes may come from three external sources. We say that Y is dependent.

> One way is to count the number of dependencies that enter and leave that component. These counts will allow us to calculate the positional stability of the component. Fan-in: Incoming dependencies. This metric identifies the number of classes outside this component that depend on classes within the component. Fan-out: Outgoing dependencies. This metric identifies the number of classes inside this component that depend on classes outside the component. I: Instability: I = Fan-out / (Fan-in + Fan-out). This metric has the range [0, 1]. I = 0 indicates a maximally stable component. I = 1 indicates a maximally unstable component.

> The SDP says that the I metric of a component should be larger than the I metrics of the components that it depends on. That is, I metrics should decrease in the direction of dependency.

> The Stable Abstractions Principle — A component should be as abstract as it is stable.

> The SAP and the SDP combined amount to the DIP for components. This is true because the SDP says that dependencies should run in the direction of stability, and the SAP says that stability implies abstraction. Thus dependencies run in the direction of abstraction.

> The A metric is a measure of the abstractness of a component. Its value is simply the ratio of interfaces and abstract classes in a component to the total number of classes in the component. Nc: The number of classes in the component. Na: The number of abstract classes and interfaces in the component. A: Abstractness. A = Na ÷ Nc.

> If it is desirable for components to be on, or close, to the Main Sequence, then we can create a metric that measures how far away a component is from this ideal. D3: Distance. D = |A+I–1| . The range of this metric is [0, 1]. A value of 0 indicates that the component is directly on the Main Sequence. A value of 1 indicates that the component is as far away as possible from the Main Sequence.

> Software architects are the best programmers, and they continue to take programming tasks, while they also guide the rest of the team toward a design that maximizes productivity.

> There are many systems out there, with terrible architectures, that work just fine. Their troubles do not lie in their operation; rather, they occur in their deployment, maintenance, and ongoing development.

> Good architecture makes the system easy to understand, easy to develop, easy to maintain, and easy to deploy. The ultimate goal is to minimize the lifetime cost of the system and to maximize programmer productivity.

> Different team structures imply different architectural decisions.

> [A] component-per-team architecture is not likely to be the best architecture for deployment, operation, and maintenance of the system. Nevertheless, it is the architecture that a group of teams will gravitate toward if they are driven solely by development schedule.

> The fact that hardware is cheap and people are expensive means that architectures that impede operation are not as costly as architectures that impede development, deployment, and maintenance.

> By separating the system into components, and isolating those components through stable interfaces, it is possible to illuminate the pathways for future features and greatly reduce the risk of inadvertent breakage.

> A good architect pretends that the decision has not been made, and shapes the system such that those decisions can still be deferred or changed for as long as possible.

> Good architects carefully separate details from policy, and then decouple the policy from the details so thoroughly that the policy has no knowledge of the details and does not depend on the details in any way.

> [A] good architecture must support: The use cases and operation of the system. The maintenance of the system. The development of the system. The deployment of the system.

> The most important thing a good architecture can do to support behavior is to clarify and expose that behavior so that the intent of the system is visible at the architectural level.

> A shopping cart application with a good architecture will look like a shopping cart application. The use cases of that system will be plainly visible within the structure of that system.

> A good architecture helps the system to be immediately deployable after build.

> A good architecture makes the system easy to change, in all the ways that it must change, by leaving options open.

> Use cases are a very natural way to divide the system.

> Thus, as we are dividing the system in to horizontal layers, we are also dividing the system into thin vertical use cases that cut through those layers.

> Architects often fall into a trap—a trap that hinges on their fear of duplication.

> There is true duplication, in which every change to one instance necessitates the same change to every duplicate of that instance. Then there is false or accidental duplication.

> My preference is to push the decoupling to the point where a service could be formed. should it become necessary; but then to leave the components in the same address space as long as possible. This leaves the option for a service open.

> A good architecture will allow a system to be born as a monolith, deployed in a single file, but then to grow into a set of independently deployable units, and then all the way to independent services and/or micro-services.

> Which kinds of decisions are premature? Decisions that have nothing to do with the business requirements—the use cases—of the system.

> Boundaries are drawn where there is an axis of change. The components on one side of the boundary change at different rates, and for different reasons, than the components on the other side of the boundary.

> A strict definition of “level” is “the distance from the inputs and outputs.” The farther a policy is from both the inputs and the outputs of the system, the higher its level.

> The Entity is pure business and nothing else.

> You don’t need to use an object-oriented language to create an Entity. All that is required is that you bind the Critical Business Data and the Critical Business Rules together in a single and separate software module.

> Use cases contain the rules that specify how and when the Critical Business Rules within the Entities are invoked. Use cases control the dance of the Entities.

> Use cases do not describe how the system appears to the user. Instead, they describe the application-specific rules that govern the interaction between the users and the Entities. How the data gets in and out of the system is irrelevant to the use cases.

> Why are Entities high level and use cases lower level? Because use cases are specific to a single application and, therefore, are closer to the inputs and outputs of that system. Entities are generalizations that can be used in many different applications, so they are farther from the inputs and outputs of the system. Use cases depend on Entities; Entities do not depend on use cases.

> You might be tempted to have these data structures contain references to Entity objects. You might think this makes sense because the Entities and the request/response models share so much data. Avoid this temptation! The purpose of these two objects is very different. Over time they will change for very different reasons, so tying them together in any way violates the Common Closure and Single Responsibility Principles. The result would be lots of tramp data, and lots of conditionals in your code.

> So what does the architecture of your application scream?

> If your architecture is based on frameworks, then it cannot be based on your use cases.

> The first concern of the architect is to make sure that the house is usable—not to ensure that the house is made of bricks. Indeed, the architect takes pains to ensure that the homeowner can make decisions about the exterior material (bricks, stone, or cedar) later, after the plans ensure that the use cases are met.

> The web is a delivery mechanism—an IO device—and your application architecture should treat it as such.

> - Hexagonal Architecture (also known as Ports and Adapters), developed by Alistair Cockburn, and adopted by Steve Freeman and Nat Pryce
> - DCI from James Coplien and Trygve Reenskaug
> - BCE, introduced by Ivar Jacobson

> An entity can be an object with methods, or it can be a set of data structures and functions. It doesn’t matter so long as the entities can be used by many different applications in the enterprise.

> The software in the use cases layer contains application-specific business rules. It encapsulates and implements all of the use cases of the system.

> The software in the interface adapters layer is a set of adapters that convert data from the format most convenient for the use cases and entities, to the format most convenient for some external agency such as the database or the web.

> We take advantage of dynamic polymorphism to create source code dependencies that oppose the flow of control so that we can conform to the Dependency Rule, no matter which direction the flow of control travels.

> The important thing is that isolated, simple data structures are passed across the boundaries. We don’t want to cheat and pass Entity objects or database rows.

> First, let’s get something straight: There is no such thing as an object relational mapper (ORM). The reason is simple: Objects are not data structures.

> This kind of anticipatory design is often frowned upon by many in the Agile community as a violation of YAGNI: “You Aren’t Going to Need It.” Architects, however, sometimes look at the problem and think, “Yeah, but I might.”

> It is one of the functions of an architect to decide where an architectural boundary might one day exist, and whether to fully or partially implement that boundary.

> You don’t simply decide at the start of a project which boundaries to implement and which to ignore. Rather, you watch. You pay attention as the system evolves. You note where boundaries may be required, and then carefully watch for the first inkling of friction because those boundaries don’t exist.

> The architecture of a system is defined by boundaries that separate high-level policy from low-level detail and follow the Dependency Rule.

> Changes to common system components can cause hundreds, or even thousands, of tests to break. This is known as the Fragile Tests Problem.

> The goal is to decouple the structure of the tests from the structure of the application.

> Structural coupling is one of the strongest, and most insidious, forms of test coupling.

> Firmware does not mean code lives in ROM. It’s not firmware because of where it is stored; rather, it is firmware because of what it depends on and how hard it is to change as hardware evolves.

> Non-embedded engineers also write firmware! You non-embedded developers essentially write firmware whenever you bury SQL in your code or when you spread platform dependencies throughout your code. Android app developers write firmware when they don’t separate their business logic from the Android API.

> The HAL exists for the software that sits on top of it, and its API should be tailored to that software’s needs. […] The HAL provides a service, and it does not reveal to the software how it does it.

> To give your embedded code a good chance at a long life, you have to treat the operating system as a detail and protect against OS dependencies.

> The database is a utility that provides access to the data. From the architecture’s point of view, that utility is irrelevant because it’s a low-level detail—a mechanism. And a good architect does not allow low-level mechanisms to pollute the system architecture.

> Many data access frameworks allow database rows and tables to be passed around the system as objects. Allowing this is an architectural error. It couples the use cases, business rules, and in some cases even the UI to the relational structure of the data.

> Yes, we need to get the data in and out of the data store quickly, but that’s a low-level concern. We can address that concern with low-level data access mechanisms. It has nothing whatsoever to do with the overall architecture of our systems.

> The data is significant. The database is a detail.

> We can’t seem to figure out where we want the computer power. We go back and forth between centralizing it and distributing it. And, I imagine, those oscillations will continue for some time to come.

> The GUI is a detail. The web is a GUI. So the web is a detail.

> The relationship between you and the framework author is extraordinarily asymmetric. You must make a huge commitment to the framework, but the framework author makes no commitment to you whatsoever.

> I’d personally like to use the compiler to enforce my architecture if at all possible.

> My definition of a component is slightly different: “A grouping of related functionality behind a nice clean interface, which resides inside an execution environment like an application.”

> Marking all of your types as public means you’re not taking advantage of the facilities that your programming language provides with regard to encapsulation.

> Like playing pool, each shot isn’t just about sinking that ball; it’s also about lining up the next shot.

> I was left with the realization that software architectures can be wildly different, yet equally effective.

> Architecture? Are you joking? This was a startup. We didn’t have time for architecture. Just code, dammit! Code for your very lives!

> Architecture must be flexible enough to adapt to the size of the problem.

> Architecting for the enterprise, when all you really need is a cute little desktop tool, is a recipe for failure.

> You can’t make a reusable framework until you first make a usable framework. Reusable frameworks require that you build them in concert with several reusing applications.

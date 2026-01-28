A must-read for anyone working with system architecture. Packed with so many useful tools and concrete frameworks, apart from the very interesting mental models it bestows.

What made me enjoy it so much was how it connected high-level architectural principles with low-level code design, showing that the same ideas apply across different scales. It reinforced my belief that good architecture is not just about structure but about enabling future change and adaptability. While it didn't introduce a whole lot of new concepts for me, it articulated many ideas I've encountered before in a clear and structured way, connecting them into a cohesive whole.

It does an excellent job of balancing abstract principles with practical advice, making it applicable to real-world scenarios.

- It's very cool how the author connects principles on different abstraction levels, showing that functions, modules, and systems are all ruled by the same design principles.

- I liked how he put words on something I've intuitively come to feel, that good architecture creates the structure that enables and simplifies changes in the future, which are inevitable with any system.

- One thing that I kept thinking about is the old adage that an ounce of prevention is worth a pound of cure, which holds true when working with architecture. Taking shortcuts only makes it take longer in the long run.

- I loved the formulas and concrete rules to understand and control dependency.

- This is not a book for beginners; it assumes a fair amount of prior knowledge and experience in software development and architecture.

- While it talks about abstract principles, it maintains a practical focus throughout, which I appreciated.

- One powerful concept that stuck with me is "accidental duplication", which refers to the seemingly but not actually duplicated code that arises from similar but slightly different requirements. This concept really helped me approach DRY more nuancedly, to create less coupling without over-engineering.

- Robert's prescription for making your architecture "look like" the things it's supposed to represent was enlightening. A blog should look like a blog, a shop like a shop, etc. This idea of aligning architecture with domain concepts reinforces the importance of working to solve real problems.

- One concrete thing I disliked (although, admittedly, a very minor thing) was the prescription to write APIs specific for tests, for hard-to-test cases. While I understand the intent, it felt like adding unnecessary complexity to the codebase just for the sake of avoiding the real problem of shitty architecture.

- The book made me change my mind on aligning tests with the code it tests. Simply naming it "structural coupling" highlighted the problem, and I have since written my tests to independantly from the code they test.

- Moving outside of my own experience, using hardware and firmware as examples to illustrate architectural principles was eye-opening. It showed how these concepts are universal across different domains.

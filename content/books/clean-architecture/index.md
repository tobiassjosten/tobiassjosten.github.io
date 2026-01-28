---
title: "Clean Architecture"
authors:
  - "robert-c-martin"
date: "2019-10-26"
amazonURL: "https://amzn.to/4bOR3VM"
image: "clean-architecture.jpg"
rating: 5
featuredOnHomepage: true
---

Clean Architecture is a must-read for anyone working with software systems. Robert C. Martin delivers a dense but rewarding exploration of how to structure software that remains adaptable over time. The book articulates something I've felt intuitively throughout my career: that good architecture isn't about picking the right framework or database, but about enabling future change.

What makes this book stand out is how Martin connects principles across different levels of abstraction. The same ideas that govern how you write functions apply to modules, components, and entire systems. SOLID principles, familiar to most developers at the class level, reappear transformed at higher levels: the Single Responsibility Principle becomes the Common Closure Principle for components, and eventually shapes how you draw architectural boundaries. This recursive, fractal nature of software design was eye-opening.

Martin's central thesis is straightforward: the goal of architecture is to minimize the human resources required to build and maintain a system. Good architecture separates policy from detail, making the important decisions visible while keeping implementation choices deferrable. A well-architected system should scream its purpose. A shopping cart application should look like a shopping cart application, not like a Rails app or a Spring app.

One concept that changed how I think about code is "accidental duplication." Not all similar-looking code is true duplication. Sometimes two pieces of code look the same today but will evolve differently because they serve different purposes. Recognizing this distinction has helped me apply DRY more carefully and avoid coupling things that shouldn't be coupled.

The book also changed my approach to testing. Martin's term "structural coupling" crystallized a problem I'd noticed but couldn't name: tests that mirror the structure of the code become fragile and resist refactoring. I've since written tests that focus on behavior rather than implementation, giving me more confidence when changing code.

This is not a book for beginners. Martin assumes familiarity with object-oriented programming, design patterns, and real-world development experience. The examples from hardware and firmware, showing how architectural principles transcend software, require patience but reward careful reading.

The book balances abstract principles with practical advice throughout. Martin's formulas for measuring component stability and abstractness give you concrete tools to evaluate your own systems. Whether you're designing a new system or trying to understand why an existing one is painful to work with, Clean Architecture provides the vocabulary and frameworks to reason about it clearly.

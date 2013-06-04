practical-chrestomathies
========================

How to do task T in language L

As a polyglot programmer, I am a huge fan of http://hyperpolyglot.org -- it
contains phrase-level chrestomathies (see
http://en.wikipedia.org/wiki/Chrestomathy) which are useful for quick lookup,
as well as for the simple pleasure of comparison and contrast. Yet beyond that,
I often enjoy having complete programs (not just fragments) to clone, as
productivity accelerators: a common example being how to do GNU-style getopt in
various languages (see
https://github.com/johnkerl/scripts/tree/master/one-offs). Looking at a wiki
page somewhere is a nice thing; dropping working code from one file into
another is far more useful.

In this repo are complete programs doing useful tasks in various languages.
The set of languages is small and arbitrary: I already have most all these
concepts developed in Ruby, Python, C, and Java -- see
https://github.com/johnkerl/scripts and https://github.com/johnkerl/ctools.
Here I focus on **new** languages, in particular those supporting modern closure
and concurrency paradigms: the JVM languages **Scala**, **Clojure**, and **Groovy**;
the systems languages **Go** and **Rust**; and the concurrency grandfather-language
**Erlang**.

The paradigms I explore are ones that are standard for me -- sometimes these
are the fun parts, and sometimes they are the necessary scaffolding. No matter
how appealing a language is, if I can't do basic things with such as troll
through log files a line at a time or piping to a subprocess, it isn't useful
to me.  Thus the focus here is on **practicality**, with an ongoing effort toward
**idiomaticity**. I prefer to get something up and running in a reasonably
idiomatic way; I'm sure as time goes by I'll find more idiomatic ways to do
things. Example: a checksummer in Go is single-threaded (as of this writing)
but would benefit (in performance, as well as pedagogically) from overlapped
computation and I/O: each in its own goroutine.

Scope is limited to each language's **standard library**, with an exception made
for GNU-style getopt. I know there are lots of cool packages out there for lots
of languages; getting into those would unacceptably enlarge the scope of this
project.

This is and will probably always be a **work in progress**. Most of the content
at present is Scala and Go.

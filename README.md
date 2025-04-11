Critaique
=========

CritAIque. Get it? It's the same thing as copying the text of an HTML or Markdown document into some AI chat window and asking it for feedback.

This does not facilitate AI writing for you. No one wants to read that. Knock it off. But as a first pass editor? Meh, OK. So Critaique prompts an LLM to make suggestions for improving your writing, ask follow-up questions, etc.

Installation
------------

```sh
go install github.com/rcrowley/critaique@latest
```

Usage
-----

```sh
critaique <input>
```

* `<input>`: input HTML or Markdown document

See also
--------

Critaique is part of the [Mergician](https://github.com/rcrowley/mergician) suite of tools that manipulate HTML documents:

* [Deadlinks](https://github.com/rcrowley/deadlinks): Scan a document root directory for dead links
* [Electrostatic](https://github.com/rcrowley/electrostatic): Mergician-powered, pure-HTML CMS
* [Feed](https://github.com/rcrowley/feed): Scan a document root directory to construct an Atom feed
* [Frag](https://github.com/rcrowley/frag): Extract fragments of HTML documents
* [Sitesearch](https://github.com/rcrowley/sitesearch): Index a document root directory and serve queries to it in AWS Lambda

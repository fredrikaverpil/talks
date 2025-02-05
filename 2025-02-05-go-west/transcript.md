# Transcript, Go West

## PREPARE

- Disable Aerospace
- Terminal tab with `gowest` project (resize to fully show LSP page)
- Terminal tab on laptop with transcript/notes
- Arc browser (behind terminal) with opened tabs:
  - AIP
  - AEP

### Intro

Hi everyone, and welcome to Einride!

I’m Fredrik and I’m going to talk about developing with Go in Neovim - and as
part of that I’ll also talk about a pet project I’ve been working on in my spare
time. At the end of this session there’s time for Q&A but feel free to interrupt
me if you have any questions too.

Ok - who am I?

Back when I was 12 yo, I had an Amiga. My friends and I drew images in Deluxe
Paint and wrote small adventure games in a programming language called Amos.

Some years later, my family got a 486 which had really nice graphics. Apart from
playing games I remember spending an awful amount of time configuring MS-DOS the
system to play these games. I also recall messing around with BASIC but I can’t
remember why.

When I was 16, I accidentally got into web design after having sent in artwork
to be printed onto a mouse pad and got a job offer.

In 2001 I found myself living in New York City, working with Macromedia Flash
and the now long abandoned Javascript version, which would have had runtime
checking if it would’ve taken off.

Let’s cut to the chase. About one year ago I joined Einride and I’m now working
in the backend of our digital freight platform. I’ve been a Neovim user since a
few years and I’d like to share my experience developing with Go with it!

## Quick intro

- Do we have any Neovim users in here, raise hand?
- A distinct difference from Vim is being able to use Lua

### Why Neovim

- Customisable
- Fun
- Deeper understanding of tools
- Disclaimer

### Options

- Lets jump right in;
- Single init.lua or build it out
- Representing tabs
- Per-filetype basis
- Per-project basis
- Function: Include your own logic, powerful

### Lazy

- Most editors have an extensions store or plugin manager built in
- Neovim doesn’t have that, you have to pick, I recommend
- You define the plugin in Lua;
  - Github repository, omitting github.com
  - Optional configuration, if the plugin requires it

### Mason

- Let’s move on to the tooling you may need installed and on PATH - can be done
  with ”Mason”
- You might prefer other ways to do this, like using Nix
- —
- Something note worthy is this project which adds a lock file to Mason
- Gothenburg

## Essentials

### Tree-sitter

- Tree-sitter is a central piece of software in the Neovim ecosystem.
- Provides syntax highlighting, help out with indentation and folding
- Syntax highlight one language injected into another language
  - SQL inside Go strings
  - Go templates with HTML and the Go directives inside

### LSP

- Let’s dive into the language server config which means setting up the client
  and server.
- Used to be quite painful to set up, better with Neovim 0.11.
- Starting off, you have the command
- Filetypes define when the LSP should be attached to a buffer
- …
- Agree on the combined capabilities of the client and server.
- …

### Code Lenses

- Shown inline, allows you to take actions

### Code actions

- Similar ot code lens, you don’t see anything inline in the buffer. Called
  ”quickfix” in vscode.

### Completion

- We touched briefly upon completion earlier. This plugin provides the
  functionality of handling auto-completion as you type.
- There are others to pick from.
- Uses SIMD which makes it really fast

  - a type of parallel processing; Single Instruction, Multiple Data

- Expand shorthand into block of code
- Step through parts to fill in variables

### Formatting

- Up until this point I’m not sure if you’re impressed much.
- We’re halfway through the presentation and I hope this is also a turning point
  where you might be able to see where Neovim shines.
- Plugin like many other, with the ability to chain commands in sequence.
- Customize this to your heart’s content
  - Prepend args
  - Append args
  - Write functions containing your logic

### Linting

- Similar story
- Golangci lint; function searching for yaml file in the project, fallback to a
  yaml file in my dotfiles folder.
- —
- At Einride we lint our protobuf files with Google’s api-linter.
- Lints for AIP rule violations.
- AIP is API Improvement Proposals, but we treat them as rules.
- OPEN AIP WEBSITE
- OPEN AEP WEBSITE
- Approx 100 LOC

### Neotest

- Most popular test framework for Neovim.
- Files and folder tree which you can interact with
- We’ll come back to this in the demo, which we are approaching

### Debugging

- These plugins provides an API and a user interface for debugging, based on the
  DAP protocol.
- We’ll come back for this one as well, for the demo

### Extra: Goplements

- Just included this because it’s very convenient and another showcase of a
  custom plugin made by an individual.

## DEMO

- cd gowest && nvim
- Cmd
  - Shift-K on Printf
  - Gd on Printf
  - Gr on fib
  - CODE LENS
    - Go.mod
  - CODE ACTIONS:
    - Break out fib to separate file
    - Create test for fib
    - Ask Claude to make test cases.
    - Run individual table test
  - COVERAGE
  - Open test summary
  - Run all tests in file
  - Run all tests in suite
    - Go to failing test, show diagnostics
  - Check coverage of fib.go and main.go
  - Go to cmd/main.go
    - Set break point + <leader>dc (run cli)

## Conclusion

- Make it into what you need.
- Replace what I just showed you with another language and other tools, and it
  behaves the same way.
- Friendly community, lots of toy projects that bring value.
- Writing Go in Neovim is a joy.

- Other specialised IDEs will provide convenience functionality you’ll have to
  build yourself.
- Elephant in the room; the time I’ve spent on setting Neovim up like this.
  - Worth it? Depends on you enjoying yourself while doing so or not.
- Full neovim config in dotfiles.

# Neotest adapter

## Intro

- Let’s shift gears and talk about my toy project.
- So you saw how we used the Neatest framework earlier to run test, and it was
  this adapter project I ran in order to run Go tests.

### Why a new adapter?

- There was an adapter available when I started working with Go professionally.
- Too many issues
- How hard can it be?

### Interface

- Leaned into it, had a look at the interface.
- It described what the adapter needs to do.
  - Step through points
- Doesn’t sound too hard?

## A glimpse into the dev process

### Detect tests

- Tree-sitter to detect the tests.
- Tree-sitter query language.
- Neotest framework takes the query and generates the file tree.

### Generate test command

- However, go test does not speak in terms of file paths.
- Solve that later, find workaround?
- You can give it test filenames, using a regular expression.

### Executing individual test

- Can I come up with just one command to rule them all?
- Might end up doing excessive compilation of test binaries.

#### Interactive

- Here is a test file, in a go project. Looks like this.
- Two tests in there. I just want to run one.
- Change the CWD to the dir of the file, I can limit the compilation a bit.
- Nice, that worked - but not very structured output.

- Run passing test
- Run failing test

### v0.1.0 can ship

- Nice, the MVP works.
- The good: it's quite robust. Simple implementation. No major issues.
- Room for improvements: No optimization, no debugging, customizations (build
  tags).
- Next: optimizations (run file/folder).

### Go list

- Translate path to import path.

#### Interactive

- Show file tree
- Run `go list` on it

- Adding `-json` gives me what I need.

### Current status / future plans

- Today, the adapter is in really good shape.
- Provides everything I personally need (and more)
- I had requests; I implemented and some opened PRs.

- Future additions could include runtime generated tests.
- If you would like to try it out, find a bug or would like to file a request,
  please do!

### Thanks

- Thank you so much for listening.

- Questions?
